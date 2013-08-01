package client

import (
	"net/http"
)

type GithubAuth interface {
	Authenticate(request *http.Request)
}

type GithubBasicAuth struct {
	Username string
	Password string
}

func (ga *GithubBasicAuth) Authenticate(request *http.Request) {
	request.SetBasicAuth(ga.Username, ga.Password)
}
