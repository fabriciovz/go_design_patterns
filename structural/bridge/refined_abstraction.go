package bridge

import "fmt"

type Mac struct {
	*Computer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func NewMac(printer Printer) *Mac {
	computer := NewComputer(printer)
	return &Mac{
		Computer: computer,
	}
}
