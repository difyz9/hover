
mac_arm64:
	go env -w  GOOS=darwin GOARCH=arm64
	go build  -o out/hover main.go
	mv out/hover /Users/apple/go/bin/darwin_amd64/hover

