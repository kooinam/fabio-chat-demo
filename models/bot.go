package models

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/actors"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

var BotsCollection *models.Collection

// Bot used to represent room data
type Bot struct {
	models.Base
	Actor  *actors.Actor
	roomID string
}

// MakeBot used to instantiate bot
func MakeBot(collection *models.Collection, args ...interface{}) models.Modellable {
	bot := &Bot{}

	bot.Initialize(collection)

	bot.Actor = fab.ActorManager().RegisterActor(bot.GetCollectionName(), bot)
	bot.roomID = args[0].(string)

	return bot
}

func (bot *Bot) RegisterActions(actionsHandler *actors.ActionsHandler) {
	actionsHandler.RegisterAction("Update", bot.update)
}

func (bot *Bot) update(context *actors.Context) error {
	var err error

	// broadcast message to room
	fab.ControllerManager().BroadcastEvent("chat", bot.roomID, "Message", nil, helpers.H{
		"message": "Update Here...",
	})

	return err
}
