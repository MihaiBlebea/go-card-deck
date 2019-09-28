package deck

type DeckError struct {
	message string
}

func (err DeckError) Error() string {
	return err.message
}

func newDeckError(message string) DeckError {
	return DeckError{message}
}