package soket_controller

import (
	"github.com/fobus1289/marshrudka/socket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"v2/app/model"
	"v2/app/service"
)

type Main struct {
	*gorm.DB
	*service.Auth
	*service.Logger
	*service.User
}

type Vector struct {
	X float32
	Y float32
	Z float32
}

type Hero struct {
	Id       int64   `json:"id"`
	Position *Vector `json:"position"`
}

//Default if event not found run this
func (m *Main) Default(client *socket.Client, data interface{}) {
	client.BroadcastClients("pos", data)
}

func (m *Main) Connection(currentClient *socket.Client, r *http.Request) {

	user, err := m.Auth.Decode(r.URL.Query().Get("token"))

	if err != nil {
		currentClient.Delete()
		_ = currentClient.Conn.Close()
		m.Logger.Println("error auth socket invalid token reason: ", err)
		return
	}

	log.Println("client connected")

	currentClient.SetId(user.Id)
	currentClient.SetOwner(user)
	JoinChannel(currentClient, user)

}

func JoinChannel(currentClient *socket.Client, jwtUser *model.JwtUser) {
	for _, role := range jwtUser.Roles {
		if role == nil {
			continue
		}
		switch role.Name {
		case "chef":
			currentClient.JoinChannel("chef")
		case "waiter":
			currentClient.JoinChannel("waiter")
		}
	}
}

func (m *Main) Disconnection(client *socket.Client) {
	//log.Println(client.GetId(), "disconnected")
	client.BroadcastClients("live", client.GetId())
}
