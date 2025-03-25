# Conditionals 🔠 
Just like any other programming language, Go also has conditional operators, `if`, `else if`, `else`, `switch` and other short circuiting techniques.


## :sparkles: `if`, `else if` and `else` statement
These are conditional statements to _control_ program flow based on **conditions**.

- `if` – executes a block of code if a condition is _true_
- `else if` – checks multiple conditions if previous ones fail
- `else` – runs if no previous conditions are met

:warning: **Note:** These statements, unlike some languages, do not use **parentheses**.

```go
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 75 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: C")
}
```

## :sparkles:  Short Variable Declaration in `if`
We can _declare_ a variable inside `if`, making it available only _within the block_. 

This limits the scope of the variable to the `if` block. The variable won't be available to the _parent scope_.

This code :point_down:

```go
length := getArrayLength(messages)
if length < 5 {
    fmt.Println("Not sufficient message length")
}
```

can be re-written as :point_down:

```go
if length := getArrayLength(messages); length < 5 {
    fmt.Println("Not sufficient message length")
}
```
We can add multiple variables too but they should be of **same type**.

```go
if x, y := 10, 20; x < y {
    fmt.Println("x is smaller than y")
}
```

## :sparkles: `switch` statement
Switch helps to compare a value against _multiple options_. In Go, we don't have to add `break` statement for each `case`. `break` is implicit in Go. It automatically stops after a **match**.

```go
day := "Monday"

switch day {
  case "Monday":
    fmt.Println("Start of the week!")
  case "Friday":
    fmt.Println("Weekend is near!")
  case "Saturday", "Sunday":
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("A regular day")
}
```

We can combine multiple cases by separating them with `commas (,)` :point_down:

```go
day := "Saturday"

switch day {
  case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("It's a weekday!")
  case "Saturday", "Sunday":
    fmt.Println("It's the weekend!")
  default:
    fmt.Println("Invalid day")
}
```