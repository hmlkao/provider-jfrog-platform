#!/usr/bin/env bash

# Bump version of Go language

# Usage:
#   bump_go_version.sh <current_version> <new_version>

set -euo pipefail
IFS=$'\n\t'
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_version=${1}
if [[ -z "${current_version}" ]]; then
  echo "::error Current version is empty"
  exit 1
fi
current_version_short=${current_version%.*}
new_version=${2}
if [[ -z "${new_version}" ]]; then
  echo "::error New version is empty"
  exit 1
fi
new_version_short=${new_version%.*}

main () {
  # Use of sed -i'.tmp' is to make it compatible with both GNU and BSD sed. BSD sed requires an argument for the -i option, so we provide a temporary file extension '.tmp' to create a backup of the original file before making in-place edits. After the edits are made, we remove the temporary backup file.
  sed -i'.tmp' "s|GO_REQUIRED_VERSION ?= ${current_version_short}|GO_REQUIRED_VERSION ?= ${new_version_short}|" "${dir}/../Makefile" && rm "${dir}/../Makefile.tmp"
  sed -i'.tmp' "s|GO_VERSION: '${current_version_short}'|GO_VERSION: '${new_version_short}'|" "${dir}/../.github/workflows/bump_versions.yaml" && rm "${dir}/../.github/workflows/bump_versions.yaml.tmp"
  sed -i'.tmp' "s|GO_VERSION: '${current_version_short}'|GO_VERSION: '${new_version_short}'|" "${dir}/../.github/workflows/ci.yml" && rm "${dir}/../.github/workflows/ci.yml.tmp"
  sed -i'.tmp' "s|GO_VERSION: \"${current_version_short}\"|GO_VERSION: \"${new_version_short}\"|" "${dir}/../.github/workflows/e2e.yaml" && rm "${dir}/../.github/workflows/e2e.yaml.tmp"
  sed -i'.tmp' "s|default: '${current_version_short}'|default: '${new_version_short}'|" "${dir}/../.github/workflows/publish-provider-package.yml" && rm "${dir}/../.github/workflows/publish-provider-package.yml.tmp"
  sed -i'.tmp' "s|go ${current_version}|go ${new_version}|" "${dir}/../go.mod" && rm "${dir}/../go.mod.tmp"
}

main "$@"
