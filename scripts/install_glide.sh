# This script installs Glide (a dependency management tool) for you.
# Run with `make glide`

go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/Masterminds/glide
go install

glide --version

cd $GOPATH/src/github.com/bobheadxi/calories