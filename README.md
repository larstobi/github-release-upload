Create github release and upload asset
======================================

Example:

    export GITHUB_OWNER=larstobi
    export GITHUB_REPO=github-release-upload
    export GITHUB_AUTH_TOKEN=XXXXXXXXXXXXXXX
    export GITHUB_RELEASE_ASSET=github
    export GITHUB_RELEASE_NAME=1.0.0
    ./github

Compile and statically link go binary
=====================================

    docker run --rm -e CGO_ENABLED=true -e LDFLAGS='-extldflags "-static"' -e COMPRESS_BINARY=true -v $(pwd):/src centurylink/golang-builder
