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
		},
	}
	if len(os.Args) > 1 {
		if err := app.Run(os.Args);err != nil{
			log.Fatal(err)
		}
	}else{
		_,err := config.DBConnection()
		if err != nil{
			log.Fatal("failed")
		}
		server.Server()
	}
}