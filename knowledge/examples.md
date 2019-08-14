
# Basic Structure

```
package main
import "fmt"
func main() { }
```

# Constants

```
const b = "another string"
```

# Variables

[Variables](../examples/03-a-variables.go)

[Conversions & Casting](../examples/03-b-convert-and-cast.go)


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

[Multiple Return Values](../examples/12-functions-multiple-returns.go)

[Variadic Functions](../examples/13-avariadic-functions.go)

[Optional Function Params](../examples/13-b-optional-function-params.go)

[Closures](../examples/14-closures.go)

[Recursive Functions](../examples/15-recursive-functions.go)


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
quietKitty(&myCat.meowIndex)
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

# Errors & Panic

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


[Panic](../examples/41-panic.go)

# Goroutines & Channels

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

Buffered Channels:
```
resultChan := make(chan int, 2)
resultChan <- 1
resultChan <- 2
// resultChan <- 3 // DANGER WILL ROBINSON
fmt.Println(<-resultChan)
fmt.Println(<-resultChan)
resultChan <- 4
fmt.Println(<-resultChan)
```

Channel Syncronization:
```
func action(done chan bool) { 
    time.Sleep(time.Second)
    done <- true
}
done := make(chan bool, 1)
go action(done)
// blocking
result := <-done
go somethingElse(result)
```

Channel Direction:
```
func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
pings := make(chan string, 1)
pongs := make(chan string, 1)
ping(pings, "passed message")
pong(pings, pongs)
fmt.Println(<-pongs)
```

Channel Multiple Return:
```
jobs := make(chan string, 5)
jobs <- "job one"
val, more := <- jobs
if more {
	fmt.Println("more jobs exist")
} else {
	fmt.Println("no more jobs")
}
```

# Range

Range over Array/Slice/Map:
```
for val, idx := range []int{7, 42} {
	fmt.Printf("Value: %d, Index: %d\n", val, idx)
}
```

Range over Channel (Includes Timers/Tickers):
```
queue := make(chan string, 2)
queue <- "one"
queue <- "two"
close(queue)
for elem := range queue {
	fmt.Println(elem)
}
```

# Select

Basic Usage:
```
c1 := make(chan string)
go func() {
	time.Sleep(1 * time.Second)
	c1 <- "one"
}()
c2 := make(chan string)
go func() {
	time.Sleep(2 * time.Second)
	c2 <- "two"
}()
for i := 0; i < 2; i++ {
	select {
	case msg1 := <-c1:
		fmt.Println("received 1", msg1)
	case msg2 := <-c2:
		fmt.Println("received 2", msg2)
	}
}
```

Timeout:
```
c1 := make(chan string, 1)
go func() {
	time.Sleep(2 * time.Second)
	c1 <- "result 1"
}()
select {
case res := <-c1:
	fmt.Println("c1", res)
case <-time.After(1 * time.Second):
	fmt.Println("timeout 1")
}
c2 := make(chan string, 1)
go func() {
	time.Sleep(2 * time.Second)
	c2 <- "result 2"
}()
select {
case res := <-c2:
	fmt.Println("c2", res)
case <-time.After(3 * time.Second):
	fmt.Println("timeout 2")
}
```

# Timers & Tickers

Timer:
```
timer1 := time.NewTimer(2 * time.Second)
timer2 := time.NewTimer(3 * time.Second)
timer3 := time.NewTimer(1 * time.Second)
fmt.Println("Timer 1 started")
<-timer1.C
fmt.Println("Timer 1 stopped")
go func() {
	fmt.Println("Timer 2 started")
	<-timer2.C
	fmt.Println("Timer 2 expired")
}()
fmt.Println("Timer 3 started")
<-timer3.C
fmt.Println("Timer 3 stopped")
stopped2 := timer2.Stop()
if stopped2 {
	fmt.Println("Timer 2 stopped")
}
```

Ticker:
```
ticker := time.NewTicker(500 * time.Millisecond)
go func() {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}()
time.Sleep(2600 * time.Millisecond)
ticker.Stop()
fmt.Println("Ticker stopped")
```


# Managing Goroutines

[Worker Pools](../examples/33-worker-pools.go)

[Wait Groups](../examples/34-wait-groups.go)

[Rate Limiting](../examples/35-rate-limiting.go)


# State Management

[Atomic Counters](../examples/36-atomic-counters.go)

[Mutexes](../examples/37-mutexes.go)

[Stateful Goroutines](../examples/38-stateful-goroutines.go)


# Sort

[Basic Sorting](../examples/39-sorting.go)

[Sort by Function](../examples/40-sorting-by-functions.go)


# Function Extensions

[Collection Functions](../examples/43-collection-functions.go)

[String Functions](../examples/44-string-functions.go)


# Other

[Defer](../examples/42-defer.go)


# Formatting & JSON

[String Formatting ("fmt")](../examples/45-string-formatting.go)

[Regular Expressions ("regexp")](../examples/46-regular-expressions.go)

[Json ("encoding/json")](../examples/47-json.go)




