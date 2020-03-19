# go-optional
![badge](https://github.com/ThisIsClark/go-optional/workflows/Go/badge.svg)

The optional type of Golang

 ## Usage

Get it by "go get":
```go
go get github.com/ThisIsClark/go-optional
```
Import it in your project:
```go
import "github.com/ThisIsClark/go-optional"
```
Create a empty optional by:
```go
op := optional.OptionalEmpty()
```
Or create an optional with value
```go
optionalData := 1
op := optional.Optionalof(optionalData)
```