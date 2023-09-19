#!/bin/bash

# ANSI color codes
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No color

ROOT_DIR=`pwd`
OUT_DIR=bin

GOOS=linux 
GOARCH=amd64

# Function to display menu
show_menu() {
    clear
    echo -e "${GREEN}==================================${NC}"
    echo -e "${GREEN}   Multi-Service Golang Project   ${NC}"
    echo -e "${GREEN}==================================${NC}"
    echo -e "${YELLOW}1. Build API${NC}"
    echo -e "${YELLOW}2. Exit${NC}"
    echo -e "${GREEN}==================================${NC}"
}

# Function to create folder
create_folder() {
    if [ ! -d ${ROOT_DIR}/${OUT_DIR} ]; then
        mkdir ${ROOT_DIR}/${OUT_DIR}
    fi
}

while true; do
    show_menu
    create_folder
    read -p "Select an option [1-2]: " choice
    case $choice in
        1)
            cd ${ROOT_DIR}/cmd/api && GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${ROOT_DIR}/${OUT_DIR}/api
            cd ..
            echo -e "${BLUE}API built successfully.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
        2)
            echo -e "${BLUE}Exiting.${NC}"
            exit 0
            ;;
        *)
            echo -e "${RED}Invalid option. Please choose a valid option.${NC}"
            read -n 1 -s -r -p "Press any key to continue..."
            ;;
    esac
done
