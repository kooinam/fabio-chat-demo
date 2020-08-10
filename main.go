package main

import (
	"fabio-chat-demo/controllers"
	"net/http"

	fab "github.com/kooinam/fabio"
)

func main() {
	fab.Setup()

	fab.RegisterController("chat", &controllers.ChatController{})

	fab.Serve(func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
