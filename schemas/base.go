package schemas

import "github.com/google/uuid"

type IResponseBase struct {
	LogID   uuid.UUID
	Success int
	Message string
	Code    string
	Data    interface{}
}

// NewResponse creates a new IResponseBase with an auto-generated UUID
func NewResponse() IResponseBase {
	return IResponseBase{
		LogID:   uuid.New(),
		Success: 1,
		Message: "",
		Code:    "",
		Data:    nil,
	}
}

// NewErrorResponse creates a new error IResponseBase with an auto-generated UUID
func NewErrorResponse(err error) IResponseBase {
	return IResponseBase{
		LogID:   uuid.New(),
		Success: 0,
		Message: err.Error(),
		Code:    "500",
		Data:    nil,
	}
}
