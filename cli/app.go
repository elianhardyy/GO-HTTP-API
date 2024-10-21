package cli

import (
	"fmt"
	"log"
	"os"
	"server/config"
	"server/migration"
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
					migration.Migration(db)
					fmt.Println("success migrate")
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
					fmt.Println("success seed")
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