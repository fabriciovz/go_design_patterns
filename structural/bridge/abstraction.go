package bridge

type Computer struct {
	printer Printer
}

func (m *Computer) Print() {
	m.printer.PrintFile()
}

func (m *Computer) SetPrinter(p Printer) {
	m.printer = p
}

func NewComputer(printer Printer) *Computer {
	return &Computer{printer: printer}
}
