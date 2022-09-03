package entity

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Struct that represent chapter entity.
type Chapter struct {
	Id            uint64
	StoryID       uint64
	Story         Story
	User          User
	Title         string
	Slug          string
	Body          string
	AuthorComment string
	WordCounts    uint64
	ReadingTime   string
	IsPublished   bool
	CreatedAt     uint64
	UpdatedAt     uint64
}

// Generate random ID.
func (s *Chapter) GenerateID() int {
	rand.Seed(time.Now().UnixNano() / 1000000)

	return rand.Intn(999999)
}

// Convert to slug format.
func (s *Chapter) ToSlug(str string) string {
	l := strings.ToLower(str)

	return strings.ReplaceAll(l, " ", "-")
}

// Counting how many characters on body property.
func (s *Chapter) CountChars() int {
	return len([]rune(s.Body))
}

// Calculating how many reading time needed on chapter.
func (s *Chapter) CalculateReadingTime() string {
	result := float64((s.WordCounts / 200))

	return fmt.Sprintf("%s Minutes", strconv.Itoa(int(result)))
}
