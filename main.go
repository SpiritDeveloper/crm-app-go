package main

import (
	"encoding/json"
	config "crm-app-go/config"
	"crm-app-go/controller"
	"crm-app-go/db"
	_ "crm-app-go/docs" // This line is necessary for go-swagger to find docs!
	router "crm-app-go/http"
	"crm-app-go/repository"
	"crm-app-go/service"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	c 		   config.Config
	httpRouter router.IRouter
	gormDb     db.IDatabaseEngine
	gDb        *gorm.DB
)


// User
var (
	userRepository 	repository.IUserRepository
	userService		service.IUserService
	userController	controller.IUserController
)

func main() {
	initConfig()
	httpRouter = router.NewMuxRouter()
	httpRouter.ADDVERSION("/api/v1")
	gormDb = db.NewGormDatabase()
	gDb = gormDb.GetDatabase(c.Database)
	gormDb.RunMigration()
	initUserServiceContainer()
	httpRouter.SERVE(c.App.Port)
}

func initConfig() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Printf("No ./config.json file found!! Terminating the server, error: %s\n", err.Error())
		panic("No config file found! Error : " + err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Printf("Error occurred while decoding json to config model, error: %s\n", err.Error())
		panic(err.Error())
	}
}


func initUserServiceContainer() {
	userRepository = repository.NewUserRepository(gDb)
	userService = service.NewUserService(userRepository)
	userController = controller.NewUserController(userService)

	httpRouter.GET("/user/{id}", userController.GetUser)
	httpRouter.GET("/user", userController.GetUsers)
	httpRouter.POST("/user", userController.PostUser)
	httpRouter.PUT("/user/{id}", userController.PutUser)
	httpRouter.DELETE("/user/{id}", userController.DeleteUser)
}

