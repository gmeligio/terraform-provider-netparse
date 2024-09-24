// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/gmeligio/terraform-provider-netparse/internal/netparse"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = ContainsIPFunction{}

type ContainsIPFunction struct{}

type ContainsIPFunctionReturnModel struct {
	IP      string `tfsdk:"ip"`
	Network string `tfsdk:"network"`
}

func NewContainsIPFunction() function.Function {
	return ContainsIPFunction{}
}

func ContainsIPFromCIDRModel(d *netparse.CidrModel) ContainsIPFunctionReturnModel {
	return ContainsIPFunctionReturnModel{
		IP:      d.IP,
		Network: d.Network,
	}
}

func (f ContainsIPFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "contains_ip"
}

func (f ContainsIPFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             containsIPMarkdownDescription,
		MarkdownDescription: containsIPMarkdownDescription,
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "ip",
				MarkdownDescription: ipAttrMarkdownDescription,
			},
			function.StringParameter{
				Name:                "network",
				MarkdownDescription: networkAttrMarkdownDescription,
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f ContainsIPFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var (
		network string
		ip      string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &network, &ip))
	if resp.Error != nil {
		return
	}

	contains, err := netparse.ContainsIP(network, ip)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, contains))
}
