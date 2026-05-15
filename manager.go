package main

import (
	"errors"
	"sort"
	"strings"
)

var (
	ErrHouseNotFound    = errors.New("house not found")
	ErrInvalidTitle     = errors.New("title is required")
	ErrInvalidLocation  = errors.New("location is required")
	ErrInvalidPrice     = errors.New("price per night must be greater than 0")
	ErrInvalidMaxGuests = errors.New("max guests must be greater than 0")
	ErrInvalidGuest     = errors.New("guest name is required")
	ErrInvalidNights    = errors.New("nights must be greater than 0")
	ErrHouseUnavailable = errors.New("house is currently unavailable")
	ErrHouseHasBookings = errors.New("cannot delete house with existing bookings")
)

type House struct {
	ID            int
	Title         string
	Location      string
	PricePerNight float64
	MaxGuests     int
	Available     bool
}

type HouseInput struct {
	Title         string
	Location      string
	PricePerNight float64
	MaxGuests     int
	Available     bool
}

type Booking struct {
	ID         int
	HouseID    int
	GuestName  string
	Nights     int
	TotalPrice float64
}

type HouseManager struct {
	houses        map[int]House
	bookings      map[int]Booking
	nextHouseID   int
	nextBookingID int
}

func NewHouseManager() *HouseManager {
	return &HouseManager{
		houses:        make(map[int]House),
		bookings:      make(map[int]Booking),
		nextHouseID:   1,
		nextBookingID: 1,
	}
}

func (m *HouseManager) AddHouse(input HouseInput) (House, error) {
	if err := validateHouseInput(input); err != nil {
		return House{}, err
	}

	house := House{
		ID:            m.nextHouseID,
		Title:         strings.TrimSpace(input.Title),
		Location:      strings.TrimSpace(input.Location),
		PricePerNight: input.PricePerNight,
		MaxGuests:     input.MaxGuests,
		Available:     input.Available,
	}
	m.houses[house.ID] = house
	m.nextHouseID++
	return house, nil
}

func (m *HouseManager) ListHouses() []House {
	houses := make([]House, 0, len(m.houses))
	for _, house := range m.houses {
		houses = append(houses, house)
	}
	sort.Slice(houses, func(i, j int) bool {
		return houses[i].ID < houses[j].ID
	})
	return houses
}

func (m *HouseManager) UpdateHouse(id int, input HouseInput) (House, error) {
	if err := validateHouseInput(input); err != nil {
		return House{}, err
	}

	house, exists := m.houses[id]
	if !exists {
		return House{}, ErrHouseNotFound
	}

	house.Title = strings.TrimSpace(input.Title)
	house.Location = strings.TrimSpace(input.Location)
	house.PricePerNight = input.PricePerNight
	house.MaxGuests = input.MaxGuests
	house.Available = input.Available
	m.houses[id] = house
	return house, nil
}

func (m *HouseManager) DeleteHouse(id int) error {
	if _, exists := m.houses[id]; !exists {
		return ErrHouseNotFound
	}

	for _, booking := range m.bookings {
		if booking.HouseID == id {
			return ErrHouseHasBookings
		}
	}

	delete(m.houses, id)
	return nil
}

func (m *HouseManager) BookHouse(houseID int, guestName string, nights int) (Booking, error) {
	house, exists := m.houses[houseID]
	if !exists {
		return Booking{}, ErrHouseNotFound
	}
	if !house.Available {
		return Booking{}, ErrHouseUnavailable
	}
	if strings.TrimSpace(guestName) == "" {
		return Booking{}, ErrInvalidGuest
	}
	if nights <= 0 {
		return Booking{}, ErrInvalidNights
	}

	booking := Booking{
		ID:         m.nextBookingID,
		HouseID:    houseID,
		GuestName:  strings.TrimSpace(guestName),
		Nights:     nights,
		TotalPrice: house.PricePerNight * float64(nights),
	}
	m.bookings[booking.ID] = booking
	m.nextBookingID++
	return booking, nil
}

func (m *HouseManager) ListBookings() []Booking {
	bookings := make([]Booking, 0, len(m.bookings))
	for _, booking := range m.bookings {
		bookings = append(bookings, booking)
	}
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].ID < bookings[j].ID
	})
	return bookings
}

func validateHouseInput(input HouseInput) error {
	if strings.TrimSpace(input.Title) == "" {
		return ErrInvalidTitle
	}
	if strings.TrimSpace(input.Location) == "" {
		return ErrInvalidLocation
	}
	if input.PricePerNight <= 0 {
		return ErrInvalidPrice
	}
	if input.MaxGuests <= 0 {
		return ErrInvalidMaxGuests
	}
	return nil
}
