package constants

// ### [ Constants to log ] ### //
const (
	TagInit                  = "Init"
	TagCreateCluster         = "CreateCluster"
	TagValidateCreateCluster = "ValidateCreateCluster"
	TagValidateUpdateCluster = "ValidateUpdateCluster"
	TagGetClusterStatus      = "GetClusterStatus"
	TagUpdateCluster         = "UpdateCluster"
	TagGetCluster            = "GetCluster"
	TagDeleteCluster         = "DeleteCluster"
	TagDeleteDeployment      = "DeleteDeployment"
	TagCreateDeployment      = "CreateDeployment"
	TagListDeployments       = "ListDeployments"
	TagPrometheus            = "Prometheus"
	TagListClusters          = "ListClusters"
	TagGetClusterInfo        = "GetClusterInfo"
	TagFetchClusterConfig    = "FetchClusterConfig"
	TagGetTillerStatus       = "GetTillerStatus"
	TagFetchDeploymentStatus = "FetchDeploymentStatus"
	TagStatus                = "Status"
	TagSlack                 = "Slack"
	TagAuth                  = "Auth"
	TagDatabase              = "Database"
	TagKubernetes            = "Kubernetes"
)

// ### [ Constants to Azure cluster default values ] ### //
const (
	AzureDefaultAgentCount        = 1
	AzureDefaultAgentName         = "agentpool1"
	AzureDefaultKubernetesVersion = "1.7.7"
)

// ### [ Constants to Amazon cluster default values ] ### //
const (
	AmazonDefaultNodeImage          = "ami-bdba13c4"
	AmazonDefaultMasterImage        = "ami-bdba13c4"
	AmazonDefaultMasterInstanceType = "m4.xlarge"
	AmazonDefaultNodeMinCount       = 1
	AmazonDefaultNodeMaxCount       = 1
	AmazonDefaultNodeSpotPrice      = "0.2"
)
