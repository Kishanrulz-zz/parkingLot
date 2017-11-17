package vehicle


//vehicle holds the meta of a vehicle
type Vehicle struct {
	RegistrationNumber string //Registration Number of the vehicle
	Color              string //Color of the vehicle
}

// CreateVehicle creates a new vehicle whenever a request comes for parking a vehicle
func CreateVehicle(number string, color string) *Vehicle {
	return &Vehicle{
		RegistrationNumber: number,
		Color:              color,
	}
}
