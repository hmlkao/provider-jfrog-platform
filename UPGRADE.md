# Upgrade instructions

Follow these instructions for breaking changes

## `v0.5.0`

The scope for `providerconfigusages.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Remove all `providerconfigusages.artifactory.jfrog.m.crossplane.io` resources
4. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources
5. Remove CRD `providerconfigusages.artifactory.jfrog.m.crossplane.io`
6. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`
7. Upgrade the provider to the `v0.5.0` version
8. Check if the provider started (is in `Running` state)
9. Create some resource using this provider to ensure it works properly
