package soket_controller

import (
	"github.com/fobus1289/marshrudka/socket"
	"v2/app/model"
	"v2/app/service"
)

type Order struct {
}

func (o *Order) Create(currentClient *socket.Client, orderService *service.Order, descriptionOrder *model.DescriptionOrder) {
	jwtUser := currentClient.GetOwner().(*model.JwtUser)

	order, err := orderService.Create(descriptionOrder, jwtUser)

	if err != nil {
		return
	}

	currentClient.Emit("my-order", order)

	currentClient.BroadcastChannel("order", "chef", order)
	currentClient.BroadcastMeToo("queue-order", order)
}

func (o *Order) ForClients(currentClient *socket.Client, orderService *service.Order) {
	orders, err := orderService.ForClients()

	if err != nil {
		return
	}

	currentClient.Emit("queue-orders", orders)
}

func (o *Order) All(currentClient *socket.Client, orderService *service.Order) {
	orders, err := orderService.All()

	if err != nil {
		return
	}

	currentClient.Emit("my-orders", orders)
}

func (o *Order) Status(currentClient *socket.Client, order *model.Order, orderService *service.Order) {
	order, err := orderService.StatusUpdate(order)

	if err != nil {
		return
	}

	currentClient.Emit("queue-order", order)
}
