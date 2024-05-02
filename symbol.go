package ttt

type Symbol uint8

const (
	Cross Symbol = iota
	Circle
	Dot
)

func (s Symbol) Other() Symbol {
	if s == Cross {
		return Circle
	}
	return Cross
}
