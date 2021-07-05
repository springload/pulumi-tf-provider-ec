# Pulumi Bridge Provider for Elastic Cloud

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/xyx

or `yarn`:

    $ yarn add @pulumi/xyx

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_xyx

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/springload/pulumi-tf-provider-ec/sdk/go/...

## Configuration

The following configuration points are available for the `ec` provider:

-   `ec:apiKey` (environment: `XYZ_API_KEY`) - the API key for `ec`
-   `ec:region` (environment: `XYZ_REGION`) - the region in which to deploy resources

## Update terraform provider version

To update the underlying terraform provider version:

1. Go to provider and change the version in go.mod
2. Run `go mod download && go mod tidy`
3. Commit the changes, push.
4. Tag the commit, push. We use the same version as the underlying EC provider, i.e. if the official EC provider has version 0.1.1 we use 0.1.1 too.
