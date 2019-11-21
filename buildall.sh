set -e

version=$(make version)
git_state=$(make gitstate)
verstr=$version-$git_state

rm -r release
mkdir release

# Generates uncompressed artifacts in the `work/wlrand` directory.
#
# $1: OS name
# $2: executable filename
build_one() {
    os=$1
    exe=$2

    rm -rf release/work
    mkdir -p release/work/wlrand

    make build GOOS="$os"
    mv "$exe" release/work/wlrand
    cp README.md release/work/wlrand
}

build_one linux wlrand
(
    artifact=wlrand-$verstr-linux.tgz
    cd release/work
    tar -czvf "$artifact" wlrand
    mv "$artifact" ..
)

build_one darwin wlrand
(
    artifact=wlrand-$verstr-macos.tgz
    cd release/work
    tar -czvf "$artifact" wlrand
    mv "$artifact" ..
)

build_one windows wlrand.exe
(
    artifact=wlrand-$verstr-windows.zip
    cd release/work
    zip -r "$artifact" wlrand
    mv "$artifact" ..
)
