#!/bin/bash

set -e

if [ $# != 1 ] ; then
echo "USAGE: $0 version"
echo " e.g.: $0 1.1"
exit 1;
fi

VERSION="$1"

wails build --clean --platform darwin/amd64
zip -r EasyClash-macos-amd-v"${VERSION}".zip build/bin/EasyClash.app

wails build --clean --platform darwin/arm64
zip -r EasyClash-macos-arm-v"${VERSION}".zip build/bin/EasyClash.app

wails build --clean --platform windows/arm64
zip -r EasyClash-win-arm-v"${VERSION}".zip build/bin/EasyClash.exe

wails build --clean --platform windows/amd64
zip -r EasyClash-win-amd-v"${VERSION}".zip build/bin/EasyClash.exe

echo "Done!"