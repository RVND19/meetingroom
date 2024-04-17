package server

import (
	"fmt"

	"github.com/RVND19/meetingroom/controllers"
	"github.com/RVND19/meetingroom/db"
	"github.com/RVND19/meetingroom/repostitories"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
func Bootstrap(app *fiber.App) {
	
	//load .env
	err:= godotenv.Load("dev.env")
	if err != nil {
		panic(err)
	}

	db := db.GetDB()
	fmt.Println("DB connected" + db.Dialector.Name())
	

	//setup  repository
	userRepo:=repostitories.NewUserRepository(db)
	roomRepo:=repostitories.NewRoomRepository(db)
	reservationRepo:=repostitories.NewReservationRepository(db)

	

	// //setup controller
	userController := controllers.NewUserController(userRepo)
	roomController := controllers.NewRoomController(roomRepo)
	reservationController:=controllers.NewReservationController(reservationRepo, userRepo, roomRepo)


	//route
	userController.Route(app)
	roomController.Route(app)
	reservationController.Route(app)
}