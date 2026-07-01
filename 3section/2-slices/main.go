package main

import "fmt"

func main() {
	 // slices are arrays but can be dinamic (pop / append) 
	 names := []string{} // If i specify the length of the array, it will be fixed and cannot be changed
	 names = append(names, "Lucas")
	 fmt.Println( names)

	 items := make([]int, 3, 5)

	 fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	 items = append(items, 1, 2, 3)
	 fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	 fmt.Printf("Items: %+v\n", items[3:])
}
