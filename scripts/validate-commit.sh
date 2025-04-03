#!/bin/bash

file="$1"

if [ ! -f "$file" ]; then
  echo "❌ Commit message file not found."
  exit 1
fi

commit_msg=$(cat "$file")

if ! echo "$commit_msg" | grep -qE '^NANI-[0-9]+: .+'; then
  echo "❌ Invalid commit message. Use format: NANI-xxx: description"
  exit 1
fi

echo "✅ Commit message is valid."
