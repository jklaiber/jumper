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


create_configuration() {
  local path="$CONFIGURATION_DIR/.jumper.yaml"
  info "Create initial configuration file..."
  if test -f ${path}
  then
    info "Configuration already exists, skipping..."
  else
    bash -c "cat > "${path}"" << EOF
inventory_file: $CONFIGURATION_DIR/.jumper.inventory.yaml
vault_password: 
EOF
    completed "Configuration $path is generated"
  fi
  local path="$CONFIGURATION_DIR/.jumper.inventory.yaml"
  info "Create example inventory file..."
  if test -f ${path}
  then
    info "Inventory already exists, skipping..."
  else
    bash -c "cat > "${path}"" << EOF
all:
  hosts:
    foo.example.com:
    bar.example.com:
  children:
    webservers:
      hosts:
        foo.example.com:
          username: userfoo
          password: passfoo
        bar.example.com:
          username: userfoo
          password: passfoo
        foobar.example.com:
          password: passfoo
    dbservers:
      hosts:
        one.example.com:
        two.example.com:
        three.example.com:
      vars:
        username: userfoo
        password: passfoo
  vars:
    sshkey: ${HOME}/.ssh/id_rsa
    username: globalusername
EOF
    completed "Example inventory $path is generated"
  fi
}

install_binary () {
  info "Install jumper binary..."
  printf '\n'

  $(curl -s -H "Accept: application/vnd.github.v3+json" https://api.github.com/repos/jklaiber/jumper/releases \
  | grep "browser_download_url" \
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
  create_configuration
}

if [ -z "${CONFIGURATION_DIR-}" ]; then
  CONFIGURATION_DIR=${HOME}
fi

# parse argv variables
while [ "$#" -gt 0 ]; do
  case "$1" in
  -c | --configuration-dir)
    CONFIGURATION_DIR="$2"
    shift 2
    ;;
  -V | --verbose)
    VERBOSE=1
    shift 1
    ;;
  -f | -y | --force | --yes)
    FORCE=1
    shift 1
    ;;
  -c=* | --configuration-dir=*)
    CONFIGURATION_DIR="${1#*=}"
    shift 1
    ;;
  -V=* | --verbose=*)
    VERBOSE="${1#*=}"
    shift 1
    ;;
  -f=* | -y=* | --force=* | --yes=*)
    FORCE="${1#*=}"
    shift 1
    ;;
  *)
    error "Unknown option: $1"
    exit 1
    ;;
  esac
done

printf "
    _                                 
   (_)                                
    _ _   _ _ __ ___  _ __   ___ _ __ 
   | | | | |  _   _ \|  _ \ / _ \  __|
   | | |_| | | | | | | |_) |  __/ |   
   | |\__ _|_| |_| |_|  __/ \___|_|   
  _/ |               | |              
 |__/                |_|              
"
printf "\n  %s\n" "${UNDERLINE}Configuration${NO_COLOR}"
info "${BOLD}Configuration directory${NO_COLOR}: ${GREEN}${CONFIGURATION_DIR}${NO_COLOR}"

# non-empty VERBOSE enables verbose untarring
if [ -n "${VERBOSE-}" ]; then
  VERBOSE=v
  info "${BOLD}Verbose${NO_COLOR}: yes"
else
  VERBOSE=
fi

printf '\n'

confirm "Install Jumper SSH CLI Manager?"
install
printf '\n'
completed "Jumper installed"

URL="https://github.com/jklaiber/jumper"

printf '\n'
info "Please follow the steps to use Jumper on your machine:
  ${BOLD}${UNDERLINE}Create configuration${NO_COLOR}
  You can edit the configuration file ${BOLD}${CONFIGURATION_DIR}/.jumper.yaml${NO_COLOR} default is:
      inventory_file: ${HOME}/.jumper.inventory.yaml
      vault_password:
  ${BOLD}${UNDERLINE}Documentation${NO_COLOR}
  To check out the documentation go to:
      ${UNDERLINE}${BLUE}${URL}${NO_COLOR}
"