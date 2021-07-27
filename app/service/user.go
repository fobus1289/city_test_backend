package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"v2/app/model"
)

type User struct {
	*gorm.DB
}

func NewUserService(DB *gorm.DB) *User {
	return &User{DB: DB}
}

func (u *User) Users(page, limit int, scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Response, error) {

	response := model.NewResponse("users", page, limit, &model.Users{})

	if err := Paginate(u.DB, response, scopes...); err != nil {
		return nil, err
	}

	return response, nil
}

func (u *User) FindById(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.User, error) {
	var user *model.User

	if err := u.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").
		Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) FindByUsername(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.User, error) {
	var user *model.User

	if err := u.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").
		Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) FindByFIO(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Users, error) {
	var users model.Users

	if err := u.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) Create(user *model.User) error {

	var (
		tx  = u.Begin()
		err error
	)

	defer commit(tx, func() error {
		return err
	})

	if err = user.EncryptPassword(); err != nil {
		return err
	}

	if err = tx.Table("users").
		Omit("Id", "Roles", "CreatedAt", "UpdatedAt", "DeletedAt").
		Create(&user).Error; err != nil {
		return err
	}

	userRoles := user.Roles.UserRoles(user.Id)

	if len(userRoles) > 0 {
		err = tx.Table("user_roles").
			Clauses(clause.Insert{Modifier: "IGNORE"}).
			Create(userRoles).Error
		if err != nil {
			return err
		}
	}

	if err = tx.Table("users").Omit("Password").Preload("Roles").Find(&user).Error; err != nil {
		return err
	}

	user.Sanitize()

	return nil
}

func (u *User) Update(user *model.User, scopes ...func(tx *gorm.DB) *gorm.DB) error {

	var (
		tx  = u.Begin()
		err error
	)

	defer commit(tx, func() error {
		return err
	})

	if err = user.EncryptPassword(); err != nil {
		return err
	}

	if err = tx.Table("users").
		Scopes(scopes...).
		Omit("Id", "Roles", "CreatedAt", "UpdatedAt", "DeletedAt").
		Updates(&user).Error; err != nil {
		return err
	}

	userRoles := user.Roles.UserRoles(user.Id)

	if len(userRoles) > 0 {
		err = tx.Table("user_roles").
			Clauses(clause.Insert{Modifier: "IGNORE"}).
			Create(userRoles).Error
		if err != nil {
			return err
		}
	}

	if err = tx.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").Find(&user).Error; err != nil {
		return err
	}

	user.Sanitize()

	return nil
}

func (u *User) Delete(scopes ...func(tx *gorm.DB) *gorm.DB) error {
	return u.Table("users").Scopes(scopes...).Update("deleted_at", time.Now()).Error
}

func (u *User) FindByCompanyId(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Users, error) {
	var users model.Users

	if err := u.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) FindByBranchId(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Users, error) {
	var users model.Users

	if err := u.Table("users").
		Scopes(scopes...).
		Omit("Password").
		Preload("Roles").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) DeleteRole(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.User, error) {

	var (
		user model.User
		tx   = u.Begin()
		err  error
	)

	defer commit(tx, func() error {
		return err
	})

	var userRole *model.UserRole

	if err = tx.Table("user_roles").Scopes(scopes...).Find(&userRole).Error; err != nil {
		return nil, err
	}

	if err = tx.Table("user_roles").
		Scopes(scopes...).
		Delete(&userRole).Error; err != nil {
		return nil, err
	}

	if err = tx.Where("id=?", userRole.UserId).Preload("Roles").
		Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Exists(scopes ...func(tx *gorm.DB) *gorm.DB) (bool, error) {
	var id int64

	if err := u.Table("users").Scopes(scopes...).Select("id").Scan(&id).Error; err != nil {
		return false, err
	}

	return id != 0, nil
}
