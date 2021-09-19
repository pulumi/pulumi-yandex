# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .alb_backend_group import *
from .alb_http_router import *
from .alb_load_balancer import *
from .alb_target_group import *
from .alb_virtual_host import *
from .api_gateway import *
from .compute_disk import *
from .compute_disk_placement_group import *
from .compute_image import *
from .compute_instance import *
from .compute_instance_group import *
from .compute_placement_group import *
from .compute_snapshot import *
from .container_registry import *
from .container_registry_iam_binding import *
from .container_repository import *
from .container_repository_iam_binding import *
from .dataproc_cluster import *
from .dns_record_set import *
from .dns_zone import *
from .function import *
from .function_iam_binding import *
from .function_scaling_policy import *
from .function_trigger import *
from .get_alb_backend_group import *
from .get_alb_http_router import *
from .get_alb_load_balancer import *
from .get_alb_target_group import *
from .get_alb_virtual_host import *
from .get_api_gateway import *
from .get_client_config import *
from .get_compute_disk import *
from .get_compute_disk_placement_group import *
from .get_compute_image import *
from .get_compute_instance import *
from .get_compute_instance_group import *
from .get_compute_placement_group import *
from .get_compute_snapshot import *
from .get_container_registry import *
from .get_container_repository import *
from .get_dataproc_cluster import *
from .get_dns_zone import *
from .get_function import *
from .get_function_scaling_policy import *
from .get_function_trigger import *
from .get_iam_policy import *
from .get_iam_role import *
from .get_iam_service_account import *
from .get_iam_user import *
from .get_iot_core_device import *
from .get_iot_core_registry import *
from .get_kubernetes_cluster import *
from .get_kubernetes_node_group import *
from .get_lb_network_load_balancer import *
from .get_lb_target_group import *
from .get_mdb_clickhouse_cluster import *
from .get_mdb_elastic_search_cluster import *
from .get_mdb_greenplum_cluster import *
from .get_mdb_kafka_cluster import *
from .get_mdb_mongodb_cluster import *
from .get_mdb_mysql_cluster import *
from .get_mdb_postgresql_cluster import *
from .get_mdb_redis_cluster import *
from .get_mdb_sqlserver_cluster import *
from .get_message_queue import *
from .get_resourcemanager_cloud import *
from .get_resourcemanager_folder import *
from .get_vpc_address import *
from .get_vpc_network import *
from .get_vpc_route_table import *
from .get_vpc_security_group import *
from .get_vpc_security_group_rule import *
from .get_vpc_subnet import *
from .get_ydb_database_dedicated import *
from .get_ydb_database_serverless import *
from .iam_service_account import *
from .iam_service_account_api_key import *
from .iam_service_account_iam_binding import *
from .iam_service_account_iam_member import *
from .iam_service_account_iam_policy import *
from .iam_service_account_key import *
from .iam_service_account_static_access_key import *
from .iot_core_device import *
from .iot_core_registry import *
from .kms_secret_ciphertext import *
from .kms_symmetric_key import *
from .kubernetes_cluster import *
from .kubernetes_node_group import *
from .lb_network_load_balancer import *
from .lb_target_group import *
from .mdb_clickhouse_cluster import *
from .mdb_elastic_search_cluster import *
from .mdb_greenplum_cluster import *
from .mdb_kafka_cluster import *
from .mdb_mongodb_cluster import *
from .mdb_mysql_cluster import *
from .mdb_redis_cluster import *
from .mdb_sql_server_cluster import *
from .message_queue import *
from .organization_manager_organization_iam_binding import *
from .organization_manager_organization_iam_member import *
from .provider import *
from .resourcemanager_cloud_iam_binding import *
from .resourcemanager_cloud_iam_member import *
from .resourcemanager_folder import *
from .resourcemanager_folder_iam_binding import *
from .resourcemanager_folder_iam_member import *
from .resourcemanager_folder_iam_policy import *
from .storage_bucket import *
from .storage_object import *
from .vpc_address import *
from .vpc_default_security_group import *
from .vpc_network import *
from .vpc_route_table import *
from .vpc_security_group import *
from .vpc_security_group_rule import *
from .vpc_subnet import *
from .ydb_database_dedicated import *
from .ydb_database_serverless import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "yandex:index/albBackendGroup:AlbBackendGroup":
                return AlbBackendGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/albHttpRouter:AlbHttpRouter":
                return AlbHttpRouter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/albLoadBalancer:AlbLoadBalancer":
                return AlbLoadBalancer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/albTargetGroup:AlbTargetGroup":
                return AlbTargetGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/albVirtualHost:AlbVirtualHost":
                return AlbVirtualHost(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/apiGateway:ApiGateway":
                return ApiGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeDisk:ComputeDisk":
                return ComputeDisk(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup":
                return ComputeDiskPlacementGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeImage:ComputeImage":
                return ComputeImage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeInstance:ComputeInstance":
                return ComputeInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeInstanceGroup:ComputeInstanceGroup":
                return ComputeInstanceGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computePlacementGroup:ComputePlacementGroup":
                return ComputePlacementGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/computeSnapshot:ComputeSnapshot":
                return ComputeSnapshot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/containerRegistry:ContainerRegistry":
                return ContainerRegistry(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding":
                return ContainerRegistryIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/containerRepository:ContainerRepository":
                return ContainerRepository(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding":
                return ContainerRepositoryIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/dataprocCluster:DataprocCluster":
                return DataprocCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/dnsRecordSet:DnsRecordSet":
                return DnsRecordSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/dnsZone:DnsZone":
                return DnsZone(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/function:Function":
                return Function(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/functionIamBinding:FunctionIamBinding":
                return FunctionIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/functionScalingPolicy:FunctionScalingPolicy":
                return FunctionScalingPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/functionTrigger:FunctionTrigger":
                return FunctionTrigger(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccount:IamServiceAccount":
                return IamServiceAccount(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey":
                return IamServiceAccountApiKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountIamBinding:IamServiceAccountIamBinding":
                return IamServiceAccountIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember":
                return IamServiceAccountIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy":
                return IamServiceAccountIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountKey:IamServiceAccountKey":
                return IamServiceAccountKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iamServiceAccountStaticAccessKey:IamServiceAccountStaticAccessKey":
                return IamServiceAccountStaticAccessKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iotCoreDevice:IotCoreDevice":
                return IotCoreDevice(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/iotCoreRegistry:IotCoreRegistry":
                return IotCoreRegistry(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/kmsSecretCiphertext:KmsSecretCiphertext":
                return KmsSecretCiphertext(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/kmsSymmetricKey:KmsSymmetricKey":
                return KmsSymmetricKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/kubernetesCluster:KubernetesCluster":
                return KubernetesCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/kubernetesNodeGroup:KubernetesNodeGroup":
                return KubernetesNodeGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer":
                return LbNetworkLoadBalancer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/lbTargetGroup:LbTargetGroup":
                return LbTargetGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbClickhouseCluster:MdbClickhouseCluster":
                return MdbClickhouseCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbElasticSearchCluster:MdbElasticSearchCluster":
                return MdbElasticSearchCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbGreenplumCluster:MdbGreenplumCluster":
                return MdbGreenplumCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbKafkaCluster:MdbKafkaCluster":
                return MdbKafkaCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbMongodbCluster:MdbMongodbCluster":
                return MdbMongodbCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbMysqlCluster:MdbMysqlCluster":
                return MdbMysqlCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbRedisCluster:MdbRedisCluster":
                return MdbRedisCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/mdbSqlServerCluster:MdbSqlServerCluster":
                return MdbSqlServerCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/messageQueue:MessageQueue":
                return MessageQueue(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding":
                return OrganizationManagerOrganizationIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember":
                return OrganizationManagerOrganizationIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding":
                return ResourcemanagerCloudIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember":
                return ResourcemanagerCloudIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerFolder:ResourcemanagerFolder":
                return ResourcemanagerFolder(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding":
                return ResourcemanagerFolderIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember":
                return ResourcemanagerFolderIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy":
                return ResourcemanagerFolderIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/storageBucket:StorageBucket":
                return StorageBucket(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/storageObject:StorageObject":
                return StorageObject(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcAddress:VpcAddress":
                return VpcAddress(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcDefaultSecurityGroup:VpcDefaultSecurityGroup":
                return VpcDefaultSecurityGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcNetwork:VpcNetwork":
                return VpcNetwork(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcRouteTable:VpcRouteTable":
                return VpcRouteTable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcSecurityGroup:VpcSecurityGroup":
                return VpcSecurityGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcSecurityGroupRule:VpcSecurityGroupRule":
                return VpcSecurityGroupRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/vpcSubnet:VpcSubnet":
                return VpcSubnet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated":
                return YdbDatabaseDedicated(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "yandex:index/ydbDatabaseServerless:YdbDatabaseServerless":
                return YdbDatabaseServerless(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("yandex", "index/albBackendGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/albHttpRouter", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/albLoadBalancer", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/albTargetGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/albVirtualHost", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/apiGateway", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeDisk", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeDiskPlacementGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeImage", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeInstance", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeInstanceGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computePlacementGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/computeSnapshot", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/containerRegistry", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/containerRegistryIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/containerRepository", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/containerRepositoryIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/dataprocCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/dnsRecordSet", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/dnsZone", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/function", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/functionIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/functionScalingPolicy", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/functionTrigger", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccount", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountApiKey", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountIamMember", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountKey", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iamServiceAccountStaticAccessKey", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iotCoreDevice", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/iotCoreRegistry", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/kmsSecretCiphertext", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/kmsSymmetricKey", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/kubernetesCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/kubernetesNodeGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/lbNetworkLoadBalancer", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/lbTargetGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbClickhouseCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbElasticSearchCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbGreenplumCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbKafkaCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbMongodbCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbMysqlCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbRedisCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/mdbSqlServerCluster", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/messageQueue", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/organizationManagerOrganizationIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/organizationManagerOrganizationIamMember", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerCloudIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerCloudIamMember", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerFolder", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerFolderIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerFolderIamMember", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/resourcemanagerFolderIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/storageBucket", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/storageObject", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcAddress", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcDefaultSecurityGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcNetwork", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcRouteTable", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcSecurityGroup", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcSecurityGroupRule", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/vpcSubnet", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/ydbDatabaseDedicated", _module_instance)
    pulumi.runtime.register_resource_module("yandex", "index/ydbDatabaseServerless", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:yandex":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("yandex", Package())

_register_module()
