package main

import (
	"strings"
	"log"
	"encoding/json"
)

type AuthToken struct {
	Token string
	Validity string
}

func getAuthToken() AuthToken {
	token := new(AuthToken)
	if data, err := getContentJson("http://auth:3005/auth"); err != nil {
		log.Printf("Failed to get auth-token: %v", err)
	} else {    
		json.NewDecoder(strings.NewReader(data)).Decode(token)
	}
	return *token
}

