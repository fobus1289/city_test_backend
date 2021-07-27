package server

import (
	"fmt"
	"github.com/fobus1289/marshrudka/router"
	"github.com/fobus1289/marshrudka/socket"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
	"v2/app/soket_controller"
	_ "v2/docs"
	"v2/repository"
	"v2/route"
	web_socket "v2/route/web-socket"
)

func NewServer() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db := repository.NewMariaDBGorm(&repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	})

	drive := router.NewRouter()

	drive.Use(middleware.Cross)

	socketServer := sockets()

	main := &soket_controller.Main{}

	service.Init(drive, socketServer, db)

	route.Init(drive)

	socketServer.Dep(main)
	socketServer.Default(main.Default)
	socketServer.Connection(main.Connection)
	socketServer.Disconnection(main.Disconnection)

	drive.GET("ws", func(w http.ResponseWriter, r *http.Request) {
		_, err := socketServer.NewClient(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
	})

	web_socket.Init(socketServer)
	testes(drive)
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT"))
	go docs()
	drive.Run(url)
}

func saveFile(db *gorm.DB) {
	//bytes, err := ioutil.ReadFile("migration/Паспортpdf.pdf")
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}

	type Test struct {
		Id   int
		Data []byte
	}

	var test Test

	db.Table("test").Find(&test)
	log.Println(test)
	ioutil.WriteFile("migration/test.pdf", test.Data, 0644)
	//if err := db.Exec("insert into file (owner,data) values ('company',?)", bytes).Error; err != nil {
	//	log.Fatalln(err)
	//}
	//
}

func testes(drive *router.Drive) {

	testGroup := drive.Group("test")
	{
		testGroup.GET("category", func(category *service.Category) interface{} {
			categories, err := category.All()

			if err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}

			return categories
		})

		testGroup.POST("category", func(category *model.Category, categoryService *service.Category) interface{} {
			if err := categoryService.Create(category); err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}
			return category
		})

		testGroup.GET("category-component", func(categoryComponentService *service.CategoryComponent) interface{} {

			categoryComponents, err := categoryComponentService.All()

			if err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}

			return categoryComponents
		})

		testGroup.POST("category-component", func(categoryComponent *model.CategoryComponent, categoryComponentService *service.CategoryComponent) interface{} {
			if err := categoryComponentService.Create(categoryComponent); err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}
			return categoryComponent
		})

		testGroup.GET("component", func(componentService *service.Component) interface{} {

			components, err := componentService.All()

			if err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}

			return components
		})

		testGroup.POST("component", func(component *model.Component, componentService *service.Component) interface{} {
			if err := componentService.Create(component); err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}
			return component
		})

		testGroup.GET("product", func(productService *service.Product) interface{} {

			products, err := productService.All()

			if err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}

			return products
		})

		testGroup.POST("product", func(product *model.Product, productService *service.Product) interface{} {
			if err := productService.Create(product); err != nil {
				return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
					"message": err.Error(),
				})
			}
			return product
		})

	}

}

func t(db *gorm.DB) {
	type Location struct {
		X float64
		Y float64
	}

	type T struct {
		Id       int
		Location Location
	}

	var tt T

	db.Table("t").Find(&tt)

	log.Println(tt)
}

func te(db *gorm.DB) {

	type AbortClient struct {
		Name  string `json:"name" gorm:"column:VARIABLE_NAME"`
		Value string `json:"value" gorm:"column:VARIABLE_VALUE"`
	}

	var quas []AbortClient

	find := db.Table("information_schema.GLOBAL_STATUS").Where("VARIABLE_NAME like 'ABORTED%'").Find

	for {
		if err := find(&quas).Error; err != nil {
			log.Println(err)
		} else {
			//log.Println(quas)
		}
		time.Sleep(2 * time.Second)
	}

}

func sockets() *socket.WebSocket {

	webSocket := socket.NewWebSocket(&socket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			log.Println(status)
			log.Println(reason)
		},
	})

	//go func() {
	//	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDR"), os.Getenv("SOCKET_PORT")),
	//		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			_, err := webSocket.NewClient(w, r, nil)
	//			if err != nil {
	//				log.Println(err)
	//				return
	//			}
	//		}),
	//	))
	//}()

	return webSocket
}

func docs() {
	if ok, _ := strconv.ParseBool(os.Getenv("DOCS")); ok {
		log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDR"), os.Getenv("DOC_PORT")), httpSwagger.WrapHandler))
	}
}
