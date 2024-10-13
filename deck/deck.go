package deck

import (
	"fmt"
	"strconv"
)

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Clubs:
		return "CLUBS"
	case Diamonds:
		return "DIAMONDS"
	case Harts:
		return "HARTS"
	default:
		panic("Invalid card suite")
	}
}

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)

type Card struct {
	suite Suit
	value int
}

func (c Card) String() string {

	value := strconv.Itoa(c.value)

	if c.value == 1 {
		value = "ACE"
	}

	return fmt.Sprintf("%s of %s %s", value, c.suite, suiteToUnicode(c.suite))
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("This value of card cannot be higher than 13")
	}

	return Card{
		suite: s,
		value: v,
	}
}

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0

	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return d
}

func suiteToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Harts:
		return "♥"
	default:
		panic("Invalid card suite")
	}
}
