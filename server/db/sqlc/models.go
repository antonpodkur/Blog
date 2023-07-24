// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Photo     string
	Verified  bool
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
