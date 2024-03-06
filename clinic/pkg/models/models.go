package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
	Balance        int
	IsAdmin        bool
	Phone          string
}

type Product struct {
	ID          int
	Name        string
	ImagePath   string
	Description string
	Price       float64
	Category    string
}

type News struct {
	ID        int
	Title     string
	Content   string
	Category  string
	ImagePath string
	Created   time.Time
}
