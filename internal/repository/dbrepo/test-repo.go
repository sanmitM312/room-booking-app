package dbrepo

import (
	"time"
	"errors"
	"github.com/sanmitM312/room-booking-app/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database and returns the id entered
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// if the room id is 2, then fail otherwise pass
	// 2 for triggering the error
	if res.RoomID == 2{
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID  and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms for given start and end date if any
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets the room's details because it is needed to render
// room name in the make reservation after choosing the room
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil

}

func (m *testDBRepo)GetUserByID(id int)(models.User,error){
	var u models.User

	return u,nil
}

func (m *testDBRepo)UpdateUser(u models.User) error{
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string)(int,string,error){
	return 1,"",nil
}
// AllReservations returns a slice of reservations
func (m *testDBRepo) AllReservations()([]models.Reservation, error){
	var reservations []models.Reservation
	return reservations,nil
}

// AllNewReservations returns a slice of reservations where processed = 0/ new reservations
func (m *testDBRepo) AllNewReservations()([]models.Reservation, error){
	
	var reservations []models.Reservation

	return reservations,nil 
}

// GetReservationByID returns one reservation by iD
func (m *testDBRepo) GetReservationByID(id int)(models.Reservation, error){
	var res models.Reservation

	return res, nil 
}

// UpdateReservation updates a user in the database
func (m *testDBRepo)UpdateReservation(u models.Reservation) error{

	return nil 
}

// DeleteReservation deletes one reservation by id
func (m *testDBRepo)DeleteReservation(id int) error{
	
	return nil 
}

//UpdateProcessedForReservation updates processed for a reservation by id
func (m *testDBRepo)UpdateProcessedForReservation(id,processed int) error {
	return nil
}