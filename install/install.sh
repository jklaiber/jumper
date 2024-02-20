#!/usr/bin/env sh

# shellcheck disable=SC2039

# Options
#
#   -V, --verbose
#     Enable verbose output for the installer
#
#   -f, -y, --force, --yes
#     Skip the confirmation prompt during installation
#
#   Example:
#     sh install.sh -y

set -eu
printf '\n'

BOLD="$(tput bold 2>/dev/null || printf '')"
GREY="$(tput setaf 0 2>/dev/null || printf '')"
UNDERLINE="$(tput smul 2>/dev/null || printf '')"
RED="$(tput setaf 1 2>/dev/null || printf '')"
GREEN="$(tput setaf 2 2>/dev/null || printf '')"
YELLOW="$(tput setaf 3 2>/dev/null || printf '')"
BLUE="$(tput setaf 4 2>/dev/null || printf '')"
MAGENTA="$(tput setaf 5 2>/dev/null || printf '')"
NO_COLOR="$(tput sgr0 2>/dev/null || printf '')"

help() {
   echo "Script to install jumper (a simple cli ssh manager)"
   echo
   echo "Syntax: install.sh [-h|f]"
   echo "options:"
   echo "-h               Print this Help."
   echo "-f               Force install."
   echo
}

info() {
  printf '%s\n' "${BOLD}${GREY}>${NO_COLOR} $*"
}

warn() {
  printf '%s\n' "${YELLOW}! $*${NO_COLOR}"
}

error() {
  printf '%s\n' "${RED}x $*${NO_COLOR}" >&2
}

completed() {
  printf '%s\n' "${GREEN}✓${NO_COLOR} $*"
}

has() {
  command -v "$1" 1>/dev/null 2>&1
}

confirm() {
  if [ -z "${FORCE-}" ]; then
    printf "%s " "${MAGENTA}?${NO_COLOR} $* ${BOLD}[y/N]${NO_COLOR}"
    set +e
    read -r yn </dev/tty
    rc=$?
    set -e
    if [ $rc -ne 0 ]; then
      error "Error reading from prompt (please re-run with the '--yes' option)"
      exit 1
    fi
    if [ "$yn" != "y" ] && [ "$yn" != "yes" ]; then
      error 'Aborting (please answer "yes" to continue)'
      exit 1
    fi
  fi
}

detect_platform_package_manager() {
  if [ -f /etc/os-release ]; then
    . /etc/os-release
    case $ID in
      debian|ubuntu) PACKAGE_MANAGER="apt";;
      fedora|centos|rhel) PACKAGE_MANAGER="rpm";;
      alpine) PACKAGE_MANAGER="apk";;
      *) error "Unsupported OS for this installation script"; exit 1;;
    esac
  else
    error 'Unable to detect OS or package manager'
    exit 1
  fi
}

detect_arch() {
  ARCH=$(uname -m)
  case $ARCH in
    x86_64) ARCH="amd64";;
    armv7l|armv6l) ARCH="arm";;
    aarch64) ARCH="arm64";;
    *) error "Unsupported architecture: $ARCH"; exit 1;;
  esac
}

fetch_latest_version() {
  VERSION=$(curl -s https://api.github.com/repos/jklaiber/jumper/releases/latest | grep 'tag_name' | cut -d\" -f4)
  if [ -z "$VERSION" ]; then
    error "Unable to fetch latest version"
    exit 1
  fi
  # Strip the 'v' prefix from version tag if present
  VERSION=${VERSION#v}
}

download_and_install() {
  local base_url="https://github.com/jklaiber/jumper/releases/download/v${VERSION}"
  local package_format
  local package_name="jumper_${VERSION}_linux_${ARCH}"

  case $PACKAGE_MANAGER in
    apt) package_format=".deb";;
    rpm) package_format=".rpm";;
    apk) package_format=".apk";;
  esac

  local url="${base_url}/${package_name}${package_format}"

  temp_dir=$(mktemp -d)
  cd "$temp_dir"

  info "Downloading ${package_name}${package_format}..."
  if ! curl -LO "$url"; then
    error "Failed to download ${package_name}${package_format}"
    exit 1
  fi

  info "Installing ${package_name}${package_format}..."
  case $PACKAGE_MANAGER in
    apt)
      sudo apt install "./${package_name}${package_format}" -y
      ;;
    rpm)
      sudo rpm -i "${package_name}${package_format}"
      ;;
    apk)
      sudo apk add --allow-untrusted "${package_name}${package_format}"
      ;;
  esac

  cd - > /dev/null
  rm -rf "$temp_dir"

  completed "Jumper installed successfully."
}

while getopts "hfyV" option; do
  case $option in
    h)
      help
      exit;;
    f|y)
      FORCE=1
      ;;
    V)
      VERBOSE=1
      ;;
    ?)
      error "Invalid option: -$OPTARG"
      exit;;
  esac
done

print_banner() {
  printf '%s\n' \
    '    _                                 '\
    '   (_)                                '\
    '    _ _   _ _ __ ___  _ __   ___ _ __ '\
    '   | | | | |  _   _ \|  _ \ / _ \  __|'\
    '   | | |_| | | | | | | |_) |  __/ |   '\
    '   | |\__ _|_| |_| |_|  __/ \___|_|   '\
    '  _/ |               | |              '\
    ' |__/                |_|              '
  printf '\n'
}

print_banner
detect_platform_package_manager
detect_arch
fetch_latest_version
confirm "Install Jumper SSH CLI Manager?"
download_and_install
finish_message
