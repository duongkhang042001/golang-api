#!/bin/bash

# ANSI color codes
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No color

# Function to display menu
show_menu() {
    clear
    echo -e "${GREEN}==================================${NC}"
    echo -e "${GREEN}   Multi-Service Golang Project   ${NC}"
    echo -e "${GREEN}==================================${NC}"
    echo -e "${YELLOW}1. Build API${NC}"
    echo -e "${YELLOW}2. Build Worker${NC}"
    echo -e "${YELLOW}3. Build Gateway${NC}"
    echo -e "${YELLOW}4. Run API${NC}"
    echo -e "${YELLOW}5. Run Worker${NC}"
    echo -e "${YELLOW}6. Run Gateway${NC}"
    echo -e "${YELLOW}7. Clean Projects${NC}"
    echo -e "${YELLOW}8. Exit${NC}"
    echo -e "${GREEN}==================================${NC}"
}

while true; do
    show_menu
    read -p "Select an option [1-8]: " choice
    case $choice in
        1)
            go build  -o bin/api cmd/api/main.go
            cd ..
            echo -e "${BLUE}API built successfully.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
        2)
            cd worker
            go build -o worker
            cd ..
            echo -e "${BLUE}Worker built successfully.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
        3)
            cd gateway
            go build -o gateway
            cd ..
            echo -e "${BLUE}Gateway built successfully.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
        4)
            cd api
            ./api
            cd ..
            ;;
        5)
            cd worker
            ./worker
            cd ..
            ;;
        6)
            cd gateway
            ./gateway
            cd ..
            ;;
        7)
            rm -f api/api worker/worker gateway/gateway
            echo -e "${BLUE}Projects cleaned.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
        8)
            echo -e "${BLUE}Exiting.${NC}"
            exit 0
            ;;
        *)
            echo -e "${RED}Invalid option. Please choose a valid option.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
    esac
done
