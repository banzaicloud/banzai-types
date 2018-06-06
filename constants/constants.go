package constants

import "errors"

// ### [ Constants to log ] ### //
const (
	TagInit                    = "Init"
	TagCreateCluster           = "CreateCluster"
	TagValidateCreateCluster   = "ValidateCreateCluster"
	TagValidateUpdateCluster   = "ValidateUpdateCluster"
	TagGetClusterStatus        = "GetClusterStatus"
	TagUpdateCluster           = "UpdateCluster"
	TagGetCluster              = "GetCluster"
	TagDeleteCluster           = "DeleteCluster"
	TagDeleteDeployment        = "DeleteDeployment"
	TagCreateDeployment        = "CreateDeployment"
	TagListDeployments         = "ListDeployments"
	TagPrometheus              = "Prometheus"
	TagListClusters            = "ListClusters"
	TagGetClusterInfo          = "GetClusterInfo"
	TagFetchClusterConfig      = "FetchClusterConfig"
	TagInstallSecretsToCluster = "InstallSecretsToCluster"
	TagGetTillerStatus         = "GetTillerStatus"
	TagFetchDeploymentStatus   = "FetchDeploymentStatus"
	TagStatus                  = "Status"
	TagSlack                   = "Slack"
	TagAuth                    = "Auth"
	TagDatabase                = "Database"
	TagKubernetes              = "Kubernetes"
	TagFormat                  = "Format"
	TagHelmInstall             = "HelmInstall"
	TagGetClusterProfile       = "GetClusterProfile"
	TagSetClusterProfile       = "SetClusterProfile"
	TagUpdateClusterProfile    = "UpdateClusterProfile"
	TagDeleteClusterProfile    = "DeleteClusterProfile"
)

// ### [ Constants to Azure cluster default values ] ### //
const (
	AzureDefaultAgentName         = "agentpool1"
	AzureDefaultKubernetesVersion = "1.9.2"
)

// ### [ Constants to common cluster default values ] ### //
const (
	DefaultNodeMinCount = 1
	DefaultNodeMaxCount = 2
)

// ### [ Constants to Amazon cluster default values ] ### //
const (
	AmazonDefaultMasterInstanceType = "m4.xlarge"
	AmazonDefaultNodeSpotPrice      = "0.2"
)

// ### [ Constants to Google cluster default values ] ### //
const (
	GoogleDefaultNodePoolName = "default-pool"
)

// ### [ Constants to helm]
const (
	HELM_RETRY_ATTEMPT_CONFIG = "helm.retryAttempt"
	HELM_RETRY_SLEEP_SECONDS  = "helm.retrySleepSeconds"
)

// ### [ Constants to cloud types ] ### //
const (
	Amazon     = "amazon"
	Azure      = "azure"
	Google     = "google"
	Dummy      = "dummy"
	Kubernetes = "kubernetes"
)

// ### [ Constants to table names ] ### //
const (
	TableNameClusters             = "clusters"
	TableNameAmazonProperties     = "amazon_cluster_properties"
	TableNameAmazonNodePools      = "amazon_node_pools"
	TableNameAzureProperties      = "azure_cluster_properties"
	TableNameAzureNodePools       = "azure_node_pools"
	TableNameGoogleProperties     = "google_cluster_properties"
	TableNameGoogleNodePools      = "google_node_pools"
	TableNameDummyProperties      = "dummy_cluster_properties"
	TableNameKubernetesProperties = "kubernetes_cluster_properties"
)

// ### [ Errors ] ### //
var (
	ErrorNotSupportedCloudType          = errors.New("Not supported cloud type")
	ErrorAmazonClusterNameRegexp        = errors.New("Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.")
	ErrorAmazonFieldIsEmpty             = errors.New("Required field 'amazon' is empty.")
	ErrorAmazonMasterFieldIsEmpty       = errors.New("Required field 'master' is empty.")
	ErrorAmazonImageFieldIsEmpty        = errors.New("Required field 'image' is empty ")
	ErrorAmazonNodePoolFieldIsEmpty     = errors.New("At least one 'nodePool' is required.")
	ErrorAmazonInstancetypeFieldIsEmpty = errors.New("Required field 'instanceType' is empty ")
	ErrorNodePoolMinMaxFieldError       = errors.New("'maxCount' must be greater than 'minCount'")
	ErrorMinFieldRequiredError          = errors.New("'minCount' must be set in case 'autoscaling' is set to true")
	ErrorMaxFieldRequiredError          = errors.New("'maxCount' must be set in case 'autoscaling' is set to true")
	ErrorGoogleClusterNameRegexp        = errors.New("Name must start with a lowercase letter followed by up to 40 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.")
	ErrorAzureClusterNameRegexp         = errors.New("Only numbers, lowercase letters and underscores are allowed under name property. In addition, the value cannot end with an underscore, and must also be less than 32 characters long.")
	ErrorAzureClusterNameEmpty          = errors.New("The name should not be empty.")
	ErrorAzureClusterNameTooLong        = errors.New("Cluster name is greater than or equal 32")
	ErrorAzureCLusterStageFailed        = errors.New("cluster stage is 'Failed'")
	ErrorNotDifferentInterfaces         = errors.New("There is no change in data")
	ErrorReconcile                      = errors.New("Error during reconcile")
	ErrorEmptyUpdateRequest             = errors.New("Empty update cluster request")
	ErrorClusterNotReady                = errors.New("Cluster not ready yet")
	ErrorNilCluster                     = errors.New("<nil> cluster")
	ErrorWrongKubernetesVersion         = errors.New("Wrong kubernetes version for master/nodes. The required minimum kubernetes version is 1.8.x ")
	ErrorDifferentKubernetesVersion     = errors.New("Different kubernetes version for master and nodes")
	ErrorLocationEmpty                  = errors.New("Location field is empty")
	ErrorNodeInstanceTypeEmpty          = errors.New("NodeInstanceType field is empty")
	ErrorRequiredLocation               = errors.New("location is required")
	ErrorRequiredSecretId               = errors.New("Secret id is required")
	ErrorCloudInfoK8SNotSupported       = errors.New("Not supported key in case of amazon")
	ErrorNodePoolNotProvided            = errors.New("At least one 'nodepool' is required for creating or updating a cluster")
	ErrorOnlyOneNodeModify              = errors.New("only one node can be modified at a time")
	ErrorNotValidLocation               = errors.New("not valid location")
	ErrorNotValidMasterImage            = errors.New("not valid master image")
	ErrorNotValidNodeImage              = errors.New("not valid node image")
	ErrorNotValidNodeInstanceType       = errors.New("not valid nodeInstanceType")
	ErrorNotValidMasterVersion          = errors.New("not valid master version")
	ErrorNotValidNodeVersion            = errors.New("not valid node version")
	ErrorNotValidKubernetesVersion      = errors.New("not valid kubernetesVersion")
	ErrorResourceGroupRequired          = errors.New("resource group is required")
	ErrorProjectRequired                = errors.New("project is required")
	ErrorNodePoolNotFoundByName         = errors.New("nodepool not found by name")
	ErrorNoInfrastructureRG             = errors.New("no infrastructure resource group found")
	ErrStateStorePathEmpty              = errors.New("statestore path cannot be empty")
)

// ### [ Keywords ] ###
const (
	KeyWordLocation          = "location"
	KeyWordInstanceType      = "instanceType"
	KeyWordKubernetesVersion = "k8sVersion"
	KeyWordImage             = "image"
)

// ### [ Regexps for cluster names ] ### //
const (
	RegexpAWSName = `^[A-z0-9-_]{1,255}$`
	RegexpAKSName = `^[a-z0-9_]{0,31}[a-z0-9]$`
	RegexpGKEName = `^[a-z]$|^[a-z][a-z0-9-]{0,38}[a-z0-9]$`
)

// ### [ Cluster statuses ] ### //
const (
	Creating = "CREATING"
	Running  = "RUNNING"
	Updating = "UPDATING"
	Deleting = "DELETING"
	Error    = "ERROR"

	CreatingMessage = "Cluster is creating"
	RunningMessage  = "Cluster is running"
	UpdatingMessage = "Cluster is updating"
	DeletingMessage = "Cluster is deleting"
)

const (
	// GenericSecret represents generic secret types, without schema
	GenericSecret = "generic"
	// AllSecrets represents generic secret types which selects all secrets
	AllSecrets = ""
	// SshSecretType marks secrets as of type "ssh"
	SshSecretType = "ssh"
)

// DefaultRules key matching for types
var DefaultRules = map[string][]string{
	Amazon: {
		AwsAccessKeyId,
		AwsSecretAccessKey,
	},
	Azure: {
		AzureClientId,
		AzureClientSecret,
		AzureTenantId,
		AzureSubscriptionId,
	},
	Google: {
		Type,
		ProjectId,
		PrivateKeyId,
		PrivateKey,
		ClientEmail,
		ClientId,
		AuthUri,
		TokenUri,
		AuthX509Url,
		ClientX509Url,
	},
	Kubernetes: {
		K8SConfig,
	},
	SshSecretType: {
		User,
		Identifier,
		PublicKeyData,
		PublicKeyFingerprint,
		PrivateKeyData,
	},

	GenericSecret: {},
}

// Amazon keys
const (
	AwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

// Azure keys
const (
	AzureClientId       = "AZURE_CLIENT_ID"
	AzureClientSecret   = "AZURE_CLIENT_SECRET"
	AzureTenantId       = "AZURE_TENANT_ID"
	AzureSubscriptionId = "AZURE_SUBSCRIPTION_ID"
)

// Google keys
const (
	Type          = "type"
	ProjectId     = "project_id"
	PrivateKeyId  = "private_key_id"
	PrivateKey    = "private_key"
	ClientEmail   = "client_email"
	ClientId      = "client_id"
	AuthUri       = "auth_uri"
	TokenUri      = "token_uri"
	AuthX509Url   = "auth_provider_x509_cert_url"
	ClientX509Url = "client_x509_cert_url"
)

// Kubernetes keys
const (
	K8SConfig = "K8Sconfig"
)

// Ssh keys
const (
	User                 = "user"
	Identifier           = "identifier"
	PublicKeyData        = "public_key_data"
	PublicKeyFingerprint = "public_key_fingerprint"
	PrivateKeyData       = "private_key_data"
)

// Internal usage
const (
	TagKubeConfig = "KubeConfig"
)

// ForbiddenTags are not supported in secret creation
var ForbiddenTags = []string{
	TagKubeConfig,
}
