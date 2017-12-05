package component

type Component interface {
	Template() string
}

type Identity interface {
	ID() string
}
