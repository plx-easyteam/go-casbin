package main

import (
	"encoding/json"
	"os"
)

type Resources []string

type Actions map[string]Resources

type Roles map[string]Actions

func LoadRoles() (Roles, error) {
	var roles Roles

	if err := LoadJson("./roles.json", &roles); err != nil {
		return nil, err
	}

	return roles, nil
}

func LoadJson(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}