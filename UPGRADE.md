# Upgrade instructions

Follow these instructions for breaking changes.

## `v0.5.0`

The scope for `providerconfigusages.platform.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.platform.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `spec.providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Remove all `providerconfigusages.platform.jfrog.m.crossplane.io` resources
4. Remove all `providerconfigs.platform.jfrog.m.crossplane.io` resources
5. Remove CRD `providerconfigusages.platform.jfrog.m.crossplane.io`
6. Remove CRD `providerconfigs.platform.jfrog.m.crossplane.io`
7. Upgrade the provider to the `v0.5.0` version
8. Check if is the provider healthy
9. Create some resource using this provider to ensure it works properly
