#!/bin/bash

NEW_NAME=$1
OLD_NAME="go-base-end"

cat hack/logo.txt

if [ -z "$NEW_NAME" ]; then
  echo "Usage: make init name=<module name>"
  exit 1
fi

echo "--- Initializing project: ${NEW_NAME} ---"

# modify module name in go.mod
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s|${OLD_NAME}|${NEW_NAME}|g" go.mod
else
    sed -i "s|${OLD_NAME}|${NEW_NAME}|g" go.mod
fi

# modify module name in .go files
find . -type f -name "*.go" -exec sed -i.bak "s|${OLD_NAME}|${NEW_NAME}|g" {} +
find . -type f -name "*.bak" -delete

echo "Resetting git history..."
rm -rf .git
git init
git add .
git commit -m "Initial commit from scaffold"

echo "--- Initialized  project: ${NEW_NAME} ---"


