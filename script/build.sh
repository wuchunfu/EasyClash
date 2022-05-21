#!/bin/bash

set -e

if [ $# != 1 ] ; then
echo "USAGE: $0 version"
echo " e.g.: $0 1.1"
exit 1;
fi

VERSION="$1"

wails build --clean --platform darwin/amd64
cd build/bin
zip -r EasyClash-macos-amd-v"${VERSION}".zip EasyClash.app
mv EasyClash-macos-amd-v"${VERSION}".zip ../../
cd ../../

wails build --clean --platform darwin/arm64
cd build/bin
zip -r EasyClash-macos-arm-v"${VERSION}".zip EasyClash.app
mv EasyClash-macos-arm-v"${VERSION}".zip ../../ 
cd ../../

wails build --clean --platform windows/arm64
cd build/bin
zip -r EasyClash-win-arm-v"${VERSION}".zip EasyClash.exe
mv EasyClash-win-arm-v"${VERSION}".zip ../../
cd ../../

wails build --clean --platform windows/amd64
cd build/bin
zip -r EasyClash-win-amd-v"${VERSION}".zip EasyClash.exe
mv EasyClash-win-amd-v"${VERSION}".zip ../../
cd ../../

echo "Done!"