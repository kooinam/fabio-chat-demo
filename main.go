package main

import (
	"fabio-chat-demo/controllers"
	"fabio-chat-demo/models"
	"net/http"

	fab "github.com/kooinam/fabio"
)

func main() {
	fab.Setup()

	models.BotsCollection = fab.ModelManager().CreateCollection("bot", models.MakeBot)

	models.BotsCollection.Create("1")
	models.BotsCollection.Create("2")
	models.BotsCollection.Create("3")

	fab.ControllerManager().RegisterController("chat", &controllers.ChatController{})

	fab.ControllerManager().Serve("8000", func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
