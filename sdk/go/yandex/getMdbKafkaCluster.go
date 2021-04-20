// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed Kafka cluster. For more information, see
// [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-yandex/sdk/go/yandex"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "test"
// 		foo, err := yandex.LookupMdbKafkaCluster(ctx, &yandex.LookupMdbKafkaClusterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("networkId", foo.NetworkId)
// 		return nil
// 	})
// }
// ```
func LookupMdbKafkaCluster(ctx *pulumi.Context, args *LookupMdbKafkaClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbKafkaClusterResult, error) {
	var rv LookupMdbKafkaClusterResult
	err := ctx.Invoke("yandex:index/getMdbKafkaCluster:getMdbKafkaCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbKafkaCluster.
type LookupMdbKafkaClusterArgs struct {
	// The ID of the Kafka cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Configuration of the Kafka cluster. The structure is documented below.
	Config *GetMdbKafkaClusterConfig `pulumi:"config"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// The name of the Kafka cluster.
	Name      *string  `pulumi:"name"`
	SubnetIds []string `pulumi:"subnetIds"`
	// A topic of the Kafka cluster. The structure is documented below.
	Topics []GetMdbKafkaClusterTopic `pulumi:"topics"`
	// A user of the Kafka cluster. The structure is documented below.
	Users []GetMdbKafkaClusterUser `pulumi:"users"`
}

// A collection of values returned by getMdbKafkaCluster.
type LookupMdbKafkaClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Configuration of the Kafka cluster. The structure is documented below.
	Config *GetMdbKafkaClusterConfig `pulumi:"config"`
	// Creation timestamp of the key.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the Kafka cluster.
	Description string `pulumi:"description"`
	// Deployment environment of the Kafka cluster.
	Environment string `pulumi:"environment"`
	FolderId    string `pulumi:"folderId"`
	// Health of the host.
	Health string `pulumi:"health"`
	// A host of the Kafka cluster. The structure is documented below.
	Hosts []GetMdbKafkaClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the Kafka cluster.
	Labels map[string]string `pulumi:"labels"`
	// The fully qualified domain name of the host.
	Name string `pulumi:"name"`
	// ID of the network, to which the Kafka cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// A list of security groups IDs of the Kafka cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Status of the cluster.
	Status    string   `pulumi:"status"`
	SubnetIds []string `pulumi:"subnetIds"`
	// A topic of the Kafka cluster. The structure is documented below.
	Topics []GetMdbKafkaClusterTopic `pulumi:"topics"`
	// A user of the Kafka cluster. The structure is documented below.
	Users []GetMdbKafkaClusterUser `pulumi:"users"`
}
