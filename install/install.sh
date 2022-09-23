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
#   -c, --configuration-dir
#     Override the bin installation directory
#
#   Example:
#     sh install.sh -i ens2 -y

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

detect_platform() {
  if has uname; then
    PLATFORM="$(uname -s | tr '[:upper:]' '[:lower:]')"
  elif has os; then
    PLATFORM="$(os | tr '[:upper:]' '[:lower:]')"
  else
    error 'Unable to detect platform'
    exit 1
  fi
}

detect_arch() {
  if has uname; then
    ARCH="$(uname -m | tr '[:upper:]' '[:lower:]')"
  elif has arch; then
    ARCH="$(arch | tr '[:upper:]' '[:lower:]')"
  else
    error 'Unable to detect architecture'
    exit 1
  fi
  if [ "$ARCH" = "x86_64" ]; then
    ARCH='amd64'
  fi
}

remove_old_binary() {
  local path="${HOME}/.local/bin/jumper"
  if test -f ${path}
  then
    info "Removing old jumper binary..."
    rm ${path}
    completed "Old jumper binary is removed"
  fi
}

install_binary () {
  info "Install jumper binary..."
  printf '\n'

  $(curl -s -H "Accept: application/vnd.github.v3+json" https://api.github.com/repos/jklaiber/jumper/releases \
  | grep "browser_download_url" \
  | grep "$PLATFORM-$ARCH.tar.gz" \
  | head -n 1 \
  | cut -d : -f 2,3 \
  | tr -d \" \
  | wget -O ${HOME}/.local/bin/jumper.tar.gz -qi -)
  cd ${HOME}/.local/bin
  tar xfvz jumper.tar.gz
  rm -f jumper.tar.gz
  printf '\n'
  completed "Jumper binary installed"
}

install() {
  printf '\n'
  install_binary
}

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

start_message() {
  detect_platform
  detect_arch
  info "${BOLD}Detected platform${NO_COLOR}: ${GREEN}${PLATFORM}${NO_COLOR}"
  info "${BOLD}Detected architecture${NO_COLOR}: ${GREEN}${ARCH}${NO_COLOR}"
  printf '\n'
  confirm "Install Jumper SSH CLI Manager?"
  if has jumper; then
    info "Jumper is already installed, upgrading..."
    remove_old_binary
    install_binary
    printf '\n'
    completed "Jumper is upgraded"
  else
    install
    printf '\n'
    completed "Jumper is installed"
  fi
}

finish_message() {
  URL="https://github.com/jklaiber/jumper"

  printf '\n'
  info "Please follow the steps to use Jumper on your machine:
    ${BOLD}${UNDERLINE}Setup${NO_COLOR}
    Jumper will force to setup the application before the first use.
    Please follow the instructions to setup the application.
    ${BOLD}${UNDERLINE}Documentation${NO_COLOR}
    To check out the documentation go to:
        ${UNDERLINE}${BLUE}${URL}${NO_COLOR}
  "
}

while getopts "hf" option; do
  case $option in
    h)
      help
      exit;;
    f)
      FORCE=1
      ;;
    ?)
      error "Invalid option: -$OPTARG"
      exit;;
  esac
done

print_banner
start_message
finish_message