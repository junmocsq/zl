package monitor

type Parser interface {
	Add()
	Parse()
	Storage()
}
