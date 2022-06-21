package routes

import (
	"encoding/json"
	"fmt"
	"go-casbin/controllers"
	"log"
	"net/http"

	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"
	authz "github.com/casbin/mux-authz"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) *mux.Router {
	// Load rbac model and policy
	c := new(authz.CasbinAuthorizer)
	err := c.Load("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		panic(fmt.Sprintf("failed to load: %v", err))
	}

	// init router
	r := mux.NewRouter().StrictSlash(true)
	r.Use(jsonHeader) // set default Content-Type

	// Here was archiveBackup()

	// Route handles & endpoints
	r.Use(c.Middleware)
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/greetings", greetings).Methods("GET")

	log.Println("::: ::")
	return r
}

func getHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
}

func greetings(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Greetings & Salutations"))
	var g = map[string]interface{}{
		"int":     1,
		"string":  "two",
		"boolean": true,
	}

	json.NewEncoder(w).Encode(g)
}

func about(w http.ResponseWriter, r *http.Request) {
	// structs needs to be in capital letters
	type About struct {
		Who  string
		Year int
	}

	about := About{"Some Dev", 2022}
	// b, _ := json.Marshal(about)

	json.NewEncoder(w).Encode(about)
}

// func getPolicy(enforcer *casbin.Enforcer){
type Policy struct {
	sub string // the user that wants to access a resource.
	obj string // the resource that is going to be accessed.
	act string // the operation that the user performs on the resource.
}

func getPolicy() []Policy {
	data := "data"

	policies := []Policy{
		{
			sub: "ADMIN",
			obj: data,
			act: "read",
		},
		{
			sub: "ADMIN",
			obj: data,
			act: "write",
		},
		{
			sub: "CLIENT",
			obj: data,
			act: "read",
		},
	}

	return policies
}

/*
func archiveBackup() {

	// Init casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	// adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("failed to initialize casbin adapter: %v", err)
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	// enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	// enforcer.LoadPolicy()

	for _, pol := range getPolicy() {
		//add policy
		if hasPolicy := enforcer.HasPolicy(pol.sub, pol.obj, pol.act); !hasPolicy {
			enforcer.AddPolicy(pol.sub, pol.obj, pol.act)
		}
	}
}
*/