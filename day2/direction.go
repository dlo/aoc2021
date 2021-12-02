package day2

type Direction int

const (
	Up Direction = iota
	Down
	Forward
)

func DirectionFromRawDirection(raw string) Direction {
	switch raw {
	case "up":
		return Up
	case "down":
		return Down
	default:
		return Forward
	}
}
