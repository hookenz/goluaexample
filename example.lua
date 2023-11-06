-- Call the Go function from Lua
HelloGo()

-- Call the Go function from Lua
result = HelloGoNumber()

-- Print the returned integer value
print("Returned value from Go:", result)


-- This is restricted.  Lets see if it works
file = io.open("example.lua", "r")

-- Print the file contents
print(file:read("*a"))