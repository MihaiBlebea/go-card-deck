package card

// EqualColor checks if the color of two cards is the same. Returns a boolean
func (c Card) EqualColor(card CardType) bool {
	suit1, _ := c.GetSuit()
	suit2, _ := card.GetSuit()
	if suit1 == suit2 {
		return true
	}
	return false
}

// EqualSuit checks if the suit of two cards is the same. Returns a boolean
func (c Card) EqualSuit(card CardType) bool {
	return c.EqualColor(card)
}

// EqualRank checks if the rank of two cards is the same. Returns a boolean
func (c Card) EqualRank(card CardType) bool {
	if c.GetRank() == card.GetRank() {
		return true
	}
	return false
}

// IsBiggerThen returns true if the supplied card is smaller in value then this card
func (c Card) IsBiggerThen(card CardType) bool {
	if c.GetRank() > card.GetRank() {
		return true
	}
	return false
}

// IsBiggerThen returns true if the supplied card is bigger in value then this card
func (c Card) IsSmallerThen(card CardType) bool {
	if c.GetRank() < card.GetRank() {
		return true
	}
	return false
}
