# Known issues

- [Known issues](#known-issues)
  - [Resource Import Not Implemented](#resource-import-not-implemented)
  - [Nested Schema](#nested-schema)
  - [`check-diff` is failing in CI workflow](#check-diff-is-failing-in-ci-workflow)
  - [`platform_group`](#platform_group)
  - [`platform_scim_user`](#platform_scim_user)

## Resource Import Not Implemented

This Terraform resource does not support import, and it is not possible to set an external name because no ID is set in the Terraform state, for example:

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
> The issue crossplane/upjet#372 has been resolved, so nested schemas are no longer a problem.

This Terraform resource contains a nested schema, and Upjet was previously unable to generate a provider; it failed with the following error:

```bash
# Example
$ make generate
...
panic: cannot generate crd for resource platform_permission: cannot build types for Permission: cannot build the Types for resource "platform_permission": cannot infer type from schema of field artifact: invalid schema type TypeInvalid
```

## `check-diff` is failing in CI workflow

A new Terraform provider version was probably merged into the repository, but you still have an older one locally.

Remove the `.work` folder and run `make generate` (the new version will be downloaded) to fix the issue.

```shell
rm -rf .work
make generate
```

## `platform_group`

If you omit the `useGroupMembersResource` parameter, the provider will reconcile this resource repeatedly because a Terraform refresh sets this parameter to `null`, and `terraform apply` then attempts to set it back to `true` (its default value).

## `platform_scim_user`

Refreshing this resource ends with an error:

```text
platform_scim_user.my-scim-user: Refreshing state...
╷
│ Error: Unable to Refresh Resource
│
│   with platform_scim_user.my-scim-user,
│   on main.tf.json line 1, in resource.platform_scim_user.my-scim-user:
│    1: {"provider":{"platform":{"access_token":"REDADTED","url":"REDACTED"}},"resource":{"platform_scim_user":{"my-scim-user":{"emails":[{"primary":true,"value":"test@tempurl.org"}],"lifecycle":{"prevent_destroy":true},"username":"test@tempurl.org"}}},"terraform":{"required_providers":{"platform":{"source":"jfrog/platform","version":"2.2.6"}}}}
│
│ An unexpected error occurred while attempting to refresh resource state. Please retry the operation or report this issue to the provider developers.
│
│ Error: test@tempurl.org isn't found
╵
```

The problem is that the provider tries to refresh the resource first and then run `terraform apply`. This means the apply never happens because the refresh fails.

If the user exists, it works well and user is removed properly when is the resource deleted - so, the SCIM user cannot be managed by Crossplane now.
