package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Author struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Website  string `json:"website" bson:"website"`
}

type Course struct {
	CourseId   string  `json:"courseId" bson:"courseId"`
	CourseName string  `json:"courseName" bson:"courseName"`
	Price      int     `json:"price" bson:"price"`
	Author     *Author `json:"author" bson:"author"`
}

func main() {
	connectDB()

	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/course", AddCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courses []Course
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, map[string]interface{}{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var course Course
		cursor.Decode(&course)
		courses = append(courses, course)
	}
	json.NewEncoder(w).Encode(courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var course Course
	err := collection.FindOne(ctx,
		map[string]interface{}{"courseId": params["id"]},
	).Decode(&course)

	if err != nil {
		json.NewEncoder(w).Encode("Not Found")
		return
	}

	json.NewEncoder(w).Encode(course)
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, course)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := map[string]interface{}{
		"$set": course,
	}
	result, err := collection.UpdateOne(
		ctx,
		map[string]interface{}{"courseId": params["id"]},
		update,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.DeleteOne(ctx,
		map[string]interface{}{"courseId": params["id"]},
	)
	json.NewEncoder(w).Encode(result)
}
