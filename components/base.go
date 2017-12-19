package components

import "github.com/banzaicloud/banzai-types/utils"

type BanzaiResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *BanzaiResponse) String() string {
	return utils.Convert2Json(e)
}