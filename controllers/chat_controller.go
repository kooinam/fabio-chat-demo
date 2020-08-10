package controllers

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/controllers"
)

// ChatController used for chat purposes
type ChatController struct {
}

// AddBeforeActions used to add before actions callbacks
func (controller *ChatController) AddBeforeActions(callbacksHandler *controllers.CallbacksHandler) {
}

// AddActions used to add actions
func (controller *ChatController) AddActions(actionsHandler *controllers.ActionsHandler) {
	actionsHandler.AddAction("Join", controller.join)
	actionsHandler.AddAction("Message", controller.message)
}

// join used for player to join a room
func (controller *ChatController) join(connection *controllers.Connection) (interface{}, error) {
	var err error
	roomID := connection.ParamsStr("roomID")

	// leave all previously joined rooms, and join new room
	connection.SingleJoin(roomID)

	return nil, err
}

// message used for player to join a room
func (controller *ChatController) message(connection *controllers.Connection) (interface{}, error) {
	var err error
	roomID := connection.ParamsStr("roomID")
	message := connection.ParamsStr("message")

	// broadcast message to room
	fab.BroadcastEvent("chat", roomID, "Message", nil, fab.H{
		"message": message,
	})

	return nil, err
}
