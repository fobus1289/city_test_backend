package soket_controller

import (
	"github.com/fobus1289/marshrudka/socket"
	"gorm.io/gorm"
	"log"
	"v2/app/service"
)

type Company struct {
}

func (c *Company) GetAllUsers(user *service.User, currentClient *socket.Client) {

	users, err := user.Users(0, 0, func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Password")
	})

	if err != nil {
		return
	}

	_ = users

	//return filterClients
}

func (c *Company) TestAll(client *socket.Client, data interface{}) {
	log.Println(data)
}
