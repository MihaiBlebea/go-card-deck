package card

type Suit int

// All the suit types with a value of integer starting from 1
const (
	_ Suit = iota
	Heart
	Spade
	Club
	Diamond
)

type Rank int

// All the rank types with a value of integer starting from 1
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func GetRankTypes() [13]Rank {
	return [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

func GetSuitTypes() [4]Suit {
	return [4]Suit{Heart, Diamond, Club, Spade}
}

// Card model is the main actor of our package
type Card struct {
	Rank
	Suit
}

// GetSuit returns the suit and color of the card in literal, string form
func (c *Card) GetSuit() int {
	return int(c.Suit)
}

// GetRank returns an integer value of the card
func (c *Card) GetRank() int {
	return int(c.Rank)
}
