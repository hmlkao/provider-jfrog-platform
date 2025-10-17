# Examples

1. Create K8s Secret from template

    ```bash
    # Example
    $ PLATFORM_URL=https://artifactory.site.com/artifactory
    $ read -r PLATFORM_TOKEN
    <paste-your-token-here + enter>
    $ cat examples/namespaced/clusterproviderconfig/secret.yaml.tmpl | sed -e "s/y0ur-t0k3n/${PLATFORM_TOKEN}/g" -e "s^y0ur-url^${PLATFORM_URL}^g" > examples/namespaced/clusterproviderconfig/secret.yaml
    $ kubectl -n crossplane-system apply -f examples/namespaced/clusterproviderconfig/secret.yaml
    ```
