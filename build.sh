#!/usr/bin/env bash
osArray=("linux" "darwin" "freebsd" "windows")
archs=("amd64" "386")
version=${1-"0.0.1-preview1"}
out_file="yoyocli"

build() {
  os=$1
  arch=$2

  GOOS=$os GOARCH=$arch go build -o "${out_file}" Cli.go
  tgzName="${out_file}_${version}_${os}_${arch}.tar.gz"
  rm -f "${tgzName}"
  tar -czf "${tgzName}" "${out_file}"
  rm -f "${out_file}"

  echo $(shasum -a 256 "${tgzName}")
}

main() {
  for os in "${osArray[@]}"; do
    for arch in "${archs[@]}"; do
      build "${os}" "${arch}"
    done
  done
}

main