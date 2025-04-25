#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_version=$(grep 'export TERRAFORM_PROVIDER_VERSION' < "$dir/../Makefile" | cut -d' ' -f4)
new_version=$(curl -sSL https://api.github.com/repos/jfrog/terraform-provider-platform/releases/latest | jq -r '.tag_name' | tr -d 'v')

sed -i '' "s|export TERRAFORM_PROVIDER_VERSION ?= ${current_version}|export TERRAFORM_PROVIDER_VERSION ?= ${new_version}|" "$dir/../Makefile"
sed -i '' "s|terraform-provider-platform_v${current_version}|terraform-provider-platform_v${new_version}|" "$dir/../Makefile"
sed -i '' "s|jfrog/platform v${current_version}|jfrog/platform v${new_version}|g" "$dir/../README.md"
sed -i '' "s|Terraform provider v${current_version}|Terraform provider v${new_version}|g" "$dir/../README.md"
sed -i '' "s|jfrog/platform/${current_version}|jfrog/platform/${new_version}|" "$dir/../README.md"
