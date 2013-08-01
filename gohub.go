package gohub

import (
	"gohub/client"
	"gohub/user"
	"net/http"
)

type Gohub struct {
	*client.GithubClient
	Users user.UserController
}

// TODO: Add other authentication methods.
// Oauth, Token, etc...
func BasicAuthInitialize(username string, password string) *Gohub {
	auth := &client.GithubBasicAuth{username, password}
	client := &client.GithubClient{auth, http.Client{}}

	return &Gohub{client, user.UserController{client}}
}
