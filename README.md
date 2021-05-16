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

### 3. Variables

Variable declaration : var varName DataType = value
Variable declaration with expression assignment operator(:=) : varName := value
There is no need for the "var" keyword in this case

There are a number of datatypes represented in GO depending on signed/unsigned, for 16,32 or 64 bit representation. Examples are integers, float, boolean or string

**Explicit declaration** : var name int64 = 100
int64 is specificlly telling Go to use int for a 64bit scenario

**Implicit declaration** : var name int = 0
This means GO will implicitly determine which datatype it is based on the value assigned

### 4. Printing to Console with fmt

Tested a number of formatting tips for printing data to console
