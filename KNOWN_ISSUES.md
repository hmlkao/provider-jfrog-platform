# Known issues

- [Known issues](#known-issues)
  - [`platform_group`](#platform_group)

## `platform_group`

When you omit parameter `useGroupMembersResource`, the provider will reconcile this resource still again and again, because the terraform refresh sets this parameter to `null` and terraform apply wants to set it to `true` (default value) again.
