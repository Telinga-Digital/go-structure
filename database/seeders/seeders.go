package seeders

import (
	"fmt"

	"github.com/Telinga-Digital/go-structure/app/models"
)

func init() {
	models.MakeConnection()
}

func Execute() {
	fmt.Println("Seeding...")

	// Put All Seeder Method Below...
	UserSeeder()

	fmt.Println("Seeding complete.")
}
