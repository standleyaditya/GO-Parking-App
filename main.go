package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	RegistrationNumber string
}

type ParkingLot struct {
	Slots [](*Car)
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Slots: make([](*Car), capacity),
	}
}

func (pl *ParkingLot) Park(carNumber string) string {
	for i := 0; i < len(pl.Slots); i++ {
		if pl.Slots[i] == nil {
			pl.Slots[i] = &Car{RegistrationNumber: carNumber}
			return fmt.Sprintf("Allocated slot number: %d", i+1)
		}
	}
	return "Sorry, parking lot is full"
}

func (pl *ParkingLot) Leave(carNumber string, hours int) string {
	for i, car := range pl.Slots {
		if car != nil && car.RegistrationNumber == carNumber {
			pl.Slots[i] = nil
			charge := 10 // Base charge for first 2 hours
			if hours > 2 {
				charge += 10 * (hours - 2)
			}
			return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d", carNumber, i+1, charge)
		}
	}
	return fmt.Sprintf("Registration number %s not found", carNumber)
}

func (pl *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i, car := range pl.Slots {
		if car != nil {
			fmt.Printf("%d %s\n", i+1, car.RegistrationNumber)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	var parkingLot *ParkingLot
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch strings.ToLower(parts[0]) {
		case "create_parking_lot":
			if len(parts) != 2 {
				fmt.Println("Invalid command: create_parking_lot <capacity>")
				continue
			}
			capacity, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Invalid capacity: %s\n", parts[1])
				continue
			}
			parkingLot = NewParkingLot(capacity)
			fmt.Printf("Created parking lot with %d slots\n", capacity)

		case "park":
			if parkingLot == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			if len(parts) != 2 {
				fmt.Println("Invalid command: park <car_number>")
				continue
			}
			fmt.Println(parkingLot.Park(parts[1]))

		case "leave":
			if parkingLot == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			if len(parts) != 3 {
				fmt.Println("Invalid command: leave <car_number> <hours>")
				continue
			}
			hours, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Printf("Invalid hours: %s\n", parts[2])
				continue
			}
			fmt.Println(parkingLot.Leave(parts[1], hours))

		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			parkingLot.Status()

		default:
			fmt.Printf("Unknown command: %s\n", parts[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
}
