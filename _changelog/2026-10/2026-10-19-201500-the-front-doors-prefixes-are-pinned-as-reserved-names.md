# The Front Door's Prefixes Are Pinned as Reserved Names

**Date**: October 19, 2026

## Summary

The operator's front-door route table now has a test that pins its handle-shaped first path segments (`idp`, `rpc`, `storage`, `webhooks`). Organizations and people live at the root of an installation's address, so every prefix the front door sends to a backend before the console is a name no organization or person may hold. The platform reserves each one, and its reserved-names guard pins the same list because it cannot read this repository. A new prefix now fails this test until it is reserved there too.

## Problem Statement

A new front-door prefix could be added here without the platform reserving it, and an organization or person who held that name would then never be reachable at their own address.

## Solution

`front_door_routes_test.go` derives the first segments from `FrontDoorRoutes()` itself and compares them with the pinned list. Its failure message names both places to update.

## Impact

None for a running installation: no route changes. It is a guard for whoever adds the next route.
