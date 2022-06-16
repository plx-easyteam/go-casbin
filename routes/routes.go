package routes

import (
	"encoding/json"
	"go-casbin/controllers"
	"net/http"

	"github.com/gorilla/mux"
)


func Handler() *mux.Router{
	// init router
	r := mux.NewRouter().StrictSlash(true)
	r.Use(jsonHeader) // set default Content-Type

	// Route handles & endpoints
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	return r
}


func getHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func jsonHeader(next http.Handler)  http.Handler{
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
}


func greetings(w http.ResponseWriter, r *http.Request){
	// w.Write([]byte("Greetings & Salutations"))
	var g = map[string]interface{}{
		"int": 1,
		"string": "two",
		"boolean": true,
	}

	json.NewEncoder(w).Encode(g)
}

func about(w http.ResponseWriter, r *http.Request){
	// structs needs to be in capital letters
	type About struct {
		Who string
		Year int
	}

	about := About{"Some Dev", 2022}
	// b, _ := json.Marshal(about)

	json.NewEncoder(w).Encode(about)
}