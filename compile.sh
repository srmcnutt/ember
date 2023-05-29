# this file is used to compile the fire-t project.
# output is in the under the binaries/$os directory. i.e binaries/linux/fire-t
GOOS=windows GOARCH=386 go build -o ./binaries/windows/fire-t
GOOS=linux GOARCH=386 go build -o ./binaries/linux/fire-t
GOOS=darwin GOARCH=386 go build -o ./binaries/darwin/fire-t
