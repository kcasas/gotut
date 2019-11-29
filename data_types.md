# Data Types

## Basic types

https://tour.golang.org/basics/11

## Zero Values

- `0` for numeric types
- `false` for boolean
- `""` (empty string) for strings
- `nil` for other data types

---
#### Coding exercise
- Create an app that will print out these details about you `firstname`, `lastname`, `age`, `gender`, `birthdate`, `seniorCitizen`, `deathCount`.
- Don't just print them out, declare variables and constants as needed.
---

## Pointers

Data type that holds the address of another variable.

syntax:
- `*{type}`, declare any variable to hold address of type
- `&{variable}`, get address of variable 
- `*{variable}`, set/get the value the variable pointer is pointing to

play: https://tour.golang.org/moretypes/1

## Arrays

Arrays stores a fixed length list of values with the same type.

syntax:
- declaration
    ```go
    var {name} [{length}]{type}
    ```
- nested
    ```go
        var {name} [{length}][{length}]{type}
    ```
- initialization
    ```go
    {name} := [{length}]{type}{ {value, value1, ...} }
    ```
- set/get element
    ```go
    {name}[{index}]
    ```

### Slices

- Slices are references to an underlying array.
- Syntax for arrays minus the length. 

syntax:
- slice another array/slice. both are optional w/ defaults low=0 and high={array_length}.
    ```go
    {name}[{low}:{high}]
    ```
- append element
    ```go
    {name} = append({name}, {value}, {value1})
    ```
- delete element/s at index n to m
    ```go
    {name} = append({name}[:{n}], {name}[{m+1}:]...)
    ```
- unshift element/s
    ```go
    {name} = append([]{type}{ {value}, {value1} }, {name}...)
    ```
- length and capacity
    ```go
        len({name})
        cap({name})
    ```
- make syntax: define capacity to avoid resizing of underlying array 
    ```go
        {name} = make([]{type}, {length}, {capacity})
    ```

play: https://play.golang.org/p/8DeztkO-w_7

## Maps

a list of keys mapped a value

syntax:
- declaration
    ```go
    var {name} map[{key_type}][{value_type}]{}
    ```
- initialization
    ```go
    {name} := map[{key_type}][{value_type}]{
        {key}: {value},
        {key2}: {value1},
        ...
    }
    ```
- set/get element
    ```go
    {name}[{key}]
    ```
- delete element
    ```go
    delete({name}, {key})
    ```
- check if key is present. If it is `ok` will have boolean value `true`.
    ```go
    _, ok = {name}[{key}]
    ```

play: https://play.golang.org/p/l7zacVQwA3A

## Structs

A `struct` is a **user defined type** that holds a collection of fields w/ different data types.

syntax:
- defining a struct type
    ```go
    type {struct_name} struct {
        {field_name} {type}
        {field_name1} {type1}
        ...
    }
    ```
- declaring a struct variable
    ```go
    var {name} {struct_name}
    ```
- iniatializing a struct variable
    ```go
    {name} := {struct_name}{value, value1, ...}
    ```
- initializing w/ explicit fields, some fields can be omitted
    ```go
    {name} := {struct_name}{
        {field_name}: {value},
        {field_name1}: {value1},
    }
    ```
- dot notation for setting/getting fields
    ```go
    {name}.{field_name}
    ```

play: https://play.golang.org/p/kDShFRDkxcM

---
#### Coding exercise
- Create a struct type that will store your profile details w/ fields like name, age, and skills
- name should also be a struct w/ first, middle, and last name.
- skills should be stored as a map w/ key as skill and then value as an array of flavors of that skill.
  - example: javascript [react, webpack]

[Answer](https://play.golang.org/p/bAFvZM6GHNS)

---


## Type conversion

- operations in Go only works on same data types
- arguments on function calls need to follow the data type
- Go does not automatically convert data types even if value would be the same. e.g. `int` and `int64`
- cannot convert multiple values at once

syntax:
```go
{type}({ value || name })
```

---
#### Coding exercise
- Calculate the area of circle with radius 7

[Answer](https://play.golang.org/p/bygQ2Q-4Q7k)

---


## Type aliasing

Any data type, even aliases, can be aliased in Go.

syntax:
```go
type {alias} {type}
```

---
#### Coding exercise
- Fix bug: https://play.golang.org/p/Lj64Xob-Gbw
---


