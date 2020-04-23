package controller

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPatter *regexp.Regexp
}

func (uc userController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from the User Controller!"))
}

func newUserController() *userController {
	return &userController{
		userIDPatter: regexp.MustCompile(`^/user/(\d+)/?`),
	}
}
