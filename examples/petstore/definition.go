package petstore

import (
	"github.com/charlesvdv/intapi/http"
	"github.com/charlesvdv/intapi/schemas"
)

var limitParameter = http.QueryString("limit", schemas.Integer()).
	Description("How many items to return at one time (max 100)")

var errorCatchAllResponse = http.ResponseCatchAll().
	Description("unexpected error").
	Contents(
		http.JsonContent(errorSchema),
	)

var listPetsOperation = http.Get("/pets").
	Description("List all pets").
	Request(
		http.Request().
			QueryStrings(limitParameter),
	).
	Responses(
		http.ResponseOk().
			Description("A paged array of pets").
			Contents(
				http.JsonContent(schemas.Array(petSchema)),
			).Headers(http.Header("x-next", schemas.String())),
		errorCatchAllResponse,
	)

var createPetsOperation = http.Post("/pets").
	Description("Create a pet").
	Responses(
		http.ResponseCreated().
			Description("Null response"),
		errorCatchAllResponse,
	)

var getPetOperation = http.Get("/pets/{petId}",
	http.PathExpression("petId", schemas.String()).
		Description("The id of the pet to retrieve")).
	Description("Info for a specific pet").
	Responses(
		http.ResponseOk().
			Description("Expected response to a valid request").
			Contents(http.JsonContent(petSchema)),
		errorCatchAllResponse,
	)
