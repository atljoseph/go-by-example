
# Constants

```
const b = "another string"
```

# Variables

```
var a = "a string"
c := "yet another string"
var d, e int = 1, 2
t := []string{"g", "h", "i"}
```

# Format

```
fmt.Println("fragment1", "fragment2")
fmt.Printf("%T\n", myVar)
```

# If / Else

```
if a, b := 4, 5; a > b {
	fmt.Println("if")
} else if a < b {
	fmt.Println("else if")
} else {
	fmt.Println("else")
}
```

# Switch

```
switch i := 2 {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}
```

# For Loop

```
for {
    fmt.Println("loop")
    break
}
```
```
for j := 7; j <= 9; j++ {
    if (j == 8) {
        continue
    }
    fmt.Println(j)
}
```

# Arrays

```
var a [5]int
a[4] = 100
fmt.Println(a[0])
```

# Slices

```
t := []string{"g", "h", "i"}
c := make([]string, 5)
append(t, "j")
```

# Maps

```
n := map[string]int{"foo": 1, "bar": 2}
n["new-Key"] = 3
delete(n, "foo")
val, isPresent := n["foo"]
```



