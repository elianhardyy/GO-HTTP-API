package cli

import (
	"log"
	"os"
	"server/config"
	"server/models"
	"server/seeders"
	"server/server"

	"github.com/urfave/cli/v2"
)
func App(){
	// config.DBConnection()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "migrate",
				Usage: "Migrate model to database",
				Action: func(ctx *cli.Context) error {
					db,err := config.DBConnection()
					if err != nil{
						log.Fatal("failed")
					}
					db.AutoMigrate(&models.Role{})
					db.AutoMigrate(&models.User{})
					db.AutoMigrate(&models.Category{})
					db.AutoMigrate(&models.Product{})
					return nil
				},
			},
			{
				Name: "db:seed",
				Usage: "Run the database seeders",
				Action: func(ctx *cli.Context) error {
					db,err := config.DBConnection()
					if err != nil{
						log.Fatal("failed")
					}
					seeders.Seeder(db)
					return nil
				},
			},
			{
				Name: "server",
				Usage: "Run the server",
				Action: func(ctx *cli.Context) error {
					server.Server()
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args);err != nil{
		log.Fatal(err)
	}
}