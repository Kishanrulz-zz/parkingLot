package parkingAreaTest

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	vehicle "parking/vehicle"
	parkingArea "parking/parkingArea"
	"testing"
)

func TestParking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParkingArea Suite")
}

var _ = Describe("Parking", func() {
	Context("When parking area does not exist", func() {
		/*Context("When user tries to park a car", func() {
			FIt("Should return error: Sorry, please create the parking", func() {
				var parkingAggregator parkingArea.ParkingAggregator
				vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				res := parkingAggregator.Park(*vehicle)
				Expect(res).To(Equal("Sorry, please create the parkingAggregator"))
			})
		})*/
	})
	Context("When parking area with 6 slots exist", func() {
		Context("When all the parking slots are occupied", func() {
			It("Should return error: Sorry, parking lot is full", func() {
				parkingAggregator, err := parkingArea.CreateParking(6)
				Expect(err).To(BeNil())
				for j := 0; j<3; j++ {
					for i := 1; i <= 6; i++ {
						vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
						parkingAggregator.Park(*vehicle)
					}
				}

				vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				res := parkingAggregator.Park(*vehicle)
				Expect(res).To(Equal("Sorry, parking lot is full"))
			})
		})
		Context("When all the slots are not full", func() {
			It("Should park the vehicle", func() {
				parkingAggregator, err := parkingArea.CreateParking(6)
				fmt.Println("error", err)
				Expect(err).To(BeNil())
				i := 1
				for i = 1; i <= 5; i++ {
					vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
					parkingAggregator.Park(*vehicle)
				}
				vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				res := parkingAggregator.Park(*vehicle)
				Expect(res).To(Equal(fmt.Sprint("Allocated slot number: ", i-1)))
			})
		})
	})
	/*Describe("Remove vehicle", func() {
		Context("When there is no vehicle parked at a specified spot", func() {
			Context("When remove vehicle command is fired", func() {
				It("Should return the error: Slot was already empty", func() {
					parking := parkingArea.CreateParking(6)
					res := parking.RemoveVehicle(2)
					Expect(res).To(Equal("Slot was already empty"))
				})
			})
		})
		Context("When there is a vehicle parked at a specified spot", func() {
			Context("When remove vehicle command is fired", func() {
				It("Should return: Slot number free", func() {
					parking := parkingArea.CreateParking(6)
					vehicle := vehicle.CreateVehicle("KA-01-HH-1234", "White")
					parking.Park(*vehicle)
					res := parking.RemoveVehicle(0)
					Expect(res).To(Equal("Slot number 0 free"))
				})
			})
		})
	})

	Describe("Get registration number from color", func() {
		Context("When there are cars parked with color white", func() {
			It("Should return the registration number of all the cars parked with color white", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "White")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				carRegNumber := []string{"KA-01-HH-1234", "KA-01-HH-1235"}
				expectedCardRegNumber := parking.GetRegistrationNumberFromColor("White")
				Expect(carRegNumber).To(Equal(expectedCardRegNumber))

			})
		})
		Context("When there are no cars parked with color white", func() {
			It("Should return an empty array", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "Blue")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "Black")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				carRegNumber := []string{}
				expectedCarRegNumber := parking.GetRegistrationNumberFromColor("White")
				Expect(carRegNumber).To(Equal(expectedCarRegNumber))

			})
		})
	})


	Describe("Get slot number from color", func() {
		Context("When there are cars parked with color white", func() {
			It("Should return the slot where all the cars parked with color white is parked", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "White")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				carSlot := []int{0, 1}
				expectedCarSlot := parking.GetSlotNumberFromColor("White")
				Expect(carSlot).To(Equal(expectedCarSlot))

			})
		})
		Context("When there are no cars parked with color white", func() {
			It("Should return an empty array", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "Blue")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "Black")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				carRegNumber := []int{}
				expectedCardRegNumber := parking.GetSlotNumberFromColor("White")
				Expect(carRegNumber).To(Equal(expectedCardRegNumber))

			})
		})
	})

	Describe("Get slot number from registration number", func() {
		Context("When there are cars parked with registration number: KA-01-HH-1234", func() {
			It("Should return the slot where the car is parked", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "White")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "White")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				carSlot := 2
				expectedCarSlot, err := parking.GetSlotNumberFromRegNumber("KA-01-HH-1236")
				Expect(err).To(BeNil())
				Expect(carSlot).To(Equal(expectedCarSlot))
			})
		})
		Context("When there are no cars parked with color white", func() {
			It("Should return an empty array", func() {
				parking := parkingArea.CreateParking(6)
				vehicle1 := vehicle.CreateVehicle("KA-01-HH-1234", "Blue")
				vehicle2 := vehicle.CreateVehicle("KA-01-HH-1235", "Black")
				vehicle3 := vehicle.CreateVehicle("KA-01-HH-1236", "Blue")
				parking.Park(*vehicle1)
				parking.Park(*vehicle2)
				parking.Park(*vehicle3)

				expectedCarSlot, err := parking.GetSlotNumberFromRegNumber("KA-01-HH-1237")
				Expect(err).To(HaveOccurred())
				Expect(expectedCarSlot).To(Equal(-1))

			})
		})
	})

	Describe("Call Commands", func() {
		parking := parkingArea.CreateParking(6)
		Context("park", func() {
			Context("When incorrect command is passed", func() {
				It("Should return: Insufficient arguments", func() {
					parking.CallCommands([]string{"park", "KA-01-HH-1234"})
				})
			})

		})
	})*/

})
