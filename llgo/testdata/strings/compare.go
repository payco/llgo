package main

func main() {
    x := "abc"
    y := "def"
    z := "abcd"

    println(x == x) // true
    println(x == y) // false
    println(x != x) // false
    println(x != y) // true
    println(x < x) // false
    println(x < y) // true
    println(y < x) // false
    println(x > x) // false
    println(x > y) // false
    println(y > x) // true

    println(x == z) // false
    println(z == x) // false
    println(x < z) // true
    println(x > z) // false
    println(z < x) // false
    println(z > x) // true

    println(x <= x) // true
    println(x <= y) // true
    println(x >= x) // true
    println(y >= x) // true
}

