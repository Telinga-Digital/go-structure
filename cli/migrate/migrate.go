package migrate

import (
	"database/sql"
	"fmt"

	"github.com/Telinga-Digital/go-structure/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = cli.Command{
		Name:   "migrate",
		Usage:  "migrate database",
		Action: Execute,
	}
	driver = config.GetDBConfig().Connection
	host   = config.GetDBConfig().Host
	port   = config.GetDBConfig().Port
	user   = config.GetDBConfig().Username
	pass   = config.GetDBConfig().Password
	data   = config.GetDBConfig().Database
)

func Execute(c *cli.Context) error {
	arg := c.Args().Get(0)

	db, err := getSql()

	if err != nil {
		panic(err)
	}

	instance, err := getDriver(db)

	if err != nil {
		panic(err)
	}

	switch arg {
	case "up":
		up(instance)
		fmt.Println("Migration success!")
	case "down":
		down(instance)
		fmt.Println("Migration success!")
	case "refresh":
		refresh(instance)
		fmt.Println("Migration success!")
	default:
		fmt.Println("migrate up|down|refresh")
	}

	return nil
}

func getSql() (*sql.DB, error) {
	switch driver {
	case "mysql":
		return sql.Open("mysql", fmt.Sprintf("%s%s@tcp(%s:%s)/%s", user, ":"+pass, host, port, data))
	case "postgres":
		return sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, data))
	default:
		return nil, fmt.Errorf("driver not supported")
	}
}

func getDriver(db *sql.DB) (database.Driver, error) {
	switch driver {
	case "mysql":
		return mysql.WithInstance(db, &mysql.Config{})
	case "postgres":
		return postgres.WithInstance(db, &postgres.Config{})
	default:
		return nil, fmt.Errorf("driver not supported")
	}
}

func up(instance database.Driver) {
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		driver,
		instance,
	)

	if err != nil {
		panic(err)
	}

	m.Up()
}

func down(instance database.Driver) {
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		driver,
		instance,
	)

	if err != nil {
		panic(err)
	}

	m.Down()
}

func refresh(instance database.Driver) {
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		driver,
		instance,
	)

	if err != nil {
		panic(err)
	}

	m.Down()
	m.Up()
}
