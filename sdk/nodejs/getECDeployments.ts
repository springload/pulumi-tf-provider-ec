// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getECDeployments(args?: GetECDeploymentsArgs, opts?: pulumi.InvokeOptions): Promise<GetECDeploymentsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ec:index/getECDeployments:getECDeployments", {
        "apm": args.apm,
        "deploymentTemplateId": args.deploymentTemplateId,
        "elasticsearch": args.elasticsearch,
        "enterpriseSearch": args.enterpriseSearch,
        "healthy": args.healthy,
        "kibana": args.kibana,
        "namePrefix": args.namePrefix,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getECDeployments.
 */
export interface GetECDeploymentsArgs {
    readonly apm?: inputs.GetECDeploymentsApm;
    readonly deploymentTemplateId?: string;
    readonly elasticsearch?: inputs.GetECDeploymentsElasticsearch;
    readonly enterpriseSearch?: inputs.GetECDeploymentsEnterpriseSearch;
    readonly healthy?: string;
    readonly kibana?: inputs.GetECDeploymentsKibana;
    readonly namePrefix?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getECDeployments.
 */
export interface GetECDeploymentsResult {
    readonly apm?: outputs.GetECDeploymentsApm;
    readonly deploymentTemplateId?: string;
    readonly deployments: outputs.GetECDeploymentsDeployment[];
    readonly elasticsearch?: outputs.GetECDeploymentsElasticsearch;
    readonly enterpriseSearch?: outputs.GetECDeploymentsEnterpriseSearch;
    readonly healthy?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kibana?: outputs.GetECDeploymentsKibana;
    readonly namePrefix?: string;
    readonly returnCount: number;
    readonly tags?: {[key: string]: string};
}
