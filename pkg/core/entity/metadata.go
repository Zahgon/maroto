package entity

import (
	"time"
)

// Metadata is the representation of a PDF metadata.
type Metadata struct {
	Author       *Utf8Text
	Creator      *Utf8Text
	Subject      *Utf8Text
	Title        *Utf8Text
	CreationDate *time.Time
	KeywordsStr  *Utf8Text
}

// AppendMap appends the metadata to a map.
func (m *Metadata) AppendMap(mp map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Utf8Text is the representation of a text with a flag to indicate if it's UTF8.
type Utf8Text struct {
	Text string
	UTF8 bool
}

// ToString returns a string representation of the text.
func (u *Utf8Text) ToString() string { _ = "STUB: not implemented"; return "" }
