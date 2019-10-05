package card

// EqualColor checks if the color of two cards is the same. Returns a boolean
func (c Card) EqualColor(card Card) bool {
	if c.GetSuit() == card.GetSuit() {
		return true
	}
	return false
}

// EqualSuit checks if the suit of two cards is the same. Returns a boolean
func (c Card) EqualSuit(card Card) bool {
	if c.GetSuit() == card.GetSuit() {
		return true
	}
	return false
}

// EqualRank checks if the rank of two cards is the same. Returns a boolean
func (c Card) EqualRank(card Card) bool {
	if c.GetRank() == card.GetRank() {
		return true
	}
	return false
}

// IsBiggerThen returns true if the supplied card is smaller in value then this card
func (c Card) IsBiggerThen(card Card) bool {
	if c.GetRank() > card.GetRank() {
		return true
	}
	return false
}

// IsBiggerThen returns true if the supplied card is bigger in value then this card
func (c Card) IsSmallerThen(card Card) bool {
	if c.GetRank() < card.GetRank() {
		return true
	}
	return false
}
