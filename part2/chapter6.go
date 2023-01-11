package main

import (
	"errors"
	"fmt"
	"strconv"
)

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than zero")
	}

	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than zero")
	}

	if index >= len(orig) {
		return orig, nil
	}

	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

func chapter6() {
	/*---Chapter 6: slicing and dicing using arrays and slices---*/

	// Declaring an array
	var array_ [10]int
	fmt.Println(array_) // just like python :)

	// Initializing an array
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// you can omit the size with ...
	nums2 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(len(nums2))

	// working with multidimensional arrays
	var table [5][6]string
	for row := 0; row < 5; row++ {
		for col := 0; col < 6; col++ {
			table[row][col] = strconv.Itoa(row) + "," + strconv.Itoa(col)
		}
	}

	fmt.Println(table)

	// Creating and initializing a slice
	t := []int{1, 2, 3, 4, 5}
	fmt.Println(len(t))
	fmt.Println(cap(t))

	t = append(t, 6, 7, 8)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	t = append(t, 10, 9, 8)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	u := t //u and t are pointing to the same underlying array
	// Note: When your array exceeds its capacity, go will allocate the array to a new
	// segment in space and add the new element there.
	fmt.Println(u)
	fmt.Println(t)

	u[10] = 100
	fmt.Println(u)
	fmt.Println(t)

	// Slices: like array, but have the ability to grow and shrink in size.
	// also, slices are a reference type, not value type like array
	s := make([]int, 5)
	fmt.Println(s)

	// Creating and initializing a slice
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(v)

	// Slicing and Ranging
	var Operating_systems = [3]string{"IOS", "Android", "MAC"}
	fmt.Println(Operating_systems[0:2]) //First argument is starting index, next is range

	// Just like python :)
	for i, c := range v {
		fmt.Println(i, c)
	}

	// Making copies of an array or slice
	nums1 := [5]int{1, 2, 3, 4, 5}
	nums_2 := nums1

	fmt.Println(nums1)
	fmt.Println(nums_2)

	// Inserting into an array
	queue_ := []int{1, 2, 3, 4, 5}
	fmt.Println(queue_)
	queue_, err := insert(queue_, 2, 10)
	if err == nil {
		fmt.Println(queue_)
	} else {
		fmt.Println(err)
	}

	queue_, err = delete(queue_, 1)
	fmt.Println(queue_, err)
}
