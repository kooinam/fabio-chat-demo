package main

import (
	"fabio-chat-demo/controllers"
	"fabio-chat-demo/models"
	"net/http"

	Models "fabio-chat-demo/models"

	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/simplerecords"
)

func main() {
	fab.Setup()

	adapter := simplerecords.MakeAdapter()
	fab.ModelManager().RegisterAdapter("simple", adapter)

	models.BotsCollection = fab.ModelManager().RegisterCollection("simple", "bot", models.MakeBot)

	roomIDs := []string{"1", "2", "3"}
	for _, roomID := range roomIDs {
		result := models.BotsCollection.Create(helpers.H{
			"roomID": roomID,
		})

		if result.StatusSuccess() {
			bot := result.Item().(*Models.Bot)

			bot.Memoize()
		}
	}

	fab.ControllerManager().RegisterController("chat", &controllers.ChatController{})

	fab.ControllerManager().Serve("9000", func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
