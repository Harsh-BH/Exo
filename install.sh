#!/usr/bin/env bash
set -euo pipefail

# EXO Installer
# Usage: curl -sSL https://raw.githubusercontent.com/Harsh-BH/Exo/main/install.sh | bash

REPO="Harsh-BH/Exo"
BINARY="exo"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# ── Detect OS and arch ────────────────────────────────────────────────────────
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "${ARCH}" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "❌ Unsupported architecture: ${ARCH}"
    exit 1
    ;;
esac

case "${OS}" in
  linux|darwin) ;;
  *)
    echo "❌ Unsupported OS: ${OS}"
    echo "   Download manually: https://github.com/${REPO}/releases"
    exit 1
    ;;
esac

ASSET_NAME="${BINARY}-${OS}-${ARCH}"

# ── Fetch latest version ──────────────────────────────────────────────────────
echo "→ Fetching latest EXO release..."
LATEST=$(curl -sSf "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\(.*\)".*/\1/')

if [ -z "${LATEST}" ]; then
  echo "❌ Could not determine latest version. Check https://github.com/${REPO}/releases"
  exit 1
fi

echo "→ Installing EXO ${LATEST} (${OS}/${ARCH})..."

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST}/${ASSET_NAME}"

# ── Download ──────────────────────────────────────────────────────────────────
TMP_FILE="$(mktemp)"
trap 'rm -f "${TMP_FILE}"' EXIT

if ! curl -sSfL "${DOWNLOAD_URL}" -o "${TMP_FILE}"; then
  echo "❌ Download failed: ${DOWNLOAD_URL}"
  echo "   Check https://github.com/${REPO}/releases for available assets."
  exit 1
fi

chmod +x "${TMP_FILE}"

# ── Install ───────────────────────────────────────────────────────────────────
if [ -w "${INSTALL_DIR}" ]; then
  mv "${TMP_FILE}" "${INSTALL_DIR}/${BINARY}"
else
  echo "→ Requesting sudo to install to ${INSTALL_DIR}..."
  sudo mv "${TMP_FILE}" "${INSTALL_DIR}/${BINARY}"
fi

# ── Verify ────────────────────────────────────────────────────────────────────
if command -v exo &>/dev/null; then
  echo ""
  echo "✓ EXO installed successfully!"
  exo version
  echo ""
  echo "  Get started: exo init"
  echo "  Help:        exo help"
else
  echo "✓ EXO installed to ${INSTALL_DIR}/${BINARY}"
  echo "  Make sure ${INSTALL_DIR} is in your PATH."
fi
