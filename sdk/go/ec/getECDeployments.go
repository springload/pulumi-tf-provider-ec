// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetECDeployments(ctx *pulumi.Context, args *GetECDeploymentsArgs, opts ...pulumi.InvokeOption) (*GetECDeploymentsResult, error) {
	var rv GetECDeploymentsResult
	err := ctx.Invoke("ec:index/getECDeployments:getECDeployments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getECDeployments.
type GetECDeploymentsArgs struct {
	Apm                  *GetECDeploymentsApm              `pulumi:"apm"`
	DeploymentTemplateId *string                           `pulumi:"deploymentTemplateId"`
	Elasticsearch        *GetECDeploymentsElasticsearch    `pulumi:"elasticsearch"`
	EnterpriseSearch     *GetECDeploymentsEnterpriseSearch `pulumi:"enterpriseSearch"`
	Healthy              *string                           `pulumi:"healthy"`
	Kibana               *GetECDeploymentsKibana           `pulumi:"kibana"`
	NamePrefix           *string                           `pulumi:"namePrefix"`
	Tags                 map[string]string                 `pulumi:"tags"`
}

// A collection of values returned by getECDeployments.
type GetECDeploymentsResult struct {
	Apm                  *GetECDeploymentsApm              `pulumi:"apm"`
	DeploymentTemplateId *string                           `pulumi:"deploymentTemplateId"`
	Deployments          []GetECDeploymentsDeployment      `pulumi:"deployments"`
	Elasticsearch        *GetECDeploymentsElasticsearch    `pulumi:"elasticsearch"`
	EnterpriseSearch     *GetECDeploymentsEnterpriseSearch `pulumi:"enterpriseSearch"`
	Healthy              *string                           `pulumi:"healthy"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                  `pulumi:"id"`
	Kibana      *GetECDeploymentsKibana `pulumi:"kibana"`
	NamePrefix  *string                 `pulumi:"namePrefix"`
	ReturnCount int                     `pulumi:"returnCount"`
	Tags        map[string]string       `pulumi:"tags"`
}
