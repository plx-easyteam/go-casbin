package middleware

import (
	"fmt"
	"net/http"
)

// func Authorize(obj string, act string, enforcer *casbin.Enforcer) http.Handler {
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		fmt.Println(":: Authorize ::")
		return
	})
}