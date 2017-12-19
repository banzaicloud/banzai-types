package azure

import (
	"github.com/banzaicloud/banzai-types/utils"
)

type Values struct {
	Value []Value `json:"value"`
}

type Value struct {
	Id         string     `json:"id"`
	Location   string     `json:"location"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	ProvisioningState string    `json:"provisioningState"`
	AgentPoolProfiles []Profile `json:"agentPoolProfiles"`
}

type Profile struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ResponseWithValue struct {
	StatusCode int   `json:"status_code"`
	Value      Value `json:"message,omitempty"`
}

type ListResponse struct {
	StatusCode int    `json:"status_code"`
	Value      Values `json:"message"`
}

func (r *ResponseWithValue) String() string {
	return utils.Convert2Json(r)
}

func (r *ResponseWithValue) Update(code int, Value Value) {
	r.Value = Value
	r.StatusCode = code
}

func (v *Value) String() string {
	return utils.Convert2Json(v)
}

func (v *Values) String() string {
	return utils.Convert2Json(v)
}

func (l *ListResponse) String() string {
	return utils.Convert2Json(l)
}

func (p *Properties) String() string {
	return utils.Convert2Json(p)
}

func (p *Profile) String() string {
	return utils.Convert2Json(p)
}
