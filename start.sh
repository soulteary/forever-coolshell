#! /bin/bash

CUR_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

cd "${CUR_DIR}/coolshell" || exit 1

PORT="${PORT:-8080}"

python3 -m http.server "${PORT}"
