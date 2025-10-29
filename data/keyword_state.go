package data

type keywordState int

const (
	FIRST_WORD_IN_LINE keywordState = iota
	AFTER_SPACE
	AFTER_DOT
	AFTER_COMMA
)
