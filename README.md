## Topics

### 1. GO Installation

I installed using the go installer for a Windows OS.
Set GOPATH in the environment variables as well.

### 2. Go Modules

This allows go to manage dependencies especially when creating projects outside the Go Path. Originally, before version 1.11 projects could only be created within the gopath. This required constantly changing the path per project (outside the default go path) or to create all projects within the go path.

Go modules helps to create repos wherever necessary and handle dependencies alot better.

**GO111MODULE=on** -> means that go recoginized your project outside the default go path with the help of a go.mod file

**GO111MODULE=auto** -> means that projects outside the gopath and also within the gopath are recognized

**GO111MODULE=off** -> means that it can only recognize projects within the gopath

In order to change the value of GO111MODULE on Windows, locate the file in the following path "%HOME%/Appdata/Roaming/go/env" and edit accordingly

"go build" helps to download dependencies and reference them in the go.mod

"go test" recognizes all test files ("in this format \*\_test.go") and runs them.

"go tidy" adds all necessary dependencies to run tests in your project

### 3.
