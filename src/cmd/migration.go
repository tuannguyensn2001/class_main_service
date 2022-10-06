package cmd

import (
	"class_main_service/src/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func migrateUp() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-up",
		Run: func(cmd *cobra.Command, args []string) {
			m := getMigrate()

			err := m.Up()
			if err != nil {
				zap.S().Fatalln(err)
			}
		},
	}
}

func getMigrate() *migrate.Migrate {
	cfg, err := config.GetConfig()
	if err != nil {
		zap.S().Fatalln(err)
	}

	db, err := cfg.Db.DB()
	if err != nil {
		zap.S().Fatalln(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	dir, _ := os.Getwd()
	path := "file://" + dir + "/src/database/migrations"
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		zap.S().Fatalln(err)
	}

	return m
}

func migrateDown() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {
			m := getMigrate()

			err := m.Steps(-1)
			if err != nil {
				zap.S().Fatalln(err)
			}

		},
	}
}
