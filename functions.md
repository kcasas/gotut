# Functions

## Functions in Go
- can be anonymous w/ the option of immediately invoking it
- are types and thus can be assigned to a variables, use as parameters, and use as return values
- can be attached to custom types, like type aliases and structs

`func ({receiver}) {name}({parameters}) ({returns})`

- every component is optional, anonymous functions.
- `({receiver})` - a single variable declaration that attaches the function to a custom type
- `{name}` - name of the function
- `{parameters}` is a comma delimited list of variables
  - `{name} {type}, {name1} {type1}`
- `({returns})` - comma delimited list of returns w/ optional variable name and required type
  - `{name} {type}, {name1} {type1}`

## Examples

- Simple: https://gobyexample.com/functions
- w/ Receiver: https://gobyexample.com/methods

---
#### Coding Exercise
- create an interactive command line tool to manage parking lot spaces
- commands
  - `create_parking_lot {space}`
    - initialize parking lot space
  - `park {plate_number} {color}`
    - should print error message if parking space is full
    - return {slot}
  - `leave {slot}`
    - return {plate_number} {color}
  - `status`
    - prints the parking slots along with plate number and color of car parked
- Resources
  - https://golang.org/pkg/
    - String manipulation: https://golang.org/pkg/strings/
    - String conversion: https://golang.org/pkg/strconv/
  - simple REPL echoer: https://play.golang.org/p/UDA_HWBMhfW
---

