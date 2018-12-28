package core

import "fmt"

var (
	ranks = [13]string{
		"A",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
	}

	// TODO: suits as runes
	suits = [4]string{
		"\u2665", // ♥ hearts
		"\u2663", // ♣ clovers
		"\u2666", // ♦ tiles
		"\u2660", // ♠ pikes
	}
)

// Card represents a simple priimtive in matrices
type Card struct {
	Number uint8
	Suit   string
	Rank   string
	Short  string
	Long   string
}

// NewCardFromNumber returns type *Card from the given number (0, 52]
func NewCardFromNumber(n uint8) (*Card, error) {
	if n <= 0 || n > 52 {
		return nil, fmt.Errorf("Unable to create card: invalid number -> %v", n)
	}

	return &Card{
		Number: uint8(n),
		Suit:   suits[(n-1)/13],
		Rank:   ranks[(n-1)%13],
		Short:  "",
		Long:   "",
	}, nil
}
