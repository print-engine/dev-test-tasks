#!/bin/bash

# Check if script is being sourced
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    echo "❌ ERROR: This script must be sourced, not executed directly!"
    echo ""
    echo "Please run it like this:"
    echo "  source setup.sh"
    echo ""
    echo "Or:"
    echo "  . setup.sh"
    exit 1
fi

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

clear

echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║                                                            ║${NC}"
echo -e "${BLUE}║     Senior Backend Developer Assessment Setup              ║${NC}"
echo -e "${BLUE}║     WebSocket Notifications Service in Go                  ║${NC}"
echo -e "${BLUE}║                                                            ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${YELLOW}Welcome to the technical assessment!${NC}"
echo ""
echo "This setup script will configure your environment with the"
echo "Anthropic API key needed for Claude Code assistance."
echo ""
echo -e "${YELLOW}Please paste the API key provided by the interviewer:${NC}"
read -r ANTHROPIC_API_KEY

# Validate that something was entered
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo ""
    echo -e "${YELLOW}⚠️  No API key provided. Please run the setup again.${NC}"
    return 1 2>/dev/null || exit 1
fi

# Export the API key
export ANTHROPIC_API_KEY

echo ""

echo -e "${GREEN}✓ API key configured successfully!${NC}"
echo ""
echo -e "${GREEN}Verifying... API key is set in your environment.${NC}"
echo ""
echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║                    You're All Set! 🚀                      ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${GREEN}Next steps:${NC}"
echo ""
echo -e "  1. Type: ${YELLOW}claude${NC}"
echo -e "  2. Start vibing! ✨"
echo ""
echo -e "${BLUE}Tips:${NC}"
echo -e "  • Read the ${YELLOW}TASK.md${NC} file for requirements"
echo -e "  • The ${YELLOW}CLAUDE.md${NC} file has helpful context"
echo -e "  • Ask Claude questions - this is about AI collaboration!"
echo -e "  • Focus on core features first, then stretch goals"
echo ""
echo -e "${GREEN}Good luck! 🎯${NC}"
echo ""