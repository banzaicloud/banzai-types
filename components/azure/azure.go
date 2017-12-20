package azure

import (
	"github.com/banzaicloud/banzai-types/utils"
	"github.com/banzaicloud/banzai-types/constants"
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

type CreateClusterAzure struct {
	Node *CreateAzureNode `json:"node,omitempty"`
}

type CreateAzureNode struct {
	ResourceGroup     string `json:"resourceGroup"`
	AgentCount        int    `json:"agentCount"`
	AgentName         string `json:"agentName"`
	KubernetesVersion string `json:"kubernetesVersion"`
}

// Validate validates azure cluster create request
func (azure *CreateClusterAzure) Validate() (bool, string) {

	utils.LogInfo(constants.TagValidateCreateCluster, "Start validate create request (azure)")

	if azure == nil {
		utils.LogInfo(constants.TagValidateCreateCluster, "Azure is <nil>")
		return false, ""
	}

	if azure == nil {
		msg := "Required field 'azure' is empty."
		utils.LogInfo(constants.TagValidateCreateCluster, msg)
		return false, msg
	}

	// ---- [ Node check ] ---- //
	if azure.Node == nil {
		msg := "Required field 'node' is empty."
		utils.LogInfo(constants.TagValidateCreateCluster, msg)
		return false, msg
	}

	if len(azure.Node.ResourceGroup) == 0 {
		msg := "Required field 'resourceGroup' is empty."
		utils.LogInfo(constants.TagValidateCreateCluster, msg)
		return false, msg
	}

	if azure.Node.AgentCount == 0 {
		utils.LogInfo(constants.TagValidateCreateCluster, "Node agentCount set to default value: ", constants.AzureDefaultAgentCount)
		azure.Node.AgentCount = constants.AzureDefaultAgentCount
	}

	if len(azure.Node.AgentName) == 0 {
		utils.LogInfo(constants.TagValidateCreateCluster, "Node agentName set to default value: ", constants.AzureDefaultAgentName)
		azure.Node.AgentName = constants.AzureDefaultAgentName
	}

	if len(azure.Node.KubernetesVersion) == 0 {
		utils.LogInfo(constants.TagValidateCreateCluster, "Node kubernetesVersion set to default value: ", constants.AzureDefaultKubernetesVersion)
		azure.Node.KubernetesVersion = constants.AzureDefaultKubernetesVersion
	}

	return true, ""
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
