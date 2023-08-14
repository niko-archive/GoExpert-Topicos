package populate

import (
	"fmt"
	"log"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
)

// commit transaction in 100 users
func CreateUsersInBatches(db *gorm.DB, quantity int) {
	// start transaction
	tx := db.Begin()
	if tx.Error != nil {
		log.Default().Println("Error creating transaction")
	}
	defer tx.Rollback()

	// create users
	for i := 0; i < quantity; i++ {
		username := "user" + fmt.Sprintf("%d", i)
		u, err := entity.NewUser(username, username+"@mail.com", "user")
		if err != nil {
			log.Default().Println("Error creating user: NewUser")
		}
		err = tx.Create(&u).Error
		if err != nil {
			log.Default().Println("Error creating user: Create")
		}
		// commit transaction in 100 users
		if i%100 == 0 {
			log.Println("Commiting transaction: +100 users: ", i)
			tx.Commit()
			tx = db.Begin()
		}
	}
	// commit remaining users
	log.Println("Commiting transaction")
	tx.Commit()

}
