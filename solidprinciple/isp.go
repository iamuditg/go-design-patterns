package solidprinciple

/*
	Interface Segregation Principle
*/

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {

}

func (o OldFashionedPrinter) Fax(d Document) {

}

// Scan Deprecated ..
func (o OldFashionedPrinter) Scan(d Document) {

}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct {
}

func (p Photocopier) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

func (p Photocopier) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

type MultiFunctionMachine struct {
	Printer
	Scanner
}

func main() {

}
