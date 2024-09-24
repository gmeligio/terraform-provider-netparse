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

var _ function.Function = ParseUrlFunction{}

type ParseUrlFunction struct{}

type parseUrlFunctionReturnModel struct {
	Authority   string `tfsdk:"authority"`
	Protocol    string `tfsdk:"protocol"`
	Scheme      string `tfsdk:"scheme"`
	Credentials string `tfsdk:"credentials"`
	Username    string `tfsdk:"username"`
	Password    string `tfsdk:"password"`
	Host        string `tfsdk:"host"`
	Port        string `tfsdk:"port"`
	Path        string `tfsdk:"path"`
	Search      string `tfsdk:"search"`
	Query       string `tfsdk:"query"`
	Hash        string `tfsdk:"hash"`
	Fragment    string `tfsdk:"fragment"`
}

func NewParseUrlFunction() function.Function {
	return ParseUrlFunction{}
}

func FromUrlModel(u *netparse.UrlModel) parseUrlFunctionReturnModel {
	return parseUrlFunctionReturnModel{
		Authority:   u.Authority,
		Protocol:    u.Protocol,
		Scheme:      u.Scheme,
		Credentials: u.Credentials,
		Username:    u.Username,
		Password:    u.Password,
		Host:        u.Host,
		Port:        u.Port,
		Path:        u.Path,
		Search:      u.Search,
		Query:       u.Query,
		Hash:        u.Hash,
		Fragment:    u.Fragment,
	}
}

func (f ParseUrlFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "parse_url"
}

func (f ParseUrlFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             urlMarkdownDescription,
		MarkdownDescription: urlMarkdownDescription,
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "url",
				MarkdownDescription: urlAttributeMarkdownDescription,
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"authority":   types.StringType,
				"scheme":      types.StringType,
				"protocol":    types.StringType,
				"credentials": types.StringType,
				"username":    types.StringType,
				"password":    types.StringType,
				"host":        types.StringType,
				"port":        types.StringType,
				"path":        types.StringType,
				"search":      types.StringType,
				"query":       types.StringType,
				"fragment":    types.StringType,
				"hash":        types.StringType,
			},
		},
	}
}

func (f ParseUrlFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var (
		url string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &url))
	if resp.Error != nil {
		return
	}

	urlModel, err := netparse.ParseUrl(url)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)
		return
	}

	result := FromUrlModel(urlModel)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
