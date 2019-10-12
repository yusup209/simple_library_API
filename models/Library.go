package models

import (
	"time"
)

type Users struct {
	ID        uint       `gorm:"size:10;primary_key;not null;auto_increment" json:"id"`
	Email     string     `gorm:"size:255;"json:"email"`
	Username  string     `gorm:"size:100;not null" json:"username"`
	Password  string     `gorm:"size:100;not null" json:"password"`
	Avatar    string     `gorm:"size:255" json:"avatar"`
	Role      string     `gorm:"size:100" json:"role"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Category struct {
	ID        uint       `gorm:"size:10;primary_key;not null;auto_increment" json:"id"`
	Name      string     `gorm:"size:125;not null" json:"name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Bookshelf struct {
	ID        uint       `gorm:"size:10;primary_key;not null;auto_increment" json:"id"`
	Code      string     `gorm:"size: 100;not null" json:"bookshelf_code"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Book struct {
	ID          uint       `gorm:"size:10;primary_key;not null;auto_increment" json:"id"`
	ISBN        uint       `gorm:"size:100" json:"isbn"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Author      string     `gorm:"size:100" json:"author"`
	CategoryID  uint       `gorm:"size:10;ForeignKey:ID" json:"category_id"`
	Category    Category   `json:"category"`
	BookshelfID uint       `gorm:"size:10;ForeignKey:ID" json:"bookshelf_id"`
	Bookshelf   Bookshelf  `json:"bookshelf"`
	Publisher   string     `gorm:"size:255;" json:"publisher"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type BookTransaction struct {
	ID        uint       `gorm:"size:10;primary_key;not null;auto_increment "json:"id"`
	Name      string     `gorm:"size: 125;not null" "json:"name"`
	NIS       uint       `gorm:"size: 100;not null" "json:"nis"`
	BookID    uint       `gorm:"size:10;ForeignKey:ID" json:"book_id"`
	Book      Book       `json:"book"`
	Status    string     `gorm:"size: 100;not null" "json:"status"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Guestbook struct {
	ID        uint       `gorm:"size:10;primary_key;not null;auto_increment" json:"id"`
	Name      string     `gorm:"size:125;not null" json:"name"`
	NIS       uint       `gorm:"size:100;not null" json:"nis"`
	Purpose   string     `gorm:"size:255" json:"purpose"`
	CreatedAt time.Time  `json:"tanggal"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
