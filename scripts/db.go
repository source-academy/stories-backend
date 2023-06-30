package main

import (
	"flag"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
)

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	// Connect to the database
	d, err := database.Connect(conf.Database)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	defer database.Close(d)

	db, err := d.DB()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	migrations := (migrate.FileMigrationSource{
		Dir: "migrations",
	})

	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		migrate.ExecMax(db, "postgres", migrations, migrate.Up, 0)
		fmt.Println("  ✔", "Migration complete")
	case "rollback":
		migrate.ExecMax(db, "postgres", migrations, migrate.Down, 1)
		fmt.Println("  ✔", "Rollback complete")
	case "status":
		completed, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Printf("✔ %d migrations previously done:\n", len(completed))
		for _, c := range completed {
			fmt.Println("  ✔", c.Id)
		}

		pending, _, err := migrate.PlanMigration(db, "postgres", migrations, migrate.Up, 0)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}

		fmt.Printf("❯ Pending %d migrations:\n", len(pending))
		for _, p := range pending {
			fmt.Println("  ❯", p.Migration.Id)
		}
	default:
		logrus.Errorln("Invalid command")
	}
}
