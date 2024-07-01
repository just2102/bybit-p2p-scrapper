package data

type Side string

const (
	SELL Side = "0"
	BUY  Side = "1"
)

var validSides = map[Side]bool{
	SELL: true,
	BUY:  true,
}

func IsValidSide(side Side) bool {
	return validSides[side]
}
