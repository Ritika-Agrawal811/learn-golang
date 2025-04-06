# Structs 📂
Structs are collections of **named fields** of _different types_. It helps us to define **custom data types** that can hold multiple related values.

## :sparkles: Defining a `struct`
A `struct` is defined using the `type` keyword, followed by the _name_ of the struct and its **fields with types** inside `curly braces {}`.

```go
type Person struct {
    Name string
    Age  int
}

type Student struct {
    Name   string
    Age    int
    Grade  string
    Marks  float64
}

type Car struct {
    Brand   string
    Model   string
    Year    int
    Mileage float64
}
```

## :sparkles: Initializing a `struct`
To _use_ a struct, we need to create an **instance** of it. Go allows _multiple_ ways to initialize a struct.

### a) Using a Struct Literal

```go
p2 := Person{"Bob", 30}
fmt.Println(p2) // Output: {Bob 30}
```
- Here, the order of values must match the struct fields.

### b) Using Named Fields

```go
p3 := Person{Name: "Charlie", Age: 35}
fmt.Println(p3) // Output: {Charlie 35}
```
- This is the recommended way
- We can omit fields and initialize them later.

### c) Using `new` Keyword

```go
p4 := new(Person) // Returns a pointer to a zero-initialized struct
p4.Name = "David"
p4.Age = 40
fmt.Println(*p4) // Output: {David 40}
```

- Here, `new(Person)` returns a pointer `*Person` to a zero-initialized struct

## :sparkles: Anonymous Structs
These are used to define a quick one-time struct without explicity defining a _type_.

```go
student := struct {
    Name  string
    Grade int
}{"Eve", 90}

fmt.Println(student) // Output: {Eve 90}
```
- useful only when the struct is used at a  _single place_. When we want the same struct in **many files**, it is best to create a _separate type_ for it.

## :sparkles: Nested Structs
A struct can have another struct inside it.

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Age     int
    Address Address
}
```

- Here, `Employee` struct uses `Address` struct to define its field **Address**

When we _initialize_ such a struct it looks like this :point_down:

```go
 emp := Employee{
        Name: "John Doe",
        Age:  35,
        Address: Address{
            City:  "New York",
            State: "NY",
        },
    }
```
To access a field inside **Address** field, we _again_ use the `dot (.) operator` :point_down:

```go
fmt.Println(emp.Name, "lives in", emp.Address.City, emp.Address.State) 
```

## :sparkles: Embedded Structs
Struct embedding allows one struct to be included _inside_ another. This enables **composition** which is a form of code reuse that is an alternative to _inheritance_ in other languages.

Embedding allows **embedded struct fields** to be accessed _directly_ without explicit nesting.

```go
type Person struct {
    Name string
    Age  int
}

type Student struct {
    Person   // Embedded struct
    Grade    string
    School   string
}
```
- Here, `Person` struct is directly embedded within `Student` struct.
- We can access `Name` and `Age` directly from `Student` without writing `s.Person.Name` because Go **promotes** the fields of embedded structs.

```go
/* initializing part is similar to nested structs
 only the field name is same as embedded struct type */
s := Student{
        Person: Person{Name: "Alice", Age: 16},
        Grade:  "10th",
        School: "Greenwood High",
    }

fmt.Println(s.Name, "is", s.Age, "years old and studies in", s.Grade) 
```
:warning: **Note:** Embedding and Nested structs both are initialized in the _same_ way. Embedded fields are _promoted_ to top-level whereas nested fields require an _extra level_ to access them.

## :sparkles: Struct with Methods
In Go, **methods** are functions that are _associated_ with a specific struct type. They allow us to **add behavior** to structs, similar to how methods work in object-oriented programming.

### a) Defining methods for a struct
A **method** in Go is just a function with a **receiver**. The _receiver_ is the struct type on which the method operates.

```go
type Person struct {
    Name string
    Age  int
}

// Method for Person struct
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}
```
- Here, `(p Person)` before the function name _associates_ the method `Greet()` with `Person` struct.

### b) Using the methods defined for a struct
We can _call_ a method defined for a struct by using its _instance_.

```go
 person := Person{Name: "Alice", Age: 25}
 person.Greet() // Calling the method
```
- Here, _person_ is an instance of struct `Person`. We call its method `Greet()` by calling it on the _instance_ like in any OOP language.

<p align="center">· · · · · · · · · · · · ·</p>

# Composition 🧱 
Composition is a **design principle** where _smaller_, _reusable components_ are combined to build complex systems.

In Go, composition is _preferred_ over inheritance, meaning that instead of defining a _hierarchy_ (like in object-oriented languages), we build complex types by **combining** multiple independent components.

 ### Key Benefits :
✅ More flexible than inheritance.
✅ Encourages reusability and modularity.
✅ Avoids deep hierarchies and tight coupling.
✅ Promotes cleaner, more maintainable code.

<p align="center">· · · · · · · · · · · · ·</p>

# Pointer vs. Value Receiver in Methods ⚙️
The _receivers_ defined in the struct methods can be of **two types** &mdash; _Pointer_ and _Value_

## :sparkles: Pointer Receiver
A pointer receiver `(*Type)` in a method allows the function to _modify_ the struct's fields and _avoid_ unnecessary copying of large structs.

```go
type File struct {
    content string
}

func (f *File) Write(data string) {
    f.content = data
}

func (f *File) Read() string {
    return f.content
}
```

- Here, both methods `Writer` and `Read` use _pointer receiver_.

## :sparkles: Value Receiver
A _value receiver_ `(Type)` means the method gets a **copy of the struct** instead of modifying the original instance.

```go
type Point struct {
    x, y int
}

func (p Point) Display() {
    fmt.Println("Coordinates:", p.x, p.y)
}
```
- Here, we use a value receiver for method `Display`.

## :sparkles: When to use which?

| **When to Use**                                      | **Pointer Receiver (`*Type`)** | **Value Receiver (`Type`)** |
|------------------------------------------------------|----------------------------------|------------------------------|
| Modifies struct fields?                         | ✅ Yes                           | ❌ No                         |
| Struct is large?                                 | ✅ Yes                           | ❌ No                         |
| Implements an interface requiring pointer methods? | ✅ Yes                           | ❌ No                         |
| Method only reads data?                          | ❌ No                            | ✅ Yes                         |
| Struct is small?                                 | ❌ No                            | ✅ Yes                         |


### 🔑 Quick Rule of Thumb

- Use **pointer** receiver when _modifying_ data or for large structs.
- Use **value** receiver when the method **only reads** data and doesn’t need to _modify_ the struct.

<p align="center">· · · · · · · · · · · · ·</p>

# JSON Struct Tags 🧬 
In Go, struct tags are **metadata** attached to struct fields. When working with JSON, these tags tell the `encoding/json` package how to _encode_ (marshal) and _decode_ (unmarshal) data.

### Syntax :
```go
FieldName Type `json:"jsonKeyName,option1,option2,..."`
```
- **jsonKeyName :** name to use in the JSON output/input
- **Options like :**
    - **omitempty :** omit the field if it's empty/zero.
    - **dash(-) :** exclude the field from JSON entirely.

### Example :
```go
type Person struct {
    Name string `json:"name"`
}
```
- Here, `json:"name"` tells the **JSON encoder/decoder** to map the Go `Name` field to/from a JSON field called `name`.

When we marshal this `Person` struct, we get:
```json
{"name": "Sara"}
```

## :sparkles: JSON Struct Tag Options

### a) omitempty
Omits the field if it has the _zero value_ for its **type**.

| Type     | Zero Value |
|----------|------------|
| `string` | `""`       |
| `int`    | `0`        |
| `float`  | `0.0`      |
| `bool`   | `false`    |
| `slice`  | `nil`      |
| `map`    | `nil`      |
| `pointer`| `nil`      |

#### Example :

```go
type Movie struct {
    Title string `json:"title,omitempty"`
    Year  int    `json:"year,omitempty"`
}
```
Suppose `Year` is 0, then output becomes :
```json
{"title": "Inception"}
```
#### 🧩 Special Considerations with `omitempty` and Structs
**omitempty** works as expected for most types — but **not** plain structs.

Suppose we have :
```go
type Response struct {
    Data Movie `json:"data,omitempty"`
}
```

Where `Movie` is a struct type as :
```go
type Movie struct {
    Title string `json:"title,omitempty"`
    Year  int    `json:"year,omitempty"`
}
```

If Data is **empty (Movie{})**, it will still be _included_ in the output as :
```json
{
    "data": {
        "title": "",
        "year": 0
    }
}
```

To solve this, we can use a **pointer** :
```go
type Response struct {
    Data *Movie `json:"data,omitempty"`
}
```
Here we can mention Data as `nil` and it will be entirely omitted by `encoding/json` package.