package main

import (
	"fmt"
	"http-server/internal/core/service/usersrv"
	"http-server/internal/handler/userhdl"
	"http-server/internal/repositories/userrepo"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("start http-config-server!!!")

	userRepository := userrepo.NewMemKVS()
	userService := usersrv.New(userRepository)
	userHanlder := userhdl.NewHTTPHandler(userService)

	app := fiber.New()

	app.Post("/api/v1/config/admin/login", userHanlder.Login)
	app.Post("/api/v1/config/admins", userHanlder.Create)
	app.Get("/api/v1/config/admins", userHanlder.List)
	app.Get("/api/v1/config/admins:id", userHanlder.Get)
	app.Post("/api/v1/config/admins:id", userHanlder.Update)
	app.Delete("/api/v1/config/admins:id", userHanlder.Delete)

	log.Fatal(app.Listen(":6020"))
}
