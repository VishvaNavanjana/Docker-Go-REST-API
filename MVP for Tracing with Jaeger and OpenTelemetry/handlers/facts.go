package handlers

import (
	"context"
	"time"

	"github.com/VishvaNavanjana/Docker-Go-REST-API/database"
	"github.com/VishvaNavanjana/Docker-Go-REST-API/models"
	"github.com/gofiber/fiber/v2"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
    "go.opentelemetry.io/otel/trace"
)


const name = "Facts"


// demo functions to increase span count
func demoService(ctx context.Context) error {

    var span trace.Span

    ctx, span = otel.Tracer(name).Start(ctx, "demo service")

    //call demo service two
    demoServiceTwo(ctx)

    //add a delay
    time.Sleep(200 * time.Millisecond)

    //Set error code as true for testing
    span.SetStatus(codes.Error, "test error in demo service")

    // End the span
    defer span.End()

    return nil
}


func demoServiceTwo(ctx context.Context) error {

    var span trace.Span

    ctx, span = otel.Tracer(name).Start(ctx, "demo service two")

    //add a delay
    time.Sleep(150 * time.Millisecond)

    // End the span
    defer span.End()

    return nil
}



func ListFacts(c *fiber.Ctx) error {
    // Create a span for the operation
    newCtx, span := otel.Tracer(name).Start(c.Context(), "ListFacts")

    //call demo service
    demoService(newCtx)

    //set error code as true for testing
    span.SetStatus(codes.Error, "test error in list facts")

    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    //add a delay to simulate a slow operation
    time.Sleep(800 * time.Millisecond)

    // End the span
    defer span.End()

    return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
    // Create a span for the operation
    newCtx, span := otel.Tracer(name).Start(c.Context(), "CreateFacts")

    // call demo service
    demoServiceTwo(newCtx) 

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