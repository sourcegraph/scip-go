#!/usr/bin/env bash

fromGolangLines="$(grep 'FROM golang' Dockerfile | wc -l)"
if [[ "$fromGolangLines" -ne 2 ]]; then
  echo "Expected exactly 2 FROM lines in Dockerfile"
  exit 1
fi

uniqueGolangImages="$(grep 'FROM golang' Dockerfile | cut -d " " -f 2 | uniq | wc -l)"
if [[ "$uniqueGolangImages" -ne 1 ]]; then
  echo "Expected same image to be repeated in Dockerfile"
  exit 1
fi

golangTag="$(grep 'FROM golang' Dockerfile | head -n 1 | cut -d " " -f 2 | cut -d "@" -f 1)"
golangSha="$(grep 'FROM golang' Dockerfile | head -n 1 | cut -d " " -f 2 | cut -d "@" -f 2)"
digestLine="$(docker buildx imagetools inspect "$golangTag" | grep -E "^Digest:")"
if [[ "$digestLine" =~ *"golangSha"* ]]; then
  echo "SHA for $golangTag image ($golangSha) does not match $digestLine"
  exit 1
fi
