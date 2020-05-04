package handler

import (
	"github.com/labstack/echo"
	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/interfaces/response"
	"github.com/shuufujita/data-api/usecases"
)

// TrainingHandler training handler
type TrainingHandler interface {
	RetrieveLogs(c echo.Context) error
	CreateLog(c echo.Context) error
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

// RetrieveLogsRequest request struct
type RetrieveLogsRequest struct {
	Kind string `query:"kind"`
}

// RetrieveLogsResponse response struct
type RetrieveLogsResponse struct {
	Records []*model.Training `json:"records"`
}

// RetrieveTrainingList return training list
func (th trainingHandler) RetrieveLogs(c echo.Context) error {

	request := &RetrieveLogsRequest{}
	if err := c.Bind(request); err != nil {
		return response.ErrorResponse(c, "INVALID_PARAMETER", err.Error())
	}

	if request.Kind == "" {
		trainings, err := th.trainingUseCase.GetLogAll()
		if err != nil {
			return response.ErrorResponse(c, "NOT_FOUND", err.Error())
		}
		return c.JSON(200, &RetrieveLogsResponse{Records: trainings})
	}

	trainings, err := th.trainingUseCase.GetLogByKind(request.Kind)
	if err != nil {
		return response.ErrorResponse(c, "NOT_FOUND", err.Error())
	}
	return c.JSON(200, &RetrieveLogsResponse{Records: trainings})
}

// CreateLogRequest request struct
type CreateLogRequest struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
	Date  string `json:"date"`
}

// RetrieveTrainingList return training list
func (th trainingHandler) CreateLog(c echo.Context) error {

	request := &CreateLogRequest{}
	if err := c.Bind(request); err != nil {
		return response.ErrorResponse(c, "INVALID_PARAMETER", err.Error())
	}

	kind := request.Kind
	trainingKind, err := th.trainingUseCase.GetKindByKindTag(kind)
	if err != nil {
		return response.ErrorResponse(c, "DB_NOT_FOUND", err.Error())
	}

	date := request.Date
	count := request.Count
	err = th.trainingUseCase.CreateLog(trainingKind.TrainingKindID, date, count)
	if err != nil {
		return response.ErrorResponse(c, "DB_REQUEST_ERROR", err.Error())
	}

	return c.NoContent(201)
}
