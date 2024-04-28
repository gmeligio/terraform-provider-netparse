# Terraform Provider: netparse

## Using the provider

Official documentation on how to use this provider can be found on the
[Terraform Registry](https://registry.terraform.io/providers/gmeligio/url/latest/docs).
In case of specific questions or discussions, please use the

We also provide:

- [Contributing](.github/CONTRIBUTING.md) guidelines in case you want to help this project

The remainder of this document will focus on the development aspects of the provider.

## Development

### Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

### Building The Provider

1. Clone the repository.
1. Enter the repository directory.
1. Build the provider using the Go `install` command:

```shell
# This will build the provider and put the provider binary in the `$GOPATH/bin` directory
go install
```

### Generating documentation

To generate or update documentation, run `go generate`.

### Testing

In order to run the full suite of Acceptance tests, run `make testacc`.

```shell
make testacc
```

### TODO

1. Migrate to GitHub action [ghaction-terraform-provider-release](https://github.com/hashicorp/ghaction-terraform-provider-release)
1. Refer in documentation: [What is a URL](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL)
1. Use Table Driven tests for Acceptance tests

    - <https://go.dev/wiki/TableDrivenTests>
    - <https://github.com/northwood-labs/terraform-provider-corefunc/blob/main/testfixtures/url_parse.go>
