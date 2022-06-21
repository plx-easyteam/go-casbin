package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// An Authorizer must implement a single function that,
// given a userID, an action and an asset, returns a bool
// indicating whether the user has permission to perform the action on the asset.
type Authorizer interface {
	HasPermission(userID, action, asset string) bool
}

func Middleware(a Authorizer) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			username, _, ok := r.BasicAuth()
			// This is where the password would normally be verified

			asset := mux.Vars(r)["asset"]
			action := actionFromMethod(r.Method)
			if !ok || !a.HasPermission(username, action, asset) {
				log.Printf("User '%v' not allowed '%v' resource '%v'", username, action, asset)
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}


// used to map HTTP methods to actions
func actionFromMethod(httpMethod string) string{
	switch httpMethod {
	case "GET":
		return "read"
	case "POST":
		return "create"
	case "PUT":
		return "update"
	case "DELETE":
		return "delete"
	default:
		return ""
	}
}

type tempTest struct {
	role string
	methods []string
}

func test(){
	newAdmin := tempTest{
		role: "ADMIN",
		methods: []string{"GET", "POST", "PUT", "DELETE"},
	}

	log.Println(newAdmin)
}