package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)


func UrlHandler(w http.ResponseWriter, r *http.Request) {
	result := mux.Vars(r)
	myname := result["name"]
	myage := result["age"]
	fmt.Println(myname)
	fmt.Println(result)
	res := []byte("My name is "+myname+" and i am known as kingpin in dark world and i am "+myage + " years old")
	_, err := w.Write(res)
	if err != nil {
		fmt.Println("couldn't write the response")
		os.Exit(1)
	}
}
// handler function for /. func ()
func HomeHandler(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

}

