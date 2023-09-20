package main

import "fmt"

func main() {
    var num int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scanln(&num)

    if isPrime(num) {
        fmt.Printf("%d adalah bilangan prima.\n", num)
    } else {
        fmt.Printf("%d bukan bilangan prima.\n", num)
    }
}

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    if num <= 3 {
        return true
    }
    if num%2 == 0 || num%3 == 0 {
        return false
    }

    for i := 5; i*i <= num; i += 6 {
        if num%i == 0 || num%(i+2) == 0 {
            return false
        }
    }

    return true
}