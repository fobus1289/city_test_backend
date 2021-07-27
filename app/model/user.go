package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	CompanyId *int64    `json:"companyId"`
	BranchId  *int64    `json:"branchId"`
	FIO       string    `json:"fio"`
	PinCode   int       `json:"pinCode"`
	Rate      float32   `json:"rate"`
	Phone     string    `json:"phone"`
	Companies Companies `json:"companies" gorm:"foreignKey:user_id"`
	Branches  Branches  `json:"branches" gorm:"foreignKey:user_id"`
	Roles     Roles     `json:"roles" gorm:"many2many:user_roles;"`
	Active    bool      `json:"active"`
	CreatedAt *string   `json:"createdAt"`
	UpdatedAt *string   `json:"updatedAt"`
	DeletedAt *string   `json:"deletedAt"`
}

type Users []User

type SigInUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtUser struct {
	Id                  int64  `json:"id"`
	Username            string `json:"username"`
	CompanyId           *int64 `json:"companyId"`
	BranchId            *int64 `json:"branchId"`
	Roles               Roles  `json:"roles"`
	*jwt.StandardClaims `json:"claims,omitempty"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ValidatePassword(inPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inPassword)) == nil
}

func (u *User) EncryptPassword() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
