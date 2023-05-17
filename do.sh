#!/usr/bin/env sh
# Do - The Simplest Build Tool on Earth.
# Documentation and examples see https://github.com/8gears/do

set -e -u # -e "Automatic exit from bash shell script on error"  -u "Treat unset variables and parameters as errors"


build() {
  CGO_ENABLED=0 go build -o .bin/proxy .
}

build_image() {
  git_commit=$(git rev-parse --short HEAD)
  docker buildx build --platform linux/amd64 --push -t 8gears.container-registry.com/library/helm-charts-oci-proxy:$git_commit .
}

build_push_chart() {
  git_commit=$(git rev-parse --short HEAD)
  helm package chart
  helm push helm-charts-oci-proxy-1.0.0.tgz oci://8gears.container-registry.com/library/helm-charts-oci-proxy
}


deploy() {
   helm upgrade -i --namespace ocip-staging --create-namespace ocip-staging ./chart
}

run() {
  USE_TLS=1 DEBUG=1 DB_LOCATION=/tmp/fstore CACHE_TTL=30 .bin/proxy registry serve
}

tests() {
  go test -v ./...
}
"$@" # <- execute the task

[ "$#" -gt 0 ] || printf "Usage:\n\t./do.sh %s\n" "($(compgen -A function | grep '^[^_]' | paste -sd '|' -))"
