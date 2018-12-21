package blance

type Blancer interface {
	Doblance([]*Instance) (*Instance,error)
}