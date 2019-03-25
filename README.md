# gosharp
> Expose static definitions from go in C#
Expose constants from your go code as C# code

## Installation
```sh
go install github.com/mastern2k3/gosharp/cmd/gosharp
```

## Example

Given the following go code

```go
//go:generate gosharp -o ./opcodes.cs -classname Opcodes opcodes.go

package messages

const (
	SomeOpcode      = 101
	SomeOtherOpcode = 102
)
```

A csharp that looks as follows will be generated

```csharp
static class Opcodes {
    public const int SomeOpcode = 101;
    public const int SomeOtherOpcode = 102;
}
```
