package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      string  `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname  string `json:"fullname"`
	Portfolio string `json:"portfolio"`
}

// sample database (seeding - process of seed fake values to work with)
var courses []Course

// middleware or helper functions
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("apiBuilding in golang")

	// seeding
	courses = append(courses, Course{
		CourseId: "1", CourseName: "React Js", Price: "2999", Author: &Author{Fullname: "Prathya Thakur", Portfolio: "https://thakurprathya-portfolio.netlify.app"},
	})
	courses = append(courses, Course{
		CourseId: "2", CourseName: "Node Js", Price: "2999", Author: &Author{Fullname: "Prathya Thakur", Portfolio: "https://thakurprathya-portfolio.netlify.app"},
	})

	// routing
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// setting up web server
	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go backend</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "applcation/json")

	json.NewEncoder(w).Encode(courses) // whatever's inside encode will be treated as json value and throwed back to writer
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "applcation/json")

	// extracting parameters from request
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode("No Course ID provided")
		return
	}

	// looping through courses & maping params id to return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding a course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Body found")
		return
	}

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse) // as passing newCourse as reference will have value inside it, so not storing in some variable

	// if body = {}
	if newCourse.isEmpty() {
		json.NewEncoder(w).Encode("Empty object found")
		return
	}

	// checking for duplicate entry
	for _, course := range courses {
		if course.CourseName == newCourse.CourseName {
			json.NewEncoder(w).Encode("Course Already Exists")
			return
		}
	}

	// generating a uniqueId and converting it to string
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	newCourse.CourseId = strconv.Itoa(rd.Intn(100))

	// appending new newCourse
	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Body found")
		return
	}

	var updatedCourse Course
	_ = json.NewDecoder(r.Body).Decode(&updatedCourse) // as passing updatedCourse as reference will have value inside it, so not storing in some variable

	// if body = {}
	if updatedCourse.isEmpty() {
		json.NewEncoder(w).Encode("Empty object found")
		return
	}

	// extracting parameters from request
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode("No Course ID provided")
		return
	}

	// looping through courses & maping params id to update record at matched id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			updatedCourse.CourseId = params["id"]
			courses[index] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting a course")
	w.Header().Set("Content-Type", "application/json")

	// extracting parameters from request
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode("No Course ID provided")
		return
	}

	// looping through courses & maping params id to remove current record
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// removing current record
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(fmt.Sprintf("course with id=%v is successfully deleted", params["id"]))
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
}
