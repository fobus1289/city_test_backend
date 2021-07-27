package service

import (
	"github.com/fobus1289/marshrudka/router"
	"github.com/fobus1289/marshrudka/socket"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
	"v2/app/model"
)

func Init(drive *router.Drive, socket *socket.WebSocket, db *gorm.DB) {

	logService := NewLoggerService(os.Stdout)
	userService := NewUserService(db)
	companyService := NewCompanyService(db)
	branchService := NewBranchService(db)
	categoryService := NewCategoryService(db)
	categoryComponentService := NewCategoryComponentService(db)
	componentService := NewComponentService(db)
	productService := NewProductService(db)
	orderService := NewOrderService(db)
	realizationService := NewRealizationService(db)

	expired, err := strconv.ParseInt(os.Getenv("EXPIRED"), 10, 64)

	if err != nil {
		expired = 15
	}

	authService := NewAuthService(db, []byte(os.Getenv("SECRET")), time.Duration(expired))

	drive.Register(db)
	drive.Register(logService)
	drive.Register(authService)
	drive.Register(userService)
	drive.Register(companyService)
	drive.Register(branchService)
	drive.Register(categoryService)
	drive.Register(categoryComponentService)
	drive.Register(componentService)
	drive.Register(productService)
	drive.Register(orderService)
	drive.Register(realizationService)

	socket.Register(db)
	socket.Register(logService)
	socket.Register(authService)
	socket.Register(userService)
	socket.Register(companyService)
	socket.Register(branchService)
	socket.Register(categoryService)
	socket.Register(categoryComponentService)
	socket.Register(componentService)
	socket.Register(productService)
	socket.Register(orderService)
	socket.Register(realizationService)

	drive.Dep(logService)
	drive.Dep(authService)
	drive.Dep(userService)
	drive.Dep(companyService)
	drive.Dep(branchService)
	drive.Dep(categoryService)
	drive.Dep(categoryComponentService)
	drive.Dep(componentService)
	drive.Dep(productService)
	drive.Dep(orderService)
	drive.Dep(realizationService)

	if err = userService.Create(&model.User{
		Username: os.Getenv("ADMIN_LOGIN"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Roles: model.Roles{
			{
				Id: 1,
			},
		},
		Active: true,
	}); err != nil {
		logService.Println(err)
	}
}

func commit(tx *gorm.DB, fn func() error) {
	if fn() != nil {
		tx.Rollback()
		return
	}
	if err := &tx.Commit().Error; err != nil {
		tx.Rollback()
	}
}

func Paginate(db *gorm.DB, response *model.Response, scopes ...func(db *gorm.DB) *gorm.DB) error {

	var (
		err error
	)

	tx := db.Begin().Scopes(scopes...).Table(response.Name)

	defer commit(tx, func() error {
		return err
	})

	if response.Page < 1 {
		response.Page = 1
	}

	switch {
	case response.Limit > 100:
		response.Limit = 100
	case response.Limit < 1:
		response.Limit = 10
	}

	if err = tx.Count(&response.Count).Error; err != nil {
		return err
	}

	offset := (response.Page - 1) * response.Limit

	if err = tx.Table(response.Name).
		Limit(response.Limit).
		Offset(offset).
		Find(response.Data).Error; err != nil {
		return err
	}

	return nil
}
