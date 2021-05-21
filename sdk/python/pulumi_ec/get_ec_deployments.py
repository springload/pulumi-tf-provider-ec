# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetECDeploymentsResult',
    'AwaitableGetECDeploymentsResult',
    'get_ec_deployments',
]

@pulumi.output_type
class GetECDeploymentsResult:
    """
    A collection of values returned by getECDeployments.
    """
    def __init__(__self__, apm=None, deployment_template_id=None, deployments=None, elasticsearch=None, enterprise_search=None, healthy=None, id=None, kibana=None, name_prefix=None, return_count=None, tags=None):
        if apm and not isinstance(apm, dict):
            raise TypeError("Expected argument 'apm' to be a dict")
        pulumi.set(__self__, "apm", apm)
        if deployment_template_id and not isinstance(deployment_template_id, str):
            raise TypeError("Expected argument 'deployment_template_id' to be a str")
        pulumi.set(__self__, "deployment_template_id", deployment_template_id)
        if deployments and not isinstance(deployments, list):
            raise TypeError("Expected argument 'deployments' to be a list")
        pulumi.set(__self__, "deployments", deployments)
        if elasticsearch and not isinstance(elasticsearch, dict):
            raise TypeError("Expected argument 'elasticsearch' to be a dict")
        pulumi.set(__self__, "elasticsearch", elasticsearch)
        if enterprise_search and not isinstance(enterprise_search, dict):
            raise TypeError("Expected argument 'enterprise_search' to be a dict")
        pulumi.set(__self__, "enterprise_search", enterprise_search)
        if healthy and not isinstance(healthy, str):
            raise TypeError("Expected argument 'healthy' to be a str")
        pulumi.set(__self__, "healthy", healthy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kibana and not isinstance(kibana, dict):
            raise TypeError("Expected argument 'kibana' to be a dict")
        pulumi.set(__self__, "kibana", kibana)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if return_count and not isinstance(return_count, int):
            raise TypeError("Expected argument 'return_count' to be a int")
        pulumi.set(__self__, "return_count", return_count)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def apm(self) -> Optional['outputs.GetECDeploymentsApmResult']:
        return pulumi.get(self, "apm")

    @property
    @pulumi.getter(name="deploymentTemplateId")
    def deployment_template_id(self) -> Optional[str]:
        return pulumi.get(self, "deployment_template_id")

    @property
    @pulumi.getter
    def deployments(self) -> Sequence['outputs.GetECDeploymentsDeploymentResult']:
        return pulumi.get(self, "deployments")

    @property
    @pulumi.getter
    def elasticsearch(self) -> Optional['outputs.GetECDeploymentsElasticsearchResult']:
        return pulumi.get(self, "elasticsearch")

    @property
    @pulumi.getter(name="enterpriseSearch")
    def enterprise_search(self) -> Optional['outputs.GetECDeploymentsEnterpriseSearchResult']:
        return pulumi.get(self, "enterprise_search")

    @property
    @pulumi.getter
    def healthy(self) -> Optional[str]:
        return pulumi.get(self, "healthy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kibana(self) -> Optional['outputs.GetECDeploymentsKibanaResult']:
        return pulumi.get(self, "kibana")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="returnCount")
    def return_count(self) -> int:
        return pulumi.get(self, "return_count")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")


class AwaitableGetECDeploymentsResult(GetECDeploymentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetECDeploymentsResult(
            apm=self.apm,
            deployment_template_id=self.deployment_template_id,
            deployments=self.deployments,
            elasticsearch=self.elasticsearch,
            enterprise_search=self.enterprise_search,
            healthy=self.healthy,
            id=self.id,
            kibana=self.kibana,
            name_prefix=self.name_prefix,
            return_count=self.return_count,
            tags=self.tags)


def get_ec_deployments(apm: Optional[pulumi.InputType['GetECDeploymentsApmArgs']] = None,
                       deployment_template_id: Optional[str] = None,
                       elasticsearch: Optional[pulumi.InputType['GetECDeploymentsElasticsearchArgs']] = None,
                       enterprise_search: Optional[pulumi.InputType['GetECDeploymentsEnterpriseSearchArgs']] = None,
                       healthy: Optional[str] = None,
                       kibana: Optional[pulumi.InputType['GetECDeploymentsKibanaArgs']] = None,
                       name_prefix: Optional[str] = None,
                       tags: Optional[Mapping[str, str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetECDeploymentsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['apm'] = apm
    __args__['deploymentTemplateId'] = deployment_template_id
    __args__['elasticsearch'] = elasticsearch
    __args__['enterpriseSearch'] = enterprise_search
    __args__['healthy'] = healthy
    __args__['kibana'] = kibana
    __args__['namePrefix'] = name_prefix
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ec:index/getECDeployments:getECDeployments', __args__, opts=opts, typ=GetECDeploymentsResult).value

    return AwaitableGetECDeploymentsResult(
        apm=__ret__.apm,
        deployment_template_id=__ret__.deployment_template_id,
        deployments=__ret__.deployments,
        elasticsearch=__ret__.elasticsearch,
        enterprise_search=__ret__.enterprise_search,
        healthy=__ret__.healthy,
        id=__ret__.id,
        kibana=__ret__.kibana,
        name_prefix=__ret__.name_prefix,
        return_count=__ret__.return_count,
        tags=__ret__.tags)
