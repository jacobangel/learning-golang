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

### 1.3 Finding Duplicate Lines

derp, := means you don't need the var heh.

bufio is a funny looking name for a project.
Unclear how i get things to print out when i run the command though.

Map declaration is... od `make(map[string]int)` I guess it makes sense if you say it out loud but wow. I guess the 0 value thing is pretty awesome. Does it have no concept of null?

Printf stuff seems rather standard.
It's odd, this example doesn't seem to work? How do I break out of `input.Scan()`?  hmmm oh thanks [google/SO `Ctrl+D`](https://stackoverflow.com/questions/34481065/break-out-of-input-scan/34481857). Huh, my first item always does `<N>D`  instead of just printing the number. Kind of odd. Might be a consequence of how I have to send the end of input.

While we can declare functions in any order, I think they should have put the countLines first if this were a real program. I can see that from a teaching standpoint though, it's probably better to introduce the function second. What's a `*` mean exactly.

I did the exercise for 2, but i made it more interesting for splitting on spaces.

### 1.4 Animaged GIFs

Import symantics. `image/color` => color.White. Curious how namespace issues are resolved.
constnts can only be numbers, strings, or booleans.

[]color.Color{} wtf is this. 'Composite literals', for noting slicese and structs. Wish i had learned more C lol. Watch out for (§4.2, §4.4.1)








