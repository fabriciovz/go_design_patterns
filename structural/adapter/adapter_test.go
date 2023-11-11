package adapter

import "testing"

func TestAdapter(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {

		client := &Client{}
		mac := &Mac{}

		client.InsertLightningConnectorIntoComputer(mac)

		windowsMachine := &Windows{}
		windowsMachineAdapter := &WindowsAdapter{
			windowMachine: windowsMachine,
		}

		client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	})
}
