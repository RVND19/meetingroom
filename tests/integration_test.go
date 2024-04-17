package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/RVND19/meetingroom/db"
	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"github.com/RVND19/meetingroom/repostitories"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var userRepo interfaces.IUserRepo
var roomRepo interfaces.IRoomRepo
var reservationRepo interfaces.IReservationRepo
func setup() {
   //load env file
	err := godotenv.Load("test.env")
	if err != nil {
		panic(err)
	}
	db := db.GetDB()
	userRepo=repostitories.NewUserRepository(db)
	roomRepo=repostitories.NewRoomRepository(db)
	reservationRepo=repostitories.NewReservationRepository(db)
}
func createDummyUser() *entities.User {
	dummy:=faker.New()
	return &entities.User{
		Name: dummy.Person().FirstName(),
		Email: dummy.Person().Contact().Email,
		Password: dummy.Internet().Password(),
		Salt: dummy.Internet().Password(),
		IsVerified: true,
		IsAdmin: true,
	}
}

func createDummyReservation(idRoom int, idUser int) *entities.Reservation {
	return &entities.Reservation{
		RoomId: idRoom,
		UserId: idUser,
		StartDate: time.Now(),
		EndDate: time.Now().Add(time.Hour * 1),
	}
}

func createReservation(t *testing.T) *entities.Reservation {
	

	user:=createDummyUser()
	err:=userRepo.CreateUser(user)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Printf("%+v\n",*user)
	room:=&entities.Room{
		Name: faker.New().Address().City(),
	}
	err=roomRepo.CreateRoom(room)
	if err != nil {
		t.Error(err)
		t.Fail()

	}

	rs:=createDummyReservation(room.Id, user.Id)
	err=reservationRepo.CreateReservation(rs)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Printf("%+v\n",*rs)
	return rs
}

func TestReservationRepository_CreateReservation(t *testing.T) {
	setup()
	createReservation(t)

}

func TestReservationRepository_GetById(t *testing.T) {
	setup()
	rsrv:=createReservation(t)
	if rsrv == nil {
		t.Fail()
	}

	dataCheck,err:=reservationRepo.GetById(rsrv.Id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, rsrv.Id, dataCheck.Id,"Id not match")

}

func TestReservationRepository_DeleteReservation(t *testing.T) {
	setup()
	rsrv:=createReservation(t)
	if rsrv == nil {
		t.Fail()
	}

	err:=reservationRepo.DeleteReservation(rsrv.Id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}

func TestReservationRepository_GetById_notFound(t *testing.T) {
	setup()
	rsrv:=createReservation(t)
	if rsrv == nil {
		t.Fail()
	}

	err:=reservationRepo.DeleteReservation(rsrv.Id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_,err=reservationRepo.GetById(rsrv.Id)
	assert.ErrorAs(t, gorm.ErrRecordNotFound, &err)

}

func TestReservationRepository_IsRoomAvailable(t *testing.T) {
	setup()
	rsrv:=createReservation(t)
	if rsrv == nil {
		t.Fail()
	}

	check,err:=reservationRepo.IsRoomAvailable(rsrv.RoomId, rsrv.StartDate, rsrv.EndDate)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, false, check)

	check,err=reservationRepo.IsRoomAvailable(rsrv.RoomId, rsrv.StartDate.Add(time.Minute), rsrv.EndDate.Add(time.Hour *2))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, false, check)

	check,err=reservationRepo.IsRoomAvailable(rsrv.RoomId, rsrv.StartDate.Add(-time.Hour), rsrv.StartDate.Add(time.Minute *2))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, false, check)

	check,err=reservationRepo.IsRoomAvailable(rsrv.RoomId, rsrv.EndDate.Add(time.Second), rsrv.EndDate.Add(time.Hour))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, true, check)

	check,err=reservationRepo.IsRoomAvailable(rsrv.RoomId, rsrv.StartDate.Add(-time.Hour), rsrv.StartDate.Add(-time.Second))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, true,check)

}