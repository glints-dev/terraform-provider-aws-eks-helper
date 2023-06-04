# Terraform Provider AWS EKS Helper

This repository provides helper data sources for use with [Amazon EKS](https://aws.amazon.com/eks/).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `build` command: 
```sh
$ go build
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

You can use the provider via the [Terraform provider registry](https://registry.terraform.io/providers/glints/aws-eks-helper).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go build`.

To generate or update documentation, run `go generate`.
