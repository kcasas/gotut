# Control Statements

## Looping

- The `for` loop is Go's only looping construct.
- `for` loop has three components separated by semicolons
    - init statement: executed before the first iteration
    - condition expression: evaluated before every iteration
    - post statement: executed at the end of every iteration
- every component is optional

syntax:
- normal
    ```go
    sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
    ```
- while
    ```go
    sum := 1
    for sum < 100 {
		sum += sum
	}
    ```
- do while
	```go
    sum := 10
    for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
    ```

## If-Else

- Go's `if` statements are similar to other languages.
- Go does __not__ have a `ternary` operator, use `if-else` statements instead.

syntax:
```go
age := 20
if age < 13 {
    fmt.Println("child")
} else if age < 20 {
    fmt.Println("teenager")
} else {
    fmt.Println("adult")
}
```

## Switch

`switch` statements in Go are just a shorter way to write `if-else` statements.

- `break` statement is not needed at the end of the case.
- `case` need not be constants. Expressions and function calls are evaluated.
- `value` to compare with each `case` is optional. If not provided, `cases` should evaluate to a `boolean` value.

syntax:
- w/o value to compare
	```go
    age := 20
    switch {
        case age < 13:
            fmt.Println("child")
        case age < 20:
            fmt.Println("teenager")
        case age < 30:
            fmt.Println("young adult")
        case age < 50
            fmt.Println("adult")
        default:
            fmt.Println("boomer")
        }
    }
    ```
- w/ value to compare
    ```go
    position := 1
    switch position {
        case 1:
            fmt.Prinln("Hard Carry")
        case 2:
            fmt.Println("Mid")
        case 3:
            fmt.Println("Offlane")
        case 4:
            fmt.Println("Roamer Support")
        case 5:
            fmt.Println("Hard Support")
        default:
            fmt.Println("Spectator")
    }
    ```

## Defer

- Defer delays function calls until the surrounding function returns.
- Multiple defer statements in a function are pushed onto a stack (LIFO).

syntax:
```go
defer {function_call/s}
```

_play_: https://play.golang.org/p/Y-XzOOM5rUd
