package simulator

import (
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"os"
	parking "parking/parkingArea"
	utils "parking/utils"
)

// simulates parking by reading commands from terminal
func SimulateParkingInteractive() {
	reader := bufio.NewReader(os.Stdin)
	parking, err := createParkingInteractive(reader)
	if err != nil {
		panic(err)
	}
	for {
		text, _ := reader.ReadString('\n')
		cmd := []string{text}
		parking.CallCommands(cmd)
	}
}

// helper function to create parking
func createParkingInteractive(reader *bufio.Reader) (*parking.Parking, error) {
	parkingArea := &parking.Parking{}
	cmd := []string{}
	for {
		text, _ := reader.ReadString('\n')
		cmd = strings.Split(text, " ")
		if cmd[0] != "create_parking_lot" {
			fmt.Println("Please create a parking first")
			continue
		}
		break
	}
	slots, err := strconv.Atoi(strings.Split(cmd[1], "\n")[0])
	if err != nil {
		fmt.Println("err", err)
		return parkingArea, err
	}
	parkingArea = parking.CreateParking(slots)
	fmt.Println("Created a parking lot with ", len(parkingArea.Slots), " slots")
	return parkingArea, nil
}

// simulates parking by reading commands from a file
func SimulateParkingFromFile(fileName string) {
	commands, err := GetCommands(fileName)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	slots, err := GetSlotsCount(commands[0])
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	parkingArea := parking.CreateParking(slots)
	fmt.Println("Created a parking lot with ", len(parkingArea.Slots), " slots")
	parkingArea.CallCommands(commands[1:])
}


func GetSlotsCount(cmd string) (int, error) {
	slotsString := strings.Split(cmd, " ")
	slots, err := strconv.Atoi(slotsString[1])
	return slots, err
}

func GetCommands(fileName string) ([]string, error) {
	commands := []string{}
	file, err := utils.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file",err)
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	return commands, nil
}