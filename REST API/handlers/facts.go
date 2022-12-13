package handlers

import (
	"time"

	"github.com/VishvaNavanjana/Docker-Go-REST-API/database"
	"github.com/VishvaNavanjana/Docker-Go-REST-API/models"
	"github.com/gofiber/fiber/v2"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)


const name = "Facts"

func ListFacts(c *fiber.Ctx) error {
    // Create a span for the operation
    ctx := c.Context()
    tracer := otel.Tracer(name)
    _, span := tracer.Start(ctx, "ListFacts")

    //aet error code as true for testing
    span.SetStatus(codes.Error, "test error in list facts")

    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    //add a delay to simulate a slow operation
    time.Sleep(2 * time.Second)

    // End the span
    defer span.End()

    return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
    // Create a span for the operation
    ctx := c.Context()
    tracer := otel.Tracer(name)
    _, span := tracer.Start(ctx, "CreateFact")

    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        //record error in the span
        span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)


    // End the span
    defer span.End()

    return c.Status(200).JSON(fact)
}