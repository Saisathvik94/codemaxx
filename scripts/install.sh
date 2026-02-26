#!/bin/bash

set -euo pipefail

# ---------------- CONFIG ----------------
REPO="Saisathvik94/codemaxx"
BINARY="codemaxx"
INSTALL_DIR="/usr/local/bin"
TMP_DIR="$(mktemp -d -t codemaxx-install-XXXXXXXX)"

GREEN="\033[0;32m"
CYAN="\033[0;36m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
RESET="\033[0m"

# ---------------- FUNCTIONS ----------------
show_ascii() {
cat << "EOF"
                           /$$                                                      
                          | $$                                                      
  /$$$$$$$  /$$$$$$   /$$$$$$$  /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$   /$$ /$$   /$$
 /$$_____/ /$$__  $$ /$$__  $$ /$$__  $$| $$_  $$_  $$ |____  $$|  $$ /$$/|  $$ /$$/
| $$      | $$  \ $$| $$  | $$| $$$$$$$$| $$ \ $$ \ $$  /$$$$$$$ \  $$$$/  \  $$$$/ 
| $$      | $$  | $$| $$  | $$| $$_____/| $$ | $$ | $$ /$$__  $$  >$$  $$   >$$  $$ 
|  $$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$ /$$/\  $$ /$$/\  $$
 \_______/ \______/  \_______/ \_______/|__/ |__/ |__/ \_______/|__/  \__/|__/  \__/
EOF
}

check_root() {
    if [[ $EUID -ne 0 ]]; then
        echo -e "${RED}Please run as root (sudo)${RESET}"
        exit 1
    fi
}

get_latest_release() {
    echo -e "${YELLOW}Fetching latest release...${RESET}"
    local latest_json
    latest_json=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest")
    VERSION=$(echo "$latest_json" | sed -n 's/.*"tag_name": "\(.*\)".*/\1/p')
}

determine_os_arch() {
    OS="$(uname | tr '[:upper:]' '[:lower:]')"
    ARCH="$(uname -m)"
    
    case "$ARCH" in
        x86_64) ARCH="amd64" ;;
        arm64|aarch64) ARCH="arm64" ;;
        *) echo -e "${RED}Unsupported architecture: $ARCH${RESET}" ; exit 1 ;;
    esac

    ZIP_NAME="codemaxx_${VERSION}_${OS}_${ARCH}.zip"
    URL="https://github.com/$REPO/releases/download/$VERSION/$ZIP_NAME"
    ZIP_PATH="$TMP_DIR/$ZIP_NAME"
}

download_release() {
    echo -e "${GREEN}Downloading $ZIP_NAME...${RESET}"
    curl -fL -o "$ZIP_PATH" "$URL" || { echo -e "${RED}Download failed${RESET}"; exit 1; }
}

extract_and_install() {
    echo -e "${CYAN}Extracting files...${RESET}"
    unzip -o "$ZIP_PATH" -d "$TMP_DIR"

    if [[ ! -f "$TMP_DIR/$BINARY" ]]; then
        echo -e "${RED}Binary not found in archive${RESET}"
        exit 1
    fi

    rm -f "$INSTALL_DIR/$BINARY"
    mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY"

    rm -rf "$TMP_DIR"
}

show_success() {
    echo -e "${GREEN}codemaxx installed successfully!${RESET}"
    echo -e "${CYAN}Version: $VERSION${RESET}"
    echo -e "${CYAN}Location: $INSTALL_DIR/$BINARY${RESET}"
    echo -e "${YELLOW}Run:${RESET} codemaxx --help"
}

# ---------------- SCRIPT EXECUTION ----------------
check_root
show_ascii
echo -e "${CYAN}Installing CodeMaxx CLI Tool...${RESET}"

get_latest_release
determine_os_arch
download_release
extract_and_install
show_success