// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./albBackendGroup";
export * from "./albHttpRouter";
export * from "./albTargetGroup";
export * from "./apiGateway";
export * from "./computeDisk";
export * from "./computeDiskPlacementGroup";
export * from "./computeImage";
export * from "./computeInstance";
export * from "./computeInstanceGroup";
export * from "./computePlacementGroup";
export * from "./computeSnapshot";
export * from "./containerRegistry";
export * from "./containerRegistryIamBinding";
export * from "./containerRepository";
export * from "./containerRepositoryIamBinding";
export * from "./dataprocCluster";
export * from "./dnsRecordSet";
export * from "./dnsZone";
export * from "./function";
export * from "./functionIamBinding";
export * from "./functionTrigger";
export * from "./getAlbBackendGroup";
export * from "./getAlbHttpRouter";
export * from "./getAlbTargetGroup";
export * from "./getApiGateway";
export * from "./getClientConfig";
export * from "./getComputeDisk";
export * from "./getComputeDiskPlacementGroup";
export * from "./getComputeImage";
export * from "./getComputeInstance";
export * from "./getComputeInstanceGroup";
export * from "./getComputePlacementGroup";
export * from "./getComputeSnapshot";
export * from "./getContainerRegistry";
export * from "./getContainerRepository";
export * from "./getDataprocCluster";
export * from "./getDnsZone";
export * from "./getFunction";
export * from "./getFunctionTrigger";
export * from "./getIamPolicy";
export * from "./getIamRole";
export * from "./getIamServiceAccount";
export * from "./getIamUser";
export * from "./getIotCoreDevice";
export * from "./getIotCoreRegistry";
export * from "./getKubernetesCluster";
export * from "./getKubernetesNodeGroup";
export * from "./getLbNetworkLoadBalancer";
export * from "./getLbTargetGroup";
export * from "./getMdbClickhouseCluster";
export * from "./getMdbKafkaCluster";
export * from "./getMdbMongodbCluster";
export * from "./getMdbMysqlCluster";
export * from "./getMdbPostgresqlCluster";
export * from "./getMdbRedisCluster";
export * from "./getMdbSqlserverCluster";
export * from "./getMessageQueue";
export * from "./getResourcemanagerCloud";
export * from "./getResourcemanagerFolder";
export * from "./getVpcAddress";
export * from "./getVpcNetwork";
export * from "./getVpcRouteTable";
export * from "./getVpcSecurityGroup";
export * from "./getVpcSubnet";
export * from "./getYdbDatabaseDedicated";
export * from "./getYdbDatabaseServerless";
export * from "./iamServiceAccount";
export * from "./iamServiceAccountApiKey";
export * from "./iamServiceAccountIamBinding";
export * from "./iamServiceAccountIamMember";
export * from "./iamServiceAccountIamPolicy";
export * from "./iamServiceAccountKey";
export * from "./iamServiceAccountStaticAccessKey";
export * from "./iotCoreDevice";
export * from "./iotCoreRegistry";
export * from "./kmsSecretCiphertext";
export * from "./kmsSymmetricKey";
export * from "./kubernetesCluster";
export * from "./kubernetesNodeGroup";
export * from "./lbNetworkLoadBalancer";
export * from "./lbTargetGroup";
export * from "./mdbClickhouseCluster";
export * from "./mdbKafkaCluster";
export * from "./mdbMongodbCluster";
export * from "./mdbMysqlCluster";
export * from "./mdbRedisCluster";
export * from "./mdbSqlServerCluster";
export * from "./messageQueue";
export * from "./provider";
export * from "./resourcemanagerCloudIamBinding";
export * from "./resourcemanagerCloudIamMember";
export * from "./resourcemanagerFolderIamBinding";
export * from "./resourcemanagerFolderIamMember";
export * from "./resourcemanagerFolderIamPolicy";
export * from "./storageBucket";
export * from "./storageObject";
export * from "./vpcAddress";
export * from "./vpcDefaultSecurityGroup";
export * from "./vpcNetwork";
export * from "./vpcRouteTable";
export * from "./vpcSecurityGroup";
export * from "./vpcSubnet";
export * from "./ydbDatabaseDedicated";
export * from "./ydbDatabaseServerless";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { AlbBackendGroup } from "./albBackendGroup";
import { AlbHttpRouter } from "./albHttpRouter";
import { AlbTargetGroup } from "./albTargetGroup";
import { ApiGateway } from "./apiGateway";
import { ComputeDisk } from "./computeDisk";
import { ComputeDiskPlacementGroup } from "./computeDiskPlacementGroup";
import { ComputeImage } from "./computeImage";
import { ComputeInstance } from "./computeInstance";
import { ComputeInstanceGroup } from "./computeInstanceGroup";
import { ComputePlacementGroup } from "./computePlacementGroup";
import { ComputeSnapshot } from "./computeSnapshot";
import { ContainerRegistry } from "./containerRegistry";
import { ContainerRegistryIamBinding } from "./containerRegistryIamBinding";
import { ContainerRepository } from "./containerRepository";
import { ContainerRepositoryIamBinding } from "./containerRepositoryIamBinding";
import { DataprocCluster } from "./dataprocCluster";
import { DnsRecordSet } from "./dnsRecordSet";
import { DnsZone } from "./dnsZone";
import { Function } from "./function";
import { FunctionIamBinding } from "./functionIamBinding";
import { FunctionTrigger } from "./functionTrigger";
import { IamServiceAccount } from "./iamServiceAccount";
import { IamServiceAccountApiKey } from "./iamServiceAccountApiKey";
import { IamServiceAccountIamBinding } from "./iamServiceAccountIamBinding";
import { IamServiceAccountIamMember } from "./iamServiceAccountIamMember";
import { IamServiceAccountIamPolicy } from "./iamServiceAccountIamPolicy";
import { IamServiceAccountKey } from "./iamServiceAccountKey";
import { IamServiceAccountStaticAccessKey } from "./iamServiceAccountStaticAccessKey";
import { IotCoreDevice } from "./iotCoreDevice";
import { IotCoreRegistry } from "./iotCoreRegistry";
import { KmsSecretCiphertext } from "./kmsSecretCiphertext";
import { KmsSymmetricKey } from "./kmsSymmetricKey";
import { KubernetesCluster } from "./kubernetesCluster";
import { KubernetesNodeGroup } from "./kubernetesNodeGroup";
import { LbNetworkLoadBalancer } from "./lbNetworkLoadBalancer";
import { LbTargetGroup } from "./lbTargetGroup";
import { MdbClickhouseCluster } from "./mdbClickhouseCluster";
import { MdbKafkaCluster } from "./mdbKafkaCluster";
import { MdbMongodbCluster } from "./mdbMongodbCluster";
import { MdbMysqlCluster } from "./mdbMysqlCluster";
import { MdbRedisCluster } from "./mdbRedisCluster";
import { MdbSqlServerCluster } from "./mdbSqlServerCluster";
import { MessageQueue } from "./messageQueue";
import { ResourcemanagerCloudIamBinding } from "./resourcemanagerCloudIamBinding";
import { ResourcemanagerCloudIamMember } from "./resourcemanagerCloudIamMember";
import { ResourcemanagerFolderIamBinding } from "./resourcemanagerFolderIamBinding";
import { ResourcemanagerFolderIamMember } from "./resourcemanagerFolderIamMember";
import { ResourcemanagerFolderIamPolicy } from "./resourcemanagerFolderIamPolicy";
import { StorageBucket } from "./storageBucket";
import { StorageObject } from "./storageObject";
import { VpcAddress } from "./vpcAddress";
import { VpcDefaultSecurityGroup } from "./vpcDefaultSecurityGroup";
import { VpcNetwork } from "./vpcNetwork";
import { VpcRouteTable } from "./vpcRouteTable";
import { VpcSecurityGroup } from "./vpcSecurityGroup";
import { VpcSubnet } from "./vpcSubnet";
import { YdbDatabaseDedicated } from "./ydbDatabaseDedicated";
import { YdbDatabaseServerless } from "./ydbDatabaseServerless";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "yandex:index/albBackendGroup:AlbBackendGroup":
                return new AlbBackendGroup(name, <any>undefined, { urn })
            case "yandex:index/albHttpRouter:AlbHttpRouter":
                return new AlbHttpRouter(name, <any>undefined, { urn })
            case "yandex:index/albTargetGroup:AlbTargetGroup":
                return new AlbTargetGroup(name, <any>undefined, { urn })
            case "yandex:index/apiGateway:ApiGateway":
                return new ApiGateway(name, <any>undefined, { urn })
            case "yandex:index/computeDisk:ComputeDisk":
                return new ComputeDisk(name, <any>undefined, { urn })
            case "yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup":
                return new ComputeDiskPlacementGroup(name, <any>undefined, { urn })
            case "yandex:index/computeImage:ComputeImage":
                return new ComputeImage(name, <any>undefined, { urn })
            case "yandex:index/computeInstance:ComputeInstance":
                return new ComputeInstance(name, <any>undefined, { urn })
            case "yandex:index/computeInstanceGroup:ComputeInstanceGroup":
                return new ComputeInstanceGroup(name, <any>undefined, { urn })
            case "yandex:index/computePlacementGroup:ComputePlacementGroup":
                return new ComputePlacementGroup(name, <any>undefined, { urn })
            case "yandex:index/computeSnapshot:ComputeSnapshot":
                return new ComputeSnapshot(name, <any>undefined, { urn })
            case "yandex:index/containerRegistry:ContainerRegistry":
                return new ContainerRegistry(name, <any>undefined, { urn })
            case "yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding":
                return new ContainerRegistryIamBinding(name, <any>undefined, { urn })
            case "yandex:index/containerRepository:ContainerRepository":
                return new ContainerRepository(name, <any>undefined, { urn })
            case "yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding":
                return new ContainerRepositoryIamBinding(name, <any>undefined, { urn })
            case "yandex:index/dataprocCluster:DataprocCluster":
                return new DataprocCluster(name, <any>undefined, { urn })
            case "yandex:index/dnsRecordSet:DnsRecordSet":
                return new DnsRecordSet(name, <any>undefined, { urn })
            case "yandex:index/dnsZone:DnsZone":
                return new DnsZone(name, <any>undefined, { urn })
            case "yandex:index/function:Function":
                return new Function(name, <any>undefined, { urn })
            case "yandex:index/functionIamBinding:FunctionIamBinding":
                return new FunctionIamBinding(name, <any>undefined, { urn })
            case "yandex:index/functionTrigger:FunctionTrigger":
                return new FunctionTrigger(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccount:IamServiceAccount":
                return new IamServiceAccount(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey":
                return new IamServiceAccountApiKey(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountIamBinding:IamServiceAccountIamBinding":
                return new IamServiceAccountIamBinding(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember":
                return new IamServiceAccountIamMember(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy":
                return new IamServiceAccountIamPolicy(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountKey:IamServiceAccountKey":
                return new IamServiceAccountKey(name, <any>undefined, { urn })
            case "yandex:index/iamServiceAccountStaticAccessKey:IamServiceAccountStaticAccessKey":
                return new IamServiceAccountStaticAccessKey(name, <any>undefined, { urn })
            case "yandex:index/iotCoreDevice:IotCoreDevice":
                return new IotCoreDevice(name, <any>undefined, { urn })
            case "yandex:index/iotCoreRegistry:IotCoreRegistry":
                return new IotCoreRegistry(name, <any>undefined, { urn })
            case "yandex:index/kmsSecretCiphertext:KmsSecretCiphertext":
                return new KmsSecretCiphertext(name, <any>undefined, { urn })
            case "yandex:index/kmsSymmetricKey:KmsSymmetricKey":
                return new KmsSymmetricKey(name, <any>undefined, { urn })
            case "yandex:index/kubernetesCluster:KubernetesCluster":
                return new KubernetesCluster(name, <any>undefined, { urn })
            case "yandex:index/kubernetesNodeGroup:KubernetesNodeGroup":
                return new KubernetesNodeGroup(name, <any>undefined, { urn })
            case "yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer":
                return new LbNetworkLoadBalancer(name, <any>undefined, { urn })
            case "yandex:index/lbTargetGroup:LbTargetGroup":
                return new LbTargetGroup(name, <any>undefined, { urn })
            case "yandex:index/mdbClickhouseCluster:MdbClickhouseCluster":
                return new MdbClickhouseCluster(name, <any>undefined, { urn })
            case "yandex:index/mdbKafkaCluster:MdbKafkaCluster":
                return new MdbKafkaCluster(name, <any>undefined, { urn })
            case "yandex:index/mdbMongodbCluster:MdbMongodbCluster":
                return new MdbMongodbCluster(name, <any>undefined, { urn })
            case "yandex:index/mdbMysqlCluster:MdbMysqlCluster":
                return new MdbMysqlCluster(name, <any>undefined, { urn })
            case "yandex:index/mdbRedisCluster:MdbRedisCluster":
                return new MdbRedisCluster(name, <any>undefined, { urn })
            case "yandex:index/mdbSqlServerCluster:MdbSqlServerCluster":
                return new MdbSqlServerCluster(name, <any>undefined, { urn })
            case "yandex:index/messageQueue:MessageQueue":
                return new MessageQueue(name, <any>undefined, { urn })
            case "yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding":
                return new ResourcemanagerCloudIamBinding(name, <any>undefined, { urn })
            case "yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember":
                return new ResourcemanagerCloudIamMember(name, <any>undefined, { urn })
            case "yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding":
                return new ResourcemanagerFolderIamBinding(name, <any>undefined, { urn })
            case "yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember":
                return new ResourcemanagerFolderIamMember(name, <any>undefined, { urn })
            case "yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy":
                return new ResourcemanagerFolderIamPolicy(name, <any>undefined, { urn })
            case "yandex:index/storageBucket:StorageBucket":
                return new StorageBucket(name, <any>undefined, { urn })
            case "yandex:index/storageObject:StorageObject":
                return new StorageObject(name, <any>undefined, { urn })
            case "yandex:index/vpcAddress:VpcAddress":
                return new VpcAddress(name, <any>undefined, { urn })
            case "yandex:index/vpcDefaultSecurityGroup:VpcDefaultSecurityGroup":
                return new VpcDefaultSecurityGroup(name, <any>undefined, { urn })
            case "yandex:index/vpcNetwork:VpcNetwork":
                return new VpcNetwork(name, <any>undefined, { urn })
            case "yandex:index/vpcRouteTable:VpcRouteTable":
                return new VpcRouteTable(name, <any>undefined, { urn })
            case "yandex:index/vpcSecurityGroup:VpcSecurityGroup":
                return new VpcSecurityGroup(name, <any>undefined, { urn })
            case "yandex:index/vpcSubnet:VpcSubnet":
                return new VpcSubnet(name, <any>undefined, { urn })
            case "yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated":
                return new YdbDatabaseDedicated(name, <any>undefined, { urn })
            case "yandex:index/ydbDatabaseServerless:YdbDatabaseServerless":
                return new YdbDatabaseServerless(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("yandex", "index/albBackendGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/albHttpRouter", _module)
pulumi.runtime.registerResourceModule("yandex", "index/albTargetGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/apiGateway", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeDisk", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeDiskPlacementGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeImage", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeInstance", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeInstanceGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computePlacementGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/computeSnapshot", _module)
pulumi.runtime.registerResourceModule("yandex", "index/containerRegistry", _module)
pulumi.runtime.registerResourceModule("yandex", "index/containerRegistryIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/containerRepository", _module)
pulumi.runtime.registerResourceModule("yandex", "index/containerRepositoryIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/dataprocCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/dnsRecordSet", _module)
pulumi.runtime.registerResourceModule("yandex", "index/dnsZone", _module)
pulumi.runtime.registerResourceModule("yandex", "index/function", _module)
pulumi.runtime.registerResourceModule("yandex", "index/functionIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/functionTrigger", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccount", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountApiKey", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountIamMember", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountIamPolicy", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountKey", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iamServiceAccountStaticAccessKey", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iotCoreDevice", _module)
pulumi.runtime.registerResourceModule("yandex", "index/iotCoreRegistry", _module)
pulumi.runtime.registerResourceModule("yandex", "index/kmsSecretCiphertext", _module)
pulumi.runtime.registerResourceModule("yandex", "index/kmsSymmetricKey", _module)
pulumi.runtime.registerResourceModule("yandex", "index/kubernetesCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/kubernetesNodeGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/lbNetworkLoadBalancer", _module)
pulumi.runtime.registerResourceModule("yandex", "index/lbTargetGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbClickhouseCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbKafkaCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbMongodbCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbMysqlCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbRedisCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/mdbSqlServerCluster", _module)
pulumi.runtime.registerResourceModule("yandex", "index/messageQueue", _module)
pulumi.runtime.registerResourceModule("yandex", "index/resourcemanagerCloudIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/resourcemanagerCloudIamMember", _module)
pulumi.runtime.registerResourceModule("yandex", "index/resourcemanagerFolderIamBinding", _module)
pulumi.runtime.registerResourceModule("yandex", "index/resourcemanagerFolderIamMember", _module)
pulumi.runtime.registerResourceModule("yandex", "index/resourcemanagerFolderIamPolicy", _module)
pulumi.runtime.registerResourceModule("yandex", "index/storageBucket", _module)
pulumi.runtime.registerResourceModule("yandex", "index/storageObject", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcAddress", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcDefaultSecurityGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcNetwork", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcRouteTable", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcSecurityGroup", _module)
pulumi.runtime.registerResourceModule("yandex", "index/vpcSubnet", _module)
pulumi.runtime.registerResourceModule("yandex", "index/ydbDatabaseDedicated", _module)
pulumi.runtime.registerResourceModule("yandex", "index/ydbDatabaseServerless", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("yandex", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:yandex") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
