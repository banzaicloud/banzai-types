package components

import (
	"github.com/banzaicloud/banzai-types/utils"
	"github.com/banzaicloud/banzai-types/components/amazon"
	"github.com/banzaicloud/banzai-types/components/azure"
)

type BanzaiResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

type CreateClusterRequest struct {
	Name             string `json:"name" binding:"required"`
	Location         string `json:"location" binding:"required"`
	Cloud            string `json:"cloud" binding:"required"`
	NodeInstanceType string `json:"nodeInstanceType" binding:"required"`
	Properties struct {
		CreateClusterAmazon *amazon.CreateClusterAmazon `json:"amazon,omitempty"`
		CreateClusterAzure  *azure.CreateClusterAzure   `json:"azure,omitempty"`
	} `json:"properties" binding:"required"`
}

func (e *BanzaiResponse) String() string {
	return utils.Convert2Json(e)
}
