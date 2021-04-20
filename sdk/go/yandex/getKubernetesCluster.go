// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Kubernetes Cluster.
func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("yandex:index/getKubernetesCluster:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterArgs struct {
	// ID of a specific Kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Name of a specific Kubernetes cluster.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// IP range for allocating pod addresses.
	ClusterIpv4Range string `pulumi:"clusterIpv4Range"`
	// The Kubernetes cluster creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// A description of the Kubernetes cluster.
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// Health of the Kubernetes cluster.
	Health string `pulumi:"health"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// cluster KMS provider parameters.
	KmsProvider GetKubernetesClusterKmsProvider `pulumi:"kmsProvider"`
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]string `pulumi:"labels"`
	// Kubernetes master configuration options. The structure is documented below.
	Master GetKubernetesClusterMaster `pulumi:"master"`
	Name   string                     `pulumi:"name"`
	// The ID of the cluster network.
	NetworkId string `pulumi:"networkId"`
	// Network policy provider for the cluster, if present. Possible values: `CALICO`.
	NetworkPolicyProvider string `pulumi:"networkPolicyProvider"`
	// Size of the masks that are assigned to each node in the cluster.
	NodeIpv4CidrMaskSize int `pulumi:"nodeIpv4CidrMaskSize"`
	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId string `pulumi:"nodeServiceAccountId"`
	// Cluster release channel.
	ReleaseChannel string `pulumi:"releaseChannel"`
	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from
	ServiceIpv4Range string `pulumi:"serviceIpv4Range"`
	// Status of the Kubernetes cluster.
	Status string `pulumi:"status"`
}
