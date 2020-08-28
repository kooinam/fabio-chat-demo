package main

import (
	"fabio-chat-demo/controllers"
	"fabio-chat-demo/models"
	"net/http"

	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/simplerecords"
)

func main() {
	fab.Setup()

	adapter := simplerecords.MakeAdapter()
	fab.ModelManager().RegisterAdapter("simple", adapter)

	models.BotsCollection = fab.ModelManager().RegisterCollection("simple", "bot", models.MakeBot)

	botsValues := []helpers.H{
		helpers.H{
			"roomID": "1",
		},
		helpers.H{
			"roomID": "2",
		},
		helpers.H{
			"roomID": "3",
		},
	}

	for _, botValues := range botsValues {
		bot, _ := models.BotsCollection.Create(botValues)

		bot.Memoize()
	}

	fab.ControllerManager().RegisterController("chat", &controllers.ChatController{})

	fab.ControllerManager().Serve("8000", func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
