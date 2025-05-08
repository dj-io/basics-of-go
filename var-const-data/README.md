<h1 align="center"><img style="width: 50px; height: 50px;" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmiro.medium.com%2Fv2%2Fresize%3Afit%3A1200%2F1*i2skbfmDsHayHhqPfwt6pA.png&f=1&nofb=1&ipt=a4c2f7e24c09760336433fc4e06c7f6c3f82d0fc57ab93c14699dec61d419415" /> </br> Variables, Constants & Data Types </h1>

<h3 align="center">This chapter covers fundamental syntax rules and project structure of go, creating and running a multi-module, go's capability to infer data types, working with packages (function scope, accessing data from imported packages), finally it touches on arrays, slices and maps</h3>

###  Table of Contents
  - [Go Basics Overview](#go-basics-overview)
  - [Create a Go Module](#create-a-go-module)
  - [Variables, Types, Const](#variables-types-const)
  - [Data Types](#data-types)
  - [Packages](#packages)
  - [Packages Q&A](#packages-q-a)
  - [Visibility](#visibility)
  - [Numbers & Collections](#numbers-&-collections)

---

## Go Basics Overview

<img src="../assets/go-basic-rules.png" />

<details>
 <summary>Learn More..</summary>

 - **Language Types**
    - Interpreted Languages -> Source Code `JS` `Python` `Ruby`
        - You write code and ship source code.
    - Intermediate Compiled Languages -> Bytecode `Java`
        - These languages compile and then ship to bytecode, this allows it to run on any os without specification `JVM`
    - **Compiled Languages -> Machine Code** `Go`
        - These languages compile and ship machine code for one cpu within the context of one os, these languages require you to specify the cpu & os to run it across platforms
    - **Go Language Types**
        - Go *ALSO* can be shipped in all three formats
            - Source code (transpiling into javascript)
            - bytecode (webassembly code)
            - machine code (what this course covers and most go applications in the wild will ship machine code)
    - **Note**
        - Go is more comparable to C than javascript overrall, thus the machine code use over source or bytecode
- **Basic Rules**
    - we use .go files
    - we use code blocks {}
        - we dont care about identation, spaces or tabs like python
    - no styling freedom
    - We do have semi-colon to seperate statements
        - only used if defining statements in the same line (optional)
        - new lines do not automatically mean new statement (the compiler understands the difference usually if rules are followed generally speaking)
    - case sensitive
        - named variables with lower case or upper case have different meanings
    - strongly typed
        - few use cases for dynamic typing
    - NOT an object-oriented language
        - no classes, no objects, no inheritance etc.
    - No exceptions
        - no try catch
        - go has a way to throw errors but not using exceptions
    - We have one file acting as the entry point (the main function)
    - A folder is a package
        - packages can have simple names, services, utily, data, llm (typical folder names)
    - Within one .go file we can have
        - variables
        - functions
        - type declarations
        - method declarations

- **Modules and CLI**
    - A module is a group of packages
        - a go module is simply a project or complete go application
    - Its our project
    - It contains a go.mod file with configuration and metadata

    - CLI manipulates the module
        - go mod init
        - go build
        - go run
        - go test
        - go get

</details>

## Create a go module

<img src="../assets/" />

<details>
 <summary>Learn More..</summary>

 - **How do you create a module?**
    - Create a new folder using your IDE or terminal
        - if using terminal run
        ```bash
        mkdir <module-name>
        ```
    - Now we are going to use the go cli to create the module
        - cd into the folder and run (terminal required)
        ```bash
        go mod init <module-name>
        ```
        - Naming convention can follow typical server application naming structure <app-services> <server> <backend> etc.
        - Optionally we can use a url to set the name, and this url does not have to exist on a network it is just a name set as a url i.e all of your modules can be prefixed with a company name `stratumlabs.ai/go/io`

    - Once you execute this command, a `go.mod` txt file will be created
        - This file contains all the metadata for our module, it should look something like this when initialized ->
            ```txt
            module stratumlabs.ai/go/io
            go 1.24
            ```
            - The name of the module
            - The version of go you want to use

        - Want to change the name of the module?
            - We can change the name of the module from this file by simply changing it!
        - Want to change the version?
            - We can change the version from this file by simply changing it!

    - *Summary*
        - So how do you create a go module?
            - you simply need a text file named go.mod with some go module metadata in the root of your project directory
        - The go.mod is similar to package.json in the javascript world, it will include any dependencies we bring in to our application similar to package.json

- *Whats the difference because we have a module*
    - To run our file from the first chapter, we used `go run <filename>` and the file name was mandatory to run the script/code in that file
    - Because we have a module we have more flexibility with how our code is run
        - i.e we can say I want to run code in this module, we can use `go run .` the . specifies the current module and go cli looks through all the files and looks for a package main with a function main. *this can be run from anywhere in the project and the main function can live in a file with any name, go cli will find the main function and run the module*
        - It is best practice to call the main file main.go but it is not mandatory, the only thing needed is `package main` `fun main() {}`

- *Workspaces and CLI*
    - Overview
        - In the first version of go, workspaces existed but it was deprecated
        - 1.18 workspaces were reintroduced

    - Workspaces
        - A way to manage multi module go applications, workspace is like a super module
        - It is a simple structure, you have the folder with several folders that will have several modules
        - Most applications in the wild are not multi-module so they dont have a workspace
        - it contains a go.work file with config metadata including which module to use

    - CLI
        - cli manipulates the worksapce
            - go work init

 - [inputoutput](../var-const-data/INPUTOUTPUT/) is the module example created for this section

</details>

## Variables, Types, Const

<img src="../assets/vars-types-const.png" />

<details>
 <summary>Learn More..</summary>

 - **Variables**
    - How do you declare a variable in go?
        - `var x int` (var keyword, x as name, int is type)
        - `var name string` (var keyword, name as name, string is type)
            - you create variables with var keyword, that is the only keyword go uses for variables
            - var keyword is followed with the identifier or name of the variable
            - the identifier is followed by the data type (no colon between name and type needed just a space)
            - any variable without an assignment (var x = 2, var name = "go"), they will be assigned nil by default (null or None)
        - `const y = 2`
            - declaring constants is similar to variables as it needs the const keyword, that is the only keyword go uses for constants
            - the const keyword is followed with the identifier or the name of the constant
            - then you assign the constant, when declaring constants we HAVE to include a value
            - constants can only be bool, string or numbers nothing else
                - constants in go are different than in js
                    - in js constants are set as immutable variables, which means it is a variable (a space in memory) that can not be changed (some engine checking that the value is not changed). Because it is just a space in memory you can assign any value, and the value simply can not be changed
                    - in go constants are constants which are fixed values, not a space in memory.
                    - when the compiler finds a constans reference it goes to the definition copies the value and pastes the value, with an exception for strings
                - you will not use const alot in go (urls, domains). Only things that will not change after compilation as a constant can not set its value at run time
        - `var z int = 2`
            - we can declare variables with initialization
        - `var text string`
        - `text = "hello!"`
            - we can declare variables without initialization and assign the variable later
            - strings with use double quotes, go does not recognize single quotes

        - `otherText := "Bye!`
            - we can use a variable initialization shortcut
            - using := (walrus operator), this will create the variable for you
            - we dont need the var keyword
            - this is used when you want to create the variable and assign a value
            - this syntax only works within functions, you can not set a global variable this way

[variable code examples](../var-const-data/INPUTOUTPUT/variable_declaration_examples.go)

</details>

## Data Types

<img src="../assets/data-types.png" />

<details>
 <summary>Learn More..</summary>

 - **string**
    - strings are declared with lower case s

- **integer values**
    - int (alias for int32)
    - int8, int16, int32, int64 (signed integer - can store from -127 to 128. if you want to set the amount of bits it will use in memory, this also sets the min and max that you can store there)
    - uint (alias for unit32)
    - uint8, uint16, uint32, uint64 (unsigned integer - can only store positive numbers from 0 - 155)

- **floating point values**
    - float32, float64 (in go and web, json every number in json will be converted to float64 so this will be commonly used with json data)

- **bool**
    - accepts true of false
    - *Boolean Operators*
        - == (equals),
        - != (not equals),
        - < (less than),
        - (> greater than),
        - <= (less than or equals)
        - (>= greater than or equals)
        - && (and)
        - || (or)
        - ! (not)

- **pointers**
    - lets say we have a variable we declared, then we create another variable that is pointing to the original this creates a link between our new variable and the old
        - pointers are simple in go, not as low level as c or c++

</details>

## Packages

<img src="../assets/" />

<details>
 <summary>Learn More..</summary>

 - go

</details>
