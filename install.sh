#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

REPO_URL="https://github.com/DevMatheusSilva/palantir"
BINARY_NAME="palantir"
INSTALL_DIR="/usr/local/bin"

echo -e "${BLUE}"
echo "🔮 =============================================="
echo "   Palantír - The Seeing Stone"
echo "   Installation Script"
echo "============================================== 🔮"
echo -e "${NC}"

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go 1.25.1 or higher first.${NC}"
    echo -e "${YELLOW}Visit: https://golang.org/dl/${NC}"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "${GREEN}✓ Go ${GO_VERSION} found${NC}"

if ! command -v code &> /dev/null; then
    echo -e "${YELLOW}⚠️  Warning: VSCode command 'code' not found in PATH${NC}"
    echo -e "${YELLOW}   Palantír requires VSCode to be installed and accessible via 'code' command${NC}"
    read -p "Continue anyway? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

echo -e "${BLUE}📦 Cloning Palantír repository...${NC}"
git clone --depth 1 "$REPO_URL" palantir 

cd palantir

echo -e "${BLUE}📚 Downloading dependencies...${NC}"
go mod download

echo -e "${BLUE}🔨 Building Palantír...${NC}"
go build -o "$BINARY_NAME" .

echo -e "${BLUE}📍 Installing to ${INSTALL_DIR}...${NC}"
if [ -w "$INSTALL_DIR" ]; then
    mv "$BINARY_NAME" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

cd ~
rm -rf "$TEMP_DIR"

echo -e "${GREEN}"
echo "✨ =============================================="
echo "   Installation Complete!"
echo "============================================== ✨"
echo -e "${NC}"

SHELL_RC=""
if [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
elif [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
fi

echo -e "${YELLOW}⚠️  IMPORTANT: Configure your environment variables${NC}"
echo ""
echo -e "Add these lines to your shell profile (${SHELL_RC}):"
echo ""
echo -e "${BLUE}export PALANTIR_ROOT_FOLDER=Projects${NC}"
echo -e "${BLUE}export PALANTIR_PROJECTS_DEPTH=1${NC}"
echo ""
echo -e "Adjust the values according to your project structure:"
echo -e "  • ${GREEN}PALANTIR_ROOT_FOLDER${NC}: Folder name where your projects are located (relative to HOME)"
echo -e "  • ${GREEN}PALANTIR_PROJECTS_DEPTH${NC}: How many levels deep your projects are"
echo ""
echo -e "Example project structure:"
echo -e "  ~/Projects/my-project → ${BLUE}DEPTH=1${NC}"
echo -e "  ~/Projects/work/my-project → ${BLUE}DEPTH=2${NC}"
echo ""
echo -e "${GREEN}After adding the variables, reload your shell:${NC}"
echo -e "  source ${SHELL_RC}"
echo ""
echo -e "${GREEN}Then try:${NC}"
echo -e "  ${BLUE}palantir list${NC}      - List all your projects"
echo -e "  ${BLUE}palantir open <name>${NC} - Open a project in VSCode"
echo ""
echo -e "${GREEN}🧙‍♂️ The Seeing Stone is now ready to serve you!${NC}"
