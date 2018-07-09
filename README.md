# Godash Slice Utilities

An implementation of utility functions inspired by [Lodash](https://lodash.com) for the Go language, focusing on efficiency without the use of reflection.

Some of the motivation behind this library is discussed in this [blog post](https://medium.freecodecamp.org/lodash-in-go-language-without-reflection-1d64b5115486).

## Methods

* `Reverse`
* `Uniq`
* `Filter`
* `Map`
* `Reduce`
* `Concat`
* `First`
* `Last`
* `Chain`
* `Value`

&nbsp;
#### `_.Reverse(slice)`

Returns a new slice in reverse order.

```go
_int.Reverse([]int{1, 2, 3})
// => []int{3, 2, 1}
```

#### `_.Uniq(slice)`

Returns a new slice without duplicates (all elements are unique).

```go
_int.Uniq([]int{1, 2, 1, 3, 3})
// => []int{1, 2, 3}
```

#### `_.Filter(slice, func)`

Returns a new slice of all elements which the function predicate returns true for.

```go
even := func (element int, index int) bool {
  return element % 2 == 0
}
_int.Filter([]int{1, 2, 3, 4}, even)
// => []int{2, 4}
```

#### `_.Map(slice, func)`

Returns a new slice of all elements after the function has been executed on them.

```go
double := func (element int, index int) int {
  return 2 * element
}
_int.Map([]int{1, 2, 3, 4}, double)
// => []int{2, 4, 6, 8}
```

#### `_.Reduce(slice, func, initial)`

Reduces the slice to a single value which is the accumulated result of running each element through the function.

```go
sum := func (acc int, element int, index int) int {
  return acc + element
}
_int.Reduce([]int{1, 2, 3, 4}, sum, 0)
// => int(10)
```

#### `_.Concat(slice, slice)`

Returns a new slice which is the first slice with the second concatenated at its end.

```go
_int.Concat([]int{1, 2, 3}, []int{4, 5})
// => []int{1, 2, 3, 4, 5}
```

#### `_.First(slice)`

Returns the first element in the slice.

```go
_int.First([]int{1, 2, 3, 4})
// => int(1)
```

#### `_.Last(slice)`

Returns the last element in the slice.

```go
_int.Last([]int{1, 2, 3, 4})
// => int(4)
```

#### `_.Chain(slice).Action().Action().Value()`

Chains multiple actions together and runs each on the result of the previous one. `Value()` returns the final result.

```go
_int.Chain([]int{1, 2, 1, 3}).Uniq().Reverse().Value()
// => []int{3, 2, 1}
```

&nbsp;
## Working with different types

In order to avoid inefficient reflection, the library creates dedicated implementations for each type you need.

In the original [Lodash](https://lodash.com), the library is used through the underscore character `_`. For example: `_.uniq()`.

We keep the same convention, except that the underscore is followed by the type. For example: `_int.Uniq()` for integers, `_string.Uniq()` for strings.

#### Primitive types (int, string, etc)

Simply import the relevant subset of the library with the type appearing after the underscore:

```go
import "github.com/go-dash/slice/_string"

func main() {
    _string.Uniq([]string{"aa", "bb", "aa"})
    // => []string{"aa", "bb"}
}
```

#### Custom types (structs)

Do the same thing, just add a comment afterwards of where the struct is defined.

```go
import "github.com/go-dash/slice/_Person" // github.com/my-user/my-repo/person

func main() {
    _Person.Uniq([]Person{Person{"John", 18}, Person{"Rachel", 17}, Person{"John", 18}})
    // => []Person{Person{"John", 18}, Person{"Rachel", 17}}
}
```

&nbsp;
## Installation

1. Install the `_gen` command line tool which generates code for the types you need:

    ```
    brew install go-dash/tools/gen
    ```
    > Verify with `_gen --version`
  
2. Go get the library:

    ```
    go get github.com/go-dash/slice
    ```
  
3. In your project, import the relevant subsets for the types you need:

    ```go
    import (
        "github.com/go-dash/slice/_string"
        "github.com/go-dash/slice/_int"
    )
    ```
  
4. Once, generate the code for the library subsets by running the following in your project root:

    ```
    cd my-project
    _gen
    ```
  
A working example is available in the [test suite](test.sh).

&nbsp;
## Running the tests

From the root directory run `./test.sh`.

&nbsp;
## Motivation

This is mostly a thought experiment coming from the love of lodash and missing basic functional operations for slices in golang. Implementations based on [reflection](https://github.com/robpike/filter) miss the point IMO since they're inefficient in runtime. Since golang doesn't currently support generics, code generation is the closest we can come.

If you like this approach and want to contribute to porting all of lodash to golang, don't hesitate to come help!  
