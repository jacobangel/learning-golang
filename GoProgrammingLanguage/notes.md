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

Omg types are so annoying sometimes.

### 1.5 Fetching a URL

Easy peasy; liking the error structure so far.

### 1.6 Concurrently Fetching URLs

Nothing too big, though it's funny how easy reddit will time you out.

### 1.7 A Web Server
Webserver stuff seems nice and straight forward so far.
Hmm... what is a mutex. I have always heard it as a mutex lock, but I honestly have no idea what that is.
Ugh auto save trimming packages while you're writing out an example is super annoying. Starting with ampersand.
Huh my example seems to increment the count a ton of times when i swap between count and othres.
`err := r.ParseForm(); err != nil` is an interesting idiom, though unclear how i feel with it.

The interface stuff looks really nice; great example with combining lissajous with the webserver.

### 1.8 Loose Ends

Control flow; switches are a bit interesting here, i like the tagless part. also t he no fall through, like duh.

"simple statements" need to dig in here for the differentaition.
Pointers are visisble. & gives the address, and *retrieves the varaible, but there is no pointer arithmetic.

golang.org/pkg and godoc.

## Chapter 2: Program Structure

Names begin with letters and case matters.
Something names like "int" you can redeclare but woof.
scope is function local, though package level variables may be accessed anywhere

Keywords:

- break -
- case -
- chan -
- const  -
- continue  -
- default -
- defer -
- else  -
- fallthrough  -
- for  -
- func -
- go -
- goto -
- break -
- case -
- chan -
- const  -
- continue  -
- default -
- defer -
- else  -
- fallthrough  -
- for  -
- func -
- go -
- goto -
- if -
- import -
- interface -
- map -
- package -
- range -
- return -
- select -
- struct -
- switch -
- type -
- var -

Constnats: true, false, iota, nil
Types: int, int8-64, uint, uint8-64, uintptr, float32/64, complext, bool, byte rune string error
Functions:
- make
- len
- cap
- new
- append
- copy
- close
- delete
- complex
- real
- imag
- panic
- recover

Whoa, names that start with a capital letter are exported. fmt.Print...

lowerCamelCase, with acronyms always uppercase

### 2.2 Declarations

four major types: var, const, type, and func

Oh things return value / error tuples a lot. I guess that's what you get with no exceptions. ALos it's really noisy when you ignore an error value.

I'm still a bit iffy on var i = asfd; vs i := asdf
tuple expansion is nice.

Short variable declaration are mostly for local variables that don't need to be explicitly types. `:=` is declaration `=` is assignment. So `var i = 100`  is both declaration and assignment.

sometimes short decl doesn't assign either. Also, short decl cannot reuse all the values:
`in, err :=` followd by `out, err` will reuse `err`
`f, err :=` followed by `f, err :=` will error though.

# 2.3.2 Pointers

Refresher:

variables are storage containign a value. They are referenced by a name or by an expression, `x` or `x[0]`.
A _pointer_ value is the _address_ of a variable. Not every value has an address, but all variables do. Pointers allow us to read or update a value of a variable indirectly, knowing nothing else about it.

```go
var x int
var xPointer *int
xPointer = &x
// &x = address of x yeilds a pointer
// *int = pointer to integer
```

Addressable values.
Oh shit flags are awesome.

### 2.3 Variables
The whole reassignment ruling is a bit odd sometimes. Also can see := vs = being confusing wihtout assistance from the IDE

### 2.4 Assignments
Tuples tuples tuples

### 2.5 Type Declarations
man i love named types


### 2.6 Packages and Files
Man, super magical in some regards.
Still not sure how I feel about the capitalization thing.
Wow package assembly is wild. Just crazy.
Imports are sensible, but I'll be honest,
I don't entirely understand how they go about just generally working.
Derp you have to put them in folders lol.
### 2.7 Scope
Lexical scoping with a Universal scope. The whole block scoping kind of throws me:

if:

if x := f(); x == 0 {
  fmt.PrintLine
} else if y := g(); y == x {
  fmt.Println(x, y)
} else {
  fmt.Println(x, y)
}
=> is the same as lik:
if x := f(); x == 0 {
  fmt.PrintLine
} else {
  if y := g(); y == x {
    fmt.Println(x, y)
  } else {
    fmt.Println(x, y)
  }
}

That said, its good they favor style that keeps code narrow and not marching off to the right.

## Chapter 3: Basic Data Types

### 3.1 Integers
Nothing too crazy here.
integers, 8 through 64, signed and unsigned.

`rune`s are synonyms for `int32`
`uintprt` unsepcified width but holds a pionter.

Operator precence goes in five tiers
1. "* / & << >> & &^"
2. "+ - | ^"
3. "== != < <= > =>"
4. "&&"
5. "||"

Similar groupings go left to right.
Overflow is silently discarded
All basic types can be compared.
Hah, I never realized why unsigned ints would be so hilariously bad for loops. Could never be less than 0 hah.
Conversion using types helps make things work better together.

### 3.2 Floating Point
Nothing surprising here.
Float32 / Float64, prefer 64 unless you have a good reason. NaN works like expected, though I am starting to really like the pattern of returning tuples.

boy typing up that example made me feel smart.



