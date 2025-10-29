#!/usr/bin/env bash

WORKSPACE="/home/taw/workspace"

find ".." -type f -maxdepth 5 | while read -r file; do
  dir="$(dirname "$file")"
  (
    cd "$dir" || exit

    if [ -f .git/config ]; then
      BANG="$(git ls-files --cached --others --exclude-standard)"

      current_dir=$(pwd)
      echo "$BANG" | while IFS= read -r line; do
      echo "$current_dir/$line" >> "$(WORKSPACE)/taw-language-server/scan.txt"
      done
    fi
  )
done
