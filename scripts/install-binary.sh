# // Copyright Meshery Authors
# //
# // Licensed under the Apache License, Version 2.0 (the "License");
# // you may not use this file except in compliance with the License.
# // You may obtain a copy of the License at
# //
# //     http://www.apache.org/licenses/LICENSE-2.0
# //
# // Unless required by applicable law or agreed to in writing, software
# // distributed under the License is distributed on an "AS IS" BASIS,
# // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# // See the License for the specific language governing permissions and
# // limitations under the License.

# This script installs the Helm Kanvas Snapshot plugin.

#!/usr/bin/env sh

echo "Installing Helm Kanvas Snapshot plugin..."

CLI="kanvas-snapshot"
REPO_NAME="helm-kanvas-snapshot"
PROJECT_ORG="${PROJECT_ORG:-meshery}"
PROJECT_GH="$PROJECT_ORG/$REPO_NAME"
HELM_BIN="/usr/local/bin/helm"
export GREP_COLOR="never"

HELM_MAJOR_VERSION=$("${HELM_BIN}" version --client --short | awk -F '.' '{print $1}')

# : ${HELM_PLUGIN_DIR:="$("${HELM_BIN}" home --debug=false)/plugins/helm-diff"}

# Handle HELM_PLUGIN_DIR filepath based on OS. Use *nix-based filepathing

if type cygpath >/dev/null 2>&1; then
  HELM_PLUGIN_DIR=$(cygpath -u $HELM_PLUGIN_DIR)
fi

if [ "$SKIP_BIN_INSTALL" = "1" ]; then
  echo "Skipping binary install"
  exit
fi

# Identify systm architecture
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
  armv5*) ARCH="armv5" ;;
  armv6*) ARCH="armv6" ;;
  armv7*) ARCH="armv7" ;;
  aarch64) ARCH="arm64" ;;
  x86) ARCH="386" ;;
  x86_64) ARCH="x86_64" ;;
  i686) ARCH="386" ;;
  i386) ARCH="386" ;;
  esac
  echo "ARCH: $ARCH"
}

# Identify operating system
initOS() {
  OS=$(uname | tr '[:upper:]' '[:lower:]')

  case "$OS" in
  # Msys support
  msys*) OS='windows' ;;
  # Minimalist GNU for Windows
  mingw*) OS='windows' ;;
  darwin) OS='darwin' ;;
  esac
  echo "OS: $OS"
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  supported="linux_amd64\ndarwin_x86_64\nlinux_arm64\ndarwin_arm64\nwindows_amd64"
  if ! echo "${supported}" | grep -q "${OS}_${ARCH}"; then
    echo "No prebuilt binary for ${OS}_${ARCH}."
    exit 1
  fi

  if ! type "curl" >/dev/null && ! type "wget" >/dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# getDownloadURL checks the latest available version.
getDownloadURL() {
  #version=$(git -C "$HELM_PLUGIN_DIR" describe --tags --exact-match 2>/dev/null || :)
    echo "OS: $OS"

  version="$(cat $HELM_PLUGIN_DIR/plugin.yaml | grep "version" | cut -d '"' -f 2)"
  if [ -n "$version" ]; then
DOWNLOAD_URL="https://github.com/$PROJECT_GH/releases/download/v$version/$CLI_$version_$OS_$ARCH.tar.gz"
    echo "DOWNLOAD_URL1: $DOWNLOAD_URL"
    # https://github.com/meshery/helm-kanvas-snapshot/releases/download/v0.2.0/kanvas-snapshot_0.2.0_Darwin_x86_64.tar.gz
  else
    # Use the GitHub API to find the download url for this project.
    url="https://api.github.com/repos/$PROJECT_GH/releases/latest"
    if type "curl" >/dev/null; then
      DOWNLOAD_URL=$(curl -s $url | grep $OS_$ARCH\" | awk '/\"browser_download_url\":/{gsub( /[,\"]/,"", $2); print $2}')
          echo "DOWNLOAD_URL2: $DOWNLOAD_URL"

    elif type "wget" >/dev/null; then
      DOWNLOAD_URL=$(wget -q -O - $url | grep $OS_$ARCH\" | awk '/\"browser_download_url\":/{gsub( /[,\"]/,"", $2); print $2}')
          echo "DOWNLOAD_URL3: $DOWNLOAD_URL"

    fi
  fi

}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  BINDIR="$HELM_PLUGIN_DIR/bin"
  rm -rf "$BINDIR"
  mkdir -p "$BINDIR"
  echo "Downloading $DOWNLOAD_URL"
  if type "curl" >/dev/null; then
    HTTP_CODE=$(curl -sL --write-out "%{http_code}" "$DOWNLOAD_URL" --output "$BINDIR/$CLI")
    if [ ${HTTP_CODE} -ne 200 ]; then
      exit 1
    fi
  elif type "wget" >/dev/null; then
    wget -q -O "$BINDIR/$CLI" "$DOWNLOAD_URL"
  fi

  chmod +x "$BINDIR/$CLI"

}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    echo "Failed to install $CLI"
    printf "\tFor support, go to https://discuss.layer5.io.\n"
  fi
  exit $result
}

# testVersion tests the installed client to make sure it is working.
testVersion() {
  set +e
  echo "$CLI installed into $HELM_PLUGIN_DIR/$CLI"
  "${HELM_PLUGIN_DIR}/bin/$CLI" version
  echo "`helm $CLI --help` to get started."
  set -e
}


# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e
initArch
initOS
verifySupported
getDownloadURL
downloadFile
testVersion