# Examples

1. Create K8s Secret from template

    ```bash
    # Example
    $ PLATFORM_URL=https://artifactory.site.com/artifactory
    $ read -r PLATFORM_TOKEN
    <put-your-token-here + enter>
    $ cat examples/providerconfig/secret.yaml.tmpl | sed -e "s/y0ur-t0k3n/${PLATFORM_TOKEN}/g" -e "s^y0ur-url^${PLATFORM_URL}^g" > examples/providerconfig/secret.yaml
    ```
