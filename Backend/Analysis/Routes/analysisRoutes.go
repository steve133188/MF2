package Routes

import (
	"mf-analysis-services/Services"

	"github.com/gofiber/fiber/v2"
)

func BroadCastRoute(route fiber.Router) {
	route.Get("/", Services.GetAllAnalysisRecords)
	route.Get("/id/:id", Services.GetAnalysisRecordById)
	route.Post("/", Services.AddAnalysis)
	route.Put("/id/:id", Services.UpdateAnalysisRecordByID)
	route.Delete("/id/:id", Services.DeleteAnalysisById)
}
