#!/usr/bin/env bash

PROJECT_NAME="helm-kanvas-snapshot"
PROJECT_GH="meshery/$PROJECT_NAME"

# Convert the HELM_PLUGIN_PATH to unix if cygpath is
# available. This is the case when using MSYS2 or Cygwin
# on Windows where helm returns a Windows path but we
# need a Unix path
if command -v cygpath >/dev/null 2>&1; then
  HELM_BIN="$(cygpath -u "${HELM_BIN}")"
  HELM_PLUGIN_DIR="$(cygpath -u "${HELM_PLUGIN_DIR}")"
fi

[ -z "$HELM_BIN" ] && HELM_BIN=$(command -v helm)

[ -z "$HELM_HOME" ] && HELM_HOME=$(helm env | grep 'HELM_DATA_HOME' | cut -d '=' -f2 | tr -d '"')

mkdir -p "$HELM_HOME"

: "${HELM_PLUGIN_DIR:="$HELM_HOME/plugins/$PROJECT_NAME"}"

if [ "$SKIP_BIN_INSTALL" = "1" ]; then
  echo "Skipping binary install"
  exit
fi


# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="arm";;
    armv6*) ARCH="arm";;
    armv7*) ARCH="arm";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="x86_64";;
    i686) ARCH="i386";;
    i386) ARCH="i386";;
  esac
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(uname -s)

  case "$OS" in
  Windows_NT) OS='windows' ;;
  # Msys support
  MSYS*) OS='windows' ;;
  # Minimalist GNU for Windows
  MINGW*) OS='windows' ;;
  CYGWIN*) OS='windows' ;;
  Darwin) OS='darwin' ;;
  Linux) OS='linux' ;;
  esac
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  local supported=""
  for os in darwin freebsd linux windows; do
    for arch in arm arm64 i386 x86_64 386 amd64; do
      supported+="${os}-${arch}\n"
    done
  done
  echo "supported: ${supported[@]}"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuild binary for ${OS}-${ARCH}."
    exit 1
  fi

  if ! type "curl" > /dev/null && ! type "wget" > /dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# getDownloadURL checks the latest available version.
getDownloadURL() {
  version=$(git -C "$HELM_PLUGIN_DIR" describe --tags --abbrev=0 2> /dev/null)
  # remove the 'v' at the beginning of the version
  version_without_v=$(git -C "$HELM_PLUGIN_DIR" describe --tags --abbrev=0 2> /dev/null | sed 's/^v//')
  PROJECT_NAME_WITH_VERSION="${PROJECT_NAME}_${version_without_v}_${OS}_${ARCH}"
  if [ -n "$version" ]; then
    DOWNLOAD_URL="https://github.com/$PROJECT_GH/releases/download/$version/$PROJECT_NAME_WITH_VERSION.tar.gz"
  else
    echo "No release found. "
    exit 1
  fi
}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  PLUGIN_TMP_FILE="/tmp/${PROJECT_NAME}.tar.gz"
  echo "Downloading $DOWNLOAD_URL"
  if type "curl" > /dev/null; then
    echo "curl -L $DOWNLOAD_URL -o $PLUGIN_TMP_FILE"
    curl -L "$DOWNLOAD_URL" -o "$PLUGIN_TMP_FILE"
  elif type "wget" > /dev/null; then
   echo "wget -q -O $PLUGIN_TMP_FILE $DOWNLOAD_URL"
    wget -q -O "$PLUGIN_TMP_FILE" "$DOWNLOAD_URL"
  fi
}

# installFile verifies the SHA256 for the file, then unpacks and
# installs it.
installFile() {
  HELM_TMP="/tmp/$PROJECT_NAME"
  HELM_TMP_BIN="/tmp/$PROJECT_NAME/$PROJECT_NAME_WITH_VERSION/$PROJECT_NAME"
  mkdir -p "$HELM_TMP"
  tar xzf "$PLUGIN_TMP_FILE" -C "$HELM_TMP"
  if [ "${OS}" = "windows" ]; then
    HELM_TMP_BIN="$HELM_TMP_BIN.exe"
  fi
  echo "Preparing to install into ${HELM_PLUGIN_DIR}"
  mkdir -p "$HELM_PLUGIN_DIR/bin"
  cp "$HELM_TMP_BIN" "$HELM_PLUGIN_DIR/bin"
}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    echo "Failed to install $PROJECT_NAME"
    echo "\tFor support, open issue at https://github.com/$PROJECT_GH."
  fi
  exit $result
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
installFile
echo
echo "helm-kanvas-snapshot is installed."
echo "${HELM_PLUGIN_DIR}/bin/helm-kanvas-snapshot" -h
echo
echo "See https://github.com/$PROJECT_GH#readme for more information on getting started."

