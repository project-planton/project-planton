# Cloudflare Pulumi Provider v6 Migration Guide

## Overview

This document tracks the migration from Cloudflare Pulumi Provider v5 to v6.10.1. The v6 release includes breaking API changes that require code updates.

## Status

### ✅ All Modules Migrated and Working
- **CloudflareWorker** ✅
- **CloudflareDnsZone** ✅
- **CloudflareLoadBalancer** ✅
- **CloudflareZeroTrustAccessApplication** ✅
- **CloudflareR2Bucket** ✅
- **CloudflareD1Database** ✅
- **CloudflareKvNamespace** ✅

## Breaking Changes by Module

### ✅ CloudflareWorker (FIXED)

**Changes:**
1. `WorkerScript` → `WorkersScript` (resource name)
2. `WorkerRoute` → `WorkersRoute` (resource name)
3. Bindings unified into single array with discriminator `Type` field
4. Field changes in `WorkersScriptArgs`:
   - `Name` → `ScriptName`
   - `PlainTextBindings` → `Bindings` (with `type: "plain_text"`)
   - `KvNamespaceBindings` → `Bindings` (with `type: "kv_namespace"`)
5. Field changes in `WorkersRouteArgs`:
   - `ScriptName` → `Script`
6. AWS S3 data source:
   - `LookupBucketObject` → `GetObject` (not Cloudflare-specific, but updated)

**Files Updated:**
- `cloudflareworker/v1/iac/pulumi/module/worker_script.go`
- `cloudflareworker/v1/iac/pulumi/module/route.go`

**Binding Examples:**

```go
// Plain text binding (environment variable)
cloudfl.WorkersScriptBindingArgs{
    Name: pulumi.String("MY_VAR"),
    Type: pulumi.String("plain_text"),
    Text: pulumi.String("value"),
}

// KV namespace binding
cloudfl.WorkersScriptBindingArgs{
    Name:        pulumi.String("MY_KV"),
    Type:        pulumi.String("kv_namespace"),
    NamespaceId: pulumi.String("namespace-id"),
}
```

### ✅ CloudflareDnsZone (FIXED)

**File:** `cloudflarednszone/v1/iac/pulumi/module/dns_zone.go`

**Changes:**
- `AccountId` → `Account` (nested structure with `Id` field)
- `Zone` → `Name`
- `Plan` → Removed (now managed separately via zone settings)

**Fixed Code:**
```go
zoneArgs := &cloudflare.ZoneArgs{
    Account: cloudflare.ZoneAccountArgs{
        Id: pulumi.String(accountId),
    },
    Name:   pulumi.String(zoneName),
    Paused: pulumi.BoolPtr(paused),
}
```

**Note:** Plan field was completely removed in v6 and must be managed via separate API or account configuration.

### ✅ CloudflareLoadBalancer (FIXED)

**File:** `cloudflareloadbalancer/v1/iac/pulumi/module/load_balancer.go`

**Changes:**
- `DefaultPoolIds` → `DefaultPools` (same type, just renamed)
- `FallbackPoolId` → `FallbackPool` (same type, just renamed)

**Fixed Code:**
```go
&cloudflare.LoadBalancerArgs{
    ZoneId:       pulumi.String(zoneId),
    Name:         pulumi.String(hostname),
    DefaultPools: pulumi.StringArray{poolId},
    FallbackPool: poolId,
    // ... other fields
}
```

### ✅ CloudflareZeroTrustAccessApplication (FIXED)

**File:** `cloudflarezerotrustaccessapplication/v1/iac/pulumi/module/application.go`

**Changes:**

1. **AccessPolicyInclude structure:**
   - `Emails` (array) → `Email` (single nested object per email)
   - `Groups` (array) → `Group` (single nested object per group)

2. **AccessPolicyArgs:**
   - `ApplicationId` → Removed (policies are now account-level)
   - `ZoneId` → Removed (replaced with `AccountId`)
   - `Precedence` → Removed
   - Added: `AccountId` (required)

3. **Account ID Lookup:**
   - Added zone lookup to get account ID: `zone.Account().Id()`

**Fixed Code:**
```go
// Lookup zone to get account ID
zone := cloudflare.LookupZoneOutput(ctx, cloudflare.LookupZoneOutputArgs{
    ZoneId: pulumi.String(zoneId),
}, pulumi.Provider(cloudflareProvider))

accountId := zone.Account().Id()

// Email includes (one per email)
for _, email := range allowedEmails {
    includeBlocks = append(includeBlocks, &cloudflare.AccessPolicyIncludeArgs{
        Email: &cloudflare.AccessPolicyIncludeEmailArgs{
            Email: pulumi.String(email),
        },
    })
}

// Group includes (one per group)
for _, groupId := range allowedGroups {
    includeBlocks = append(includeBlocks, &cloudflare.AccessPolicyIncludeArgs{
        Group: &cloudflare.AccessPolicyIncludeGroupArgs{
            Id: pulumi.String(groupId),
        },
    })
}

// Create policy with account ID
&cloudflare.AccessPolicyArgs{
    AccountId: accountId,
    Name:      pulumi.String("policy-name"),
    Decision:  pulumi.String("allow"),
    Includes:  includeBlocks,
    Requires:  requireBlocks,
}
```

## Migration Complete! 🎉

All Cloudflare provider modules have been successfully migrated to v6.10.1.

### Summary of Changes

1. **Updated go.mod:** `pulumi-cloudflare/sdk/v6 v6.10.1`
2. **Fixed 7 modules** with breaking API changes
3. **All imports updated** from v5 to v6
4. **All compilation errors resolved**
5. **All tests passing**

### Common Patterns in v6

- **Field name changes:** Many fields were renamed (e.g., `DefaultPoolIds` → `DefaultPools`)
- **Nested structures:** Account and identity fields moved to nested objects
- **Unified bindings:** Worker bindings consolidated into single array with discriminator
- **Account-level resources:** Some zone-level resources moved to account-level (e.g., Access Policies)
- **Removed fields:** Some fields like `Plan` in Zone were removed entirely

### Reference Resources

- [Pulumi Cloudflare Provider Docs](https://www.pulumi.com/registry/packages/cloudflare/)
- [Cloudflare API Documentation](https://developers.cloudflare.com/api/)
- [Pulumi Go SDK Documentation](https://pkg.go.dev/github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare)

### Next Steps

1. **Test deployments** with all updated modules
2. **Update API token permissions** if needed (see `API_TOKEN_PERMISSIONS.md`)
3. **Monitor for any runtime issues** during first deployments
4. **Update documentation** as needed based on real-world usage

