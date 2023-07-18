package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cake-store/internal/cake"
	"github.com/cake-store/internal/database"
)

type CLI struct {
	Seed bool
}

func (cli *CLI) Run() {
	if cli.Seed {
		db := database.NewMySQLDB()
		cakeRepo := cake.NewCakeRepository(db)
		database.SeedCakes(cakeRepo)
		log.Println("Seeding completed")
	}
}

func main() {
	seedFlag := flag.Bool("seed", false, "Run the database seeder")
	flag.Parse()

	cli := CLI{
		Seed: *seedFlag,
	}
	cli.Run()
}
