package job

import "time"

// JobPost struct
type JobPost struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   User      `json:"created_by"`
}

// User struct
type User struct {
	ID          string      `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	AvatarUrl   string      `json:"avatar_url"`
	JobTitle    string      `json:"job_title,omitempty"`
	Company     string      `json:"company,omitempty"`
	CreatedAt   time.Time   `json:"-"`
	Headline    string      `json:"headline,omitempty"`
	Profile     UserProfile `json:"profiles,omitempty"`
	HuntingMode string      `json:"hunting_mode"`
}

type UserProfile struct {
	LinkedIn string `json:"linkedin"`
}
