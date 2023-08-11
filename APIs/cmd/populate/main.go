package populate

import (
	"fmt"
	"log"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, quantity int) {
	tx := db.Begin()
	if tx.Error != nil {
		log.Default().Println("Error creating transaction")
	}
	defer tx.Rollback()

	for i := 0; i < quantity; i++ {
		username := "user" + fmt.Sprintf("%d", i)
		u, err := entity.NewUser(username, username+"@mail.com", "user")
		if err != nil {
			log.Fatal(err)
		}
		err = tx.Create(&u).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

}
