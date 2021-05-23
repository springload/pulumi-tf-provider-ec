// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ECDeploymentExtension struct {
	pulumi.CustomResourceState

	// Description for extension
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// download url
	DownloadUrl pulumi.StringPtrOutput `pulumi:"downloadUrl"`
	// Extension type. bundle or plugin
	ExtensionType pulumi.StringOutput `pulumi:"extensionType"`
	// file hash
	FileHash pulumi.StringPtrOutput `pulumi:"fileHash"`
	// file path
	FilePath     pulumi.StringPtrOutput `pulumi:"filePath"`
	LastModified pulumi.StringOutput    `pulumi:"lastModified"`
	// Required name of the ruleset
	Name pulumi.StringOutput `pulumi:"name"`
	Size pulumi.IntOutput    `pulumi:"size"`
	Url  pulumi.StringOutput `pulumi:"url"`
	// Eleasticsearch version
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewECDeploymentExtension registers a new resource with the given unique name, arguments, and options.
func NewECDeploymentExtension(ctx *pulumi.Context,
	name string, args *ECDeploymentExtensionArgs, opts ...pulumi.ResourceOption) (*ECDeploymentExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtensionType == nil {
		return nil, errors.New("invalid value for required argument 'ExtensionType'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource ECDeploymentExtension
	err := ctx.RegisterResource("ec:index/eCDeploymentExtension:ECDeploymentExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetECDeploymentExtension gets an existing ECDeploymentExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetECDeploymentExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ECDeploymentExtensionState, opts ...pulumi.ResourceOption) (*ECDeploymentExtension, error) {
	var resource ECDeploymentExtension
	err := ctx.ReadResource("ec:index/eCDeploymentExtension:ECDeploymentExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ECDeploymentExtension resources.
type ecdeploymentExtensionState struct {
	// Description for extension
	Description *string `pulumi:"description"`
	// download url
	DownloadUrl *string `pulumi:"downloadUrl"`
	// Extension type. bundle or plugin
	ExtensionType *string `pulumi:"extensionType"`
	// file hash
	FileHash *string `pulumi:"fileHash"`
	// file path
	FilePath     *string `pulumi:"filePath"`
	LastModified *string `pulumi:"lastModified"`
	// Required name of the ruleset
	Name *string `pulumi:"name"`
	Size *int    `pulumi:"size"`
	Url  *string `pulumi:"url"`
	// Eleasticsearch version
	Version *string `pulumi:"version"`
}

type ECDeploymentExtensionState struct {
	// Description for extension
	Description pulumi.StringPtrInput
	// download url
	DownloadUrl pulumi.StringPtrInput
	// Extension type. bundle or plugin
	ExtensionType pulumi.StringPtrInput
	// file hash
	FileHash pulumi.StringPtrInput
	// file path
	FilePath     pulumi.StringPtrInput
	LastModified pulumi.StringPtrInput
	// Required name of the ruleset
	Name pulumi.StringPtrInput
	Size pulumi.IntPtrInput
	Url  pulumi.StringPtrInput
	// Eleasticsearch version
	Version pulumi.StringPtrInput
}

func (ECDeploymentExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecdeploymentExtensionState)(nil)).Elem()
}

type ecdeploymentExtensionArgs struct {
	// Description for extension
	Description *string `pulumi:"description"`
	// download url
	DownloadUrl *string `pulumi:"downloadUrl"`
	// Extension type. bundle or plugin
	ExtensionType string `pulumi:"extensionType"`
	// file hash
	FileHash *string `pulumi:"fileHash"`
	// file path
	FilePath *string `pulumi:"filePath"`
	// Required name of the ruleset
	Name *string `pulumi:"name"`
	// Eleasticsearch version
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a ECDeploymentExtension resource.
type ECDeploymentExtensionArgs struct {
	// Description for extension
	Description pulumi.StringPtrInput
	// download url
	DownloadUrl pulumi.StringPtrInput
	// Extension type. bundle or plugin
	ExtensionType pulumi.StringInput
	// file hash
	FileHash pulumi.StringPtrInput
	// file path
	FilePath pulumi.StringPtrInput
	// Required name of the ruleset
	Name pulumi.StringPtrInput
	// Eleasticsearch version
	Version pulumi.StringInput
}

func (ECDeploymentExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecdeploymentExtensionArgs)(nil)).Elem()
}

type ECDeploymentExtensionInput interface {
	pulumi.Input

	ToECDeploymentExtensionOutput() ECDeploymentExtensionOutput
	ToECDeploymentExtensionOutputWithContext(ctx context.Context) ECDeploymentExtensionOutput
}

func (*ECDeploymentExtension) ElementType() reflect.Type {
	return reflect.TypeOf((*ECDeploymentExtension)(nil))
}

func (i *ECDeploymentExtension) ToECDeploymentExtensionOutput() ECDeploymentExtensionOutput {
	return i.ToECDeploymentExtensionOutputWithContext(context.Background())
}

func (i *ECDeploymentExtension) ToECDeploymentExtensionOutputWithContext(ctx context.Context) ECDeploymentExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentExtensionOutput)
}

func (i *ECDeploymentExtension) ToECDeploymentExtensionPtrOutput() ECDeploymentExtensionPtrOutput {
	return i.ToECDeploymentExtensionPtrOutputWithContext(context.Background())
}

func (i *ECDeploymentExtension) ToECDeploymentExtensionPtrOutputWithContext(ctx context.Context) ECDeploymentExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentExtensionPtrOutput)
}

type ECDeploymentExtensionPtrInput interface {
	pulumi.Input

	ToECDeploymentExtensionPtrOutput() ECDeploymentExtensionPtrOutput
	ToECDeploymentExtensionPtrOutputWithContext(ctx context.Context) ECDeploymentExtensionPtrOutput
}

type ecdeploymentExtensionPtrType ECDeploymentExtensionArgs

func (*ecdeploymentExtensionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ECDeploymentExtension)(nil))
}

func (i *ecdeploymentExtensionPtrType) ToECDeploymentExtensionPtrOutput() ECDeploymentExtensionPtrOutput {
	return i.ToECDeploymentExtensionPtrOutputWithContext(context.Background())
}

func (i *ecdeploymentExtensionPtrType) ToECDeploymentExtensionPtrOutputWithContext(ctx context.Context) ECDeploymentExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentExtensionPtrOutput)
}

// ECDeploymentExtensionArrayInput is an input type that accepts ECDeploymentExtensionArray and ECDeploymentExtensionArrayOutput values.
// You can construct a concrete instance of `ECDeploymentExtensionArrayInput` via:
//
//          ECDeploymentExtensionArray{ ECDeploymentExtensionArgs{...} }
type ECDeploymentExtensionArrayInput interface {
	pulumi.Input

	ToECDeploymentExtensionArrayOutput() ECDeploymentExtensionArrayOutput
	ToECDeploymentExtensionArrayOutputWithContext(context.Context) ECDeploymentExtensionArrayOutput
}

type ECDeploymentExtensionArray []ECDeploymentExtensionInput

func (ECDeploymentExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ECDeploymentExtension)(nil))
}

func (i ECDeploymentExtensionArray) ToECDeploymentExtensionArrayOutput() ECDeploymentExtensionArrayOutput {
	return i.ToECDeploymentExtensionArrayOutputWithContext(context.Background())
}

func (i ECDeploymentExtensionArray) ToECDeploymentExtensionArrayOutputWithContext(ctx context.Context) ECDeploymentExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentExtensionArrayOutput)
}

// ECDeploymentExtensionMapInput is an input type that accepts ECDeploymentExtensionMap and ECDeploymentExtensionMapOutput values.
// You can construct a concrete instance of `ECDeploymentExtensionMapInput` via:
//
//          ECDeploymentExtensionMap{ "key": ECDeploymentExtensionArgs{...} }
type ECDeploymentExtensionMapInput interface {
	pulumi.Input

	ToECDeploymentExtensionMapOutput() ECDeploymentExtensionMapOutput
	ToECDeploymentExtensionMapOutputWithContext(context.Context) ECDeploymentExtensionMapOutput
}

type ECDeploymentExtensionMap map[string]ECDeploymentExtensionInput

func (ECDeploymentExtensionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ECDeploymentExtension)(nil))
}

func (i ECDeploymentExtensionMap) ToECDeploymentExtensionMapOutput() ECDeploymentExtensionMapOutput {
	return i.ToECDeploymentExtensionMapOutputWithContext(context.Background())
}

func (i ECDeploymentExtensionMap) ToECDeploymentExtensionMapOutputWithContext(ctx context.Context) ECDeploymentExtensionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentExtensionMapOutput)
}

type ECDeploymentExtensionOutput struct {
	*pulumi.OutputState
}

func (ECDeploymentExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ECDeploymentExtension)(nil))
}

func (o ECDeploymentExtensionOutput) ToECDeploymentExtensionOutput() ECDeploymentExtensionOutput {
	return o
}

func (o ECDeploymentExtensionOutput) ToECDeploymentExtensionOutputWithContext(ctx context.Context) ECDeploymentExtensionOutput {
	return o
}

func (o ECDeploymentExtensionOutput) ToECDeploymentExtensionPtrOutput() ECDeploymentExtensionPtrOutput {
	return o.ToECDeploymentExtensionPtrOutputWithContext(context.Background())
}

func (o ECDeploymentExtensionOutput) ToECDeploymentExtensionPtrOutputWithContext(ctx context.Context) ECDeploymentExtensionPtrOutput {
	return o.ApplyT(func(v ECDeploymentExtension) *ECDeploymentExtension {
		return &v
	}).(ECDeploymentExtensionPtrOutput)
}

type ECDeploymentExtensionPtrOutput struct {
	*pulumi.OutputState
}

func (ECDeploymentExtensionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ECDeploymentExtension)(nil))
}

func (o ECDeploymentExtensionPtrOutput) ToECDeploymentExtensionPtrOutput() ECDeploymentExtensionPtrOutput {
	return o
}

func (o ECDeploymentExtensionPtrOutput) ToECDeploymentExtensionPtrOutputWithContext(ctx context.Context) ECDeploymentExtensionPtrOutput {
	return o
}

type ECDeploymentExtensionArrayOutput struct{ *pulumi.OutputState }

func (ECDeploymentExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ECDeploymentExtension)(nil))
}

func (o ECDeploymentExtensionArrayOutput) ToECDeploymentExtensionArrayOutput() ECDeploymentExtensionArrayOutput {
	return o
}

func (o ECDeploymentExtensionArrayOutput) ToECDeploymentExtensionArrayOutputWithContext(ctx context.Context) ECDeploymentExtensionArrayOutput {
	return o
}

func (o ECDeploymentExtensionArrayOutput) Index(i pulumi.IntInput) ECDeploymentExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ECDeploymentExtension {
		return vs[0].([]ECDeploymentExtension)[vs[1].(int)]
	}).(ECDeploymentExtensionOutput)
}

type ECDeploymentExtensionMapOutput struct{ *pulumi.OutputState }

func (ECDeploymentExtensionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ECDeploymentExtension)(nil))
}

func (o ECDeploymentExtensionMapOutput) ToECDeploymentExtensionMapOutput() ECDeploymentExtensionMapOutput {
	return o
}

func (o ECDeploymentExtensionMapOutput) ToECDeploymentExtensionMapOutputWithContext(ctx context.Context) ECDeploymentExtensionMapOutput {
	return o
}

func (o ECDeploymentExtensionMapOutput) MapIndex(k pulumi.StringInput) ECDeploymentExtensionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ECDeploymentExtension {
		return vs[0].(map[string]ECDeploymentExtension)[vs[1].(string)]
	}).(ECDeploymentExtensionOutput)
}

func init() {
	pulumi.RegisterOutputType(ECDeploymentExtensionOutput{})
	pulumi.RegisterOutputType(ECDeploymentExtensionPtrOutput{})
	pulumi.RegisterOutputType(ECDeploymentExtensionArrayOutput{})
	pulumi.RegisterOutputType(ECDeploymentExtensionMapOutput{})
}