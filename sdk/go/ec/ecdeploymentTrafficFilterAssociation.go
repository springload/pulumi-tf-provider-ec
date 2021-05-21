// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ECDeploymentTrafficFilterAssociation struct {
	pulumi.CustomResourceState

	// Required deployment ID where the traffic filter will be associated
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Required traffic filter ruleset ID to tie to a deployment
	TrafficFilterId pulumi.StringOutput `pulumi:"trafficFilterId"`
}

// NewECDeploymentTrafficFilterAssociation registers a new resource with the given unique name, arguments, and options.
func NewECDeploymentTrafficFilterAssociation(ctx *pulumi.Context,
	name string, args *ECDeploymentTrafficFilterAssociationArgs, opts ...pulumi.ResourceOption) (*ECDeploymentTrafficFilterAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentId'")
	}
	if args.TrafficFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficFilterId'")
	}
	var resource ECDeploymentTrafficFilterAssociation
	err := ctx.RegisterResource("ec:index/eCDeploymentTrafficFilterAssociation:ECDeploymentTrafficFilterAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetECDeploymentTrafficFilterAssociation gets an existing ECDeploymentTrafficFilterAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetECDeploymentTrafficFilterAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ECDeploymentTrafficFilterAssociationState, opts ...pulumi.ResourceOption) (*ECDeploymentTrafficFilterAssociation, error) {
	var resource ECDeploymentTrafficFilterAssociation
	err := ctx.ReadResource("ec:index/eCDeploymentTrafficFilterAssociation:ECDeploymentTrafficFilterAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ECDeploymentTrafficFilterAssociation resources.
type ecdeploymentTrafficFilterAssociationState struct {
	// Required deployment ID where the traffic filter will be associated
	DeploymentId *string `pulumi:"deploymentId"`
	// Required traffic filter ruleset ID to tie to a deployment
	TrafficFilterId *string `pulumi:"trafficFilterId"`
}

type ECDeploymentTrafficFilterAssociationState struct {
	// Required deployment ID where the traffic filter will be associated
	DeploymentId pulumi.StringPtrInput
	// Required traffic filter ruleset ID to tie to a deployment
	TrafficFilterId pulumi.StringPtrInput
}

func (ECDeploymentTrafficFilterAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecdeploymentTrafficFilterAssociationState)(nil)).Elem()
}

type ecdeploymentTrafficFilterAssociationArgs struct {
	// Required deployment ID where the traffic filter will be associated
	DeploymentId string `pulumi:"deploymentId"`
	// Required traffic filter ruleset ID to tie to a deployment
	TrafficFilterId string `pulumi:"trafficFilterId"`
}

// The set of arguments for constructing a ECDeploymentTrafficFilterAssociation resource.
type ECDeploymentTrafficFilterAssociationArgs struct {
	// Required deployment ID where the traffic filter will be associated
	DeploymentId pulumi.StringInput
	// Required traffic filter ruleset ID to tie to a deployment
	TrafficFilterId pulumi.StringInput
}

func (ECDeploymentTrafficFilterAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecdeploymentTrafficFilterAssociationArgs)(nil)).Elem()
}

type ECDeploymentTrafficFilterAssociationInput interface {
	pulumi.Input

	ToECDeploymentTrafficFilterAssociationOutput() ECDeploymentTrafficFilterAssociationOutput
	ToECDeploymentTrafficFilterAssociationOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationOutput
}

func (*ECDeploymentTrafficFilterAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ECDeploymentTrafficFilterAssociation)(nil))
}

func (i *ECDeploymentTrafficFilterAssociation) ToECDeploymentTrafficFilterAssociationOutput() ECDeploymentTrafficFilterAssociationOutput {
	return i.ToECDeploymentTrafficFilterAssociationOutputWithContext(context.Background())
}

func (i *ECDeploymentTrafficFilterAssociation) ToECDeploymentTrafficFilterAssociationOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentTrafficFilterAssociationOutput)
}

func (i *ECDeploymentTrafficFilterAssociation) ToECDeploymentTrafficFilterAssociationPtrOutput() ECDeploymentTrafficFilterAssociationPtrOutput {
	return i.ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(context.Background())
}

func (i *ECDeploymentTrafficFilterAssociation) ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentTrafficFilterAssociationPtrOutput)
}

type ECDeploymentTrafficFilterAssociationPtrInput interface {
	pulumi.Input

	ToECDeploymentTrafficFilterAssociationPtrOutput() ECDeploymentTrafficFilterAssociationPtrOutput
	ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationPtrOutput
}

type ecdeploymentTrafficFilterAssociationPtrType ECDeploymentTrafficFilterAssociationArgs

func (*ecdeploymentTrafficFilterAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ECDeploymentTrafficFilterAssociation)(nil))
}

func (i *ecdeploymentTrafficFilterAssociationPtrType) ToECDeploymentTrafficFilterAssociationPtrOutput() ECDeploymentTrafficFilterAssociationPtrOutput {
	return i.ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(context.Background())
}

func (i *ecdeploymentTrafficFilterAssociationPtrType) ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentTrafficFilterAssociationPtrOutput)
}

// ECDeploymentTrafficFilterAssociationArrayInput is an input type that accepts ECDeploymentTrafficFilterAssociationArray and ECDeploymentTrafficFilterAssociationArrayOutput values.
// You can construct a concrete instance of `ECDeploymentTrafficFilterAssociationArrayInput` via:
//
//          ECDeploymentTrafficFilterAssociationArray{ ECDeploymentTrafficFilterAssociationArgs{...} }
type ECDeploymentTrafficFilterAssociationArrayInput interface {
	pulumi.Input

	ToECDeploymentTrafficFilterAssociationArrayOutput() ECDeploymentTrafficFilterAssociationArrayOutput
	ToECDeploymentTrafficFilterAssociationArrayOutputWithContext(context.Context) ECDeploymentTrafficFilterAssociationArrayOutput
}

type ECDeploymentTrafficFilterAssociationArray []ECDeploymentTrafficFilterAssociationInput

func (ECDeploymentTrafficFilterAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ECDeploymentTrafficFilterAssociation)(nil))
}

func (i ECDeploymentTrafficFilterAssociationArray) ToECDeploymentTrafficFilterAssociationArrayOutput() ECDeploymentTrafficFilterAssociationArrayOutput {
	return i.ToECDeploymentTrafficFilterAssociationArrayOutputWithContext(context.Background())
}

func (i ECDeploymentTrafficFilterAssociationArray) ToECDeploymentTrafficFilterAssociationArrayOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentTrafficFilterAssociationArrayOutput)
}

// ECDeploymentTrafficFilterAssociationMapInput is an input type that accepts ECDeploymentTrafficFilterAssociationMap and ECDeploymentTrafficFilterAssociationMapOutput values.
// You can construct a concrete instance of `ECDeploymentTrafficFilterAssociationMapInput` via:
//
//          ECDeploymentTrafficFilterAssociationMap{ "key": ECDeploymentTrafficFilterAssociationArgs{...} }
type ECDeploymentTrafficFilterAssociationMapInput interface {
	pulumi.Input

	ToECDeploymentTrafficFilterAssociationMapOutput() ECDeploymentTrafficFilterAssociationMapOutput
	ToECDeploymentTrafficFilterAssociationMapOutputWithContext(context.Context) ECDeploymentTrafficFilterAssociationMapOutput
}

type ECDeploymentTrafficFilterAssociationMap map[string]ECDeploymentTrafficFilterAssociationInput

func (ECDeploymentTrafficFilterAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ECDeploymentTrafficFilterAssociation)(nil))
}

func (i ECDeploymentTrafficFilterAssociationMap) ToECDeploymentTrafficFilterAssociationMapOutput() ECDeploymentTrafficFilterAssociationMapOutput {
	return i.ToECDeploymentTrafficFilterAssociationMapOutputWithContext(context.Background())
}

func (i ECDeploymentTrafficFilterAssociationMap) ToECDeploymentTrafficFilterAssociationMapOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ECDeploymentTrafficFilterAssociationMapOutput)
}

type ECDeploymentTrafficFilterAssociationOutput struct {
	*pulumi.OutputState
}

func (ECDeploymentTrafficFilterAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ECDeploymentTrafficFilterAssociation)(nil))
}

func (o ECDeploymentTrafficFilterAssociationOutput) ToECDeploymentTrafficFilterAssociationOutput() ECDeploymentTrafficFilterAssociationOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationOutput) ToECDeploymentTrafficFilterAssociationOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationOutput) ToECDeploymentTrafficFilterAssociationPtrOutput() ECDeploymentTrafficFilterAssociationPtrOutput {
	return o.ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(context.Background())
}

func (o ECDeploymentTrafficFilterAssociationOutput) ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationPtrOutput {
	return o.ApplyT(func(v ECDeploymentTrafficFilterAssociation) *ECDeploymentTrafficFilterAssociation {
		return &v
	}).(ECDeploymentTrafficFilterAssociationPtrOutput)
}

type ECDeploymentTrafficFilterAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (ECDeploymentTrafficFilterAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ECDeploymentTrafficFilterAssociation)(nil))
}

func (o ECDeploymentTrafficFilterAssociationPtrOutput) ToECDeploymentTrafficFilterAssociationPtrOutput() ECDeploymentTrafficFilterAssociationPtrOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationPtrOutput) ToECDeploymentTrafficFilterAssociationPtrOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationPtrOutput {
	return o
}

type ECDeploymentTrafficFilterAssociationArrayOutput struct{ *pulumi.OutputState }

func (ECDeploymentTrafficFilterAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ECDeploymentTrafficFilterAssociation)(nil))
}

func (o ECDeploymentTrafficFilterAssociationArrayOutput) ToECDeploymentTrafficFilterAssociationArrayOutput() ECDeploymentTrafficFilterAssociationArrayOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationArrayOutput) ToECDeploymentTrafficFilterAssociationArrayOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationArrayOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationArrayOutput) Index(i pulumi.IntInput) ECDeploymentTrafficFilterAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ECDeploymentTrafficFilterAssociation {
		return vs[0].([]ECDeploymentTrafficFilterAssociation)[vs[1].(int)]
	}).(ECDeploymentTrafficFilterAssociationOutput)
}

type ECDeploymentTrafficFilterAssociationMapOutput struct{ *pulumi.OutputState }

func (ECDeploymentTrafficFilterAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ECDeploymentTrafficFilterAssociation)(nil))
}

func (o ECDeploymentTrafficFilterAssociationMapOutput) ToECDeploymentTrafficFilterAssociationMapOutput() ECDeploymentTrafficFilterAssociationMapOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationMapOutput) ToECDeploymentTrafficFilterAssociationMapOutputWithContext(ctx context.Context) ECDeploymentTrafficFilterAssociationMapOutput {
	return o
}

func (o ECDeploymentTrafficFilterAssociationMapOutput) MapIndex(k pulumi.StringInput) ECDeploymentTrafficFilterAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ECDeploymentTrafficFilterAssociation {
		return vs[0].(map[string]ECDeploymentTrafficFilterAssociation)[vs[1].(string)]
	}).(ECDeploymentTrafficFilterAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(ECDeploymentTrafficFilterAssociationOutput{})
	pulumi.RegisterOutputType(ECDeploymentTrafficFilterAssociationPtrOutput{})
	pulumi.RegisterOutputType(ECDeploymentTrafficFilterAssociationArrayOutput{})
	pulumi.RegisterOutputType(ECDeploymentTrafficFilterAssociationMapOutput{})
}
