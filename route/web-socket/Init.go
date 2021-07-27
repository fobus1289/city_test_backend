package web_socket

import (
	"github.com/fobus1289/marshrudka/socket"
	"v2/app/soket_controller"
)

func Init(webSocket *socket.WebSocket) {

	company := &soket_controller.Company{}
	order := &soket_controller.Order{}

	webSocket.Event("orders", order.All)
	webSocket.Event("order", order.Create)
	webSocket.Event("queue-orders", order.ForClients)
	webSocket.Event("order-status", order.Status)
	webSocket.Event("b", company.GetAllUsers)
	webSocket.Event("q", company.TestAll)
}
