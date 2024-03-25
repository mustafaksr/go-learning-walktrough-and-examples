# Go Learning Walkthrough and Examples

This repository contains a collection of example code and walkthroughs to help me learn Go programming language. The examples are based on the first six sections of the [Go Tour](https://go.dev/tour/welcome/) website. Additionally, there's a section named "07-examples" where you can find my go implementations.

## Table of Contents

1. [Packages, Variables, and Functions](#packages-variables-and-functions)
2. [Flow Control Statements: for, if, else, switch, defer](#flow-control-statements-for-if-else-switch-defer)
3. [More Types: Structs, Slices, Maps](#more-types-structs-slices-maps)
4. [Methods and Interfaces](#methods-and-interfaces)
5. [Generics](#generics)
6. [Concurrency](#concurrency)
7. [Examples](#Examples)

## 1. Packages, Variables, and Functions

The first section covers basic concepts such as packages, variables, functions, constants, and data types.

## 2. Flow Control Statements: for, if, else, switch, defer

This section introduces the various flow control statements in Go, including loops, if-else statements, and the defer keyword.

## 3. More Types: Structs, Slices, Maps

Learn about more complex data structures like structs, slices, and maps, along with concepts like pointers and range.

## 4. Methods and Interfaces

Explore object-oriented concepts in Go through methods and interfaces, and understand how Go implements them.

## 5. Generics

The generics section focuses on the introduction of generics in Go (as of my knowledge cutoff date), allowing for more reusable and flexible code.

## 6. Concurrency

Learn how to work with concurrency in Go using goroutines, channels, and the select statement.

## 7. Examples

In this section, you'll find a collection of your own examples and projects:

- **01-twosum_:** An example of solving the Two Sum problem.
- **02-palindrome_:** A program to check if a given string is a palindrome.
- **03-sqrt_:** Implementing the square root function using Newton's method.
- **04-strings:** Examples of working with strings, including regular expressions.
- **05-Docker:** Docker-related examples, including connecting to a PostgreSQL database.
- **06-romanrointeger:** Converting Roman numerals to integers.
- **07-removeduplicates:** Removing duplicates from a sorted array.
- **08-weightloss:** A weight loss tracking program.
- **09-remove-element:** Removing elements from an array.
- **app:** An application with various sub-projects.
- **other:** Implementations of the classic FizzBuzz problem and other related programs.
- **...**


## Usage

Each directory corresponds to a specific topic or example. Navigate to the directory you're interested in to find the associated Go source code files. You can execute these examples using the Go compiler.



## Directory Tree

```tree
.
├── 01-Packages-Variables-Functions
│   ├── 0_hello.go
│   ├── 0_packages.go
│   ├── 10_constant.go
│   ├── 11_numeric-constant.go
│   ├── 1_add2.go
│   ├── 1_add.go
│   ├── 3_named-results.go
│   ├── 4_multiple-results.go
│   ├── 5_variables2.go
│   ├── 5_variables3.go
│   ├── 5_variables.go
│   ├── 6_basic-types.go
│   ├── 7_zero.go
│   ├── 8_type-conversions.go
│   └── 9_type-inference.go
├── 02-flow-control-statements-for-if-else-switch-defer
│   ├── 0_for2.go
│   ├── 0_for.go
│   ├── 1_for-is-gos-while.go
│   ├── 2_forever.go
│   ├── 3_exercise.go
│   ├── 3_if-and-else.go
│   ├── 3_if.go
│   ├── 3_if-with-a-short-statement.go
│   ├── 4_switch-evalutaion-order.go
│   ├── 4_switch.go
│   ├── 5_switch-with-no-condition.go
│   ├── 6_defer.go
│   └── 6_defer-multi.go
├── 03-More-types-structs-slices-maps
│   ├── 10-slice-len-cap.go
│   ├── 11-nil-slice.go
│   ├── 12-making-slices.go
│   ├── 13-slice-of-slice.go
│   ├── 14-append.go
│   ├── 15-range2.go
│   ├── 15-range.go
│   ├── 16-range-continued.go
│   ├── 17-exercise.go
│   ├── 18-maps.go
│   ├── 19-maps-literals-continued.go
│   ├── 19-maps-literals.go
│   ├── 1-pointer.go
│   ├── 20-mutations-maps.go
│   ├── 21-exercise-maps.go
│   ├── 22-functions-values.go
│   ├── 23-function-closures.go
│   ├── 24-exercise-fibonacci.go
│   ├── 2-struct.go
│   ├── 3-struct-fields.go
│   ├── 4-struct-pointers.go
│   ├── 5-struct-literals.go
│   ├── 6-array.go
│   ├── 7-slices.go
│   ├── 7-slices-pointers.go
│   ├── 8-slice-literals.go
│   └── 9-slice-bounds.go
├── 04-Methods-and-interfaces
│   ├── 01-methods.go
│   ├── 02-methods-funcs.go
│   ├── 03-methods-cont.go
│   ├── 04-methods-pointers.go
│   ├── 05-methods-pointers-explained.go
│   ├── 06-indirection.go
│   ├── 07-indirectiom-values.go
│   ├── 08-methods-with-pointer-receivers.go
│   ├── 09-interfaces.go
│   ├── 10-interfac-are-satisfied-implicitly.go
│   ├── 11-interfaces.go
│   ├── 12-interface-values-with-nil.go
│   ├── 13-nil-interface-values.go
│   ├── 14-empty-interfaces.go
│   ├── 15-type-assertion.go
│   ├── 16-type-swithes.go
│   ├── 17-stringer.go
│   ├── 18-exercise-stringer.go
│   ├── 19-errors.go
│   ├── 20-exercise-errors.go
│   ├── 21-reader.go
│   ├── 22-exercise-reader.go
│   ├── 23-rot-reader.go
│   ├── 24-images.go
│   └── 25-exercises-images.go
├── 05-generics
│   ├── index.go
│   └── list.go
├── 06-concurrency
│   ├── 01-goroutines.go
│   ├── 02-chaneels.go
│   ├── 03-buffered-channels.go
│   ├── 04-range-and-close.go
│   └── 05-select.go
├── 07-examples
│   ├── 01-twosum_
│   │   ├── 01-twosum
│   │   └── go.mod
│   ├── 02-palindrome_
│   │   ├── 02-palindrome
│   │   └── go.mod
│   ├── 03-sqrt_
│   │   ├── 03-sqrt
│   │   └── go.mod
│   ├── 04-strings
│   │   ├── 01.go
│   │   └── regex.go
│   ├── 05-Docker
│   │   ├── 01-postgres-local
│   │   ├── 2.go
│   │   ├── go.mod
│   │   └── go.sum
│   ├── 06-romanrointeger
│   │   ├── go.mod
│   │   └── roman
│   ├── 07-removeduplicates
│   │   ├── go.mod
│   │   └── removeduplicates
│   ├── 08-weightloss
│   │   └── main.go
│   ├── 09-remove-element
│   │   ├── go.mod
│   │   └── remove
│   ├── app
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── image-app
│   │   └── main.go
│   └── other
│       ├── fizzbuzz.go
│       ├── flipcoing.go
│       ├── primefactors.go
│       └── tictactoe.go
└── README.md
```