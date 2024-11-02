#! /bin/bash -e

function handle_exit() {
  result=$?
  if [ "$result" != "0" ]; then
    printf "Failed to install helm-kanvas-snapshot plugin\n"
  fi
  exit $result
}

function normalize_architecture() {
    arch=$1

    case "$arch" in
      "aarch64")
      echo "arm64"
      ;;
      *)
      echo $arch
      ;;
    esac
}

function download_plugin() {
  os_name=$(uname -s)
  os_arch=$(uname -m)

  os_arch=$(normalize_architecture $os_arch)

  OUTPUT_BASENAME=helm-kanvas-snapshot
  version=$(grep version "$HELM_PLUGIN_DIR/plugin.yaml" | cut -d'"' -f2)
  DOWNLOAD_URL="https://github.com/meshery/helm-kanvas-snapshot/releases/download/v$version/helm-kanvas-snapshot_${version}_${os_name}_${os_arch}.tar.gz"
  OUTPUT_BASENAME_WITH_POSTFIX="$HELM_PLUGIN_DIR/$OUTPUT_BASENAME.tar.gz"

  echo -e "Download URL set to ${DOWNLOAD_URL}\n"
  echo -e "Artifact path: ${OUTPUT_BASENAME_WITH_POSTFIX}\n"
  echo -e "Downloading ${DOWNLOAD_URL} to ${HELM_PLUGIN_DIR}\n"

  if [ -z "${DOWNLOAD_URL}" ]; then
    echo -e "Unsupported OS / architecture: ${os_name}/${os_arch}\n"
    exit 1
  fi

  if [[ -n $(command -v curl) ]]; then
    if curl --fail -L "${DOWNLOAD_URL}" -o "${OUTPUT_BASENAME_WITH_POSTFIX}"; then
      echo -e "Successfully downloaded the archive, proceeding to install\n"
    else
      echo -e "Failed while downloading helm-kanvas-snapshot archive\n"
      exit 1
    fi
  else
    echo "curl is required to download the plugin"
    exit -1
  fi
}

function install_plugin() {
  local HELM_PLUGIN_ARTIFACT_PATH=${OUTPUT_BASENAME_WITH_POSTFIX}
  local PROJECT_NAME="helm-kanvas-snapshot"
  local HELM_PLUGIN_TEMP_PATH="/tmp/$PROJECT_NAME"

  echo -n "HELM_PLUGIN_ARTIFACT_PATH: ${HELM_PLUGIN_ARTIFACT_PATH}"
  rm -rf "${HELM_PLUGIN_TEMP_PATH}"

  echo -e "Preparing to install into ${HELM_PLUGIN_DIR}\n"
  mkdir -p "${HELM_PLUGIN_TEMP_PATH}"
  tar -xvf "${HELM_PLUGIN_ARTIFACT_PATH}" -C "${HELM_PLUGIN_TEMP_PATH}"
  mkdir -p "$HELM_PLUGIN_DIR/bin"
  mv "${HELM_PLUGIN_TEMP_PATH}/helm-kanvas-snapshot" "${HELM_PLUGIN_DIR}/bin/helm-kanvas-snapshot"
  rm -rf "${HELM_PLUGIN_TEMP_PATH}"
  rm -rf "${HELM_PLUGIN_ARTIFACT_PATH}"
}

function install() {
  echo "Installing helm-kanvas-snapshot..."

  download_plugin
  status=$?
  if [ $status -ne 0 ]; then
    echo -e "Downloading plugin failed\n"
    exit 1
  fi

  set +e
  install_plugin
  local install_status=$?
  set -e

  if [ "$install_status" != "0" ]; then
    echo "Installing helm-kanvas-snapshot plugin failed with error code: ${install_status}"
    exit 1
  fi

  echo
  echo "helm-kanvas-snapshot is installed."
  echo
  "${HELM_PLUGIN_DIR}/bin/helm-kanvas-snapshot" -h
  echo
  echo "See https://github.com/meshery/helm-kanvas-snapshot/helm-kanvas-snapshot#readme for more information on getting started."
}

trap "handle_exit" EXIT

install "$@"
