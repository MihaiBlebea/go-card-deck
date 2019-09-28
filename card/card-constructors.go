package card

func New(suit Suit, rank Rank) Card {
	return Card{
		rank,
		suit,
	}
}
