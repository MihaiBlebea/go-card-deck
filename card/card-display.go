package card

func (c *Card) Display() map[string]string {
	cardMap := make(map[string]string)
	cardMap["suit"] = c.GetSuitString()
	cardMap["rank"] = c.GetRankString()
	cardMap["full"] = c.GetRankString() + " of " + c.GetSuitString()
	return cardMap
}

func (c *Card) GetRankString() string {
	var rank string
	switch c.Rank {
	case Ace:
		rank = "Ace"
	case Two:
		rank = "Two"
	case Three:
		rank = "Three"
	case Four:
		rank = "Four"
	case Five:
		rank = "Five"
	case Six:
		rank = "Six"
	case Seven:
		rank = "Seven"
	case Eight:
		rank = "Eight"
	case Nine:
		rank = "Nine"
	case Ten:
		rank = "Ten"
	case Jack:
		rank = "Jack"
	case Queen:
		rank = "Queen"
	case King:
		rank = "King"
	}
	return rank
}

func (c *Card) GetSuitString() string {
	var suit string
	switch c.Suit {
	case Heart:
		suit = "heart"
	case Diamond:
		suit = "diamond"
	case Club:
		suit = "club"
	case Spade:
		suit = "spade"
	}
	return suit
}
