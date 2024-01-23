#!/bin/sh

cd "$(dirname "$0")" || exit
BASEDIR=$(pwd)
cd ../
ln -s -f "$BASEDIR/pre-commit" .git/hooks/pre-commit
