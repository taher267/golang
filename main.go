package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constant
const A int = 5
const (
	B = 3
	C = 4
)

func longRunningTask() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func main2() {
	// number := "35"
	// var str string = "99================"
	// var a, b int = 5, 6
	// var c, d = 1, 9
	// e, f := 4, 6
	// fmt.Print(str, str)
	// var arr1 = [...]int{1, 2, 3}
	// arr2 := []interface{}{4, 5, 6, 7, 8, "jfksfd"}
	// fmt.Printf("str =%v & type: %T, %q", arr2, str, str)
	// fmt.Println(`Bismillah`, number, str, a, b, c, d, e, f)

	// myslice1 := []interface{}{1, 3, 5, "fdjfkdjfs"}
	// arr2 := []int{1, 2, 3}
	// arr3 := append(myslice1, arr2)
	// fmt.Println(arr3)

	// myslice1 := []int{1, 2, 3}
	// myslice2 := []int{4, 5, 6}
	// myslice3 := append(myslice1, myslice2...)
	// str1 := "text somenting"
	// fmt.Println(myslice3)
	// fmt.Println("\n\n")
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// neededNumbers := numbers[:len(numbers)-10]
	// numbersCopy := make([]int, len(neededNumbers))
	// fmt.Println(copy(numbersCopy, neededNumbers))

	// var a = make(map[string]string) // The map is empty now
	// a["brand"] = "Ford"
	// a["model"] = "Mustang"
	// a["year"] = "1964"
	// // a is no longer empty
	// b := make(map[interface{}]interface{})
	// b["kjfjsdfjdskfjsd"] = 1
	// b["2"] = 2
	// b[3] = 3
	// b["fksjfskafjsdkafjsdk"] = "fjsdkfjskafjsdka"

	// fmt.Printf("a\t%v\n", b)
	// // fmt.Printf("b\t%v\n", b)
	// fmt.Printf("kfldskfld\n")

	r := <-longRunningTask()
	fmt.Println(r)

}
