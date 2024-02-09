#!/bin/sh
#
# An example hook script to verify what is about to be committed.
# Called by "git commit" with no arguments.  The hook should
# exit with non-zero status after issuing an appropriate message if
# it wants to stop the commit.
#
# To enable this hook, rename this file to "pre-commit".

echo "RUN gofmt"
## this will retrieve all of the .go files that have been 
## changed since the last commit
STAGED_GO_FILES=$(git diff --cached --name-only -- '*.go')
## we can check to see if this is empty
if [ $STAGED_GO_FILES ]
then
    echo "$STAGED_GO_FILES"
    # golangci-lint run --tests=0 ./...
    gofmt -s -w .
    git add .
## otherwise we can do stuff with these changed go files
else
    echo "No Go Files to Update"
fi