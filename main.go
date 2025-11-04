package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3 =", 7.0/3)

	var a = "initial"
	fmt.Println(a)

	var b, c = "a", 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e string
	fmt.Printf("'%s'\n", e)
	fmt.Println("len =", len(e))

	f := "asdf"
	fmt.Println(f)

	const n = 500000000
	fmt.Println(n)

	arr := [...]int{1, 2, 3}
	var arr2 [3]int
	for i := range 3 {
		arr2[i] = i
	}

	fmt.Println(arr == arr2)

	var s []string
	fmt.Println("uniinit:", s, s == nil, len(s))

}
