package controllers

import "net/http"

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Get All"))
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Get"))
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Create"))
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Update"))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Delete"))
}
