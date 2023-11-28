package domain

type SuccessModel struct {
	Status   int         `json:"status_code"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Message  string      `json:"message"`
}
