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

	dropCmd     = "drop"
	createCmd   = "create"
	migrateCmd  = "migrate"
	rollbackCmd = "rollback"
	statusCmd   = "status"
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

	targetDBName := conf.Database.DatabaseName

	// Check for command line arguments
	flag.Parse()
	switch flag.Arg(0) {
	case dropCmd, createCmd:
		// We need to connect anonymously in order
		// to drop or create the database.
		conf.Database.DatabaseName = ""
	case migrateCmd, rollbackCmd, statusCmd:
		// Do nothing
	default:
		logrus.Errorln("Invalid command")
		return
	}

	// Connect to the database
	dbConn, err := database.Connect(conf.Database)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	// Remember to close the connection
	defer (func() {
		// Ignore non-critical error
		dbName, _ := getDBName(dbConn)
		fmt.Println(blueSandwich, "Closing connection with database", dbName+".")
		database.Close(dbConn)
	})()

	dbName, err := getDBName(dbConn)
	if err != nil {
		panic(err)
	}
	fmt.Println(blueSandwich, "Connected to database", dbName+".")

	switch flag.Arg(0) {
	case dropCmd:
		err := dropDB(dbConn, conf.Database)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Dropped database:", targetDBName)
	case createCmd:
		err := createDB(dbConn, targetDBName)
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		fmt.Println(greenTick, "Created database:", targetDBName)
	case migrateCmd:
		db, err := dbConn.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		migrateDB(db)
	case rollbackCmd:
		db, err := dbConn.DB()
		if err != nil {
			logrus.Errorln(err)
			panic(err)
		}
		rollbackDB(db)
	case statusCmd:
		db, err := dbConn.DB()
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
