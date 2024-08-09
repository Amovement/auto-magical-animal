$ENV:GOOS="js"
$ENV:GOARCH="wasm"
go build -o main.wasm
