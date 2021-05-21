# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ECDeploymentTrafficFilterAssociationArgs', 'ECDeploymentTrafficFilterAssociation']

@pulumi.input_type
class ECDeploymentTrafficFilterAssociationArgs:
    def __init__(__self__, *,
                 deployment_id: pulumi.Input[str],
                 traffic_filter_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ECDeploymentTrafficFilterAssociation resource.
        :param pulumi.Input[str] deployment_id: Required deployment ID where the traffic filter will be associated
        :param pulumi.Input[str] traffic_filter_id: Required traffic filter ruleset ID to tie to a deployment
        """
        pulumi.set(__self__, "deployment_id", deployment_id)
        pulumi.set(__self__, "traffic_filter_id", traffic_filter_id)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Input[str]:
        """
        Required deployment ID where the traffic filter will be associated
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="trafficFilterId")
    def traffic_filter_id(self) -> pulumi.Input[str]:
        """
        Required traffic filter ruleset ID to tie to a deployment
        """
        return pulumi.get(self, "traffic_filter_id")

    @traffic_filter_id.setter
    def traffic_filter_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_filter_id", value)


@pulumi.input_type
class _ECDeploymentTrafficFilterAssociationState:
    def __init__(__self__, *,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 traffic_filter_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ECDeploymentTrafficFilterAssociation resources.
        :param pulumi.Input[str] deployment_id: Required deployment ID where the traffic filter will be associated
        :param pulumi.Input[str] traffic_filter_id: Required traffic filter ruleset ID to tie to a deployment
        """
        if deployment_id is not None:
            pulumi.set(__self__, "deployment_id", deployment_id)
        if traffic_filter_id is not None:
            pulumi.set(__self__, "traffic_filter_id", traffic_filter_id)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required deployment ID where the traffic filter will be associated
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="trafficFilterId")
    def traffic_filter_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required traffic filter ruleset ID to tie to a deployment
        """
        return pulumi.get(self, "traffic_filter_id")

    @traffic_filter_id.setter
    def traffic_filter_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_filter_id", value)


class ECDeploymentTrafficFilterAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 traffic_filter_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ECDeploymentTrafficFilterAssociation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_id: Required deployment ID where the traffic filter will be associated
        :param pulumi.Input[str] traffic_filter_id: Required traffic filter ruleset ID to tie to a deployment
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ECDeploymentTrafficFilterAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ECDeploymentTrafficFilterAssociation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ECDeploymentTrafficFilterAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ECDeploymentTrafficFilterAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 traffic_filter_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ECDeploymentTrafficFilterAssociationArgs.__new__(ECDeploymentTrafficFilterAssociationArgs)

            if deployment_id is None and not opts.urn:
                raise TypeError("Missing required property 'deployment_id'")
            __props__.__dict__["deployment_id"] = deployment_id
            if traffic_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_filter_id'")
            __props__.__dict__["traffic_filter_id"] = traffic_filter_id
        super(ECDeploymentTrafficFilterAssociation, __self__).__init__(
            'ec:index/eCDeploymentTrafficFilterAssociation:ECDeploymentTrafficFilterAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deployment_id: Optional[pulumi.Input[str]] = None,
            traffic_filter_id: Optional[pulumi.Input[str]] = None) -> 'ECDeploymentTrafficFilterAssociation':
        """
        Get an existing ECDeploymentTrafficFilterAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_id: Required deployment ID where the traffic filter will be associated
        :param pulumi.Input[str] traffic_filter_id: Required traffic filter ruleset ID to tie to a deployment
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ECDeploymentTrafficFilterAssociationState.__new__(_ECDeploymentTrafficFilterAssociationState)

        __props__.__dict__["deployment_id"] = deployment_id
        __props__.__dict__["traffic_filter_id"] = traffic_filter_id
        return ECDeploymentTrafficFilterAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[str]:
        """
        Required deployment ID where the traffic filter will be associated
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter(name="trafficFilterId")
    def traffic_filter_id(self) -> pulumi.Output[str]:
        """
        Required traffic filter ruleset ID to tie to a deployment
        """
        return pulumi.get(self, "traffic_filter_id")

