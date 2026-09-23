#!/usr/bin/env bash

# Bump version of Terraform provider

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_version=${1}
if [[ -z "${current_version}" ]]; then
  echo "::error Current version is empty"
  exit 1
fi
new_version=${2}
if [[ -z "${new_version}" ]]; then
  echo "::error New version is empty"
  exit 1
fi

main() {
  # Use of sed -i'.tmp' is to make it compatible with both GNU and BSD sed. BSD sed requires an argument for the -i option, so we provide a temporary file extension '.tmp' to create a backup of the original file before making in-place edits. After the edits are made, we remove the temporary backup file.
  sed -i'.tmp' "s|export TERRAFORM_PROVIDER_VERSION ?= ${current_version}|export TERRAFORM_PROVIDER_VERSION ?= ${new_version}|g" "${dir}/../Makefile" && rm "${dir}/../Makefile.tmp"
  sed -i'.tmp' "s|terraform-provider-platform_v${current_version}|terraform-provider-platform_v${new_version}|g" "${dir}/../Makefile" && rm "${dir}/../Makefile.tmp"
  sed -i'.tmp' "s|jfrog/platform v${current_version}|jfrog/platform v${new_version}|g" "${dir}/../README.md" && rm "${dir}/../README.md.tmp"
  sed -i'.tmp' "s|Terraform provider v${current_version}|Terraform provider v${new_version}|g" "${dir}/../README.md" && rm "${dir}/../README.md.tmp"
  sed -i'.tmp' "s|jfrog/platform/${current_version}|jfrog/platform/${new_version}|g" "${dir}/../README.md" && rm "${dir}/../README.md.tmp"
}

main "$@"
