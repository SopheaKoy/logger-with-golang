package schemas

import "github.com/google/uuid"

type IResponseBase struct {
	LogID    uuid.UUID
	Success  int
	Message  string
	Code     string
	Data     interface {}
}