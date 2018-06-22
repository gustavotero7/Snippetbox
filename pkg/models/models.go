package models

import (
	"time"
)

//Snippet _
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type Snippets []*Snippet
