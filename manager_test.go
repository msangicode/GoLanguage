package main

import (
	"errors"
	"testing"
)

func TestHouseManagerAddAndListHouses(t *testing.T) {
	m := NewHouseManager()

	_, err := m.AddHouse(HouseInput{
		Title:         "Cozy Apartment",
		Location:      "New York",
		PricePerNight: 120,
		MaxGuests:     2,
		Available:     true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	houses := m.ListHouses()
	if len(houses) != 1 {
		t.Fatalf("expected 1 house, got %d", len(houses))
	}

	if houses[0].Title != "Cozy Apartment" {
		t.Fatalf("unexpected house title: %s", houses[0].Title)
	}
}

func TestHouseManagerValidation(t *testing.T) {
	m := NewHouseManager()

	_, err := m.AddHouse(HouseInput{
		Title:         "",
		Location:      "NY",
		PricePerNight: 100,
		MaxGuests:     2,
		Available:     true,
	})
	if !errors.Is(err, ErrInvalidTitle) {
		t.Fatalf("expected ErrInvalidTitle, got %v", err)
	}
}

func TestHouseManagerUpdateAndDelete(t *testing.T) {
	m := NewHouseManager()

	house, err := m.AddHouse(HouseInput{
		Title:         "Old title",
		Location:      "Boston",
		PricePerNight: 95,
		MaxGuests:     3,
		Available:     true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updated, err := m.UpdateHouse(house.ID, HouseInput{
		Title:         "New title",
		Location:      "Boston",
		PricePerNight: 100,
		MaxGuests:     4,
		Available:     false,
	})
	if err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}
	if updated.Title != "New title" || updated.Available {
		t.Fatalf("update did not apply correctly: %+v", updated)
	}

	if err := m.DeleteHouse(house.ID); err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}
}

func TestBookHouse(t *testing.T) {
	m := NewHouseManager()

	house, err := m.AddHouse(HouseInput{
		Title:         "Lake Cabin",
		Location:      "Zurich",
		PricePerNight: 200,
		MaxGuests:     5,
		Available:     true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	booking, err := m.BookHouse(house.ID, "Alice", 3)
	if err != nil {
		t.Fatalf("unexpected booking error: %v", err)
	}

	if booking.TotalPrice != 600 {
		t.Fatalf("expected total price 600, got %v", booking.TotalPrice)
	}
}

func TestDeleteHouseWithBookingFails(t *testing.T) {
	m := NewHouseManager()

	house, err := m.AddHouse(HouseInput{
		Title:         "Beach House",
		Location:      "Miami",
		PricePerNight: 300,
		MaxGuests:     6,
		Available:     true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = m.BookHouse(house.ID, "Bob", 2)
	if err != nil {
		t.Fatalf("unexpected booking error: %v", err)
	}

	err = m.DeleteHouse(house.ID)
	if !errors.Is(err, ErrHouseHasBookings) {
		t.Fatalf("expected ErrHouseHasBookings, got %v", err)
	}
}
