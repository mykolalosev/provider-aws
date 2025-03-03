/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PrivateDNSNamespace.
func (mg *PrivateDNSNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCRef,
		Selector:     mg.Spec.ForProvider.VPCSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPC")
	}
	mg.Spec.ForProvider.VPC = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCRef = rsp.ResolvedReference

	return nil
}
