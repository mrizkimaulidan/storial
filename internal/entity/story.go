package entity

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Struct that represent story entity.
type Story struct {
	Id          uint64
	UserID      uint64
	User        User
	CategoryID  uint64
	Category    Category
	Title       string
	Slug        string
	Description string
	Cover       string
	IsAdult     bool
	IsPublished bool
	CreatedAt   uint64
	UpdatedAt   uint64
}

// Generate random ID.
func (s *Story) GenerateID() int {
	rand.Seed(time.Now().UnixNano() / 1000000)

	return rand.Intn(999999)
}

// Convert to slug format.
func (s *Story) ToSlug(str string) string {
	l := strings.ToLower(str)

	return strings.ReplaceAll(l, " ", "-")
}

// Get cover path.
func (s *Story) CoverPath() string {
	return fmt.Sprintf("/public/cover/%s", s.Cover)
}
