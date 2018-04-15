package amazon

import (
	"errors"
	"github.com/banzaicloud/banzai-types/constants"
)

type CreateClusterAmazon struct {
	NodePools map[string]*AmazonNodePool `json:"nodePools,omitempty"`
	Master    *CreateAmazonMaster        `json:"master,omitempty"`
}

type CreateAmazonMaster struct {
	InstanceType string `json:"instanceType"`
	Image        string `json:"image"`
}

type AmazonNodePool struct {
	InstanceType string `json:"instanceType"`
	SpotPrice string 		`json:"spotPrice"`
	MinCount  int  		  `json:"minCount"`
	MaxCount  int 	    `json:"maxCount"`
	Image     string 		`json:"image"`
}

type UpdateClusterAmazon struct {
	NodePools []*UpdateAmazonNodePool `json:"nodePools,omitempty"`
}

type UpdateAmazonNodePool struct {
	Name string		`json:"name"`
	MinCount int 	`json:"minCount"`
	MaxCount int 	`json:"maxCount"`
}

// Validate validates amazon cluster create request
func (amazon *CreateClusterAmazon) Validate() error {
	if amazon == nil {
		return errors.New("Required field 'amazon' is empty.")
	}
	if amazon.Master == nil {
		msg := "Required field 'master' is empty."
		return errors.New(msg)
	}
	if amazon.Master.Image == "" {
		return errors.New("Required field 'image' is empty ")
	}

	if amazon.Master.InstanceType == "" {
		amazon.Master.InstanceType = constants.AmazonDefaultMasterInstanceType
	}

	if len(amazon.NodePools) == 0 {
		msg := "At least one 'nodePool' is required."
		return errors.New(msg)
	}

	for _, amazonNode := range amazon.NodePools {

		// ---- [ Node image check ] ---- //
		if len(amazonNode.InstanceType) == 0 {
			return errors.New("Required field 'instanceType' is empty ")
		}

		// ---- [ Node image check ] ---- //
		if len(amazonNode.Image) == 0 {
			return errors.New("Required field 'image' is empty ")
		}

		// ---- [ Node min count check ] ---- //
		if amazonNode.MinCount == 0 {
			amazonNode.MinCount = constants.AmazonDefaultNodeMinCount
		}

		// ---- [ Node max count check ] ---- //
		if amazonNode.MaxCount == 0 {
			amazonNode.MaxCount = constants.AmazonDefaultNodeMaxCount
		}

		// ---- [ Node min count <= max count check ] ---- //
		if amazonNode.MaxCount < amazonNode.MinCount {
			return errors.New("maxCount must be greater than mintCount")
		}

		// ---- [ Node spot price ] ---- //
		if len(amazonNode.SpotPrice) == 0 {
			amazonNode.SpotPrice = constants.AmazonDefaultNodeSpotPrice
		}
	}

	return nil
}

// ValidateAmazonRequest validates the update request (only amazon part). If any of the fields is missing, the method fills
// with stored data.
// func (r *UpdateClusterRequest) ValidateAmazonRequest(defaultValue ClusterSimple) (bool, string) {
func (a *UpdateClusterAmazon) Validate() error {

	// ---- [ Amazon field check ] ---- //
	if a == nil {
		return errors.New("'amazon' field is empty")
	}

	if len(a.NodePools) == 0 {
		msg := "At least one 'nodePool' is required."
		return errors.New(msg)
	}

	for _, amazonNode := range a.NodePools {
		// ---- [ Node image check ] ---- //
		if len(amazonNode.Name) == 0 {
			return errors.New("Required field 'name' is empty ")
		}

		// ---- [ Node max count > min count check ] ---- //
		if amazonNode.MaxCount < amazonNode.MinCount {
			return errors.New("maxCount must be greater than mintCount")
		}
	}

	return nil
}

type ClusterProfileAmazon struct {
	Master           *AmazonProfileMaster       `json:"master,omitempty"`
	NodePoolProfiles map[string]*AmazonNodePool `json:"nodePoolProfiles,omitempty"`
}

type AmazonProfileMaster struct {
	InstanceType string `json:"instanceType"`
	Image        string `json:"image"`
}
