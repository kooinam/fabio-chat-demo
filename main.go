package main

import (
	"fabio-chat-demo/controllers"
	"net/http"

	fab "github.com/kooinam/fabio"
)

func main() {
	fab.Setup()

	fab.ControllerManager().RegisterController("chat", &controllers.ChatController{})

	fab.ControllerManager().Serve("8000", func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
