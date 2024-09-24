// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/gmeligio/terraform-provider-netparse/internal/netparse"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = ParseCIDRFunction{}

type ParseCIDRFunction struct{}

type parseCIDRFunctionReturnModel struct {
	IP      string `tfsdk:"ip"`
	Network string `tfsdk:"network"`
}

func NewParseCIDRFunction() function.Function {
	return ParseCIDRFunction{}
}

func FromCIDRModel(d *netparse.CidrModel) parseCIDRFunctionReturnModel {
	return parseCIDRFunctionReturnModel{
		IP:      d.IP,
		Network: d.Network,
	}
}

func (f ParseCIDRFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "parse_cidr"
}

func (f ParseCIDRFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             parseCIDRMarkdownDescription,
		MarkdownDescription: parseCIDRMarkdownDescription,
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "cidr",
				MarkdownDescription: cidrAttrMarkdownDescription,
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"ip":      types.StringType,
				"network": types.StringType,
			},
		},
	}
}

func (f ParseCIDRFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var (
		cidr string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &cidr))
	if resp.Error != nil {
		return
	}

	cidrModel, err := netparse.ParseCIDR(cidr)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)
		return
	}

	result := ContainsIPFromCIDRModel(cidrModel)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
