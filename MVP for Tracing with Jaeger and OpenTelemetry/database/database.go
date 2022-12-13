package database

import (
	"fmt"
	"log"
	"os"

	"github.com/VishvaNavanjana/Docker-Go-REST-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"context"

	"go.opentelemetry.io/otel"

)

//// name is the Tracer name used to identify this instrumentation library.
const name = "database"

type Dbinstance struct {
    Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
    // Create a span for the operation
    ctx := context.Background()
    tracer := otel.Tracer(name)
    _, span := tracer.Start(ctx, "ConnectDb")


    dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }

    log.Println("connected")

    // End the span
    defer span.End()

    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    db.AutoMigrate(&models.Fact{})

    DB = Dbinstance{
        Db: db,
    }
}