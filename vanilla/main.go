package main

/*
import (
	"go-casbin/models"
	"log"
)


type authorizer struct {
	users models.Users
	roles Roles
}

func (a *authorizer) HasPermission(userID, action, asset string) bool {
	user, ok := a.users[userID]
	if !ok {
		// Unknown userID
		log.Print("Unknown user:", userID)
		return false
	}

	for _, roleName := range user.Roles {
		if role, ok := a.roles[roleName]; ok {
			resources, ok := role[action]
			if ok {
				for _, resource := range resources {
					if resource == asset {
						return true
					}
				}
			}
		} else {
			log.Printf("User '%s' has unknown role '%s'", userID, roleName)
		}
	}

	return false
}*/