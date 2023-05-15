package main

import (
	"encoding/json"
	"net/http"
	"os/user"
)

type Info struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	usr, _ := user.Current()

	info := Info{
		Name:      usr.Name,
		AvatarURL: "https://raw.githubusercontent.com/ashleymcnamara/gophers/master/BLUE_GOPHER.png",
	}

	json.NewEncoder(w).Encode(info)
}
