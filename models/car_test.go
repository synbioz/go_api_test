package models

import (
	"github.com/synbioz/go_api/config"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewCar(t *testing.T) {

	cars := []Car{
		{},
		{Manufacturer: "Volvo", Design: "v40", Style: "urban"},
	}
	for _, c := range cars {
		timeNow := time.Now().Format(time.UnixDate)
		NewCar(&c)

		createdAt := c.CreatedAt.Format(time.UnixDate)
		updatedAt := c.UpdatedAt.Format(time.UnixDate)

		if timeNow != createdAt {
			t.Errorf("Car created_at not have correct DateTime %q != %q", createdAt, timeNow)
		}

		if timeNow != updatedAt {
			t.Errorf("Car updated_at not have correct DateTime %q != %q", createdAt, timeNow)
		}
	}

	rows, _ := config.Db().Query("SELECT * FROM cars")

	var rowsCount int

	for rows.Next() {
		rowsCount += 1
	}

	if rowsCount != 2 {
		t.Errorf("Database does not have 2 records, it has %v records", rowsCount)
	}
}

func TestFindCarById(t *testing.T) {
	car := &Car{Manufacturer: "Volvo", Design: "v40", Style: "urban"}

	NewCar(car)
	carFound := FindCarById(car.Id)

	if car.Id != carFound.Id {
		t.Error("Couldn't find car by id")
	}
}

func TestAllCars(t *testing.T) {
	ResetTableCars()

	var cars Cars

	car := Car{Manufacturer: "Volvo", Design: "v40", Style: "urban"}

	NewCar(&car)

	cars = append(cars, *FindCarById(car.Id))

	if !reflect.DeepEqual(&cars, AllCars()) {
		t.Error("Couldn't find correct car from AllCars")
	}
}

func TestUpdateCar(t *testing.T) {
	car := &Car{Manufacturer: "Volvo", Design: "v40", Style: "urban"}

	NewCar(car)

	car.Manufacturer = "Peugeot"

	UpdateCar(car)

	carFromDb := FindCarById(car.Id)

	if !reflect.DeepEqual(car.Manufacturer, carFromDb.Manufacturer) {
		t.Error("Not updated item")
	}
}

func TestDeleteCarById(t *testing.T) {
	car := &Car{Manufacturer: "Volvo", Design: "v40", Style: "urban"}

	NewCar(car)

	DeleteCarById(car.Id)

	for _, val := range *AllCars() {
		if reflect.DeepEqual(car.Id, val.Id) {
			t.Errorf("The car should be destroyed %v", val)
		}
	}
}

func ResetTableCars() {
	config.Db().Exec("DROP TABLE cars; CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
}

func TestMain(m *testing.M) {
	config.TestDatabaseInit()

	ret := m.Run()

	config.TestDatabaseDestroy()
	os.Exit(ret)
}
