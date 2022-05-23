package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Tittle    string    `gorm:"size:255;not null:unique" json:"tittle"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"authorId"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Tittle = html.EscapeString(strings.TrimSpace(p.Tittle))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {
	if p.Tittle == "" {
		return errors.New("required Title")
	}
	if p.Content == "" {
		return errors.New("required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("reqiored Author")
	}
	return nil
}

func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	posts := []Post{}
	err := db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error

	if err != nil {
		return &[]Post{}, err
	}

	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}

	return &posts, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	err := db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error

	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}

	return p, nil
}

func (p *Post) UpdatedPost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Post{}).Where("id = ? and authorId = ?", pid, uid).Take(&Post{}).Delete(&Post{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
