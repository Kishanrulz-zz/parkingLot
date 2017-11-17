package parkingArea

import (
	vehicle "parking/vehicle"
	"strconv"
	"errors"
	"strings"
	"fmt"
	"sort"
)

// Parking holds the structure defining a parking lot
type Parking struct {
	Slots   []int           //total number of slots available
	vehicle map[int]vehicle.Vehicle //map to store slotID as key and vehicles as value
	//Empty   int             //count for empty slots
}


// CreateParking creates a parking slot and returns a pointer to the slot
func CreateParking(slots int) *Parking {
	return &Parking{
		Slots:   make([]int, slots),
		vehicle: make(map[int]vehicle.Vehicle, slots),
		//Empty:   slots,
	}
}

// Park function parks a vehicle passed in the argument for a particular parking area
func (parking *Parking) Park(vehicle vehicle.Vehicle) string {
	//Checking if parking lot is full
	if len(parking.Slots) == 0 {
		return "Sorry, please create the parking"
	}
	if len(parking.Slots) == len(parking.vehicle) {
		return "Sorry, parking lot is full"
	}
	// allocating the slot which is near to the entrance
	for key, val := range parking.Slots {
		if val == 0 {
			parking.Slots[key] = 1
			parking.vehicle[key] = vehicle
			//parking.Empty--
			return "Allocated slot number: " + strconv.Itoa(key)
		}
	}
	return ""
}

// RemoveVehicle removes a vehicle from a slot passed as an argument to the function for a given parking lot
func (parking *Parking) RemoveVehicle(slot int) string {
	//Checking whether the slot was already empty or not
	if parking.Slots[slot] == 0 {
		return "Slot was already empty"
	}
	//removing the vehicle
	parking.Slots[slot] = 0
	//parking.Empty++
	delete(parking.vehicle, slot)

	return "Slot number " + strconv.Itoa(slot) + " free"
}

// GetRegistrationNumberFromColor returns the registration number of all the vehicles for
// a color passed as an argument to the function for a given parking lot
func (parking *Parking) GetRegistrationNumberFromColor(color string) []string {
	regNumber := []string{}
	for _, val := range parking.vehicle {
		if val.Color == color {
			regNumber = append(regNumber, val.RegistrationNumber)
		}
	}
	return regNumber
}


// GetSlotNumberFromColor returns the slot number of all the vehicles for
// a color passed as an argument to the function for a given parking lot
func (parking *Parking) GetSlotNumberFromColor(color string) []int {
	slotNumber := []int{}
	for key, val := range parking.vehicle {
		if val.Color == color {
			slotNumber = append(slotNumber, key)
		}
	}
	return slotNumber
}

// GetSlotNumberFromRegNumber returns the slot number of a vehicle passing
// the registration number to the function for a given parking lot
func (parking *Parking) GetSlotNumberFromRegNumber(regNumber string) (int, error) {
	for key, val := range parking.vehicle {
		if val.RegistrationNumber == regNumber {
			return key, nil
		}
	}
	return -1, errors.New("Not Found")
}

// CallCommands is a utility function to run the commands
// for a parking lot from either an input file or command line
func (parking *Parking) CallCommands(commands []string) {
	for _, command := range commands {
		newLineSplit := strings.Split(command, "\n")
		cmd := strings.Split(newLineSplit[0], " ")
		switch cmd[0] {
		case "park":
			if len(cmd) != 3 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[2], "\n")
			vehicle := vehicle.CreateVehicle(cmd[1], action[0])
			msg := parking.Park(*vehicle)
			fmt.Println(msg)
		case "leave":
			if len(cmd) != 2 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			slot, _ := strconv.Atoi(action[0])
			//Checking whether parking is full or not
			if slot > len(parking.Slots)-1 {
				fmt.Println("Slot not available")
				return
			}
			fmt.Println(parking.RemoveVehicle(slot))
		case "registration_numbers_for_cars_with_colour":
			if len(cmd) != 2 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			regNumbers := parking.GetRegistrationNumberFromColor(action[0])
			if len(regNumbers) > 0 {
				for _, regNumber := range regNumbers {
					fmt.Print(regNumber, "\t")
				}
				fmt.Println()
			} else {
				fmt.Println("No vehicle parked of color ", action[0])
			}
		case "slot_numbers_for_cars_with_colour":
			if len(cmd) != 2 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			slotNumbers := parking.GetSlotNumberFromColor(action[0])
			if len(slotNumbers) > 0 {
				for _, slot := range slotNumbers {
					fmt.Print(slot, "\t")
				}
				fmt.Println()
			} else {
				fmt.Println("No vehicle parked of color ", action[0])
			}
		case "slot_number_for_registration_number":
			if len(cmd) != 2 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			slotNumber, err := parking.GetSlotNumberFromRegNumber(action[0])
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(slotNumber)
			}
		case "status":
			fmt.Println("Slot No.\tRegistration No.\tslots")
			slots := []int{}
			for slot, _ := range parking.vehicle {
				slots = append(slots, slot)

			}
			sort.Ints(slots)
			for _, slot := range slots {
				fmt.Println(slot, "\t\t", parking.vehicle[slot].RegistrationNumber, "\t\t", parking.vehicle[slot].Color)
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}

