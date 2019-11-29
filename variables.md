# Variables

## Syntax

Go Lang variables are statically typed. Once a type is assigned to a variable it cannot be changed. Assigning a value with a different data type will error out.

- **declaration**
    ```go
    var {name} {type}
    ```
- initialization 
    ```go
    var {name} {type} = {value}
    ```
- type inference
	```go
	var {name} = {value}
	```
- **short hand w/ type assertion**
    ```go
    {name} := {value}
    ```
- **assignment after declaration/initialization**
    ```go
    {name} = {value}
    ```
- **constants**
    ```go
    const {name} = {value}
    ```
- name and value can be a comma seperated list
    ```go
    {name, name1, ...} := {value, value1, ...}
    ```
- **multiline declaration/initialization**
    ```go
    const(
        {name} = {value}
        {name1} = {value1}
        ...
    )

    var(
        {name} {type}
        {name1} := {value1}
        ...
    )
    ```

## Scope
- **function level** variables
  - can only be used inside the function they were declared.
  - Go errors out if a these variables are not used.
- **package level** variables
  - declared outside of functions
  - can be used anywhere on the package.
  - keeps its value throughout the execution of the program.
  - does not need to be used
- Go does **not** have **global variables**.

## Variable naming

- Can only contain `alphanumeric` characters and `_`.
- **Cannot** start with a `number`.
- Go recommends writing variable names as `camelCase` words.
