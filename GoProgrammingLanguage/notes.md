# Go Programming Language Notes

## Chapter 1: Tutorial

### 1.1: Hello World
packages, not modules or libraries.
Very dogmatic about certain things, seems a bit pythonic in that regard.
Can only import used dependencies.
All files need a package declaration.
`package main` and `func main()` are pretty important.

### 1.2: CLI
Slices get metioned, which they claim to be important.
Args seem straight foward.
You can initialize by type all at once.
I love that it starts with the zero value and not null. IE `var s string` initializes to an empty string. <3
`:=` is also awesome. Declares the type based on the initializer's value.
Kind of weird for loops don't have parens though.

Seems like they use for loops for whiles, which is a neat optimization.

_for initialization; condition; post {}_

Dang, no unused variables. Might be annoying for people who autosave. For all the goofiness, it definitely seems like a Serious Programming Language.

_ is used conventionally for blank identifiers. when syntax requires something but program logic does not.

`range` produes a pair of values, hence the os.Args[1:] for each loop [index, value]
 k
`_, arg := range os.Args[1:]` =>
assign the value to _ and arg gets the key value.

Generally they recommend that you use either `s := ""` or `var s string` when declaring variables with empty values. can only var-less decl in a function, not at a package level.

Oh god, the go format default uses tabs. Nooooooooo!

