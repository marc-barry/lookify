#!/bin/bash

tooling_gopath=/tmp/tooling-gopath

function abs_path {
  (cd "$(dirname '$1')" &>/dev/null && echo $(pwd))
}
root_path="$(abs_path "$(dirname "${BASH_SOURCE[0]}")")"

ensure_go_binary() {
  binary_name=$1
  url=$2

  if ! which "${binary_name}" > /dev/null ; then
    echo "${binary_name} isn't installed globally. Installing it would make this process a bit faster."
    mkdir -p "${tooling_gopath}"
    echo "+ go get ${url}"
    GOPATH="${tooling_gopath}" go get "${url}"
    export PATH="${PATH}:${tooling_gopath}/bin"
  fi
}

function config_errcheck() {
  errcheck_path="$(godep path)/src/github.com/marc-barry"
  
  if [ ! -d "${errcheck_path}" ] && mkdir -p "${errcheck_path}"; then
    ln -sf "${root_path}" "${errcheck_path}/lookify"
    trap 'rm -r "${errcheck_path}"' EXIT
  else
    ln -sf "${root_path}" "${errcheck_path}/lookify"
    trap 'unlink "${errcheck_path}/lookify"' EXIT
  fi
}

function run_errcheck() {
  GOPATH="$(godep path)" errcheck github.com/marc-barry/lookify
}
