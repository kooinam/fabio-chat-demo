package controllers

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/controllers"
)

// ChatController used for chat purposes
type ChatController struct {
}

// RegisterBeforeHooks used to register before action hooks
func (controller *ChatController) RegisterBeforeHooks(hooksHandler *controllers.HooksHandler) {
}

// RegisterActions used to add actions
func (controller *ChatController) RegisterActions(actionsHandler *controllers.ActionsHandler) {
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
	fab.ControllerManager().BroadcastEvent("chat", roomID, "Message", nil, fab.H{
		"message": message,
	})

	return nil, err
}
