#!/usr/bin/env bash

TMPFILE=$(mktemp)
pandoc --from=latex+lhs --to=html5 --no-highlight -S $1 > "$TMPFILE"
stack exec mungehtml "$TMPFILE"
pbcopy < "$TMPFILE"
echo "Copied."
