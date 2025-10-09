<!-- markdownlint-disable descriptive-link-text -->
# Provider JFrog Platform

`provider-jfrog-platform` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
JFrog Platform API.

The repo was created from the [crossplane/upjet-provider-template@7311f9f](https://github.com/crossplane/upjet-provider-template/tree/7311f9f9baa87f4431702ba209dffbc6067ce74b) template.

The provider is generated from Terraform provider [jfrog/platform v2.2.6](https://registry.terraform.io/providers/jfrog/platform/2.2.6/docs).

- [Provider JFrog Platform](#provider-jfrog-platform)
  - [Getting Started](#getting-started)
  - [Supported resources](#supported-resources)
  - [Adding new resource](#adding-new-resource)
  - [Build provider from scratch](#build-provider-from-scratch)
  - [Developing](#developing)
  - [JFrog Platform icon](#jfrog-platform-icon)
  - [Report a Bug](#report-a-bug)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/hmlkao/provider-jfrog-platform):

```bash
up ctp provider install hmlkao/provider-jfrog-platform:v0.5.0
```

Alternatively, you can use declarative installation:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-jfrog-platform
spec:
  package: hmlkao/provider-jfrog-platform:v0.5.0
EOF
```

Notice that in this example, the Provider resource is referencing a ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/hmlkao/provider-jfrog-platform).

## Supported resources

This is a list of all resources from [Terraform provider v2.2.6](https://registry.terraform.io/providers/jfrog/platform/2.2.6/docs).

The short group is `platform`, so the `apiGroup` is `platform.artifactory.jfrog.crossplane.io` for **cluster-scoped resources**.

The short group is `platform`, so the `apiGroup` is `platform.artifactory.jfrog.m.crossplane.io` for **namespace-scoped resources**.

| Resource                         | Supported                                                                                  | Kind                |
|----------------------------------|--------------------------------------------------------------------------------------------|---------------------|
| `platform_aws_iam_role`          | :heavy_check_mark:                                                                         | `AWSIAMRole`        |
| `platform_crowd_settings`        | :heavy_check_mark:                                                                         | `CrowdSettings`     |
| `platform_global_role`           | :heavy_check_mark:                                                                         | `GlobalRole`        |
| `platform_group`                 | :heavy_check_mark: ([Known Issues](./KNOWN_ISSUES.md#platform_group))                      | `Group`             |
| `platform_group_members`         | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                     |
| `platform_http_sso_settings`     | :heavy_check_mark:                                                                         | `HTTPSSOSettings`   |
| `platform_license`               | :x: ([Resource Import Not Implemented](./KNOWN_ISSUES.md#resource-import-not-implemented)) |                     |
| `platform_oidc_configuration`    | :heavy_check_mark:                                                                         | `OIDCConfiguration` |
| `platform_oidc_identity_mapping` | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                     |
| `platform_permission`            | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                     |
| `platform_reverse_proxy`         | :heavy_check_mark:                                                                         | `ReverseProxy`      |
| `platform_saml_settings`         | :heavy_check_mark:                                                                         | `SAMLSettings`      |
| `platform_scim_group`            | :heavy_check_mark: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                      | `SCIMGroup`         |
| `platform_scim_user`             | :heavy_check_mark: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                      | `SCIMUser`          |
| `platform_workers_service`       | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                     |

## Adding new resource

1. Add/Modify files related to the Terraform resource (check this reference [commit](https://github.com/hmlkao/provider-jfrog-platform/commit/a814281b232db283a6fa3846dce2e1aa8dc0e63a)), an example diff for human-modified files:

    ```diff
    $ git diff HEAD^^ HEAD^
    diff --git a/README.md b/README.md
    index 7389013..dfe425e 100644
    --- a/README.md
    +++ b/README.md
    @@ -59,7 +59,7 @@ | `platform_oidc_configuration`    | :x:                                                                                        |                  |
     | `platform_oidc_identity_mapping` | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
     | `platform_permission`            | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
    -| `platform_reverse_proxy`         | :x:                                                                                        |                  |
    +| `platform_reverse_proxy`         | :heavy_check_mark:                                                                         | `ReverseProxy`   |
     | `platform_saml_settings`         | :heavy_check_mark:                                                                         | `SAMLSettings`   |
     | `platform_scim_group`            | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |
     | `platform_scim_user`             | :x: ([Nested Schema](./KNOWN_ISSUES.md#nested-schema))                                     |                  |

    diff --git a/config/external_name.go b/config/external_name.go
    index ce80c1e..9ebaa51 100644
    --- a/config/external_name.go
    +++ b/config/external_name.go
    @@ -11,6 +11,7 @@ import "github.com/crossplane/upjet/v2/pkg/config"
    var ExternalNameConfigs = map[string]config.ExternalName{
            // Cannot use NameAsIdentifier, Name parameter can contain characters which are not allowed in Terraform resource name
            "platform_group":         config.IdentifierFromProvider,
    +       "platform_reverse_proxy": config.IdentifierFromProvider,
            "platform_saml_settings": config.IdentifierFromProvider,
    }

    diff --git a/config/provider.go b/config/provider.go
    index ef63761..a4318f5 100644
    --- a/config/provider.go
    +++ b/config/provider.go
    @@ -10,6 +10,7 @@ import (

            ujconfig "github.com/crossplane/upjet/v2/pkg/config"
            "github.com/hmlkao/provider-jfrog-platform/config/namespaced/group"
    +       reverseproxy "github.com/hmlkao/provider-jfrog-platform/config/namespaced/reverse_proxy"
            samlsettings "github.com/hmlkao/provider-jfrog-platform/config/namespaced/saml_settings"
    )

    @@ -37,6 +38,7 @@ func GetProvider() *ujconfig.Provider {
            for _, configure := range []func(provider *ujconfig.Provider){
                    // add custom config functions
                    group.Configure,
    +               reverseproxy.Configure,
                    samlsettings.Configure,
            } {
                    configure(pc)

    diff --git a/config/namespaced/reverse_proxy/config.go b/config/namespaced/reverse_proxy/config.go
    new file mode 100644
    index 0000000..3b77efb
    --- /dev/null
    +++ b/config/namespaced/reverse_proxy/config.go
    @@ -0,0 +1,25 @@
    +package reverseproxy
    +
    +import (
    +       "errors"
    +
    +       "github.com/crossplane/upjet/v2/pkg/config"
    +)
    +
    +// Configure configures individual resources by adding custom ResourceConfigurators.
    +func Configure(p *config.Provider) {
    +       p.AddResourceConfigurator("platform_reverse_proxy", func(r *config.Resource) {
    +               r.ShortGroup = "platform"
    +               r.Kind = "ReverseProxy"
    +               // Cannot use config.NameAsIdentifier variable because 'name' parameter can use characters which are invalid for Trerraform resource name
    +               //   Terraform resource name must start with a letter or underscore and may contain only letters, digits, underscores, and dashes.
    +               // Variable config.NameAsIdentifier is using IDAsExternalName func which tries to get the 'id' from the tfstate,
    +               // but there is no 'id' after the Terraform state refresh, so we specify custum function to get 'name'
    +               r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
    +                       if name, ok := tfstate["server_provider"].(string); ok && name != "" {
    +                               return name, nil
    +                       }
    +                       return "", errors.New("cannot find 'server_provider' in tfstate")
    +               }
    +       })
    +}

    diff --git a/examples/reverse_proxy/reverse_proxy.yaml b/examples/reverse_proxy/reverse_proxy.yaml
    new file mode 100644
    index 0000000..4250e93
    --- /dev/null
    +++ b/examples/reverse_proxy/reverse_proxy.yaml
    @@ -0,0 +1,11 @@
    +apiVersion: platform.jfrog.crossplane.io/v1alpha1
    +kind: ReverseProxy
    +metadata:
    +  name: my-reverse-proxy
    +spec:
    +  forProvider:
    +    serverProvider: NGINX
    +    internalHostname: localhost
    +    publicServerName: mydomain.com
    +  providerConfigRef:
    +    name: default
    ```

2. Generate the resource files

    ```bash
    make generate
    ```

3. Verify that the changes are reviewable

    ```bash
    make reviewable test
    ```

4. Create a PR with all files included

## Build provider from scratch

Check [`BUILD_FROM_SCRATCH.md`](https://github.com/hmlkao/provider-jfrog-artifactory/blob/main/BUILD_FROM_SCRATCH.md) for notes on how `provider-jfrog-artifactory` was built using the [crossplane/upjet tool](https://github.com/crossplane/upjet) step-by-step.

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

## JFrog Platform icon

The package icon was pulled from [JFrog Brand Guidelines](https://jfrog.com/brand-guidelines/).

The icon is stored in the [`extensions/icons/`](./extensions/icons/) folder according to the instructions [Add documentation, icons, and other assets to your package](https://docs.upbound.io/manuals/marketplace/packages#add-documentation-icons-and-other-assets-to-your-package).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/hmlkao/provider-jfrog-platform/issues).
