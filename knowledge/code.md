
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

# Functions

Multiple Return Values:
```
func multipleReturnVals() (int, string) {
	return 1, "two"
}
```

Variadic Functions:
```
func sum(desc string, nums ...int) int {
	fmt.Println(desc)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
nums := []int{1, 2, 3, 4}
fmt.Println(sum("Freeze, this is a stick up", nums...))
```

Closure Functions:
```
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
newInt1, newInt2 := intSeq(), intSeq()
fmt.Println(newInt1())
fmt.Println(newInt1())
fmt.Println(newInt2())
```

Recursive Functions:
```
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
fmt.Println(factorial(7))
```

# Structs

```
type cat struct {
	name      string
	meowIndex int
}
func quietKitty(meowPtr *int) {
	*meowPtr = 0
}
myCat := cat{"kitKat", 10}
yourCat := cat{name: "kittyCat", meowIndex: 3}
fmt.Println(myCat, yourCat)
makeZero(&myCat.meowIndex)
fmt.Println(myCat, yourCat)
```

# Methods

```
type rect struct {
	width, height int
}
func (r *rect) smaller() {
	r.width = r.width - 1
	r.height = r.height - 1
}
r.smaller()
```

# Interfaces

```
type geometry interface {
	area() float64
	perim() float64
}
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
c := circle{radius: 5}
measure(c)
```

# Errors

Basic Error:
```
func funcWithError() (int, error) {
	return -1, errors.New("houston, we have a problem")
}
if r, e := funcWithError(); e != nil {
    fmt.Println("error:", e)
} else {
    fmt.Println("success:", r)
}
```

Custom Error:
```
type argError struct {
	arg  int
	prob string
}
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func funcWithError() (int, error) {
	return -1, &argError{arg, "oh no, my cheese dip is cold"}
}
if r, e := funcWithError(); e != nil {
    fmt.Println("error:", e)
} else {
    fmt.Println("success:", r)
}
```

# Goroutines

Concurrency:
```
func f(from string) {
    for i := 0; i < 25; i++ {
        fmt.Println(from, ":", i)
    }
}
go f("1")
go f("2")
```

Channels:
```
resultChan := make(chan int)
go func() { resultChan <- 100 }()
result := <-resultChan
fmt.Println(result)
```