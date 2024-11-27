package main

import "fmt"

func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

func main() {
    var n int
    fmt.Println("Hello, welcome to the Fibonacci sequence!")
    fmt.Print("Please enter a number: ")
    
    fmt.Scanln(&n)

    fmt.Println("The Fibonacci sequence for", n, "is:", fib(n))
}
