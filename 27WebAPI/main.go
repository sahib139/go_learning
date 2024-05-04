package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var id int = 0

type Courses struct {
	CourseId    int     `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

var courses []Courses

func main() {
	fmt.Println("Basic API performing crud operation!")
	seeder()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	http.ListenAndServe(":4000", r)
}

func seeder() {
	id++
	courses = append(courses, Courses{CourseId: id, CourseName: "Node-JS", CoursePrice: 299,
		Author: &Author{FullName: "Sahib Singh", Website: "sahib.com"}})
	id++
	courses = append(courses, Courses{CourseId: id, CourseName: "go-lang", CoursePrice: 499,
		Author: &Author{FullName: "ABC", Website: "abc.com"}})
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to First api page</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	courseRequestedId, _ := strconv.Atoi(params["id"])
	for _, val := range courses {
		if val.CourseId == courseRequestedId {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	w.Write([]byte("No such course exist!"))
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Courses
	json.NewDecoder(r.Body).Decode(&course)
	id++
	course.CourseId = id
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseRequestId, _ := strconv.Atoi(params["id"])
	for idx, course := range courses {
		if course.CourseId == courseRequestId {
			courses = append(courses[:idx], courses[idx+1:]...)
			w.Write([]byte("Request course deleted!"))
			return
		}
	}
	w.Write([]byte("No such courseId exist!"))
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	courseRequestId, _ := strconv.Atoi(mux.Vars(r)["id"])
	var courseUpdateParameters Courses
	json.NewDecoder(r.Body).Decode(&courseUpdateParameters)
	for idx, course := range courses {
		if course.CourseId == courseRequestId {

			if courseUpdateParameters.CourseName != "" {
				courses[idx].CourseName = courseUpdateParameters.CourseName
			}
			if courseUpdateParameters.CoursePrice != 0 {
				courses[idx].CoursePrice = courseUpdateParameters.CoursePrice
			}
			if courseUpdateParameters.Author != nil {
				if courseUpdateParameters.Author.FullName != "" {
					courses[idx].Author.FullName = courseUpdateParameters.Author.FullName
				}
				if courseUpdateParameters.Author.Website != "" {
					courses[idx].Author.Website = courseUpdateParameters.Author.Website
				}
			}
			json.NewEncoder(w).Encode(courses[idx])
			return
		}
	}
	w.Write([]byte("No such courseId exist!"))
}
