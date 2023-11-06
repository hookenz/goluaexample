package main

import (
	"fmt"
	"log"

	lua "github.com/yuin/gopher-lua"
)

// Define a Go function that will be callable from Lua
func HelloGo(L *lua.LState) int {
	fmt.Println("Hello from Go!")
	return 0 // Number of return values
}

// Define a Go function that returns an integer
func HelloGoNumber(L *lua.LState) int {
	// Return an integer value to Lua
	L.Push(lua.LNumber(42))
	return 1 // Number of return values
}

func main() {
	// Create a new Lua state
	L := lua.NewState()
	defer L.Close()

	// Register the Go function in Lua
	L.SetGlobal("HelloGo", L.NewFunction(HelloGo))

	// Register our other func
	L.SetGlobal("HelloGoNumber", L.NewFunction(HelloGoNumber))

	// Set up a custom Lua environment with restricted functions
	L.SetGlobal("io", L.NewFunction(sandboxed))

	// Set up a custom Lua environment with restricted functions
	L.SetGlobal("os", L.NewFunction(sandboxed))

	// Run the Lua script
	if err := L.DoFile("example.lua"); err != nil {
		log.Fatal(err)
	}
}

func sandboxed(L *lua.LState) int {
	L.Push(lua.LString("not allowed"))
	return 1
}
