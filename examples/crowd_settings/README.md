# CrowdSettings example

1. Create K8s Secret from template

    ```bash
    # Example
    $ read -r pass
    <paste-your-password-here + enter>
    $ cat examples/crowd_settings/secret.yaml.tmpl | sed -e "s/y0ur-pass/${pass}/g" > examples/crowd_settings/secret.yaml
    $ kubectl -n crossplane-system apply -f examples/crowd_settings/secret.yaml
    ```
