package bridge

import "fmt"

type Windows struct {
	*Computer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func NewWindows(printer Printer) *Windows {
	computer := NewComputer(printer)
	return &Windows{
		Computer: computer,
	}
}
