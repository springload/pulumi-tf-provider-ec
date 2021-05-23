// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ECDeploymentExtension extends pulumi.CustomResource {
    /**
     * Get an existing ECDeploymentExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ECDeploymentExtensionState, opts?: pulumi.CustomResourceOptions): ECDeploymentExtension {
        return new ECDeploymentExtension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ec:index/eCDeploymentExtension:ECDeploymentExtension';

    /**
     * Returns true if the given object is an instance of ECDeploymentExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ECDeploymentExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ECDeploymentExtension.__pulumiType;
    }

    /**
     * Description for extension
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * download url
     */
    public readonly downloadUrl!: pulumi.Output<string | undefined>;
    /**
     * Extension type. bundle or plugin
     */
    public readonly extensionType!: pulumi.Output<string>;
    /**
     * file hash
     */
    public readonly fileHash!: pulumi.Output<string | undefined>;
    /**
     * file path
     */
    public readonly filePath!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * Required name of the ruleset
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly size!: pulumi.Output<number>;
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Eleasticsearch version
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a ECDeploymentExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ECDeploymentExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ECDeploymentExtensionArgs | ECDeploymentExtensionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ECDeploymentExtensionState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["downloadUrl"] = state ? state.downloadUrl : undefined;
            inputs["extensionType"] = state ? state.extensionType : undefined;
            inputs["fileHash"] = state ? state.fileHash : undefined;
            inputs["filePath"] = state ? state.filePath : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ECDeploymentExtensionArgs | undefined;
            if ((!args || args.extensionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extensionType'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["downloadUrl"] = args ? args.downloadUrl : undefined;
            inputs["extensionType"] = args ? args.extensionType : undefined;
            inputs["fileHash"] = args ? args.fileHash : undefined;
            inputs["filePath"] = args ? args.filePath : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["lastModified"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ECDeploymentExtension.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ECDeploymentExtension resources.
 */
export interface ECDeploymentExtensionState {
    /**
     * Description for extension
     */
    readonly description?: pulumi.Input<string>;
    /**
     * download url
     */
    readonly downloadUrl?: pulumi.Input<string>;
    /**
     * Extension type. bundle or plugin
     */
    readonly extensionType?: pulumi.Input<string>;
    /**
     * file hash
     */
    readonly fileHash?: pulumi.Input<string>;
    /**
     * file path
     */
    readonly filePath?: pulumi.Input<string>;
    readonly lastModified?: pulumi.Input<string>;
    /**
     * Required name of the ruleset
     */
    readonly name?: pulumi.Input<string>;
    readonly size?: pulumi.Input<number>;
    readonly url?: pulumi.Input<string>;
    /**
     * Eleasticsearch version
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ECDeploymentExtension resource.
 */
export interface ECDeploymentExtensionArgs {
    /**
     * Description for extension
     */
    readonly description?: pulumi.Input<string>;
    /**
     * download url
     */
    readonly downloadUrl?: pulumi.Input<string>;
    /**
     * Extension type. bundle or plugin
     */
    readonly extensionType: pulumi.Input<string>;
    /**
     * file hash
     */
    readonly fileHash?: pulumi.Input<string>;
    /**
     * file path
     */
    readonly filePath?: pulumi.Input<string>;
    /**
     * Required name of the ruleset
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Eleasticsearch version
     */
    readonly version: pulumi.Input<string>;
}