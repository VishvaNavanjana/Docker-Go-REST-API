package handlers

import (
	"github.com/VishvaNavanjana/Docker-Go-REST-API/database"
	"github.com/VishvaNavanjana/Docker-Go-REST-API/models"
	"github.com/gofiber/fiber/v2"

	"go.opentelemetry.io/otel"

)




func ListFacts(c *fiber.Ctx) error {
    // Create a span for the operation
    ctx := c.Context()
    tracer := otel.Tracer("ListFacts")
    _, span := tracer.Start(ctx, "ListFacts")

    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    // End the span
    defer span.End()

    return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
    // Create a span for the operation
    ctx := c.Context()
    tracer := otel.Tracer("CreateFact")
    _, span := tracer.Start(ctx, "CreateFact")

    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)


    // End the span
    defer span.End()

    return c.Status(200).JSON(fact)
}