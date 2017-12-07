package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/nasa9084/salamander/salamander"
)

type options struct {
	Listen        string `short:"l" long:"listen" env:"SALAMANDER_LISTEN" default:":8080" description:"listening address"`
	MySQLAddr     string `long:"mysql-addr" env:"MYSQL_ADDR" default:"localhost:3306" description:"hostname or IP address and port of MySQL"`
	MySQLDB       string `long:"mysql-db" env:"MYSQL_DB" default:"salamander" description:"database name for this application"`
	MySQLUser     string `long:"mysql-user" env:"MYSQL_USER" default:"root" description:"MySQL username"`
	MySQLPassword string `long:"mysql-password" env:"MYSQL_PASSWORD" default:"password" description:"MySQL password"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("%s", err)
		return 1
	}
	mysqlCfg := mysql.Config{
		Net:       "tcp",
		Addr:      opts.MySQLAddr,
		User:      opts.MySQLUser,
		Passwd:    opts.MySQLPassword,
		DBName:    opts.MySQLDB,
		ParseTime: true,
	}
	log.Printf("connect to MySQL DB: %s", mysqlCfg.FormatDSN())
	db, err := sql.Open(`mysql`, mysqlCfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	s := salamander.NewServer(
		salamander.ListenAddr(opts.Listen),
		salamander.Database(db),
	)
	log.Printf("server listening: %s", opts.Listen)
	if err := s.Run(); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
