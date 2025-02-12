package server

import (
	"RiskService/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
)

const (
	StateOpen          = "open"
	StateClosed        = "closed"
	StateAccepted      = "accepted"
	StateInvestigating = "investigating"
)

var validRiskStates = []string{StateOpen, StateClosed, StateAccepted, StateInvestigating}

type Risk struct {
	Id          uuid.UUID `json:"id,ommitempty"`
	State       string    `json:"state" validate:"required,oneof=open closed accepted investigating"`
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=1000"`
}

func RiskStateValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	for _, validStatus := range validRiskStates {
		if status == validStatus {
			return true
		}
	}
	return false
}

type RiskResponse struct {
	Data []Risk `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func (s *Server) getRisks(c *gin.Context) {
	risks, err := s.database.GetRisks()
	if err != nil {
		fmt.Println("failed to get risks from db", err)
		c.JSON(http.StatusInternalServerError,
			ErrorResponse{"failed to get risks due to internal server error", err.Error()})
		return
	}

	var risksResponse []Risk
	for _, risk := range risks {
		risksResponse = append(risksResponse, Risk{
			Id:          risk.Id,
			State:       risk.State,
			Title:       risk.Title,
			Description: risk.Description,
		})
	}
	if len(risksResponse) == 0 {
		fmt.Println("no risks present in the db")
		c.JSON(http.StatusNotFound, ErrorResponse{"no risks found", ""})
		return
	}

	c.JSON(http.StatusOK, RiskResponse{
		risksResponse,
	})
}

func (s *Server) addRisk(c *gin.Context) {
	var risk Risk

	err := c.ShouldBindJSON(&risk)
	if err != nil {
		log.Println("risk with invalid state received to post", err)
		c.JSON(http.StatusBadRequest,
			ErrorResponse{"invalid risk received to post", err.Error()})
		return
	}

	err = s.requestValidator.Struct(risk)
	if err != nil {
		log.Println("failed to validate risk requested to add with error", err)
		c.JSON(http.StatusBadRequest,
			ErrorResponse{"invalid risk received to add", err.Error()})
		return
	}

	_, err = s.database.GetRiskById(risk.Id)
	if err == nil {
		log.Println("risk with id already exists", risk.Id)
		c.JSON(http.StatusBadRequest,
			ErrorResponse{"risk with supplied id is already present on the server", ""})
		return
	}

	riskOut, err := s.database.AddRisk(service.Risk{
		State:       risk.State,
		Title:       risk.Title,
		Description: risk.Description,
	})

	if err != nil {
		log.Println("failed to add risk to db", err)
		c.JSON(http.StatusInternalServerError,
			ErrorResponse{"failed to add risk due to a internal server error", err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		RiskResponse{
			Data: []Risk{
				{
					Id:          riskOut.Id,
					State:       riskOut.State,
					Title:       riskOut.Title,
					Description: riskOut.Description,
				},
			},
		})
}

func (s *Server) getRiskById(c *gin.Context) {
	id := c.Param("id")
	riskUUID, err := uuid.Parse(id)
	if err != nil {
		log.Println("received invalid risk id", err)
		c.JSON(http.StatusBadRequest,
			ErrorResponse{"supplied risk id should be a valid uuid", err.Error()})
		return
	}

	risk, err := s.database.GetRiskById(riskUUID)
	if err != nil {
		log.Println("failed to get risk with error", err)
		c.JSON(http.StatusNotFound,
			ErrorResponse{"risk not found", err.Error()})
		return
	}

	c.JSON(http.StatusOK, RiskResponse{
		Data: []Risk{
			{
				Id:          risk.Id,
				State:       risk.State,
				Title:       risk.Title,
				Description: risk.Description,
			},
		},
	})
}
