package card

func (c *Card) Display() map[string]string {
	cardMap := make(map[string]string)
	cardMap["suit"] = c.GetSuitString()
	cardMap["rank"] = c.GetRankString()
	cardMap["full"] = c.GetRankString() + " of " + c.GetSuitString()
	return cardMap
}
