# Provider JFrog Platform

`provider-jfrog-platform` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
JFrog Platform API.

The repo was created from [crossplane/upjet-provider-template@7311f9f](https://github.com/crossplane/upjet-provider-template/tree/7311f9f9baa87f4431702ba209dffbc6067ce74b) template.

Provider is generated from Terraform provider [jfrog/platform v2.2.2](https://registry.terraform.io/providers/jfrog/platform/2.2.2/docs).

- [Provider JFrog Platform](#provider-jfrog-platform)
  - [Getting Started](#getting-started)
  - [Supported resources](#supported-resources)
  - [Build provider from scratch](#build-provider-from-scratch)
  - [Developing](#developing)
  - [Report a Bug](#report-a-bug)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-jfrog-platform):

```bash
up ctp provider install hmlkao/provider-jfrog-platform:v0.1.0
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-jfrog-platform
spec:
  package: hmlkao/provider-jfrog-platform:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/hmlkao/provider-jfrog-platform).

## Supported resources

List of all resources of [Terraform provider v2.2.2](https://registry.terraform.io/providers/jfrog/platform/2.2.2/docs).

| Resource                         | Supported                                                                                  | Kind             |
|----------------------------------|--------------------------------------------------------------------------------------------|------------------|
| `platform_aws_iam_role`          | :x:                                                                                        |                  |
| `platform_crowd_settings`        | :x:                                                                                        |                  |
| `platform_global_role`           | :x:                                                                                        |                  |
| `platform_group`                 | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#platform_group))                      | `Group`          |
| `platform_group_members`         | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                  |
| `platform_http_sso_settings`     | :x:                                                                                        |                  |
| `platform_license`               | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                  |
| `platform_oidc_configuration`    | :x:                                                                                        |                  |
| `platform_oidc_identity_mapping` | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
| `platform_permission`            | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
| `platform_reverse_proxy`         | :x:                                                                                        |                  |
| `platform_saml_settings`         | :heavy_check_mark:                                                                         | `SAMLSettings`   |
| `platform_scim_group`            | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
| `platform_scim_user`             | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
| `platform_workers_service`       | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |

## Build provider from scratch

Check [`BUILD_FROM_SCRATCH.md`]([./BUILD_FROM_SCRATCH.md](https://github.com/hmlkao/provider-artifactory/blob/main/BUILD_FROM_SCRATCH.md)) for notes how was `provider-artifactory` built using [crossplane/upjet tool](https://github.com/crossplane/upjet) step-by-step.

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/hmlkao/provider-jfrog-platform/issues).
