$ENV:CGO_ENABLED=0
$ENV:GOOS="windows"
$ENV:GOARCH="amd64"
go build -o auto-magical-animal-v002.exe main.go