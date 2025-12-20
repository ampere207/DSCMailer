#!/bin/bash

# DSC-Mailer Runner Script
# This script checks prerequisites and starts the DSC-Mailer server

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "🚀 Starting DSC-Mailer..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go 1.24.4 or higher.${NC}"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "${GREEN}✓${NC} Found Go version: $GO_VERSION"

# Check if .env file exists
if [ ! -f .env ]; then
    echo -e "${YELLOW}⚠️  Warning: .env file not found!${NC}"
    echo "Creating a template .env file..."
    cat > .env << EOF
SMTP_USER=your-email@example.com
SMTP_PASS=your-app-password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
EOF
    echo -e "${YELLOW}Please edit the .env file with your SMTP credentials before running again.${NC}"
    exit 1
fi

echo -e "${GREEN}✓${NC} Found .env file"

# Check if go.mod exists
if [ ! -f go.mod ]; then
    echo -e "${RED}❌ go.mod not found. Are you in the correct directory?${NC}"
    exit 1
fi

# Download dependencies if needed
echo "📦 Checking dependencies..."
go mod download
echo -e "${GREEN}✓${NC} Dependencies ready"

# Create necessary directories
mkdir -p uploads output
echo -e "${GREEN}✓${NC} Created necessary directories"

# Run the server
echo ""
echo "=========================================="
echo "  Starting DSC-Mailer Server"
echo "=========================================="
echo ""

go run cmd/server/main.go
