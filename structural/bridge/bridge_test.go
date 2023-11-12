package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		hpPrinter := &Hp{}
		epsonPrinter := &Epson{}

		macComputer := NewMac(hpPrinter)
		macComputer.Print()
		fmt.Println()

		macComputer.SetPrinter(epsonPrinter)
		macComputer.Print()
		fmt.Println()

		winComputer := NewWindows(hpPrinter)
		winComputer.Print()
		fmt.Println()

		winComputer.SetPrinter(epsonPrinter)
		winComputer.Print()
		fmt.Println()

	})
}
