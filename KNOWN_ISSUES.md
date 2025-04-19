# Known issues

- [Known issues](#known-issues)
  - [Resource Import Not Implemented](#resource-import-not-implemented)
  - [Nested Schema](#nested-schema)
  - [`platform_group`](#platform_group)

## Resource Import Not Implemented

Terraform resource doesn't allow import and it's not possible to set external name, because there is no id set in terraform state, e.g.:

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

Terraform resource contains Nested Schema and upjet is not able to generate provider, it fails with error:

```bash
# Example
$ make generate
...
panic: cannot generate crd for resource platform_permission: cannot build types for Permission: cannot build the Types for resource "platform_permission": cannot infer type from schema of field artifact: invalid schema type TypeInvalid
```

There is [opened issue](https://github.com/crossplane/upjet/issues/372) on crossplane/upjet, no workaround here for now.

## `platform_group`

If you omit parameter `useGroupMembersResource`, the provider will reconcile this resource still again and again, because the terraform refresh sets this parameter to `null` and terraform apply wants to set it to `true` (default value) again.
