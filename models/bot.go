package models

import (
	"fmt"
	"math/rand"

	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/actors"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
	"github.com/kooinam/fabio/simplerecords"
)

var BotsCollection *models.Collection

// Bot used to represent room data
type Bot struct {
	simplerecords.Base
	Actor  *actors.Actor
	roomID string
}

// MakeBot used to instantiate bot
func MakeBot(collection *models.Collection, hooksHandler *models.HooksHandler) models.Modellable {
	bot := &Bot{}

	hooksHandler.RegisterInitializeHook(bot.initialize)

	hooksHandler.RegisterAfterMemoizeHook(bot.afterMemoize)

	return bot
}

// initialize used to initialize bot
func (bot *Bot) initialize(attributes *helpers.Dictionary) {
	bot.roomID = attributes.ValueStr("roomID")
}

func (bot *Bot) afterMemoize() {
	fab.ActorManager().RegisterActor(bot.GetCollectionName(), bot)
}

// RegisterActorActions used to register actor's actions
func (bot *Bot) RegisterActorActions(actionsHandler *actors.ActionsHandler) {
	actionsHandler.RegisterAction("Update", bot.update)
}

func (bot *Bot) update(context *actors.Context) error {
	var err error

	// broadcast message to room
	fab.ControllerManager().BroadcastEvent("chat", bot.roomID, "Message", nil, helpers.H{
		"message": fmt.Sprintf("Update Here from bot #%v - %v...", bot.roomID, rand.Intn(100)),
	})

	return err
}
