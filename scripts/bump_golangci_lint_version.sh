#!/usr/bin/env bash

# Replace version of golangci-lint

# Usage:
#   bump_golangci_lint_version.sh <new_version>

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

new_version=${1}
if [[ -z "${new_version}" ]]; then
  echo "::error New version is empty"
  exit 1
fi

main () {
  # Use of sed -i'.tmp' is to make it compatible with both GNU and BSD sed. BSD sed requires an argument for the -i option, so we provide a temporary file extension '.tmp' to create a backup of the original file before making in-place edits. After the edits are made, we remove the temporary backup file.
  sed -i'.tmp' "s|GOLANGCILINT_VERSION ?= .*|GOLANGCILINT_VERSION ?= ${new_version}|" "${dir}/../Makefile" && rm "${dir}/../Makefile.tmp"
  sed -i'.tmp' "s|GOLANGCI_VERSION: '.*'|GOLANGCI_VERSION: 'v${new_version}'|" "${dir}/../.github/workflows/ci.yml" && rm "${dir}/../.github/workflows/ci.yml.tmp"
}

main "$@"
