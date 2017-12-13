package main

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/nasa9084/salamander/salamander"
	"github.com/nasa9084/salamander/salamander/log"
	"github.com/nasa9084/salamander/salamander/middleware"
	"github.com/pkg/errors"
)

type options struct {
	Listen string `short:"l" long:"listen" env:"SALAMANDER_LISTEN" default:":8080" description:"listening address"`
	MySQL  mysqlOptions
}

type mysqlOptions struct {
	Addr     string `long:"mysql-addr" env:"MYSQL_ADDR" default:"localhost:3306" description:"hostname or IP address and port of MySQL"`
	DB       string `long:"mysql-db" env:"MYSQL_DB" default:"salamander" description:"database name for this application"`
	User     string `long:"mysql-user" env:"MYSQL_USER" default:"root" description:"MySQL username"`
	Password string `long:"mysql-password" env:"MYSQL_PASSWORD" default:"password" description:"MySQL password"`
}

func main() { os.Exit(exec()) }

func exec() int {
	opts, db, err := prepare()
	if err != nil {
		log.Error.Printf("%s", err)
		return 1
	}
	middlewareOption := salamander.Middlewares(middleware.Logger())
	s := salamander.NewServer(
		db,
		salamander.ListenAddr(opts.Listen),
		middlewareOption,
	)
	return run(s)
}

func prepare() (*options, *sql.DB, error) {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		return nil, nil, err
	}
	db, err := openMySQL(opts)
	if err != nil {
		return nil, nil, err
	}
	return &opts, db, nil
}

func run(s salamander.Server) int {
	log.Info.Printf("server listening: %s", s.Listen())
	if err := s.Run(); err != nil {
		log.Error.Printf("%s", err)
		return 1
	}
	return 0
}

func openMySQL(opts options) (*sql.DB, error) {
	mysqlCfg := getMySQLConfig(opts)
	log.Info.Printf("connect to MySQL DB: %s", mysqlCfg.FormatDSN())
	db, err := sql.Open(`mysql`, mysqlCfg.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, `open`)
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, `ping`)
	}
	return db, nil
}

func getMySQLConfig(opts options) mysql.Config {
	return mysql.Config{
		Net:       "tcp",
		Addr:      opts.MySQL.Addr,
		User:      opts.MySQL.User,
		Passwd:    opts.MySQL.Password,
		DBName:    opts.MySQL.DB,
		ParseTime: true,
	}
}
