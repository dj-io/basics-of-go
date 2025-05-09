<h1 align="center"><img style="width: 50px; height: 50px;" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmiro.medium.com%2Fv2%2Fresize%3Afit%3A1200%2F1*i2skbfmDsHayHhqPfwt6pA.png&f=1&nofb=1&ipt=a4c2f7e24c09760336433fc4e06c7f6c3f82d0fc57ab93c14699dec61d419415" /> </br> Introduction </h1>

<h3 align="center">This chapter highlights the characteristics of functions in Go, the usage of pointers, less common control flow mechanims, panic and defer and error design patterns, then commonly used control structures. In this chapter Max walks through building a basic calculator project finally, we go over the use of the os std lib to read and write to local files </h3>

###  Table of Contents
  - [Functions](#functions)
  - [Pointers & References](#pointers-references)
  - [Panic, Defer, & Error Design Patterns](#panic-defer-error-design-patterns)
  - [Control Structures](#control-structures)
  - [Calculater Project](#calculator-project)
  - [Read Files](#reading-files)
  - [Writing Files](#writing-files)

---

## Functions

<img src="../assets/functions-cover.png" />

<details>
 <summary>Learn More..</summary>

 - **Functions**
    - Functions can recieve arguments
    - Function can recieve n number of arguments
    - Functions can have default values as args

    - **Functions In Go**
        - Functions can return more than one value (this is similar to function returns in python i.e, `return valOne, valTwo` -> `valOne, valTwo = funcReturningTwoVals(someArg)`)
            - can a functions return type be set to different values like return string or nil
                - No in Go you must return a type, unless you are creating your own types which is more of a design pattern than a Go feature
        - Labeled variables (seen as a bad practice in modern go development)
        - By default functions recieve arguments by value not reference like other high level languages
        - *What does 'recieve args by value and not reference' mean?*
            - by default, when passing a value as an argument to a function in Go you are not passing the reference to the variable that has the value. You are cloning the value from the variable and creating a new reference (variable) with that value (i.e, in js or python args are references to the original variable (for primitive data types))
        - *What if I dont want to copy and I want to pass in the reference? (You do)*
            - To pass a reference to the variable in a go function we have to use [*Pointers](#pointers--references)
        ```go
            // declaring a function
            func save() {} // all functions start with the func keyword

            // passing a param into a function
            func save(text string) {} // param structure: identifier type

            // returning a value
            func add(a int, b int) int { return a+b }
            // add the value type that will be returned after the param block
            // using the return keyword at the end of the function to return the desired value

            // returning multiple values
            func addAndSubstract(a int, b int) (int, int) { return a+b, a-b }
            // add the value types that will be returned after the param block wrapped in parenthesis with each value seperated by a comma

            // using the return keyword at the end of the function to return the desired values seperated by comma
        ```
    - **Pointers Preview**
        - Functions can recieve pointers instead of the value
        - Pointers allows you to pass references to a value instead of a copy of the value as mentioned in the `Function in Go` bullet point above

        -*Why do we want to reference a value instead of copy it?*
            - If the value passed in to the func will be changed

        ```go
        func increment(x *int) {
            *x++
        }

        func main() {
            i := 1
            increment(&i)
        }
        ```
 [Functions Examples](../func-control-err/main.go)

</details>

## Pointers & References

```go
func increment(x *int) {
    *x++  // dereference x and increment the value it points to
}

func main() {
    i := 1
    increment(&i)  // pass the address of i
    fmt.Println(i) // prints 2
}
```

<details>
 <summary>Learn More..</summary>

 - **Pointers**
   - A pointer is a variable that stores the **memory address** of another variable.
   - In Go, pointers are explicitly declared using the `*` symbol before a type, e.g., `*int` means "pointer to an int".
   - Pointers allow you to **reference** a variable's memory location rather than copying its value.
   - This explicit referencing is different from languages like Python or JavaScript, where references are handled implicitly.
   - You can do double pointers which is expressed by prefixing a value with two asteriks i.e `**int`
    - In this case you have a pointer that has a pointer to the value (you have a memory address pointing to a memory address that points to the value)
        - There are niche use cases for this type of behavior so it is not as commonly used
    - **What is a Reference?**
      - A reference holds the **memory address** where the data is stored.
      - In Go, references are explicitly handled using **pointers**.
      - References allow functions or variables to **access or modify** the original data without copying it.
      - This contrasts with passing by value, where a copy of the data is made and modifications do not affect the original.
        - passing by value: passing a param with prefixing with * or passing an arg without prefixing with & i.e `*int`, `&age`

 - **Why use pointers?**
   - You need to **modify** the original value inside a function (pass by reference).
        - Anytime the function will modify the incoming arguments, if you pass in args without using pointers you will make and modify a copy of the arg not the arg itself which makes no sense
   - To avoid copying **large data structures**, improving performance.
   - To share data efficiently between functions or goroutines.
   - To interface with low-level system calls or C libraries.

 - **Passing arguments in Go**
   - By default, Go passes function arguments **by value** (a copy of the variable).
   - To modify the original variable, you must pass a **pointer** to it.
   - This makes side effects explicit and code easier to reason about.

 - **Pointer syntax**
   - `*Type` in a function parameter means the function expects a pointer to `Type`.
   - `*variable` inside the function **dereferences** the pointer to access or modify the value.

 - **Example: Incrementing a value using a pointer**

    ```go
    func increment(x *int) {
        *x++  // dereference x and increment the value it points to
    }

    func main() {
        i := 1
        increment(&i)  // pass the address of i
        fmt.Println(i) // prints 2
    }
    ```
 - **Explanation:**
   - `increment` receives a pointer to an int (`*int`).
   - Inside `increment`, `*x` accesses the value at the address `x` points to.
   - `&i` gets the address of `i` to pass to the function.

 - **Summary Table**

| Concept                 | Go Syntax/Behavior                               | Description                                         |
|-------------------------|-------------------------------------------------|-----------------------------------------------------|
| Pass by value           | `func f(x int)` — copies `x`                     | Function receives a copy of the variable's value    |
| Pass by reference       | `func f(x *int)` — pointer to `x`                | Function receives a pointer (reference) to variable |
| Dereference pointer     | `*x` to access value at pointer                  | Access or modify the value pointed to by a pointer  |
| Get address of variable | `&x`                                            | Obtain the memory address of a variable              |
| Modify caller's var     | Use pointer and dereference                      | Change the original variable via pointer             |
| Passing address example | `f(&x)`                                            | Pass the address of variable `x` to function `f`    |
| Dereferencing example   | `*p = *p + 1`                                      | Increment the value pointed to by pointer `p`       |

[Pointer Examples](../func-control-err/main.go)


</details>

## Panic, Defer & Error Design Pattern

<img src="../assets/panic-defer-error.png" />

<details>
 <summary>Learn More..</summary>

 - go

</details>
