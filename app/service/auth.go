package service

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"strings"
	"time"
	"v2/app/model"
)

var invalidToken = errors.New("invalid token")
var unauthorized = errors.New("unauthorized")

type Auth struct {
	*gorm.DB
	Secret  []byte
	Expired time.Duration
}

func (a *Auth) create() {
	a.DB.Table("users").Omit("Roles")
}

func NewAuthService(DB *gorm.DB, secret []byte, expired time.Duration) *Auth {
	return &Auth{DB: DB, Secret: secret, Expired: expired}
}

func (a *Auth) Decode(token string) (*model.JwtUser, error) {

	var jwtToken, err = jwt.ParseWithClaims(token, &model.JwtUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, invalidToken
		}
		return a.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, unauthorized
	}

	hash := strings.Split(jwtToken.Raw, ".")

	var user *model.JwtUser

	hashByte, err := base64.RawStdEncoding.DecodeString(hash[1])

	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(hashByte, &user)

	if err != nil {
		return nil, unauthorized
	}

	return user, nil
}

func (a *Auth) Encode(user *model.User) (string, error) {

	claims := &model.JwtUser{
		Id: user.Id, Username: user.Username, CompanyId: user.CompanyId, BranchId: user.BranchId, Roles: user.Roles,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * a.Expired).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tok, err := token.SignedString(a.Secret)

	if err != nil {
		return "", err
	}

	return tok, nil
}

func (a *Auth) SignIn(in *model.SigInUser) (map[string]interface{}, error) {

	var (
		err   error
		user  *model.User
		token string
	)

	if err = a.Table("users").Preload("Roles").
		Where("Active=?", true).
		Where("username=?", in.Username).
		Find(&user).Error; err != nil {
		return nil, err
	}

	if !user.ValidatePassword(in.Password) {
		return nil, unauthorized
	}

	if token, err = a.Encode(user); err != nil {
		return nil, unauthorized
	}

	user.Sanitize()

	var ret = map[string]interface{}{
		"user":  user,
		"token": token,
	}

	return ret, nil
}
