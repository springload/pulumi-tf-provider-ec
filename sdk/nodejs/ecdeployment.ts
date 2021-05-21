// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class ECDeployment extends pulumi.CustomResource {
    /**
     * Get an existing ECDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ECDeploymentState, opts?: pulumi.CustomResourceOptions): ECDeployment {
        return new ECDeployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ec:index/eCDeployment:ECDeployment';

    /**
     * Returns true if the given object is an instance of ECDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ECDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ECDeployment.__pulumiType;
    }

    /**
     * Optional APM resource definition
     */
    public readonly apm!: pulumi.Output<outputs.ECDeploymentApm | undefined>;
    public /*out*/ readonly apmSecretToken!: pulumi.Output<string>;
    /**
     * Required Deployment Template identifier to create the deployment from
     */
    public readonly deploymentTemplateId!: pulumi.Output<string>;
    /**
     * Required Elasticsearch resource definition
     */
    public readonly elasticsearch!: pulumi.Output<outputs.ECDeploymentElasticsearch>;
    /**
     * Computed password obtained upon creating the Elasticsearch resource
     */
    public /*out*/ readonly elasticsearchPassword!: pulumi.Output<string>;
    /**
     * Computed username obtained upon creating the Elasticsearch resource
     */
    public /*out*/ readonly elasticsearchUsername!: pulumi.Output<string>;
    /**
     * Optional Enterprise Search resource definition
     */
    public readonly enterpriseSearch!: pulumi.Output<outputs.ECDeploymentEnterpriseSearch | undefined>;
    /**
     * Optional Kibana resource definition
     */
    public readonly kibana!: pulumi.Output<outputs.ECDeploymentKibana | undefined>;
    /**
     * Optional name for the deployment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional observability settings. Ship logs and metrics to a dedicated deployment.
     */
    public readonly observability!: pulumi.Output<outputs.ECDeploymentObservability | undefined>;
    /**
     * Required ESS region where to create the deployment, for ECE environments "ece-region" must be set
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Optional request_id to set on the create operation, only use when previous create attempts return with an error and a
     * request_id is returned as part of the error
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Optional map of deployment tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Optional list of traffic filters to apply to this deployment.
     */
    public readonly trafficFilters!: pulumi.Output<string[] | undefined>;
    /**
     * Required Elastic Stack version to use for all of the deployment resources
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a ECDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ECDeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ECDeploymentArgs | ECDeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ECDeploymentState | undefined;
            inputs["apm"] = state ? state.apm : undefined;
            inputs["apmSecretToken"] = state ? state.apmSecretToken : undefined;
            inputs["deploymentTemplateId"] = state ? state.deploymentTemplateId : undefined;
            inputs["elasticsearch"] = state ? state.elasticsearch : undefined;
            inputs["elasticsearchPassword"] = state ? state.elasticsearchPassword : undefined;
            inputs["elasticsearchUsername"] = state ? state.elasticsearchUsername : undefined;
            inputs["enterpriseSearch"] = state ? state.enterpriseSearch : undefined;
            inputs["kibana"] = state ? state.kibana : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["observability"] = state ? state.observability : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["requestId"] = state ? state.requestId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["trafficFilters"] = state ? state.trafficFilters : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ECDeploymentArgs | undefined;
            if ((!args || args.deploymentTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentTemplateId'");
            }
            if ((!args || args.elasticsearch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'elasticsearch'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["apm"] = args ? args.apm : undefined;
            inputs["deploymentTemplateId"] = args ? args.deploymentTemplateId : undefined;
            inputs["elasticsearch"] = args ? args.elasticsearch : undefined;
            inputs["enterpriseSearch"] = args ? args.enterpriseSearch : undefined;
            inputs["kibana"] = args ? args.kibana : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["observability"] = args ? args.observability : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trafficFilters"] = args ? args.trafficFilters : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["apmSecretToken"] = undefined /*out*/;
            inputs["elasticsearchPassword"] = undefined /*out*/;
            inputs["elasticsearchUsername"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ECDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ECDeployment resources.
 */
export interface ECDeploymentState {
    /**
     * Optional APM resource definition
     */
    readonly apm?: pulumi.Input<inputs.ECDeploymentApm>;
    readonly apmSecretToken?: pulumi.Input<string>;
    /**
     * Required Deployment Template identifier to create the deployment from
     */
    readonly deploymentTemplateId?: pulumi.Input<string>;
    /**
     * Required Elasticsearch resource definition
     */
    readonly elasticsearch?: pulumi.Input<inputs.ECDeploymentElasticsearch>;
    /**
     * Computed password obtained upon creating the Elasticsearch resource
     */
    readonly elasticsearchPassword?: pulumi.Input<string>;
    /**
     * Computed username obtained upon creating the Elasticsearch resource
     */
    readonly elasticsearchUsername?: pulumi.Input<string>;
    /**
     * Optional Enterprise Search resource definition
     */
    readonly enterpriseSearch?: pulumi.Input<inputs.ECDeploymentEnterpriseSearch>;
    /**
     * Optional Kibana resource definition
     */
    readonly kibana?: pulumi.Input<inputs.ECDeploymentKibana>;
    /**
     * Optional name for the deployment
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Optional observability settings. Ship logs and metrics to a dedicated deployment.
     */
    readonly observability?: pulumi.Input<inputs.ECDeploymentObservability>;
    /**
     * Required ESS region where to create the deployment, for ECE environments "ece-region" must be set
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Optional request_id to set on the create operation, only use when previous create attempts return with an error and a
     * request_id is returned as part of the error
     */
    readonly requestId?: pulumi.Input<string>;
    /**
     * Optional map of deployment tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional list of traffic filters to apply to this deployment.
     */
    readonly trafficFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required Elastic Stack version to use for all of the deployment resources
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ECDeployment resource.
 */
export interface ECDeploymentArgs {
    /**
     * Optional APM resource definition
     */
    readonly apm?: pulumi.Input<inputs.ECDeploymentApm>;
    /**
     * Required Deployment Template identifier to create the deployment from
     */
    readonly deploymentTemplateId: pulumi.Input<string>;
    /**
     * Required Elasticsearch resource definition
     */
    readonly elasticsearch: pulumi.Input<inputs.ECDeploymentElasticsearch>;
    /**
     * Optional Enterprise Search resource definition
     */
    readonly enterpriseSearch?: pulumi.Input<inputs.ECDeploymentEnterpriseSearch>;
    /**
     * Optional Kibana resource definition
     */
    readonly kibana?: pulumi.Input<inputs.ECDeploymentKibana>;
    /**
     * Optional name for the deployment
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Optional observability settings. Ship logs and metrics to a dedicated deployment.
     */
    readonly observability?: pulumi.Input<inputs.ECDeploymentObservability>;
    /**
     * Required ESS region where to create the deployment, for ECE environments "ece-region" must be set
     */
    readonly region: pulumi.Input<string>;
    /**
     * Optional request_id to set on the create operation, only use when previous create attempts return with an error and a
     * request_id is returned as part of the error
     */
    readonly requestId?: pulumi.Input<string>;
    /**
     * Optional map of deployment tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional list of traffic filters to apply to this deployment.
     */
    readonly trafficFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required Elastic Stack version to use for all of the deployment resources
     */
    readonly version: pulumi.Input<string>;
}
