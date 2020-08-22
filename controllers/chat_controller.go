package controllers

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/controllers"
	"github.com/kooinam/fabio/helpers"
)

// ChatController used for chat purposes
type ChatController struct {
}

// RegisterHooksAndActions used to register hooks and actions
func (controller *ChatController) RegisterHooksAndActions(hooksHandler *controllers.HooksHandler, actionsHandler *controllers.ActionsHandler) {
	actionsHandler.RegisterAction("Join", controller.join)
	actionsHandler.RegisterAction("Message", controller.message)
}

// join used for player to join a room
func (controller *ChatController) join(context *controllers.Context) (interface{}, error) {
	var err error
	roomID := context.ParamsStr("roomID")

	// leave all previously joined rooms, and join new room
	context.SingleJoin(roomID)

	return nil, err
}

// message used for player to send message message to room
func (controller *ChatController) message(context *controllers.Context) (interface{}, error) {
	var err error
	roomID := context.ParamsStr("roomID")
	message := context.ParamsStr("message")

	// broadcast message to room
	fab.ControllerManager().BroadcastEvent("chat", roomID, "Message", nil, helpers.H{
		"message": message,
	})

	return nil, err
}
