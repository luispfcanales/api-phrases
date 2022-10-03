package domain

// Phrase is model to phrases
type Phrase struct {
	Message string `json:"message,omitempty"`
	Author  string `json:"author,omitempty"`
}

type Phrases []Phrase
