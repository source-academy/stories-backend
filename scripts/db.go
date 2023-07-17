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
	greenTick     = color.With(color.Green, "✔")
	yellowChevron = color.With(color.Yellow, "❯")
)

func createIfNotExistAndConnect(dbserver *gorm.DB, dbconf *config.DatabaseConfig) (*sql.DB, *gorm.DB) {
	// Create if not exist and Connect to the database
	d, err := CreateAndConnect(dbserver, dbconf)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	db, err := d.DB()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	return db, d
}

func main() {
	// Load configuration
	conf, err := config.LoadFromEnvironment()
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	// Connect to the db server
	dbserver, err := ConnectToDBServer(conf.Database)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

	defer Close(dbserver)

	migrations := (migrate.FileMigrationSource{
		Dir: "migrations",
	})

	// TODO: Reduce code duplication
	flag.Parse()
	switch flag.Arg(0) {
	case "drop":
		err := Drop(dbserver, conf.Database)
		if err != nil {
			panic(err)
		}
		fmt.Println(greenTick, "Dropped database: ", conf.Database.DatabaseName)
	case "create":
		_, d := createIfNotExistAndConnect(dbserver, conf.Database)
		defer Close(d)

		fmt.Println(greenTick, "Created database: ", conf.Database.DatabaseName)
	case "migrate":
		db, d := createIfNotExistAndConnect(dbserver, conf.Database)
		defer Close(d)

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
		db, d := createIfNotExistAndConnect(dbserver, conf.Database)
		defer Close(d)

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
		db, d := createIfNotExistAndConnect(dbserver, conf.Database)
		defer Close(d)

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
