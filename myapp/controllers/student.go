package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResponse"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	//verification of cookie
	if !VerifyCookie(w, r) {
		return
	}
	fmt.Println("kkkkkkkkkkkkkkkkkkkkkk")
	// fmt.Fprintln(w,"from add student handler")
	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "invalid json data")
		fmt.Println("error in decoding", err)
		// w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(stud)
	dberr := stud.Create()
	if dberr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, dberr.Error())
		fmt.Print("couldn't insert", dberr)
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusCreated, map[string]string{"message": "added successfully"})
	fmt.Println("added")

	// w.Write([]byte("successfully added"))
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	//get url parameter
	stdID := mux.Vars(r)["stdID"]
	stdid, idErr := getUserID(stdID)
	if idErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdID: stdid}
	fmt.Println(s)
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "Student not found")

		default:
			httpResponse.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusCreated, s)
}

// convert string stdID to int
func getUserID(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	return userId, userErr
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	StdID := mux.Vars(r)["stdid"]
	stdID, idErr := getUserID(StdID)
	if idErr != nil {
		fmt.Println("Error in converting string to integer")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		fmt.Println("error in decoding the request body")
		httpResponse.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	updateErr := stud.Update(stdID)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "Student is not registered in the database")
		default:
			httpResponse.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusCreated, stud)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	StdID := mux.Vars(r)["stdid"]
	stdid, idErr := getUserID(StdID)
	if idErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "Error in converting string to integer")
		return
	}
	var stud model.Student
	deleteErr := stud.Delete(stdid)
	if deleteErr != nil {
		switch deleteErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusNotFound, "Student is not registerd to the database")
		default:
			httpResponse.ResponseWithError(w, http.StatusInternalServerError, deleteErr.Error())
		}
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, "deleted the student")
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	//cookie verification
	if !VerifyCookie(w, r) {
		return
	}

	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResponse.ResponseWithJson(w, http.StatusOK, students)
}
