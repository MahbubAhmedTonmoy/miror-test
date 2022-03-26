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
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware , helper
func (c *Course) Validation() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>welcome to our API Project</p>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	log.Println(r.Body)
	//if body empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send necessy data")
	}

	//if json empty {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if course.Validation() {
		json.NewEncoder(w).Encode("please send valid data")
		return
	}

	//generate id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("update course id not found")
	return
}

func deleteCouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course delted")
			break
		}
	}
}
func main() {

	//1 start
	// log.Println("starting API server")
	// //create a new router
	// router := mux.NewRouter()
	// log.Println("creating routes")
	// //specify endpoints
	// router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	// router.HandleFunc("/persons", Persons).Methods("GET")
	// http.Handle("/", router)

	// //start and listen to requests
	// http.ListenAndServe(":8080", router)

	//1 end/

	//2nd start
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "react",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "mahbub",
			Website:  "m.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "c#",
		CoursePrice: 399,
		Author: &Author{
			Fullname: "mahbub",
			Website:  "m.com",
		},
	})
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourse).Methods("GET")
	router.HandleFunc("/delete/{id}", deleteCouse).Methods("DELETE")
	router.HandleFunc("/create", createCourse).Methods("POST")
	router.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

	//1 end

}

type Response struct {
	Persons []Person `json:"persons"`
}
type Person struct {
	Id        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	School    *Student `json:"school"`
}
type Student struct {
	SchoolName string `json:"school"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response Response
	persons := prepareResponse()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

//middleware , helper
func (c *Person) IsEmpty() bool {
	return c.Id == 0 && c.FirstName == ""
}

//slice of person
func prepareResponse() []Person {
	var persons []Person

	var person Person

	var school Student
	school.SchoolName = "BBN"
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "N"
	person.School = &school
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "E"
	persons = append(persons, person)

	person.Id = 3
	person.FirstName = "Thomas"
	person.LastName = "E"
	persons = append(persons, person)
	return persons
}
