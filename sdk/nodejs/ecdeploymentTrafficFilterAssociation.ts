// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ECDeploymentTrafficFilterAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ECDeploymentTrafficFilterAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ECDeploymentTrafficFilterAssociationState, opts?: pulumi.CustomResourceOptions): ECDeploymentTrafficFilterAssociation {
        return new ECDeploymentTrafficFilterAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ec:index/eCDeploymentTrafficFilterAssociation:ECDeploymentTrafficFilterAssociation';

    /**
     * Returns true if the given object is an instance of ECDeploymentTrafficFilterAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ECDeploymentTrafficFilterAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ECDeploymentTrafficFilterAssociation.__pulumiType;
    }

    /**
     * Required deployment ID where the traffic filter will be associated
     */
    public readonly deploymentId!: pulumi.Output<string>;
    /**
     * Required traffic filter ruleset ID to tie to a deployment
     */
    public readonly trafficFilterId!: pulumi.Output<string>;

    /**
     * Create a ECDeploymentTrafficFilterAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ECDeploymentTrafficFilterAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ECDeploymentTrafficFilterAssociationArgs | ECDeploymentTrafficFilterAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ECDeploymentTrafficFilterAssociationState | undefined;
            inputs["deploymentId"] = state ? state.deploymentId : undefined;
            inputs["trafficFilterId"] = state ? state.trafficFilterId : undefined;
        } else {
            const args = argsOrState as ECDeploymentTrafficFilterAssociationArgs | undefined;
            if ((!args || args.deploymentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentId'");
            }
            if ((!args || args.trafficFilterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficFilterId'");
            }
            inputs["deploymentId"] = args ? args.deploymentId : undefined;
            inputs["trafficFilterId"] = args ? args.trafficFilterId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ECDeploymentTrafficFilterAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ECDeploymentTrafficFilterAssociation resources.
 */
export interface ECDeploymentTrafficFilterAssociationState {
    /**
     * Required deployment ID where the traffic filter will be associated
     */
    readonly deploymentId?: pulumi.Input<string>;
    /**
     * Required traffic filter ruleset ID to tie to a deployment
     */
    readonly trafficFilterId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ECDeploymentTrafficFilterAssociation resource.
 */
export interface ECDeploymentTrafficFilterAssociationArgs {
    /**
     * Required deployment ID where the traffic filter will be associated
     */
    readonly deploymentId: pulumi.Input<string>;
    /**
     * Required traffic filter ruleset ID to tie to a deployment
     */
    readonly trafficFilterId: pulumi.Input<string>;
}
