# Data Types 🔢 
Go is a _strongly_ and _statically_ typed language.

**Strongly typed** → Go enforces _type safety_, meaning we **cannot** implicitly **convert between different types** without explicit conversion.

**Statically typed** → The _type_ of a variable is determined at **compile time**, not at runtime.

## :sparkles: Basic/Primitive Data types
Go provides various built-in data types for storing data.

### a) bool
represents `true` or `false` data

```go
var isReady bool = true
hasComplete := false
```

### b) string
sequence of characters enclosed with _double quotes_

```go
var name string = "Ritika"
language := "Go
```
### c) int
a **signed** integer type. It's _size_ depends on the system &mdash; can be _32-bit_ or _64-bit_

```go
var age int = 25
```

### d) int8, int16, int32, int64
**signed** integers of different sizes

```go
 var a int8 = 127          
 // Range: -128 to 127

 var b int16 = 32767       
 // Range: -32,768 to 32,767

 var c int32 = 2147483647  
 // Range: -2,147,483,648 to 2,147,483,647

 var d int64 = 9223372036854775807 
 // Range: -9 quintillion to +9 quintillion
```

### e) uint8, uint16, uint32, uint64
**unsigned** integers of different sizes. They don't include negative numbers.

```go
var bigNumber uint64 = 5000
```

### f) float32, float64
floating point numbers &mdash; numbers with _decimals_

```go
var price float64 = 99.99
```

##### :thought_balloon: Why no `float8` or `float16`?
- Go follows the **IEEE 754 floating-point standard**, which defines only 32-bit and 64-bit floating points
- Most CPUs lack native float8 or float64 _support_
- Precision & accuracy are _too low_ in float8 or float64
- float32 and float64 cover _most real-world use cases_

### g) complex64, complex128
complex numbers with real & imaginary parts

```go
var c complex64 = 3 + 4i
```

### g) byte
Alias for **uint8**, represents _ASCII characters_

```go
var b byte = 'A'
```

### h) rune
Alias for **int32**, represents  _Unicode characters_

```go
var r rune = '😀'
```

#### :warning: Note: Go uses "double quotes" for strings and 'single quotes' for bytes and runes.

<p align="center">· · · · · · · · · · · · ·</p>

## :sparkles: Determining Range for Integers 
The _range_ of integer types is determined by the **number of bits** allocated for _storage_.

- The formula to calculate the range for a **signed integer type** (intN) is:

$$
\large \text{Range} = -2^{N-1} \text{ to } 2^{N-1} - 1
$$


- **Unsigned integers** do not store negative values, so their range is calculated using:

$$
\large \text{Range} = 0 \text{ to } 2^N - 1
$$

We can find the _minimum_ and _maximum_ values of different signed and unsigned integer types using the `math` package :point_down:

```go
    fmt.Println("Signed Integers:")
    fmt.Println("int8:  ", math.MinInt8, "to", math.MaxInt8)
    fmt.Println("int16: ", math.MinInt16, "to", math.MaxInt16)
    fmt.Println("int32: ", math.MinInt32, "to", math.MaxInt32)
    fmt.Println("int64: ", math.MinInt64, "to", math.MaxInt64)

    fmt.Println("\nUnsigned Integers:")
    fmt.Println("uint8:  ", 0, "to", math.MaxUint8)
    fmt.Println("uint16: ", 0, "to", math.MaxUint16)
    fmt.Println("uint32: ", 0, "to", math.MaxUint32)
    fmt.Println("uint64: ", 0, "to", math.MaxUint64)
```

**Output:**

```shell
Signed Integers:
int8:  -128 to 127
int16: -32768 to 32767
int32: -2147483648 to 2147483647
int64: -9223372036854775808 to 9223372036854775807

Unsigned Integers:
uint8:  0 to 255
uint16: 0 to 65535
uint32: 0 to 4294967295
uint64: 0 to 18446744073709551615
```

<p align="center">· · · · · · · · · · · · ·</p>

## :sparkles: Complex/Composite Data types
Composite data types are _data structures_ that can hold **multiple values**, often of different types. 

### a) Structs
Structs allows us to _group_ multiple fields of **different types** into a **single entity**.

**Syntax :**
```go
type StructName struct {
    PropertyName type
    PropertyName type
    // more such name-type pairs
}
```

**Example :**
```go
type Person struct {
    Name string
    Age  int
}
```
- Here, `Person` is a struct that holds two values &mdash; `Name` and `Age` of types _string_ and _int_ respectively.

To _use_ this struct, we create an instance of it :point_down:
```go
p := Person{Name: "Alice", Age: 30}
```

We can access individual fields using `dot (.)` operator :point_down:
```go
 fmt.Println("Name:", p.Name)
 fmt.Println("Age:", p.Age)
```