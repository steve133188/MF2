package Routes

import (
	"mf-analysis-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BoardCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllAnalysisRecords)
	route.Get("/:id", Services.GetAnalysisRecordById)
	route.Post("/", Services.AddAnalysis)
	route.Put("/:id", Services.UpdateAnalysisRecordByID)
	route.Delete("/:id", Services.DeleteAnalysisById)
}
