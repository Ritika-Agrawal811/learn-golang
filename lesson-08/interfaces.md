# Interfaces 🎭
An _interface_ in Go is a **composite data type** that can hold values of _different types_ as long as those types implement the interface.

Interfaces are **dynamically typed at runtime** — meaning an interface variable can hold different underlying types.

In Go, interfaces serve a _broader purpose_:

✅ They define _behavior_ for structs by defining a set of _method signatures_, they need to implement.
✅ They can _store values of any type_ that implements the interface (acting as dynamic containers).

## :sparkles: Why Use Interfaces?

- **Abstraction :** Allows functions to operate on different types as long as they implement _all_ the required methods.

- **Decoupling :** Reduces dependencies between components.

- **Polymorphism :** Enables writing flexible and reusable code.

## :sparkles: Defining an Interface
An _interface_ is declared using the `type` keyword, followed by the _interface name_ and then the keyword `interface`. Inside the `curly braces {}`, we mention all the _method signatures_.

```go
type Shape interface {
   area() float64
   perimeter() float64
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

## :sparkles: Using an Interface Variable
We can declare a variable for an _interface_. This allows us to store a value of _any type_ in the variable that implements that interface.

```go
var s Shape      // Declare an interface variable
s = Circle{5}    // Store a Circle inside the interface
fmt.Println(s.Area()) // Calls Circle's Area method
```

- Here, variable `s` is **explicitly** declared as interface `Shape`
- We can assign `s` value of a struct `Circle` because it implements `Shape`

Now suppose, we define another struct `Rect` which also implements `Shape` like this :point_down:

```go
/* Rect struct implements the interface - Shape */
type Rect struct {
   width  float64
   height float64
}

func (r Rect) area() float64 {
   return r.width * r.height
}

func (r Rect) perimeter() float64 {
   return 2 * (r.width + r.height)
}
```

We can assign a `Rect` struct to `s` because it also implements the interface `Shape` :point_down:

```go
s = Rect{10, 5}
fmt.Println(s.Area()) // calls Rect's Area method
```



## :sparkles: Key Properties of Interfaces

### a) Implicit Implementation
- Unlike other languages, Go does not require _explicit_ declaration that a type _implements_ an interface.

- If a type has **all methods** of an interface, it automatically satisfies the interface.

- In case of declaring a _variable_ for an interface, the interface **name** is specified just like any other _type_.

### b) Empty Interface
- The empty interface (`interface{}`) can hold _any type_ since every type implements zero or more methods.

- Useful for handling _unknown types_ (like `any` in modern Go).

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
```

### c) Composing Interfaces
- Interfaces can be _combined_ to form a **larger** interface.

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

### d) Nil Interface Behavior
An interface is **nil** if both the _type_ and _value_ inside it are nil.

```go
// Uninitialized, holds nil type and nil value
var i interface{} 

fmt.Println(i == nil) // ✅ True, because both type & value are nil
```

- Always check for `nil` before using an interface.

## :sparkles: Multiple Interfaces
Any struct can implement **multiple** interfaces. It just needs to implement all of their _methods_.

```go

type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(data string) {
    f.content = data
}
```

- Here, `File` struct defines methods for 2 interfaces &mdash; `Reader` and `Writer`.

## :sparkles: Type Assertions with Interfaces
When working with interfaces, we can _extract_ the **underlying type** using _type assertions_.

### Syntax :
```go
value, ok := interfaceVar.(ConcreteType)
```


- **interfaceVar** → The interface holding _some value_
- **ConcreteType** → The type we expect `interfaceVar` to hold
- **value** → The extracted value of type `ConcreteType`.
- **ok** →  a `bool` value which is **true** if assertion is successful and **false** otherwise

### :thought_balloon: What Does "Interface Holding Some Value" Mean?
when we declare an _interface variable_, it stores two things:

- The actual **type** of the value (type metadata).
- The actual **value** itself.

Think of it like a box:
📦 `interface{}` → (Type: int, Value: 42)

When we do _type assertion_, we are extracting the _value_ inside the interface.

```go
var i interface{} = 3.14 // i holds a float64

fmt.Printf("Type: %T, Value: %v\n", i, i)

// Output:
// Type: float64, Value: 3.14
```
- Here, variable `i` is holding a **float64** with value _3.14_

### Example :

```go
type Person struct {
    Name string
}

func main() {
    var p interface{} = Person{Name: "Alice"}

    person, ok := p.(Person)
    if ok {
        // ✅ Extracts Person struct
        fmt.Println("Person Name:", person.Name) 
    } else {
        fmt.Println("Type assertion failed")
    }
}
```


## :sparkles: Type Switch
A **type switch** is a special kind of _switch statement_ that works with _interfaces_. It allows us to check the **actual type** of a value stored inside an interface at runtime.

It is useful when we want to check for _multiple_ type assertions on an interface variable.

### Syntax :

```go
switch v := x.(type) {
case int:
    fmt.Println("x is an int:", v)
case string:
    fmt.Println("x is a string:", v)
case bool:
    fmt.Println("x is a bool:", v)
default:
    fmt.Println("Unknown type")
}
```

- `x.(type)` → extracts the _type_ of `x` inside the switch
- `v` → holds the actual _value_ if the type matches a case.

Here, each case _checks_ if `x` is a specific type.

### Example :

```go
package main

import "fmt"

func identifyType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case float64:
        fmt.Println("Float:", v)
    case string:
        fmt.Println("String:", v)
    case bool:
        fmt.Println("Boolean:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    identifyType(42)       // Integer: 42
    identifyType(3.14)     // Float: 3.14
    identifyType("hello")  // String: hello
    identifyType(true)     // Boolean: true
    identifyType([]int{1, 2, 3}) // Unknown type
}
```