package main

import "fmt"

/*
	- Array:
		* Fixed size (length must be defined).
		* Elements are of the same type.
	- Slice:
		* Dynamic size (can grow or shrink).
		* Backed by an underlying array.
		* Has length and capacity.
	- Map:
		* Key-value pairs.
		* Keys must be unique.
		* Great for lookups.
*/

func main() {

	fmt.Println("==== Array Example ====")
	// Array must have a defined length
	arr := [3]int{10, 20, 30}
	fmt.Println("Array:", arr)
	fmt.Println("Length of array:", len(arr))

	// Arrays are value types → assignment copies the entire array
	copyArr := arr
	copyArr[0] = 100
	fmt.Println("Original Array after copy modification:", arr)
	fmt.Println("Copied Array:", copyArr)

	fmt.Println("\n==== Slice Example ====")
	// Slice is dynamic and more commonly used
	sli := []int{1, 2, 3, 4}
	fmt.Println("Initial Slice:", sli)

	// Append adds new elements, automatically resizing if needed
	sli = append(sli, 5, 6)
	fmt.Println("After Append:", sli)

	// Slices can be derived from arrays
	subSli := arr[0:2] // slice of first two elements of array
	fmt.Println("Slice from Array:", subSli)

	fmt.Println("Slice Length:", len(sli))   // number of elements
	fmt.Println("Slice Capacity:", cap(sli)) // underlying array size

	fmt.Println("\n==== Map Example ====")
	// Maps store key-value pairs
	ages := make(map[string]int)
	ages["Mujammal"] = 27
	ages["Saju"] = 25
	ages["GoLang"] = 15 // just for fun 😅

	fmt.Println("Map:", ages)

	// Access a value
	fmt.Println("Mujammal's age:", ages["Mujammal"])

	// Check if key exists
	if val, ok := ages["NonExistent"]; ok {
		fmt.Println("Found key with value:", val)
	} else {
		fmt.Println("Key 'NonExistent' not found in map")
	}

	fmt.Println("\n==== Mixing Example ====")
	// Using map of slices
	studentSubjects := map[string][]string{
		"Mujammal": {"Math", "Go", "AI"},
		"Saju":     {"English", "History"},
	}
	fmt.Println("Map of Slices:", studentSubjects)

	// Add a new subject to Mujammal
	studentSubjects["Mujammal"] = append(studentSubjects["Mujammal"], "Cloud")
	fmt.Println("Updated Map of Slices:", studentSubjects)

	// Using slice of maps
	users := []map[string]string{
		{"name": "Mujammal", "role": "Developer"},
		{"name": "Saju", "role": "Designer"},
	}
	fmt.Println("Slice of Maps:", users)
}
