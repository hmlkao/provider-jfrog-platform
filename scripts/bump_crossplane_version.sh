#!/usr/bin/env bash

# Bump version of Terraform provider

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_version=${1}
if [[ -z "$current_version" ]]; then
  echo "::error Current version is empty"
  exit 1
fi
new_version=${2}
if [[ -z "$new_version" ]]; then
  echo "::error New version is empty"
  exit 1
fi

main() {
  sed -i "s|CROSSPLANE_VERSION = ${current_version}|CROSSPLANE_VERSION = ${new_version}|" "$dir/../Makefile"
  sed -i "s|CROSSPLANE_CLI_VERSION ?= v${current_version}|CROSSPLANE_CLI_VERSION ?= v${new_version}|" "$dir/../Makefile"
}

main "$@"
