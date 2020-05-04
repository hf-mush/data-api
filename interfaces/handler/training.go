package handler

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/interfaces/response"
	"github.com/shuufujita/data-api/usecases"
)

// TrainingHandler training handler
type TrainingHandler interface {
	RetrieveLogs(c echo.Context) error
	CreateLog(c echo.Context) error
	UpdateLog(c echo.Context) error
	DeleteLog(c echo.Context) error
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

// UpdateLogRequest request struct
type UpdateLogRequest struct {
	ID    int64  `json:"id"`
	Kind  string `json:"kind"`
	Count int    `json:"count"`
	Date  string `json:"date"`
}

func (th trainingHandler) UpdateLog(c echo.Context) error {
	request := &UpdateLogRequest{}
	if err := c.Bind(request); err != nil {
		return response.ErrorResponse(c, "INVALID_PARAMETER", err.Error())
	}

	kind := request.Kind
	trainingKind, err := th.trainingUseCase.GetKindByKindTag(kind)
	if err != nil {
		return response.ErrorResponse(c, "DB_NOT_FOUND", err.Error())
	}

	id := request.ID
	date := request.Date
	count := request.Count
	err = th.trainingUseCase.UpdateLog(id, trainingKind.TrainingKindID, date, count)
	if err != nil {
		return response.ErrorResponse(c, "DB_REQUEST_ERROR", err.Error())
	}

	emptyJSON, _ := json.Marshal(map[string]interface{}{})
	return c.JSONBlob(200, emptyJSON)
}

// DeleteLogRequest request struct
type DeleteLogRequest struct {
	ID int64 `query:"id"`
}

func (th trainingHandler) DeleteLog(c echo.Context) error {
	request := &DeleteLogRequest{}
	if err := c.Bind(request); err != nil {
		return response.ErrorResponse(c, "INVALID_PARAMETER", err.Error())
	}

	if request.ID == 0 {
		return response.ErrorResponse(c, "INVALID_PARAMETER", "id must be greater than or equal to 0")
	}

	err := th.trainingUseCase.DeleteLog(request.ID)
	if err != nil {
		return response.ErrorResponse(c, "DB_REQUEST_ERROR", err.Error())
	}

	emptyJSON, _ := json.Marshal(map[string]interface{}{})
	return c.JSONBlob(200, emptyJSON)
}
