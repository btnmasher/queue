# queue

Simple to use FIFO data structures.

### Usage

Here's a working example of how to use the Queue.

```Go
package main

import (
    "fmt"

    "github.com/btnmasher/queue"
)

// Declare
var unbounded queue.Queue
var bounded queue.Queue
var clearme queue.Queue

// Instantiate
func init() {
    unbounded = queue.NewUnbounded()
    bounded = queue.NewBounded(3)
    clearme = queue.NewUnbounded()
}

// Use
func main() {
    words := []string{"This", "was", "a", "triumph"}

    for _, word := range words {
        unbounded.Add(word)
        clearme.Add(word)

        if err := bounded.Add(word); err != nil {
            // Queue maximum was reached.
            fmt.Println(err)
        }
    }

    clearme.Clear() //Demonstrating clearing queues

    // Demonstrating getting queue length.
    fmt.Printf("Bounded Length: %v\n", bounded.Len())
    fmt.Printf("Unbounded Length: %v\n", unbounded.Len())
    fmt.Printf("Cleared Length: %v\n", clearme.Len())

    for {
        out1 := unbounded.Take()
        out2 := bounded.Take()

        if out1 != nil { // Take() returns nil if nothing is enqueued
            fmt.Printf("Unbounded: %v\n", out1.(string))
            // Using type coercion, nodes are of type interface{}
            // (though not necessary with fmt.Println, just
            // done as a demonstration.)
        }

        fmt.Printf("Bounded: %v\n", out2)

        fmt.Printf("Cleared: %v\n", clearme.Take())
    }
```

Output:

```
Unable to add value, queue is full.
Bounded Length: 3
Unbounded Length: 4
Cleared Length: 0
Unbounded: This
Bounded: This
Cleared: <nil>
Unbounded: was
Bounded: was
Cleared: <nil>
Unbounded: a
Bounded: a
Cleared: <nil>
Unbounded: triumph
Bounded: <nil>
Cleared: <nil>
```
