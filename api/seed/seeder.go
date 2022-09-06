package seed

import (
	"github.com/LordCeilan/crud-restful-api-go/api/models"
	"github.com/jinzhu/gorm"

	"log"
	"time"
)

var users = []models.User{
	{
		ID:       0,
		NickName: "wonder",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	{
		ID:       0,
		NickName: "martin",
		Email:    "luther@gmail.com",
		Password: "nya",
	},
}

var posts = []models.Post{
	{
		ID:        0,
		Tittle:    "Title 1",
		Content:   "Hello world 1",
		Author:    models.User{},
		AuthorID:  0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Tittle:  "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("no puede hacer drop a la tabla: %v", err)
	}

	err = db.Debug().Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("no se puede migrar tabla: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
