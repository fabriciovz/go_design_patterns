package chain_of_responsibility

import "testing"

func TestCommand(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		cashier := &Cashier{}

		//Set next for medical department
		medical := &Medical{}
		medical.setNext(cashier)

		//Set next for doctor department
		doctor := &Doctor{}
		doctor.setNext(medical)

		//Set next for reception department
		reception := &Reception{}
		reception.setNext(doctor)

		patient := &Patient{name: "abc"}
		//Patient visiting
		reception.execute(patient)
	})
}
