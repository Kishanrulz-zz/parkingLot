package parkingArea

import (
	"errors"
	"fmt"
	vehicle "parking/vehicle"
	"sort"
	"strconv"
	"strings"
)

type ParkingAggregator struct {
	Parkings       []Parking
	SlotsAvailable map[int]int
}

// Parkings holds the structure defining a parking lot
type Parking struct {
	Slots   []int                   //total number of slots available
	vehicle map[int]vehicle.Vehicle //map to store slotID as key and vehicles as value
	//Empty   int             //count for empty slots
}

// CreateParking creates a parking slot and returns a pointer to the slot
func CreateParking(slots int) (*ParkingAggregator, error) {
	if slots == 0 {
		return nil, errors.New("Sorry, please pass the slots count")
	}
	parkingAggregator := &ParkingAggregator{
		SlotsAvailable: make(map[int]int, 3),
	}

	parking := []Parking{}
	for i := 0; i < 3; i++ {
		parking = append(parking, Parking{
			Slots:   make([]int, slots),
			vehicle: make(map[int]vehicle.Vehicle, slots),
		})
	}
	parkingAggregator.Parkings = parking
	return parkingAggregator, nil
}

// Park function parks a vehicle passed in the argument for a particular parking area
func (parkingAggregator *ParkingAggregator) Park(vehicle vehicle.Vehicle) string {
	//Checking if parking lot is full
	var id int

	slotsAvailable := []int{}
	for _, available := range parkingAggregator.SlotsAvailable {
		slotsAvailable = append(slotsAvailable, available)
	}

	sort.Ints(slotsAvailable)

	for parkingId, available := range parkingAggregator.SlotsAvailable {
		if available == slotsAvailable[len(slotsAvailable)-1] {
			id = parkingId
		}
	}

	/*if len(parking.Slots) == len(parking.vehicle) {
		continue
	}*/
	// allocating the slot which is near to the entrance
	parking := parkingAggregator.Parkings[id]
	for key, val := range parking.Slots {
		if val == 0 {
			parking.Slots[key] = 1
			parking.vehicle[key] = vehicle
			return "Allocated slot number: " + strconv.Itoa(key)
		}
	}

	return "Sorry, parking lot is full"
}

// RemoveVehicle removes a vehicle from a slot passed as an argument to the function for a given parking lot
func (parkingAggregator *ParkingAggregator) RemoveVehicle(parkingLotId, slot int) string {
	//Checking whether the slot was already empty or not
	if len(parkingAggregator.Parkings) < parkingLotId {
		return "Sorry, invalid parking lot id"
	}
	parking := parkingAggregator.Parkings[parkingLotId]

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
func (parkingAggregator *ParkingAggregator) GetRegistrationNumberFromColor(color string) []string {
	regNumber := []string{}
	for _, parking := range parkingAggregator.Parkings {
		for _, val := range parking.vehicle {
			if val.Color == color {
				regNumber = append(regNumber, val.RegistrationNumber)
			}
		}
	}

	return regNumber
}

// GetSlotNumberFromColor returns the slot number of all the vehicles for
// a color passed as an argument to the function for a given parking lot
func (parkingAggregator *ParkingAggregator) GetSlotNumberFromColor(color string) []int {
	slotNumber := []int{}
	for _, parking := range parkingAggregator.Parkings {
		for key, val := range parking.vehicle {
			if val.Color == color {
				slotNumber = append(slotNumber, key)
			}
		}
	}
	return slotNumber
}

// GetSlotNumberFromRegNumber returns the slot number of a vehicle passing
// the registration number to the function for a given parking lot
func (parkingAggregator *ParkingAggregator) GetSlotNumberFromRegNumber(regNumber string) (int, int, error) {
	for id, parking := range parkingAggregator.Parkings {
		for key, val := range parking.vehicle {
			if val.RegistrationNumber == regNumber {
				return id, key, nil
			}
		}
	}
	return -1, -1, errors.New("Not Found")
}

// CallCommands is a utility function to run the commands
// for a parking lot from either an input file or command line
func (parkingAggregator *ParkingAggregator) CallCommands(commands []string) {
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
			msg := parkingAggregator.Park(*vehicle)
			fmt.Println(msg)
		case "leave":
			if len(cmd) != 3 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			parkingLotId, _ := strconv.Atoi(action[0])
			slot, _ := strconv.Atoi(action[1])
			//Checking whether parking is full or not
			if slot > len(parkingAggregator.Parkings[0].Slots)-1 {
				fmt.Println("Slot not available")
				return
			}
			fmt.Println(parkingAggregator.RemoveVehicle(parkingLotId, slot))
		case "registration_numbers_for_cars_with_colour":
			if len(cmd) != 2 {
				fmt.Println("Insufficient arguments")
				return
			}
			action := strings.Split(cmd[1], "\n")
			regNumbers := parkingAggregator.GetRegistrationNumberFromColor(action[0])
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
			slotNumbers := parkingAggregator.GetSlotNumberFromColor(action[0])
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
			parkingLotId, slotNumber, err := parkingAggregator.GetSlotNumberFromRegNumber(action[0])
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(parkingLotId, slotNumber)
			}
		case "status":
			fmt.Println("Slot No.\tRegistration No.\tslots\tparkingLot")
			slots := []int{}
			for _, parking := range parkingAggregator.Parkings {
				for slot, _ := range parking.vehicle {
					slots = append(slots, slot)

				}
			}

			sort.Ints(slots)
			for parkingLotId, parking := range parkingAggregator.Parkings {
				for _, slot := range slots {
					fmt.Println(slot, "\t\t", parking.vehicle[slot].RegistrationNumber, "\t\t", parking.vehicle[slot].Color, "\t\t", parkingLotId)
				}
			}

		default:
			fmt.Println("Invalid command")
		}
	}
}
