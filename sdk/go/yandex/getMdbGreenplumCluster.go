// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMdbGreenplumCluster(ctx *pulumi.Context, args *LookupMdbGreenplumClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbGreenplumClusterResult, error) {
	var rv LookupMdbGreenplumClusterResult
	err := ctx.Invoke("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbGreenplumCluster.
type LookupMdbGreenplumClusterArgs struct {
	ClusterId *string `pulumi:"clusterId"`
	FolderId  *string `pulumi:"folderId"`
	Name      *string `pulumi:"name"`
}

// A collection of values returned by getMdbGreenplumCluster.
type LookupMdbGreenplumClusterResult struct {
	Accesses           []GetMdbGreenplumClusterAccess            `pulumi:"accesses"`
	AssignPublicIp     bool                                      `pulumi:"assignPublicIp"`
	BackupWindowStarts []GetMdbGreenplumClusterBackupWindowStart `pulumi:"backupWindowStarts"`
	ClusterId          string                                    `pulumi:"clusterId"`
	CreatedAt          string                                    `pulumi:"createdAt"`
	DeletionProtection bool                                      `pulumi:"deletionProtection"`
	Description        string                                    `pulumi:"description"`
	Environment        string                                    `pulumi:"environment"`
	FolderId           string                                    `pulumi:"folderId"`
	Health             string                                    `pulumi:"health"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                    `pulumi:"id"`
	Labels             map[string]string                         `pulumi:"labels"`
	MasterHostCount    int                                       `pulumi:"masterHostCount"`
	MasterHosts        []GetMdbGreenplumClusterMasterHost        `pulumi:"masterHosts"`
	MasterSubclusters  []GetMdbGreenplumClusterMasterSubcluster  `pulumi:"masterSubclusters"`
	Name               string                                    `pulumi:"name"`
	NetworkId          string                                    `pulumi:"networkId"`
	SecurityGroupIds   []string                                  `pulumi:"securityGroupIds"`
	SegmentHostCount   int                                       `pulumi:"segmentHostCount"`
	SegmentHosts       []GetMdbGreenplumClusterSegmentHost       `pulumi:"segmentHosts"`
	SegmentInHost      int                                       `pulumi:"segmentInHost"`
	SegmentSubclusters []GetMdbGreenplumClusterSegmentSubcluster `pulumi:"segmentSubclusters"`
	Status             string                                    `pulumi:"status"`
	SubnetId           string                                    `pulumi:"subnetId"`
	UserName           string                                    `pulumi:"userName"`
	Version            string                                    `pulumi:"version"`
	Zone               string                                    `pulumi:"zone"`
}

func LookupMdbGreenplumClusterOutput(ctx *pulumi.Context, args LookupMdbGreenplumClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbGreenplumClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMdbGreenplumClusterResult, error) {
			args := v.(LookupMdbGreenplumClusterArgs)
			r, err := LookupMdbGreenplumCluster(ctx, &args, opts...)
			return *r, err
		}).(LookupMdbGreenplumClusterResultOutput)
}

// A collection of arguments for invoking getMdbGreenplumCluster.
type LookupMdbGreenplumClusterOutputArgs struct {
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	FolderId  pulumi.StringPtrInput `pulumi:"folderId"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupMdbGreenplumClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbGreenplumClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbGreenplumCluster.
type LookupMdbGreenplumClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbGreenplumClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbGreenplumClusterResult)(nil)).Elem()
}

func (o LookupMdbGreenplumClusterResultOutput) ToLookupMdbGreenplumClusterResultOutput() LookupMdbGreenplumClusterResultOutput {
	return o
}

func (o LookupMdbGreenplumClusterResultOutput) ToLookupMdbGreenplumClusterResultOutputWithContext(ctx context.Context) LookupMdbGreenplumClusterResultOutput {
	return o
}

func (o LookupMdbGreenplumClusterResultOutput) Accesses() GetMdbGreenplumClusterAccessArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterAccess { return v.Accesses }).(GetMdbGreenplumClusterAccessArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) AssignPublicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) bool { return v.AssignPublicIp }).(pulumi.BoolOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) BackupWindowStarts() GetMdbGreenplumClusterBackupWindowStartArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterBackupWindowStart {
		return v.BackupWindowStarts
	}).(GetMdbGreenplumClusterBackupWindowStartArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbGreenplumClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) MasterHostCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) int { return v.MasterHostCount }).(pulumi.IntOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) MasterHosts() GetMdbGreenplumClusterMasterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterMasterHost { return v.MasterHosts }).(GetMdbGreenplumClusterMasterHostArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) MasterSubclusters() GetMdbGreenplumClusterMasterSubclusterArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterMasterSubcluster {
		return v.MasterSubclusters
	}).(GetMdbGreenplumClusterMasterSubclusterArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SegmentHostCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) int { return v.SegmentHostCount }).(pulumi.IntOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SegmentHosts() GetMdbGreenplumClusterSegmentHostArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterSegmentHost { return v.SegmentHosts }).(GetMdbGreenplumClusterSegmentHostArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SegmentInHost() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) int { return v.SegmentInHost }).(pulumi.IntOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SegmentSubclusters() GetMdbGreenplumClusterSegmentSubclusterArrayOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) []GetMdbGreenplumClusterSegmentSubcluster {
		return v.SegmentSubclusters
	}).(GetMdbGreenplumClusterSegmentSubclusterArrayOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupMdbGreenplumClusterResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbGreenplumClusterResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbGreenplumClusterResultOutput{})
}
