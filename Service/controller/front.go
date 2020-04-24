package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterController() {
	uc := newUserController()
	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	end := json.NewEncoder(w)
	end.Encode(data)
}
