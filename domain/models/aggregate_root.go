package models

type aggregateRoot struct {
	aggregate
	version int
}

func (a aggregateRoot) GetVersion() int {
	return a.version
}
