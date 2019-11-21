set -e

version=$(make version)

rm -r release
mkdir release

make build GOOS=linux
tar czvf "release/wlrand-$version-linux.tgz" wlrand 
rm wlrand

make build GOOS=darwin
tar czvf "release/wlrand-$version-macos.tgz" wlrand 
rm wlrand

make build GOOS=windows
tar czvf "release/wlrand-$version-windows.tgz" wlrand.exe
rm wlrand.exe
