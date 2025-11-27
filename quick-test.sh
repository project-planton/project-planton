#!/bin/bash

echo "🚀 Quick CLI Test"
echo "================"

cd "$(dirname "$0")"

# Test config
echo "Setting up config..."
go run . config set backend-url http://localhost:50051

echo -e "\nGetting config..."
go run . config get backend-url

echo -e "\nListing all configs..."
go run . config list

echo -e "\nTesting deployment components list..."
go run . list-deployment-components

echo -e "\nTesting with filter..."
go run . list-deployment-components --kind PostgresKubernetes

echo -e "\n✅ Quick test complete!"
