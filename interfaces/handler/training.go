package handler

import (
	"github.com/hf-mush/data-api/domain/model"
	"github.com/hf-mush/data-api/interfaces/response"
	"github.com/hf-mush/data-api/usecases"
	"github.com/labstack/echo"
)

// TrainingHandler training handler
type TrainingHandler interface {
	RetrieveList(c echo.Context) error
}

type trainingHandler struct {
	trainingUseCase usecases.TrainingUseCase
}

// NewTrainingHandler training handler entity
func NewTrainingHandler(tu usecases.TrainingUseCase) TrainingHandler {
	return &trainingHandler{
		trainingUseCase: tu,
	}
}

// RetrieveTrainingListRequest request struct
type RetrieveTrainingListRequest struct{}

// RetrieveTrainingListResponse response struct
type RetrieveTrainingListResponse struct {
	Records []*model.Training `json:"records"`
}

// RetrieveTrainingList return training list
func (th trainingHandler) RetrieveList(c echo.Context) error {

	request := &RetrieveTrainingListRequest{}
	if err := c.Bind(request); err != nil {
		return response.ErrorResponse(c, "INVALID_PARAMETER", err.Error())
	}

	trainings, err := th.trainingUseCase.GetByKind("pushup")
	if err != nil {
		return response.ErrorResponse(c, "NOT_FOUND", err.Error())
	}

	response := &RetrieveTrainingListResponse{
		Records: trainings,
	}

	return c.JSON(200, response)
}
