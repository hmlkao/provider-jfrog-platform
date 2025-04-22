# SAMLSettings example

1. Download Okta certifiate to `okta-saml-certificate.pem` file
2. Create K8s Secret from template

    ```bash
    # Example
    $ saml_certificate=$(cat okta-saml-certificate.pem | base64)
    $ cat examples/saml_settings/secret.yaml.tmpl | sed -e "s/y0ur-cert/${saml_certificate}/g" > examples/saml_settings/secret.yaml
    $ kubectl apply -f examples/saml_settings/secret.yaml
    ```
