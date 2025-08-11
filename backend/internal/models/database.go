package models

import (
	"time"
)

// User represents a user in the database
type User struct {
	ID                 uint      `gorm:"primarykey" json:"id"`
	Username           string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash       string    `gorm:"not null" json:"password_hash"`
	FirstLogin         bool      `gorm:"default:true" json:"first_login"`
	LoginCount         int       `gorm:"default:0" json:"login_count"`
	LastPasswordChange time.Time `gorm:"autoCreateTime" json:"last_password_change"`
	LastLoginAt        time.Time `gorm:"autoCreateTime" json:"last_login_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CVFile represents a CV file stored in the database
type CVFile struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	FileName     string    `gorm:"not null" json:"fileName"`
	OriginalName string    `gorm:"not null" json:"originalName"`
	ContentType  string    `gorm:"not null;default:'application/pdf'" json:"contentType"`
	FileSize     int64     `gorm:"not null" json:"fileSize"`
	FileData     []byte    `gorm:"not null" json:"-"` // Don't include in JSON
	IsCurrent    bool      `gorm:"default:true;index" json:"isCurrent"`
	CreatedAt    time.Time `gorm:"autoCreateTime;index" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName sets the table name for CVFile
func (CVFile) TableName() string {
	return "cv_files"
}