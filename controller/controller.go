package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bharath79/golang/model"
	"github.com/bharath79/golang/services"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	movies := services.GetAllMovies()
	json.NewEncoder(w).Encode(&movies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	id :=services.InsertMovie(movie)
	json.NewEncoder(w).Encode(id)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request)  {
w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")	

	params:=mux.Vars(r)
	services.UpdateMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	services.DeleteMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	services.DeleteMany()

	json.NewEncoder(w).Encode("All movies deleted")
}