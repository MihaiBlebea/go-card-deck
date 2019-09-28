package card

func (c Card) EqualColor(card CardType) bool {
	suit1, _ := c.GetSuit()
	suit2, _ := card.GetSuit()
	if suit1 == suit2 {
		return true
	}
	return false
}

func (c Card) EqualSuit(card CardType) bool {
	return c.EqualColor(card)
}

func (c Card) EqualRank(card CardType) bool {
	if c.GetRank() == card.GetRank() {
		return true
	}
	return false
}

func (c Card) IsBiggerThen(card CardType) bool {
	if c.GetRank() > card.GetRank() {
		return true
	}
	return false
}

func (c Card) IsSmallerThen(card CardType) bool {
	if c.GetRank() < card.GetRank() {
		return true
	}
	return false
}
