package spacing

type Spacing struct {
	Unit int
}

func New() Spacing {
	return Spacing{Unit: 8}
}
