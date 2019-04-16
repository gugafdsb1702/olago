package model

type Galinha interface {
	Cacareja() string
}

type Pata interface {
	Quack() string
}

type Ave struct {
	Nome string
}

func (a Ave) Cacareja() string {
	return "Cócoricó..."
}

func (a Ave) Quack() string {
	return "Quack, Quack..."
}
