package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strconv"

	"github.com/TwiN/go-color"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/config"
	"gorm.io/gorm"
)

const (
	defaultMaxMigrateSteps  = 0 // no limit
	defaultMaxRollbackSteps = 1
)

var (
	blueSandwich  = color.With(color.Blue, "≡")
	greenTick     = color.With(color.Green, "✔")
	yellowChevron = color.With(color.Yellow, "❯")
)

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	migrations := (migrate.FileMigrationSource{
		Dir: "migrations",
	})

	// TODO: Reduce code duplication
	flag.Parse()
	// connect to the DB
	var d *gorm.DB
	var db *sql.DB
	switch flag.Arg(0) {
	case "drop", "create":
		d, err = connectAnonDB(*conf.Database)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		defer closeDBConnection(d)
	case "migrate", "rollback", "status":
		d, err := connectDB(*conf.Database)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		defer closeDBConnection(d)

		db, err = d.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
	default:
		logrus.Errorln("Invalid command")
		return
	}

	// handling the command
	switch flag.Arg(0) {
	case "drop":
		err := dropDB(d, conf.Database)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Dropped database:", conf.Database.DatabaseName)
	case "create":
		err := createDB(d, conf.Database)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Created database:", conf.Database.DatabaseName)
	case "migrate":
		var steps int
		if flag.Arg(1) == "" {
			steps = defaultMaxMigrateSteps
		} else if steps, err = strconv.Atoi(flag.Arg(1)); err != nil {
			logrus.Warningln("Invalid number of steps", err)
			logrus.Warningf("Defaulting to %d steps\n", defaultMaxMigrateSteps)
			steps = defaultMaxMigrateSteps
		}
		_, err = migrate.ExecMax(db, "postgres", migrations, migrate.Up, steps)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Migration complete")
	case "rollback":
		var steps int
		if flag.Arg(1) == "" {
			steps = defaultMaxRollbackSteps
		} else if steps, err = strconv.Atoi(flag.Arg(1)); err != nil {
			logrus.Warningln("Invalid number of steps", err)
			logrus.Warningf("Defaulting to %d steps\n", defaultMaxRollbackSteps)
			steps = defaultMaxRollbackSteps
		}
		_, err = migrate.ExecMax(db, "postgres", migrations, migrate.Down, steps)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Rollback complete")
	case "status":
		completed, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Printf(greenTick+" %d migrations previously done:\n", len(completed))
		for _, c := range completed {
			fmt.Println("  "+greenTick, c.Id)
		}

		pending, _, err := migrate.PlanMigration(db, "postgres", migrations, migrate.Up, 0)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}

		fmt.Printf(yellowChevron+" Pending %d migrations:\n", len(pending))
		for _, p := range pending {
			fmt.Println("  "+yellowChevron, p.Migration.Id)
		}
	default:
		logrus.Errorln("Invalid command")
	}
}
