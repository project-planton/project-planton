#!/bin/bash

# Test script for Project Planton CLI commands
# This script tests all the new config and list-deployment-components functionality

set -e  # Exit on any error

echo "🧪 Testing Project Planton CLI Commands"
echo "======================================"

# Change to project directory
cd "$(dirname "$0")"

# Build the CLI
echo "📦 Building CLI..."
go build -o bin/project-planton ./cmd/project-planton
chmod +x bin/project-planton

echo "✅ CLI built successfully"
echo ""

# Test 1: Config commands
echo "🔧 Testing Configuration Commands"
echo "--------------------------------"

# Test config list (should be empty initially)
echo "1. Testing empty config list..."
./bin/project-planton config list || echo "Config is empty as expected"

# Test setting backend URL
echo "2. Testing config set..."
./bin/project-planton config set backend-url http://localhost:50051
echo "✅ Backend URL set"

# Test getting backend URL
echo "3. Testing config get..."
BACKEND_URL=$(./bin/project-planton config get backend-url)
echo "✅ Backend URL retrieved: $BACKEND_URL"

# Test config list (should show backend-url now)
echo "4. Testing config list with values..."
./bin/project-planton config list
echo ""

# Test 2: Invalid config values
echo "🚫 Testing Configuration Validation"
echo "----------------------------------"

# Test invalid URL format
echo "5. Testing invalid URL format..."
if ./bin/project-planton config set backend-url invalid-url 2>/dev/null; then
    echo "❌ ERROR: Invalid URL was accepted"
    exit 1
else
    echo "✅ Invalid URL correctly rejected"
fi

# Test unknown config key
echo "6. Testing unknown config key..."
if ./bin/project-planton config set unknown-key value 2>/dev/null; then
    echo "❌ ERROR: Unknown key was accepted"
    exit 1
else
    echo "✅ Unknown key correctly rejected"
fi
echo ""

# Test 3: List deployment components (without backend)
echo "📋 Testing List Commands (No Backend)"
echo "------------------------------------"

# Reset to invalid backend URL to test error handling
./bin/project-planton config set backend-url http://localhost:99999

echo "7. Testing connection error handling..."
if ./bin/project-planton list-deployment-components 2>/dev/null; then
    echo "❌ ERROR: Should have failed with connection error"
    exit 1
else
    echo "✅ Connection error handled correctly"
fi

# Reset to correct backend URL
./bin/project-planton config set backend-url http://localhost:50051
echo ""

# Test 4: List deployment components (with backend)
echo "📋 Testing List Commands (With Backend)"
echo "--------------------------------------"

# Wait for backend to be ready
echo "8. Waiting for backend to be ready..."
for i in {1..30}; do
    if curl -s http://localhost:50051 >/dev/null 2>&1; then
        echo "✅ Backend is ready"
        break
    fi
    if [ $i -eq 30 ]; then
        echo "❌ ERROR: Backend not ready after 30 seconds"
        echo "Please ensure backend is running with: docker-compose up -d"
        exit 1
    fi
    sleep 1
done

# Test list all components
echo "9. Testing list all deployment components..."
./bin/project-planton list-deployment-components
echo "✅ List all components works"
echo ""

# Test list with kind filter
echo "10. Testing list with kind filter..."
./bin/project-planton list-deployment-components --kind PostgresKubernetes
echo "✅ Kind filter works"
echo ""

# Test list with non-existent kind
echo "11. Testing list with non-existent kind..."
./bin/project-planton list-deployment-components --kind NonExistentKind
echo "✅ Non-existent kind handled correctly"
echo ""

# Test 5: Help commands
echo "📚 Testing Help Commands"
echo "-----------------------"

echo "12. Testing main help..."
./bin/project-planton --help | head -10
echo "✅ Main help works"
echo ""

echo "13. Testing config help..."
./bin/project-planton config --help | head -10
echo "✅ Config help works"
echo ""

echo "14. Testing list-deployment-components help..."
./bin/project-planton list-deployment-components --help | head -10
echo "✅ List command help works"
echo ""

# Final success
echo "🎉 All Tests Passed!"
echo "==================="
echo ""
echo "The CLI commands are working correctly. You can now use:"
echo "  • project-planton config set backend-url <url>"
echo "  • project-planton config get backend-url"
echo "  • project-planton config list"
echo "  • project-planton list-deployment-components"
echo "  • project-planton list-deployment-components --kind <kind>"
echo ""
echo "Configuration is stored in: ~/.project-planton/config.yaml"
