package day02

type Direction int

const (
	Up Direction = iota
	Down
	Forward
	Invalid
)

func DirectionFromRawDirection(raw string) Direction {
	switch raw {
	case "up":
		return Up
	case "down":
		return Down
	case "forward":
		return Forward
	default:
		return Invalid
	}
}
