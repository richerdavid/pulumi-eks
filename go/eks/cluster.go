// *** WARNING: this file was generated by pulumi-gen-eks. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Cluster is a component that wraps the AWS and Kubernetes resources necessary to run an EKS cluster, its worker nodes, its optional StorageClasses, and an optional deployment of the Kubernetes Dashboard.
type Cluster struct {
	pulumi.CustomResourceState

	// A kubeconfig that can be used to connect to the EKS cluster.
	Kubeconfig pulumi.AnyOutput `pulumi:"kubeconfig"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("eks:index:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("eks:index:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// A kubeconfig that can be used to connect to the EKS cluster.
	Kubeconfig interface{} `pulumi:"kubeconfig"`
}

type ClusterState struct {
	// A kubeconfig that can be used to connect to the EKS cluster.
	Kubeconfig pulumi.Input
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The tags to apply to the cluster security group.
	ClusterSecurityGroupTags map[string]string `pulumi:"clusterSecurityGroupTags"`
	// The tags to apply to the EKS cluster.
	ClusterTags map[string]string `pulumi:"clusterTags"`
	// Indicates whether an IAM OIDC Provider is created for the EKS cluster.
	//
	// The OIDC provider is used in the cluster in combination with k8s Service Account annotations to provide IAM roles at the k8s Pod level.
	//
	// See for more details:
	//  - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html
	//  - https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html
	//  - https://aws.amazon.com/blogs/opensource/introducing-fine-grained-iam-roles-service-accounts/
	//  - https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/aws/eks/#enabling-iam-roles-for-service-accounts
	CreateOidcProvider *bool `pulumi:"createOidcProvider"`
	// The number of worker nodes that should be running in the cluster. Defaults to 2.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// Enable EKS control plane logging. This sends logs to cloudwatch. Possible list of values are: ["api", "audit", "authenticator", "controllerManager", "scheduler"]. By default it is off.
	EnabledClusterLogTypes []string `pulumi:"enabledClusterLogTypes"`
	// Encrypt the root block device of the nodes in the node group.
	EncryptRootBlockDevice *bool `pulumi:"encryptRootBlockDevice"`
	// KMS Key ARN to use with the encryption configuration for the cluster.
	//
	// Only available on Kubernetes 1.13+ clusters created after March 6, 2020.
	// See for more details:
	// - https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-eks-adds-envelope-encryption-for-secrets-with-aws-kms/
	EncryptionConfigKeyArn *string `pulumi:"encryptionConfigKeyArn"`
	// Indicates whether or not the Amazon EKS private API server endpoint is enabled. Default is `false`.
	EndpointPrivateAccess *bool `pulumi:"endpointPrivateAccess"`
	// Indicates whether or not the Amazon EKS public API server endpoint is enabled. Default is `true`.
	EndpointPublicAccess *bool `pulumi:"endpointPublicAccess"`
	// Add support for launching pods in Fargate. Defaults to launching pods in the `default` namespace.  If specified, the default node group is skipped as though `skipDefaultNodeGroup: true` had been passed.
	Fargate interface{} `pulumi:"fargate"`
	// Use the latest recommended EKS Optimized Linux AMI with GPU support for the worker nodes from the AWS Systems Manager Parameter Store.
	//
	// Defaults to false.
	//
	// Note: `gpu` and `nodeAmiId` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html
	Gpu *bool `pulumi:"gpu"`
	// The default IAM InstanceProfile to use on the Worker NodeGroups, if one is not already set in the NodeGroup.
	InstanceProfileName *string `pulumi:"instanceProfileName"`
	// The maximum number of worker nodes running in the cluster. Defaults to 2.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum number of worker nodes running in the cluster. Defaults to 1.
	MinSize *int `pulumi:"minSize"`
	// The cluster's physical resource name.
	//
	// If not specified, the default is to use auto-naming for the cluster's name, resulting in a physical name with the format `${name}-eksCluster-0123abcd`.
	//
	// See for more details: https://www.pulumi.com/docs/intro/concepts/programming-model/#autonaming
	Name *string `pulumi:"name"`
	// The AMI ID to use for the worker nodes.
	//
	// Defaults to the latest recommended EKS Optimized Linux AMI from the AWS Systems Manager Parameter Store.
	//
	// Note: `nodeAmiId` and `gpu` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html.
	NodeAmiId *string `pulumi:"nodeAmiId"`
	// Whether or not to auto-assign the EKS worker nodes public IP addresses. If this toggle is set to true, the EKS workers will be auto-assigned public IPs. If false, they will not be auto-assigned public IPs.
	NodeAssociatePublicIpAddress *bool `pulumi:"nodeAssociatePublicIpAddress"`
	// Public key material for SSH access to worker nodes. See allowed formats at:
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// If not provided, no SSH access is enabled on VMs.
	NodePublicKey *string `pulumi:"nodePublicKey"`
	// The size in GiB of a cluster node's root volume. Defaults to 20.
	NodeRootVolumeSize *int `pulumi:"nodeRootVolumeSize"`
	// The tags to apply to the default `nodeSecurityGroup` created by the cluster.
	//
	// Note: The `nodeSecurityGroupTags` option and the node group option `nodeSecurityGroup` are mutually exclusive.
	NodeSecurityGroupTags map[string]string `pulumi:"nodeSecurityGroupTags"`
	// The subnets to use for worker nodes. Defaults to the value of subnetIds.
	NodeSubnetIds []string `pulumi:"nodeSubnetIds"`
	// Extra code to run on node startup. This code will run after the AWS EKS bootstrapping code and before the node signals its readiness to the managing CloudFormation stack. This code must be a typical user data script: critically it must begin with an interpreter directive (i.e. a `#!`).
	NodeUserData *string `pulumi:"nodeUserData"`
	// The set of private subnets to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// Worker network architecture options:
	//  - Private-only: Only set `privateSubnetIds`.
	//    - Default workers to run in a private subnet. In this setting, Kubernetes cannot create public, internet-facing load balancers for your pods.
	//  - Public-only: Only set `publicSubnetIds`.
	//    - Default workers to run in a public subnet.
	//  - Mixed (recommended): Set both `privateSubnetIds` and `publicSubnetIds`.
	//    - Default all worker nodes to run in private subnets, and use the public subnets for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	//
	// Also consider setting `nodeAssociatePublicIpAddress: true` for fully private workers.
	PrivateSubnetIds []string `pulumi:"privateSubnetIds"`
	// The AWS provider credential options to scope the cluster's kubeconfig authentication when using a non-default credential chain.
	//
	// This is required for certain auth scenarios. For example:
	// - Creating and using a new AWS provider instance, or
	// - Setting the AWS_PROFILE environment variable, or
	// - Using a named profile configured on the AWS provider via:
	//   `pulumi config set aws:profile <profileName>`
	//
	// See for more details:
	// - https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/aws/#Provider
	// - https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/
	// - https://www.pulumi.com/docs/intro/cloud-providers/aws/#configuration
	// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
	ProviderCredentialOpts *KubeconfigOptions `pulumi:"providerCredentialOpts"`
	// The HTTP(S) proxy to use within a proxied environment.
	//
	//  The proxy is used during cluster creation, and OIDC configuration.
	//
	// This is an alternative option to setting the proxy environment variables: HTTP(S)_PROXY and/or http(s)_proxy.
	//
	// This option is required iff the proxy environment variables are not set.
	//
	// Format:      <protocol>://<host>:<port>
	// Auth Format: <protocol>://<username>:<password>@<host>:<port>
	//
	// Ex:
	//   - "http://proxy.example.com:3128"
	//   - "https://proxy.example.com"
	//   - "http://username:password@proxy.example.com:3128"
	Proxy *string `pulumi:"proxy"`
	// Indicates which CIDR blocks can access the Amazon EKS public API server endpoint.
	PublicAccessCidrs []string `pulumi:"publicAccessCidrs"`
	// The set of public subnets to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// Worker network architecture options:
	//  - Private-only: Only set `privateSubnetIds`.
	//    - Default workers to run in a private subnet. In this setting, Kubernetes cannot create public, internet-facing load balancers for your pods.
	//  - Public-only: Only set `publicSubnetIds`.
	//    - Default workers to run in a public subnet.
	//  - Mixed (recommended): Set both `privateSubnetIds` and `publicSubnetIds`.
	//    - Default all worker nodes to run in private subnets, and use the public subnets for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	PublicSubnetIds []string `pulumi:"publicSubnetIds"`
	// Optional mappings from AWS IAM roles to Kubernetes users and groups.
	RoleMappings []RoleMapping `pulumi:"roleMappings"`
	// If this toggle is set to true, the EKS cluster will be created without node group attached. Defaults to false, unless `fargate` input is provided.
	SkipDefaultNodeGroup *bool `pulumi:"skipDefaultNodeGroup"`
	// The set of all subnets, public and private, to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// If the list of subnets includes both public and private subnets, the worker nodes will only be attached to the private subnets, and the public subnets will be used for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.
	//
	// Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value mapping of tags that are automatically applied to all AWS resources directly under management with this cluster, which support tagging.
	Tags map[string]string `pulumi:"tags"`
	// Optional mappings from AWS IAM users to Kubernetes users and groups.
	UserMappings []UserMapping `pulumi:"userMappings"`
	// Desired Kubernetes master / control plane version. If you do not specify a value, the latest available version is used.
	Version *string `pulumi:"version"`
	// The configuration of the Amazon VPC CNI plugin for this instance. Defaults are described in the documentation for the VpcCniOptions type.
	VpcCniOptions *VpcCniOptions `pulumi:"vpcCniOptions"`
	// The VPC in which to create the cluster and its worker nodes. If unset, the cluster will be created in the default VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The tags to apply to the cluster security group.
	ClusterSecurityGroupTags pulumi.StringMapInput
	// The tags to apply to the EKS cluster.
	ClusterTags pulumi.StringMapInput
	// Indicates whether an IAM OIDC Provider is created for the EKS cluster.
	//
	// The OIDC provider is used in the cluster in combination with k8s Service Account annotations to provide IAM roles at the k8s Pod level.
	//
	// See for more details:
	//  - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html
	//  - https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html
	//  - https://aws.amazon.com/blogs/opensource/introducing-fine-grained-iam-roles-service-accounts/
	//  - https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/aws/eks/#enabling-iam-roles-for-service-accounts
	CreateOidcProvider pulumi.BoolPtrInput
	// The number of worker nodes that should be running in the cluster. Defaults to 2.
	DesiredCapacity pulumi.IntPtrInput
	// Enable EKS control plane logging. This sends logs to cloudwatch. Possible list of values are: ["api", "audit", "authenticator", "controllerManager", "scheduler"]. By default it is off.
	EnabledClusterLogTypes pulumi.StringArrayInput
	// Encrypt the root block device of the nodes in the node group.
	EncryptRootBlockDevice pulumi.BoolPtrInput
	// KMS Key ARN to use with the encryption configuration for the cluster.
	//
	// Only available on Kubernetes 1.13+ clusters created after March 6, 2020.
	// See for more details:
	// - https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-eks-adds-envelope-encryption-for-secrets-with-aws-kms/
	EncryptionConfigKeyArn pulumi.StringPtrInput
	// Indicates whether or not the Amazon EKS private API server endpoint is enabled. Default is `false`.
	EndpointPrivateAccess pulumi.BoolPtrInput
	// Indicates whether or not the Amazon EKS public API server endpoint is enabled. Default is `true`.
	EndpointPublicAccess pulumi.BoolPtrInput
	// Add support for launching pods in Fargate. Defaults to launching pods in the `default` namespace.  If specified, the default node group is skipped as though `skipDefaultNodeGroup: true` had been passed.
	Fargate pulumi.Input
	// Use the latest recommended EKS Optimized Linux AMI with GPU support for the worker nodes from the AWS Systems Manager Parameter Store.
	//
	// Defaults to false.
	//
	// Note: `gpu` and `nodeAmiId` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html
	Gpu pulumi.BoolPtrInput
	// The default IAM InstanceProfile to use on the Worker NodeGroups, if one is not already set in the NodeGroup.
	InstanceProfileName pulumi.StringPtrInput
	// The maximum number of worker nodes running in the cluster. Defaults to 2.
	MaxSize pulumi.IntPtrInput
	// The minimum number of worker nodes running in the cluster. Defaults to 1.
	MinSize pulumi.IntPtrInput
	// The cluster's physical resource name.
	//
	// If not specified, the default is to use auto-naming for the cluster's name, resulting in a physical name with the format `${name}-eksCluster-0123abcd`.
	//
	// See for more details: https://www.pulumi.com/docs/intro/concepts/programming-model/#autonaming
	Name pulumi.StringPtrInput
	// The AMI ID to use for the worker nodes.
	//
	// Defaults to the latest recommended EKS Optimized Linux AMI from the AWS Systems Manager Parameter Store.
	//
	// Note: `nodeAmiId` and `gpu` are mutually exclusive.
	//
	// See for more details:
	// - https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html.
	NodeAmiId pulumi.StringPtrInput
	// Whether or not to auto-assign the EKS worker nodes public IP addresses. If this toggle is set to true, the EKS workers will be auto-assigned public IPs. If false, they will not be auto-assigned public IPs.
	NodeAssociatePublicIpAddress pulumi.BoolPtrInput
	// Public key material for SSH access to worker nodes. See allowed formats at:
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// If not provided, no SSH access is enabled on VMs.
	NodePublicKey pulumi.StringPtrInput
	// The size in GiB of a cluster node's root volume. Defaults to 20.
	NodeRootVolumeSize pulumi.IntPtrInput
	// The tags to apply to the default `nodeSecurityGroup` created by the cluster.
	//
	// Note: The `nodeSecurityGroupTags` option and the node group option `nodeSecurityGroup` are mutually exclusive.
	NodeSecurityGroupTags pulumi.StringMapInput
	// The subnets to use for worker nodes. Defaults to the value of subnetIds.
	NodeSubnetIds pulumi.StringArrayInput
	// Extra code to run on node startup. This code will run after the AWS EKS bootstrapping code and before the node signals its readiness to the managing CloudFormation stack. This code must be a typical user data script: critically it must begin with an interpreter directive (i.e. a `#!`).
	NodeUserData pulumi.StringPtrInput
	// The set of private subnets to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// Worker network architecture options:
	//  - Private-only: Only set `privateSubnetIds`.
	//    - Default workers to run in a private subnet. In this setting, Kubernetes cannot create public, internet-facing load balancers for your pods.
	//  - Public-only: Only set `publicSubnetIds`.
	//    - Default workers to run in a public subnet.
	//  - Mixed (recommended): Set both `privateSubnetIds` and `publicSubnetIds`.
	//    - Default all worker nodes to run in private subnets, and use the public subnets for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	//
	// Also consider setting `nodeAssociatePublicIpAddress: true` for fully private workers.
	PrivateSubnetIds pulumi.StringArrayInput
	// The AWS provider credential options to scope the cluster's kubeconfig authentication when using a non-default credential chain.
	//
	// This is required for certain auth scenarios. For example:
	// - Creating and using a new AWS provider instance, or
	// - Setting the AWS_PROFILE environment variable, or
	// - Using a named profile configured on the AWS provider via:
	//   `pulumi config set aws:profile <profileName>`
	//
	// See for more details:
	// - https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/aws/#Provider
	// - https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/
	// - https://www.pulumi.com/docs/intro/cloud-providers/aws/#configuration
	// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
	ProviderCredentialOpts KubeconfigOptionsPtrInput
	// The HTTP(S) proxy to use within a proxied environment.
	//
	//  The proxy is used during cluster creation, and OIDC configuration.
	//
	// This is an alternative option to setting the proxy environment variables: HTTP(S)_PROXY and/or http(s)_proxy.
	//
	// This option is required iff the proxy environment variables are not set.
	//
	// Format:      <protocol>://<host>:<port>
	// Auth Format: <protocol>://<username>:<password>@<host>:<port>
	//
	// Ex:
	//   - "http://proxy.example.com:3128"
	//   - "https://proxy.example.com"
	//   - "http://username:password@proxy.example.com:3128"
	Proxy pulumi.StringPtrInput
	// Indicates which CIDR blocks can access the Amazon EKS public API server endpoint.
	PublicAccessCidrs pulumi.StringArrayInput
	// The set of public subnets to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// Worker network architecture options:
	//  - Private-only: Only set `privateSubnetIds`.
	//    - Default workers to run in a private subnet. In this setting, Kubernetes cannot create public, internet-facing load balancers for your pods.
	//  - Public-only: Only set `publicSubnetIds`.
	//    - Default workers to run in a public subnet.
	//  - Mixed (recommended): Set both `privateSubnetIds` and `publicSubnetIds`.
	//    - Default all worker nodes to run in private subnets, and use the public subnets for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	PublicSubnetIds pulumi.StringArrayInput
	// Optional mappings from AWS IAM roles to Kubernetes users and groups.
	RoleMappings RoleMappingArrayInput
	// If this toggle is set to true, the EKS cluster will be created without node group attached. Defaults to false, unless `fargate` input is provided.
	SkipDefaultNodeGroup pulumi.BoolPtrInput
	// The set of all subnets, public and private, to use for the worker node groups on the EKS cluster. These subnets are automatically tagged by EKS for Kubernetes purposes.
	//
	// If `vpcId` is not set, the cluster will use the AWS account's default VPC subnets.
	//
	// If the list of subnets includes both public and private subnets, the worker nodes will only be attached to the private subnets, and the public subnets will be used for internet-facing load balancers.
	//
	// See for more details: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html.
	//
	// Note: The use of `subnetIds`, along with `publicSubnetIds` and/or `privateSubnetIds` is mutually exclusive. The use of `publicSubnetIds` and `privateSubnetIds` is encouraged.
	SubnetIds pulumi.StringArrayInput
	// Key-value mapping of tags that are automatically applied to all AWS resources directly under management with this cluster, which support tagging.
	Tags pulumi.StringMapInput
	// Optional mappings from AWS IAM users to Kubernetes users and groups.
	UserMappings UserMappingArrayInput
	// Desired Kubernetes master / control plane version. If you do not specify a value, the latest available version is used.
	Version pulumi.StringPtrInput
	// The configuration of the Amazon VPC CNI plugin for this instance. Defaults are described in the documentation for the VpcCniOptions type.
	VpcCniOptions VpcCniOptionsPtrInput
	// The VPC in which to create the cluster and its worker nodes. If unset, the cluster will be created in the default VPC.
	VpcId pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}