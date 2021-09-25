package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:response Health
type BodyProductListResponse struct {
	// in: body
	Body HealthResponse
}

type HealthResponse struct {
	// Required: true
	// description: indicates whether the service status is acceptable or not. API publishers
	Status string `json:"status"`
	// description: public version of the service.
	Version string `json:"version,omitempty"`
	// description:  in well-designed APIs, backwards-compatible changes in the service should not update a version number.
	RelaseID string `json:"relaseID,omitempty"`
	// description: array of notes relevant to current state of health
	Notes []string `json:"notes,omitempty"`
	// description: raw error output, in case of “fail” or “warn” states. This field SHOULD be omitted for “pass” state.
	Output string `json:"output,omitempty"`
	// description: an object representing status of sub-components of the service in question
	Details []string `json:"details,omitempty"`
	// description: an array of objects containing link relations and URIs
	Links []string `json:"links,omitempty"`
	// description: unique identifier of the service, in the application scope
	ServiceID string `json:"serviceID,omitempty"`
	// Required: true
	// description: human-friendly description of the service.
	Description string `json:"description"`
}

// swagger:route GET /health infrastructure Health
//
// This endpoint return the health of auth api
//
//	Produces:
//	- application/health+json
//
//	Schemes: http, https
//
//	Responses:
//		201: Health
//		400:
//		500:
func CheckHandler() gin.HandlerFunc {
	health := HealthResponse{
		Status:      "pass",
		Description: "Susbriptions service",
	}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, health)
	}
}
