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
	"github.com/source-academy/stories-backend/internal/database"
)

const (
	defaultMaxMigrateSteps  = 0 // no limit
	defaultMaxRollbackSteps = 1
)

var (
	blueSandwich  = color.With(color.Blue, "≡")
	greenTick     = color.With(color.Green, "✔")
	yellowChevron = color.With(color.Yellow, "❯")

	migrations = (migrate.FileMigrationSource{
		Dir: "migrations",
	})
)

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	// Check for command line arguments
	flag.Parse()
	switch flag.Arg(0) {
	case "drop", "create":
		// We need to connect anonymously in order
		// to drop or create the database.
		conf.Database.DatabaseName = ""
	case "migrate", "rollback", "status":
		// Do nothing
	default:
		logrus.Errorln("Invalid command")
		return
	}

	// Connect to the database
	d, err := database.Connect(conf.Database)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	defer (func() {
		// Ignore non-critical error
		dbName, _ := getConnectedDBName(d)
		fmt.Println(blueSandwich, "Closing connection with database", dbName+".")
		database.Close(d)
	})()

	dbName, err := getConnectedDBName(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(blueSandwich, "Connected to database", dbName+".")

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
		db, err := d.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		migrateDB(db)
	case "rollback":
		db, err := d.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		rollbackDB(db)
	case "status":
		db, err := d.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		statusDB(db)
	}
}

func migrateDB(db *sql.DB) {
	var steps int
	var err error
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
}
func rollbackDB(db *sql.DB) {
	var steps int
	var err error
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
}
func statusDB(db *sql.DB) {
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
}
