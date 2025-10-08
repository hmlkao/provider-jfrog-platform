# Known issues

- [Known issues](#known-issues)
  - [Resource Import Not Implemented](#resource-import-not-implemented)
  - [Nested Schema](#nested-schema)
  - [`check-diff` is failing in CI workflow](#check-diff-is-failing-in-ci-workflow)
  - [`platform_group`](#platform_group)

## Resource Import Not Implemented

The Terraform resource doesn't allow import and it's not possible to set an external name, because there is no ID set in the terraform state, e.g.:

```bash
# Example
$ terraform import platform_group_members.my-group-members 'My Crossplane Group'
platform_group_members.my-group-members: Importing from ID "My Crossplane Group"...
╷
│ Error: Resource Import Not Implemented
│
│ This resource does not support import. Please contact the provider developer for additional information.
╵
```

## Nested Schema

> [!NOTE]
> The issue [crossplane/upjet#372](https://github.com/crossplane/upjet/issues/372) was solved and nested schema is not a problem anymore

The Terraform resource contains a Nested Schema and upjet is not able to generate a provider, it fails with an error:

```bash
# Example
$ make generate
...
panic: cannot generate crd for resource platform_permission: cannot build types for Permission: cannot build the Types for resource "platform_permission": cannot infer type from schema of field artifact: invalid schema type TypeInvalid
```

There is an [open issue](https://github.com/crossplane/upjet/issues/372) on crossplane/upjet, with no workaround available for now.

## `check-diff` is failing in CI workflow

New Terraform provider version was probably merged into the repository, but you still have an old one on your localhost.

Just remove `.work` folder and run `make generate` (new version will be downloaded) to fix this issue.

```shell
rm -rf .work
make generate
```

## `platform_group`

If you omit the parameter `useGroupMembersResource`, the provider will reconcile this resource repeatedly, because the terraform refresh sets this parameter to `null` and terraform apply wants to set it to `true` (default value) again.
