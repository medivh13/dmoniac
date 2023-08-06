package main

import (
	"context"
	"database/sql"

	usecases "dmoniac/src/app/usecase"

	"dmoniac/src/infra/config"
	postgres "dmoniac/src/infra/persistence/postgres"

	pgRiwayat "dmoniac/src/infra/persistence/postgres/riwayat"

	"dmoniac/src/interface/rest"

	riwayatUC "dmoniac/src/app/usecase/riwayat"

	"github.com/apex/log"
	_ "github.com/joho/godotenv/autoload"
)

// reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name

	postgresdb, err := postgres.New(conf.SqlDb)

	// gracefully close connection to persistence storage
	defer func(sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			log.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			log.Errorf("sql database %s successfuly closed.", dbName)
		}
	}(postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	riwayatRepository := pgRiwayat.NewRiwayatRepository(postgresdb.Conn)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		usecases.AllUseCases{

			RiwayatUC: riwayatUC.NewRiwayatUseCase(riwayatRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
