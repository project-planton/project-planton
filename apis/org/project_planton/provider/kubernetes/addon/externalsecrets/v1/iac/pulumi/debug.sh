#!/bin/bash
go build -o build/module -gcflags "all=-N -l" .
exec dlv --listen=:2345 --headless=true --api-version=2 exec ./build/module
