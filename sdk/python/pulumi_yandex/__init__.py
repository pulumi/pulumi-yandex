# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alb_backend_group import *
from .alb_http_router import *
from .alb_load_balancer import *
from .alb_target_group import *
from .alb_virtual_host import *
from .api_gateway import *
from .cdn_origin_group import *
from .cdn_resource import *
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
from .get_cdn_origin_group import *
from .get_cdn_resource import *
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
from .get_logging_group import *
from .get_mdb_clickhouse_cluster import *
from .get_mdb_elastic_search_cluster import *
from .get_mdb_greenplum_cluster import *
from .get_mdb_kafka_cluster import *
from .get_mdb_kafka_topic import *
from .get_mdb_mongodb_cluster import *
from .get_mdb_mysql_cluster import *
from .get_mdb_postgresql_cluster import *
from .get_mdb_redis_cluster import *
from .get_mdb_sqlserver_cluster import *
from .get_message_queue import *
from .get_resourcemanager_cloud import *
from .get_resourcemanager_folder import *
from .get_serverless_container import *
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
from .kms_symmetric_key_iam_binding import *
from .kubernetes_cluster import *
from .kubernetes_node_group import *
from .lb_network_load_balancer import *
from .lb_target_group import *
from .logging_group import *
from .mdb_clickhouse_cluster import *
from .mdb_elastic_search_cluster import *
from .mdb_greenplum_cluster import *
from .mdb_kafka_cluster import *
from .mdb_kafka_topic import *
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
from .serverless_container import *
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
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_yandex.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_yandex.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "yandex",
  "mod": "index/albBackendGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albBackendGroup:AlbBackendGroup": "AlbBackendGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albHttpRouter",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albHttpRouter:AlbHttpRouter": "AlbHttpRouter"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albLoadBalancer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albLoadBalancer:AlbLoadBalancer": "AlbLoadBalancer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albTargetGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albTargetGroup:AlbTargetGroup": "AlbTargetGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albVirtualHost",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albVirtualHost:AlbVirtualHost": "AlbVirtualHost"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/apiGateway",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/apiGateway:ApiGateway": "ApiGateway"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cdnOriginGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cdnOriginGroup:CdnOriginGroup": "CdnOriginGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cdnResource",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cdnResource:CdnResource": "CdnResource"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeDisk",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeDisk:ComputeDisk": "ComputeDisk"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeDiskPlacementGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup": "ComputeDiskPlacementGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeImage",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeImage:ComputeImage": "ComputeImage"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeInstance",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeInstance:ComputeInstance": "ComputeInstance"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeInstanceGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeInstanceGroup:ComputeInstanceGroup": "ComputeInstanceGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computePlacementGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computePlacementGroup:ComputePlacementGroup": "ComputePlacementGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeSnapshot",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeSnapshot:ComputeSnapshot": "ComputeSnapshot"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRegistry",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRegistry:ContainerRegistry": "ContainerRegistry"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRegistryIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding": "ContainerRegistryIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRepository",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRepository:ContainerRepository": "ContainerRepository"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRepositoryIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding": "ContainerRepositoryIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dataprocCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dataprocCluster:DataprocCluster": "DataprocCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dnsRecordSet",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dnsRecordSet:DnsRecordSet": "DnsRecordSet"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dnsZone",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dnsZone:DnsZone": "DnsZone"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/function",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/function:Function": "Function"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionIamBinding:FunctionIamBinding": "FunctionIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionScalingPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionScalingPolicy:FunctionScalingPolicy": "FunctionScalingPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionTrigger",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionTrigger:FunctionTrigger": "FunctionTrigger"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccount",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccount:IamServiceAccount": "IamServiceAccount"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountApiKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey": "IamServiceAccountApiKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamBinding:IamServiceAccountIamBinding": "IamServiceAccountIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember": "IamServiceAccountIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy": "IamServiceAccountIamPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountKey:IamServiceAccountKey": "IamServiceAccountKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountStaticAccessKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountStaticAccessKey:IamServiceAccountStaticAccessKey": "IamServiceAccountStaticAccessKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iotCoreDevice",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iotCoreDevice:IotCoreDevice": "IotCoreDevice"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iotCoreRegistry",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iotCoreRegistry:IotCoreRegistry": "IotCoreRegistry"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSecretCiphertext",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSecretCiphertext:KmsSecretCiphertext": "KmsSecretCiphertext"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSymmetricKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSymmetricKey:KmsSymmetricKey": "KmsSymmetricKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSymmetricKeyIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSymmetricKeyIamBinding:KmsSymmetricKeyIamBinding": "KmsSymmetricKeyIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kubernetesCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kubernetesCluster:KubernetesCluster": "KubernetesCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kubernetesNodeGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kubernetesNodeGroup:KubernetesNodeGroup": "KubernetesNodeGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lbNetworkLoadBalancer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer": "LbNetworkLoadBalancer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lbTargetGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lbTargetGroup:LbTargetGroup": "LbTargetGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/loggingGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/loggingGroup:LoggingGroup": "LoggingGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbClickhouseCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbClickhouseCluster:MdbClickhouseCluster": "MdbClickhouseCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbElasticSearchCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbElasticSearchCluster:MdbElasticSearchCluster": "MdbElasticSearchCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbGreenplumCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbGreenplumCluster:MdbGreenplumCluster": "MdbGreenplumCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaCluster:MdbKafkaCluster": "MdbKafkaCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaTopic",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaTopic:MdbKafkaTopic": "MdbKafkaTopic"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMongodbCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMongodbCluster:MdbMongodbCluster": "MdbMongodbCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMysqlCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMysqlCluster:MdbMysqlCluster": "MdbMysqlCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbRedisCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbRedisCluster:MdbRedisCluster": "MdbRedisCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbSqlServerCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbSqlServerCluster:MdbSqlServerCluster": "MdbSqlServerCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/messageQueue",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/messageQueue:MessageQueue": "MessageQueue"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationManagerOrganizationIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationManagerOrganizationIamBinding:OrganizationManagerOrganizationIamBinding": "OrganizationManagerOrganizationIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationManagerOrganizationIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationManagerOrganizationIamMember:OrganizationManagerOrganizationIamMember": "OrganizationManagerOrganizationIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerCloudIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding": "ResourcemanagerCloudIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerCloudIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember": "ResourcemanagerCloudIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolder",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolder:ResourcemanagerFolder": "ResourcemanagerFolder"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding": "ResourcemanagerFolderIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember": "ResourcemanagerFolderIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy": "ResourcemanagerFolderIamPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessContainer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessContainer:ServerlessContainer": "ServerlessContainer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/storageBucket",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/storageBucket:StorageBucket": "StorageBucket"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/storageObject",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/storageObject:StorageObject": "StorageObject"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcAddress",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcAddress:VpcAddress": "VpcAddress"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcDefaultSecurityGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcDefaultSecurityGroup:VpcDefaultSecurityGroup": "VpcDefaultSecurityGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcNetwork",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcNetwork:VpcNetwork": "VpcNetwork"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcRouteTable",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcRouteTable:VpcRouteTable": "VpcRouteTable"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcSecurityGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcSecurityGroup:VpcSecurityGroup": "VpcSecurityGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcSecurityGroupRule",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcSecurityGroupRule:VpcSecurityGroupRule": "VpcSecurityGroupRule"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcSubnet",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcSubnet:VpcSubnet": "VpcSubnet"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbDatabaseDedicated",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated": "YdbDatabaseDedicated"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "yandex",
  "token": "pulumi:providers:yandex",
  "fqn": "pulumi_yandex",
  "class": "Provider"
 }
]
"""
)
