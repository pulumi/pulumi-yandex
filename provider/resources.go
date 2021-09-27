// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yandex

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi-yandex/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "yandex"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(yandex.Provider().(*schema.Provider))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "yandex",
		Description: "A Pulumi package for creating and managing yandex cloud resources.",
		Keywords:    []string{"pulumi", "yandex"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-yandex",
		GitHubOrg:   "yandex-cloud",
		Config:      map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"yandex_alb_target_group":                      {Tok: makeResource(mainMod, "AlbTargetGroup")},
			"yandex_alb_backend_group":                     {Tok: makeResource(mainMod, "AlbBackendGroup")},
			"yandex_alb_http_router":                       {Tok: makeResource(mainMod, "AlbHttpRouter")},
			"yandex_alb_virtual_host":                      {Tok: makeResource(mainMod, "AlbVirtualHost")},
			"yandex_alb_load_balancer":                     {Tok: makeResource(mainMod, "AlbLoadBalancer")},
			"yandex_api_gateway":                           {Tok: makeResource(mainMod, "ApiGateway")},
			"yandex_compute_disk":                          {Tok: makeResource(mainMod, "ComputeDisk")},
			"yandex_compute_disk_placement_group":          {Tok: makeResource(mainMod, "ComputeDiskPlacementGroup")},
			"yandex_compute_image":                         {Tok: makeResource(mainMod, "ComputeImage")},
			"yandex_compute_instance":                      {Tok: makeResource(mainMod, "ComputeInstance")},
			"yandex_compute_instance_group":                {Tok: makeResource(mainMod, "ComputeInstanceGroup")},
			"yandex_compute_placement_group":               {Tok: makeResource(mainMod, "ComputePlacementGroup")},
			"yandex_compute_snapshot":                      {Tok: makeResource(mainMod, "ComputeSnapshot")},
			"yandex_container_registry":                    {Tok: makeResource(mainMod, "ContainerRegistry")},
			"yandex_container_registry_iam_binding":        {Tok: makeResource(mainMod, "ContainerRegistryIamBinding")},
			"yandex_container_repository":                  {Tok: makeResource(mainMod, "ContainerRepository")},
			"yandex_container_repository_iam_binding":      {Tok: makeResource(mainMod, "ContainerRepositoryIamBinding")},
			"yandex_dataproc_cluster":                      {Tok: makeResource(mainMod, "DataprocCluster")},
			"yandex_dns_recordset":                         {Tok: makeResource(mainMod, "DnsRecordSet")},
			"yandex_dns_zone":                              {Tok: makeResource(mainMod, "DnsZone")},
			"yandex_function":                              {Tok: makeResource(mainMod, "Function")},
			"yandex_function_iam_binding":                  {Tok: makeResource(mainMod, "FunctionIamBinding")},
			"yandex_function_trigger":                      {Tok: makeResource(mainMod, "FunctionTrigger")},
			"yandex_function_scaling_policy":               {Tok: makeResource(mainMod, "FunctionScalingPolicy")},
			"yandex_iam_service_account":                   {Tok: makeResource(mainMod, "IamServiceAccount")},
			"yandex_iam_service_account_api_key":           {Tok: makeResource(mainMod, "IamServiceAccountApiKey")},
			"yandex_iam_service_account_iam_binding":       {Tok: makeResource(mainMod, "IamServiceAccountIamBinding")},
			"yandex_iam_service_account_iam_member":        {Tok: makeResource(mainMod, "IamServiceAccountIamMember")},
			"yandex_iam_service_account_iam_policy":        {Tok: makeResource(mainMod, "IamServiceAccountIamPolicy")},
			"yandex_iam_service_account_key":               {Tok: makeResource(mainMod, "IamServiceAccountKey")},
			"yandex_iam_service_account_static_access_key": {Tok: makeResource(mainMod, "IamServiceAccountStaticAccessKey")},
			"yandex_iot_core_device":                       {Tok: makeResource(mainMod, "IotCoreDevice")},
			"yandex_iot_core_registry":                     {Tok: makeResource(mainMod, "IotCoreRegistry")},
			"yandex_kms_secret_ciphertext":                 {Tok: makeResource(mainMod, "KmsSecretCiphertext")},
			"yandex_kms_symmetric_key":                     {Tok: makeResource(mainMod, "KmsSymmetricKey")},
			"yandex_kubernetes_cluster":                    {Tok: makeResource(mainMod, "KubernetesCluster")},
			"yandex_kubernetes_node_group":                 {Tok: makeResource(mainMod, "KubernetesNodeGroup")},
			"yandex_lb_network_load_balancer":              {Tok: makeResource(mainMod, "LbNetworkLoadBalancer")},
			"yandex_lb_target_group":                       {Tok: makeResource(mainMod, "LbTargetGroup")},
			"yandex_mdb_clickhouse_cluster":                {Tok: makeResource(mainMod, "MdbClickhouseCluster")},
			"yandex_mdb_kafka_cluster":                     {Tok: makeResource(mainMod, "MdbKafkaCluster")},
			"yandex_mdb_mongodb_cluster":                   {Tok: makeResource(mainMod, "MdbMongodbCluster")},
			"yandex_mdb_mysql_cluster":                     {Tok: makeResource(mainMod, "MdbMysqlCluster")},
			"yandex_mdb_elasticsearch_cluster":             {Tok: makeResource(mainMod, "MdbElasticSearchCluster")},
			//"yandex_mdb_postgresql_cluster": {
			//	Tok: makeResource(mainMod, "MdbPostgresqlCluster"),
			//	Docs: &tfbridge.DocInfo{
			//		Markdown: []byte{}, // there are some broken docs here that we want to exclude for the time being
			//	},
			//},
			"yandex_mdb_redis_cluster":     {Tok: makeResource(mainMod, "MdbRedisCluster")},
			"yandex_mdb_sqlserver_cluster": {Tok: makeResource(mainMod, "MdbSqlServerCluster")},
			"yandex_mdb_greenplum_cluster": {Tok: makeResource(mainMod, "MdbGreenplumCluster")},
			"yandex_message_queue":         {Tok: makeResource(mainMod, "MessageQueue")},
			"yandex_organizationmanager_organization_iam_binding": {
				Tok: makeResource(mainMod, "OrganizationManagerOrganizationIamBinding"),
			},
			"yandex_organizationmanager_organization_iam_member": {
				Tok: makeResource(mainMod, "OrganizationManagerOrganizationIamMember"),
			},
			"yandex_resourcemanager_cloud_iam_binding":  {Tok: makeResource(mainMod, "ResourcemanagerCloudIamBinding")},
			"yandex_resourcemanager_cloud_iam_member":   {Tok: makeResource(mainMod, "ResourcemanagerCloudIamMember")},
			"yandex_resourcemanager_folder_iam_binding": {Tok: makeResource(mainMod, "ResourcemanagerFolderIamBinding")},
			"yandex_resourcemanager_folder_iam_member":  {Tok: makeResource(mainMod, "ResourcemanagerFolderIamMember")},
			"yandex_resourcemanager_folder_iam_policy":  {Tok: makeResource(mainMod, "ResourcemanagerFolderIamPolicy")},
			"yandex_resourcemanager_folder":             {Tok: makeResource(mainMod, "ResourcemanagerFolder")},
			"yandex_storage_bucket":                     {Tok: makeResource(mainMod, "StorageBucket")},
			"yandex_storage_object":                     {Tok: makeResource(mainMod, "StorageObject")},
			"yandex_vpc_address":                        {Tok: makeResource(mainMod, "VpcAddress")},
			"yandex_vpc_network":                        {Tok: makeResource(mainMod, "VpcNetwork")},
			"yandex_vpc_route_table":                    {Tok: makeResource(mainMod, "VpcRouteTable")},
			"yandex_vpc_security_group":                 {Tok: makeResource(mainMod, "VpcSecurityGroup")},
			"yandex_vpc_subnet":                         {Tok: makeResource(mainMod, "VpcSubnet")},
			"yandex_vpc_default_security_group":         {Tok: makeResource(mainMod, "VpcDefaultSecurityGroup")},
			"yandex_vpc_security_group_rule":            {Tok: makeResource(mainMod, "VpcSecurityGroupRule")},
			"yandex_ydb_database_dedicated":             {Tok: makeResource(mainMod, "YdbDatabaseDedicated")},
			"yandex_ydb_database_serverless":            {Tok: makeResource(mainMod, "YdbDatabaseServerless")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"yandex_alb_target_group": {
				Tok: makeDataSource(mainMod, "getAlbTargetGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_target_group.html.markdown",
				},
			},
			"yandex_alb_backend_group": {
				Tok: makeDataSource(mainMod, "getAlbBackendGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_backend_group.html.markdown",
				},
			},
			"yandex_alb_http_router": {
				Tok: makeDataSource(mainMod, "getAlbHttpRouter"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_http_router.html.markdown",
				},
			},
			"yandex_alb_virtual_host": {
				Tok: makeDataSource(mainMod, "getAlbVirtualHost"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_alb_virtual_host.html.markdown",
				},
			},
			"yandex_alb_load_balancer": {Tok: makeDataSource(mainMod, "getAlbLoadBalancer")},
			"yandex_api_gateway":       {Tok: makeDataSource(mainMod, "getApiGateway")},
			"yandex_client_config": {
				Tok: makeDataSource(mainMod, "getClientConfig"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_client_config.html.markdown",
				},
			},
			"yandex_compute_disk": {
				Tok: makeDataSource(mainMod, "getComputeDisk"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_disk.html.markdown",
				},
			},
			"yandex_compute_disk_placement_group": {
				Tok: makeDataSource(mainMod, "getComputeDiskPlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_disk_placement_group.html.markdown",
				},
			},
			"yandex_compute_image": {
				Tok: makeDataSource(mainMod, "getComputeImage"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_image.html.markdown",
				},
			},
			"yandex_compute_instance": {
				Tok: makeDataSource(mainMod, "getComputeInstance"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_instance.html.markdown",
				},
			},
			"yandex_compute_instance_group": {
				Tok: makeDataSource(mainMod, "getComputeInstanceGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_instance_group.html.markdown",
				},
			},
			"yandex_compute_placement_group": {
				Tok: makeDataSource(mainMod, "getComputePlacementGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_placement_group.html.markdown",
				},
			},
			"yandex_compute_snapshot": {
				Tok: makeDataSource(mainMod, "getComputeSnapshot"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_compute_snapshot.html.markdown",
				},
			},
			"yandex_container_registry": {
				Tok: makeDataSource(mainMod, "getContainerRegistry"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_container_registry.html.markdown",
				},
			},
			"yandex_container_repository": {
				Tok: makeDataSource(mainMod, "getContainerRepository"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_container_repository.html.markdown",
				},
			},
			"yandex_dataproc_cluster": {
				Tok: makeDataSource(mainMod, "getDataprocCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_dataproc_cluster.html.markdown",
				},
			},
			"yandex_dns_zone": {
				Tok: makeDataSource(mainMod, "getDnsZone"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_dns_zone.html.markdown",
				},
			},
			"yandex_function": {
				Tok: makeDataSource(mainMod, "getFunction"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_function.html.markdown",
				},
			},
			"yandex_function_trigger": {
				Tok: makeDataSource(mainMod, "getFunctionTrigger"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_function_trigger.html.markdown",
				},
			},
			"yandex_function_scaling_policy": {Tok: makeDataSource(mainMod, "getFunctionScalingPolicy")},
			"yandex_iam_policy": {
				Tok: makeDataSource(mainMod, "getIamPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_policy.html.markdown",
				},
			},
			"yandex_iam_role": {
				Tok: makeDataSource(mainMod, "getIamRole"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_role.html.markdown",
				},
			},
			"yandex_iam_service_account": {
				Tok: makeDataSource(mainMod, "getIamServiceAccount"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_service_account.html.markdown",
				},
			},
			"yandex_iam_user": {
				Tok: makeDataSource(mainMod, "getIamUser"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iam_user.html.markdown",
				},
			},
			"yandex_iot_core_device": {
				Tok: makeDataSource(mainMod, "getIotCoreDevice"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iot_core_device.html.markdown",
				},
			},
			"yandex_iot_core_registry": {
				Tok: makeDataSource(mainMod, "getIotCoreRegistry"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_iot_core_registry.html.markdown",
				},
			},
			"yandex_kubernetes_cluster": {
				Tok: makeDataSource(mainMod, "getKubernetesCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_kubernetes_cluster.html.markdown",
				},
			},
			"yandex_kubernetes_node_group": {
				Tok: makeDataSource(mainMod, "getKubernetesNodeGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_kubernetes_node_group.html.markdown",
				},
			},
			"yandex_lb_network_load_balancer": {
				Tok: makeDataSource(mainMod, "getLbNetworkLoadBalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_lb_network_load_balancer.html.markdown",
				},
			},
			"yandex_lb_target_group": {
				Tok: makeDataSource(mainMod, "getLbTargetGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_lb_target_group.html.markdown",
				},
			},
			"yandex_mdb_clickhouse_cluster": {
				Tok: makeDataSource(mainMod, "getMdbClickhouseCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_clickhouse_cluster.html.markdown",
				},
			},
			"yandex_mdb_greenplum_cluster": {Tok: makeDataSource(mainMod, "getMdbGreenplumCluster")},
			"yandex_mdb_kafka_cluster": {
				Tok: makeDataSource(mainMod, "getMdbKafkaCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_kafka_cluster.html.markdown",
				},
			},
			"yandex_mdb_mongodb_cluster": {
				Tok: makeDataSource(mainMod, "getMdbMongodbCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_mongodb_cluster.html.markdown",
				},
			},
			"yandex_mdb_mysql_cluster": {
				Tok: makeDataSource(mainMod, "getMdbMysqlCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_mysql_cluster.html.markdown",
				},
			},
			"yandex_mdb_postgresql_cluster": {
				Tok: makeDataSource(mainMod, "getMdbPostgresqlCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_postgresql_cluster.html.markdown",
				},
			},
			"yandex_mdb_redis_cluster": {
				Tok: makeDataSource(mainMod, "getMdbRedisCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_redis_cluster.html.markdown",
				},
			},
			"yandex_mdb_elasticsearch_cluster": {
				Tok: makeDataSource(mainMod, "getMdbElasticSearchCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_elasticsearch_cluster.html.markdown",
				},
			},
			"yandex_mdb_sqlserver_cluster": {
				Tok: makeDataSource(mainMod, "getMdbSqlserverCluster"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_mdb_sqlserver_cluster.html.markdown",
				},
			},
			"yandex_message_queue": {
				Tok: makeDataSource(mainMod, "getMessageQueue"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_message_queue.html.markdown",
				},
			},

			"yandex_resourcemanager_cloud": {
				Tok: makeDataSource(mainMod, "getResourcemanagerCloud"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_resourcemanager_cloud.html.markdown",
				},
			},
			"yandex_resourcemanager_folder": {
				Tok: makeDataSource(mainMod, "getResourcemanagerFolder"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_resourcemanager_folder.html.markdown",
				},
			},
			"yandex_vpc_address": {
				Tok: makeDataSource(mainMod, "getVpcAddress"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_address.html.markdown",
				},
			},
			"yandex_vpc_network": {
				Tok: makeDataSource(mainMod, "getVpcNetwork"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_network.html.markdown",
				},
			},
			"yandex_vpc_route_table": {
				Tok: makeDataSource(mainMod, "getVpcRouteTable"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_route_table.html.markdown",
				},
			},
			"yandex_vpc_security_group": {
				Tok: makeDataSource(mainMod, "getVpcSecurityGroup"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_security_group.html.markdown",
				},
			},
			"yandex_vpc_subnet": {
				Tok: makeDataSource(mainMod, "getVpcSubnet"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_vpc_subnet.html.markdown",
				},
			},
			"yandex_vpc_security_group_rule": {Tok: makeDataSource(mainMod, "getVpcSecurityGroupRule")},
			"yandex_ydb_database_dedicated": {
				Tok: makeDataSource(mainMod, "getYdbDatabaseDedicated"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_ydb_database_dedicated.html.markdown",
				},
			},
			"yandex_ydb_database_serverless": {
				Tok: makeDataSource(mainMod, "getYdbDatabaseServerless"),
				Docs: &tfbridge.DocInfo{
					Source: "datasource_ydb_database_serverless.html.markdown",
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
