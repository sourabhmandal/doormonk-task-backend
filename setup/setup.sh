#!/bin/sh
HOOK_PATH="setup/hooks"
git config core.hooksPath $HOOK_PATH
ROOT=$(pwd) 
echo "SET HOOKS PATH to ::> $ROOT/$HOOK_PATH"