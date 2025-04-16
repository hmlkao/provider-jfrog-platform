# Known issues

- [Known issues](#known-issues)
  - [Resource Import Not Implemented](#resource-import-not-implemented)
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

## `platform_group`

If you omit parameter `useGroupMembersResource`, the provider will reconcile this resource still again and again, because the terraform refresh sets this parameter to `null` and terraform apply wants to set it to `true` (default value) again.
