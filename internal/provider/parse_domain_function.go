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

var _ function.Function = ParseDomainFunction{}

type ParseDomainFunction struct{}

type parseDomainFunctionReturnModel struct {
	Domain    string `tfsdk:"domain"`
	Host      string `tfsdk:"host"`
	Manager   string `tfsdk:"manager"`
	SLD       string `tfsdk:"sld"`
	Subdomain string `tfsdk:"subdomain"`
	TLD       string `tfsdk:"tld"`
}

func NewParseDomainFunction() function.Function {
	return ParseDomainFunction{}
}

func FromDomainModel(d *netparse.DomainModel) parseDomainFunctionReturnModel {
	return parseDomainFunctionReturnModel{
		Domain:    d.Domain,
		Host:      d.Host,
		Manager:   d.Manager,
		SLD:       d.SLD,
		Subdomain: d.Subdomain,
		TLD:       d.TLD,
	}
}

func (f ParseDomainFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "parse_domain"
}

func (f ParseDomainFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             domainMarkdownDescription,
		MarkdownDescription: domainMarkdownDescription,
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "host",
				MarkdownDescription: hostAttrMarkdownDescription,
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"domain":    types.StringType,
				"host":      types.StringType,
				"manager":   types.StringType,
				"sld":       types.StringType,
				"subdomain": types.StringType,
				"tld":       types.StringType,
			},
		},
	}
}

func (f ParseDomainFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var (
		host string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &host))
	if resp.Error != nil {
		return
	}

	DomainModel, err := netparse.ParseDomain(host)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)
		return
	}

	result := FromDomainModel(DomainModel)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
