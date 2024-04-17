package main

import (
	"fmt"

	"github.com/RVND19/meetingroom/server"
	"github.com/gofiber/fiber/v2"
)
func main() {
	
	//fiber route
	app := fiber.New()
	server.Bootstrap(app)

	fmt.Println("server running on port 8086")
	app.Listen(":8086")
}