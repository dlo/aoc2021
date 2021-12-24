package day16

import (
	"errors"
	"fmt"
	"github.com/dlo/aoc2021/utils"
)

type BITSTransmission struct {
	raw string
}

type Packet struct {
	version      int
	packetTypeId int
	length       int
	value        int
	subpackets   []Packet
}

func (p Packet) VersionNumberSum() int {
	result := p.version
	for _, packet := range p.subpackets {
		result += packet.VersionNumberSum()
	}
	return result
}

func (t BITSTransmission) ValueAt(index int) (int, error) {
	if len(t.raw) == 0 {
		return 0, errors.New(fmt.Sprintf("out of bounds: %d", index))
	}

	if t.raw[index] == '0' {
		return 0, nil
	} else {
		return 1, nil
	}
}

func (t *BITSTransmission) Read(len int) (int, error) {
	result := 0
	for i := 0; i < len; i++ {
		result <<= 1
		value, err := t.ValueAt(i)
		if err != nil {
			return 0, err
		}

		result += value
	}
	t.raw = t.raw[len:]
	return result, nil
}

const (
	SumPacketType = iota
	ProductPacketType
	MinimumPacketType
	MaximumPacketType
	LiteralPacketType = 4
	GreaterThanPacketType
	LessThanPacketType
	EqualPacketType
)

const (
	LengthTypeTotalLengthInBits = iota
	LengthTypeNumberOfSubPackets
)

func (t *BITSTransmission) Process() (*Packet, error) {
	version, _ := t.Read(3)
	packetTypeId, _ := t.Read(3)
	length := 6
	switch packetTypeId {
	case LiteralPacketType:
		number := 0
		for {
			shouldContinue, _ := t.Read(1)

			number <<= 4

			value, err := t.Read(4)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("literal packet error: %s", err.Error()))
			}

			number += value
			length += 5

			if shouldContinue&1 == 0 {
				break
			}
		}
		return &Packet{version, packetTypeId, length, number, []Packet{}}, nil

	default:
		lengthTypeID, _ := t.Read(1)
		length += 1
		var packets []Packet
		switch lengthTypeID {
		case LengthTypeTotalLengthInBits:
			lengthOfSubPackets, _ := t.Read(15)
			length += 15
			for lengthOfSubPackets > 0 {
				packet, err := t.Process()
				if err != nil {
					return nil, errors.New(fmt.Sprintf("sub-packet error: %s", err.Error()))
				}
				packets = append(packets, *packet)
				length += packet.length
				lengthOfSubPackets -= packet.length
			}

		case LengthTypeNumberOfSubPackets:
			subPacketCount, _ := t.Read(11)
			length += 11
			for subPacketCount > 0 {
				packet, err := t.Process()
				if err != nil {
					return nil, errors.New(fmt.Sprintf("sub-packet error: %s", err.Error()))
				}
				packets = append(packets, *packet)
				length += packet.length
				subPacketCount--
			}

		default:
			return nil, errors.New("unknown length type id")
		}
		return &Packet{version, packetTypeId, length, 0, packets}, nil
	}
}

func ImportBITSTransmission(filename string) []BITSTransmission {
	lines := utils.ReadLinesFromFile(filename)
	things := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	var transmissions []BITSTransmission
	for _, line := range lines {
		transmission := ""
		for _, character := range []rune(line) {
			transmission += things[character]
		}

		transmissions = append(transmissions, BITSTransmission{transmission})
	}

	return transmissions
}
