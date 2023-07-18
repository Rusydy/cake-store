package database

import (
	"log"

	"github.com/cake-store/internal/cake"
)

func SeedCakes(repo cake.CakeRepository) {
	// Check if the database already contains cakes
	existingCakes, err := repo.GetAll()
	if err != nil {
		log.Printf("Failed to retrieve cakes from the database: %v", err)
		return
	}

	// If cakes already exist, do not perform the seed operation
	if len(existingCakes) > 0 {
		log.Println("Cake seeder skipped as the database already contains cakes")
		return
	}

	cakes := []*cake.CreateCakeRequest{
		{
			Title:       "Chocolate Cake",
			Description: "Delicious chocolate cake",
			Rating:      5,
			Image:       "chocolate.jpg",
		},
		{
			Title:       "Strawberry Cake",
			Description: "Fresh strawberry cake",
			Rating:      4,
			Image:       "strawberry.jpg",
		},
		// Add more cake data as needed
	}

	for _, c := range cakes {
		_, err := repo.Create(c)
		if err != nil {
			log.Printf("Failed to seed cake: %v", err)
		}
	}

	log.Println("Cake seeder completed")
}
