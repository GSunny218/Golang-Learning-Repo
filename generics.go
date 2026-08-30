package main

// import (
// 	"fmt"
// )

// func printSlice[T any](items []T) {  // Can take any type as parameter
// 	for _, item := range items {
// 		fmt.Println(item);
// 	}
// }

// func printSlice[T int | string | bool](items []T) {  // Can take int, string or bool as parameter
// 	for _, item := range items {
// 		fmt.Println(item);
// 	}
// }

// func printSlice[T comparable](items []T) {  // Can take any type as parameter which is comparable
// 	for _, item := range items {
// 		fmt.Println(item);
// 	}
// }

// func printSlice[T comparable, V string](items []T, name V) {  // Can take any type as parameter which is comparable
// 	for _, item := range items {
// 		fmt.Println(item, name);
// 	}
// }

// type stack[T any] struct {
// 	elements []T
// }

// func main() {
// 	// nums := []int{1,2,3,4,5};
// 	// names := []string{"golang", "java"};
// 	// booleans := []bool{true, false, true};
// 	// printSlice(nums);
// 	// printSlice(names);
// 	// printSlice(booleans);
// 	// printSlice(nums, "John");

// 	myStack := stack[int]{
// 		elements: []int{1, 2, 3, 4},
// 	}
// 	myStringStack := stack[string]{
// 		elements: []string{"golang", "java"},
// 	}
// 	fmt.Println(myStack)
// 	fmt.Println(myStringStack)
// }
