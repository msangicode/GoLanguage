package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	manager := NewHouseManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Print("Choose an option: ")
		if !scanner.Scan() {
			fmt.Println("Exiting...")
			return
		}

		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			handleAddHouse(scanner, manager)
		case "2":
			handleListHouses(manager)
		case "3":
			handleUpdateHouse(scanner, manager)
		case "4":
			handleDeleteHouse(scanner, manager)
		case "5":
			handleBookHouse(scanner, manager)
		case "6":
			handleListBookings(manager)
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printMenu() {
	fmt.Println("\n=== Airbnb House Management ===")
	fmt.Println("1. Add house")
	fmt.Println("2. List houses")
	fmt.Println("3. Update house")
	fmt.Println("4. Delete house")
	fmt.Println("5. Book house")
	fmt.Println("6. List bookings")
	fmt.Println("0. Exit")
}

func handleAddHouse(scanner *bufio.Scanner, manager *HouseManager) {
	input, ok := collectHouseInput(scanner)
	if !ok {
		return
	}

	house, err := manager.AddHouse(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("House added with ID %d\n", house.ID)
}

func handleListHouses(manager *HouseManager) {
	houses := manager.ListHouses()
	if len(houses) == 0 {
		fmt.Println("No houses found")
		return
	}

	fmt.Println("\nHouses:")
	for _, h := range houses {
		availability := "No"
		if h.Available {
			availability = "Yes"
		}
		fmt.Printf("ID: %d | %s | %s | $%.2f/night | Max Guests: %d | Available: %s\n",
			h.ID, h.Title, h.Location, h.PricePerNight, h.MaxGuests, availability)
	}
}

func handleUpdateHouse(scanner *bufio.Scanner, manager *HouseManager) {
	id, ok := promptInt(scanner, "House ID to update")
	if !ok {
		return
	}

	input, ok := collectHouseInput(scanner)
	if !ok {
		return
	}

	_, err := manager.UpdateHouse(id, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("House updated")
}

func handleDeleteHouse(scanner *bufio.Scanner, manager *HouseManager) {
	id, ok := promptInt(scanner, "House ID to delete")
	if !ok {
		return
	}

	if err := manager.DeleteHouse(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("House deleted")
}

func handleBookHouse(scanner *bufio.Scanner, manager *HouseManager) {
	houseID, ok := promptInt(scanner, "House ID to book")
	if !ok {
		return
	}

	guestName, ok := promptString(scanner, "Guest name")
	if !ok {
		return
	}

	nights, ok := promptInt(scanner, "Nights")
	if !ok {
		return
	}

	booking, err := manager.BookHouse(houseID, guestName, nights)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Booking created with ID %d. Total price: $%.2f\n", booking.ID, booking.TotalPrice)
}

func handleListBookings(manager *HouseManager) {
	bookings := manager.ListBookings()
	if len(bookings) == 0 {
		fmt.Println("No bookings found")
		return
	}

	fmt.Println("\nBookings:")
	for _, b := range bookings {
		fmt.Printf("ID: %d | House ID: %d | Guest: %s | Nights: %d | Total: $%.2f\n",
			b.ID, b.HouseID, b.GuestName, b.Nights, b.TotalPrice)
	}
}

func collectHouseInput(scanner *bufio.Scanner) (HouseInput, bool) {
	title, ok := promptString(scanner, "Title")
	if !ok {
		return HouseInput{}, false
	}
	location, ok := promptString(scanner, "Location")
	if !ok {
		return HouseInput{}, false
	}
	price, ok := promptFloat(scanner, "Price per night")
	if !ok {
		return HouseInput{}, false
	}
	maxGuests, ok := promptInt(scanner, "Max guests")
	if !ok {
		return HouseInput{}, false
	}
	available, ok := promptBool(scanner, "Available (yes/no)")
	if !ok {
		return HouseInput{}, false
	}

	return HouseInput{
		Title:         title,
		Location:      location,
		PricePerNight: price,
		MaxGuests:     maxGuests,
		Available:     available,
	}, true
}

func promptString(scanner *bufio.Scanner, label string) (string, bool) {
	fmt.Printf("%s: ", label)
	if !scanner.Scan() {
		fmt.Println("Input ended")
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func promptInt(scanner *bufio.Scanner, label string) (int, bool) {
	value, ok := promptString(scanner, label)
	if !ok {
		return 0, false
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Invalid number")
		return 0, false
	}
	return parsed, true
}

func promptFloat(scanner *bufio.Scanner, label string) (float64, bool) {
	value, ok := promptString(scanner, label)
	if !ok {
		return 0, false
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Invalid decimal number")
		return 0, false
	}
	return parsed, true
}

func promptBool(scanner *bufio.Scanner, label string) (bool, bool) {
	value, ok := promptString(scanner, label)
	if !ok {
		return false, false
	}

	normalized := strings.ToLower(strings.TrimSpace(value))
	switch normalized {
	case "yes", "y", "true", "1":
		return true, true
	case "no", "n", "false", "0":
		return false, true
	default:
		fmt.Println("Invalid boolean value, use yes/no")
		return false, false
	}
}
