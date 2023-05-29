# this file is used to compile the ember project.
# output is in the under the binaries/$os directory. i.e binaries/linux/amd64/ember
GOOS=windows GOARCH=amd64 go build -o ./binaries/windows/amd64/ember.exe
GOOS=linux GOARCH=amd64 go build -o ./binaries/linux/amd64/ember
GOOS=darwin GOARCH=amd64 go build -o ./binaries/darwin/amd64/ember
GOOS=darwin GOARCH=arm64 go build -o ./binaries/darwin/arm64/ember

