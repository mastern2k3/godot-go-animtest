export GOROOT=/usr/local/go
# export GOPATH=/home/nitzan/Desktop/godev/go/src
export CGO_ENABLED=1

GO="$GOROOT/bin/go"

$GO build -i -v -buildmode=c-shared -o libgodot.so .
