package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResponse"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	var course model.Course

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&course)
	if err != nil {
		fmt.Printf("error in decoding the request: %s", err)
		httpResponse.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id := course.CourseId
	row := model.Check(id)
	if row == id {
		fmt.Println("already added or the course already exists")
		httpResponse.ResponseWithError(w, http.StatusConflict, "already exists in the database")
		return
	}
	addErr := course.InsertData()
	if addErr != nil {
		fmt.Println("error in inserting or adding data to the database")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, addErr.Error())
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusCreated, "added successfully")
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// Id,err := getUserID(id)
	// if err != nil {
	// 	fmt.Print("Error in converting string to integer")
	// 	httpResponse.ResponseWithError(w,http.StatusBadRequest,err.Error())
	// 	return
	// }
	var course model.Course
	getErr := course.GetInfo(id)
	if getErr != nil {
		fmt.Print("couldn't get the data from the database")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	id := mux.Vars(r)["id"]
	// Id, err := getUserID(id)
	// if err != nil {
	// 	fmt.Println("Error in converting string to integer",err)
	// 	httpResponse.ResponseWithError(w,http.StatusBadRequest,err.Error())
	// 	return
	// }
	var course model.Course
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&course)
	if err != nil {
		fmt.Print("error in decoding the request")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	updateErr := course.Update(id)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "The particular course is not registered in the database ")
		default:
			fmt.Print("Error in updating the data")
			httpResponse.ResponseWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, "updated successfully")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	id := mux.Vars(r)["id"]
	// Id, err := getUserID(id)
	// if err != nil {
	// 	fmt.Print("error in converting string to integer")
	// 	httpResponse.ResponseWithError(w,http.StatusBadRequest,err.Error())
	// 	return
	// }
	var course model.Course
	deleteErr := course.Delete(id)
	if deleteErr != nil {
		switch deleteErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "The course is not registered")
		default:
			fmt.Println("couldn't delete from the database")
			httpResponse.ResponseWithError(w, http.StatusBadRequest, deleteErr.Error())
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, "Deleted successfully")
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	courses, getErr := model.GetCourses()
	if getErr != nil {
		fmt.Print("error in getting the informaiton from the database")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, courses)
	fmt.Println("List for courses:", courses)
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["query"]
	fmt.Print(query)
	fmt.Println("it is same as we thought")
}
