package main

import "fmt"

func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

func main() {
    fmt.Println("Hello, welcome to the Fibonacci sequence!")
    fmt.Println("Please enter a number:")
    
    var n int
    fmt.Scanln(&n)

    fmt.Println("The Fibonacci sequence for", n, "is:", fib(n))
}
