package entity

import (
	"database/sql"
	"math/rand"
	"time"
)

// Struct that represent user entity.
type User struct {
	Id          uint64
	Name        string
	Username    string
	Email       string
	Password    string
	Sex         uint8
	Bio         *sql.NullString
	DateOfBirth *sql.NullInt64
	PhoneNumber *sql.NullString
	Twitter     *sql.NullString
	Instagram   *sql.NullString
	Facebook    *sql.NullString
	CreatedAt   uint64
}

// Generate random ID.
func (u *User) GenerateID() int {
	rand.Seed(time.Now().UnixNano() / 1000000)

	return rand.Intn(999999)
}

// Get gender name by int.
func (u *User) GetGenderName() string {
	switch u.Sex {
	case 0:
		return "Not known"
	case 1:
		return "Male"
	case 2:
		return "Female"
	case 9:
		return "Not applicable"
	default:
		return "unknown"
	}
}
