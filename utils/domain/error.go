package domain

import (
	"fmt"
)

type ErrModel struct {
	Status    int            `json:"status_code"`
	ErrorCode string         `json:"error_code"`
	Message   string         `json:"message"`
	Details   map[string]any `json:"detail_error,omitempty"`
}

func (e ErrModel) Error() string {
	if e.ErrorCode == "" {
		return e.Message
	}

	if len(e.Details) > 0 {
		return fmt.Sprintf("%s : %v", e.ErrorCode, e.Details)
	}

	return fmt.Sprintf("%s : %v", e.ErrorCode, e.Message)
}

func (e ErrModel) AttachDetail(detail map[string]any) ErrModel {
	e.Details = detail
	return e
}
