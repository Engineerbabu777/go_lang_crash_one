package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// DATA TYPES IN GO LANG!
	variableDaNa := 5
	var floatVariableDaNa float64 = 56.9
	var stringVariableDaNa string = "Babu Ji"
	var booleanVariableDaNa bool = true

	// ADDITION!
	fmt.Println("Addition: ", variableDaNa+int(floatVariableDaNa))
	// SUBTRACTION!
	fmt.Println("Subtraction: ", variableDaNa-int(floatVariableDaNa))
	// MULTIPLICATION!
	fmt.Println("Multiplication: ", variableDaNa*int(floatVariableDaNa))
	// DIVISION!
	fmt.Println("Division: ", floatVariableDaNa/float64(variableDaNa))
	fmt.Println("variables Di Value = ", variableDaNa, floatVariableDaNa)
	fmt.Println("variables Di String = ", stringVariableDaNa)
	fmt.Println("variables Di String = ", stringVariableDaNa+" is a good boy")
	fmt.Println("Length Of Variable String: ", len(stringVariableDaNa))
	fmt.Println("Boolean Variable: ", booleanVariableDaNa)

	var somethingGotOne, somethingGotTwo int = thisWillReturnTwoValues()
	thisWillReturnInteger()
	thisWillReturnString()

	var value, err = cannotDivideByZero(5, 5)
	if err != nil {
		println("Error: ", err.Error())
	} else {
		println("Value: ", value)
	}

	fmt.Printf("Something Got: %d and %d", somethingGotOne, somethingGotTwo)

	// fmt.Println("Something Got: ", somethingGot);

	// ARRAYS!
	// syntax!
	var myArray [3]int = [3]int{1, 2, 3}
	// var means thats the variable!
	// myArray means thats the name of array!
	// [3] means thats the size of array!
	// int means type of data stored in array!
	// {...,..} values separated by commas!

	fmt.Println("Array: ", myArray[0:3]) // ARRAY SLICING!
	fmt.Println("ADDRESS OF FIRST ELEMENT: ", &myArray[0])

	// SLICES!
	// syntax!
	// var mySlice[] string = []string{"Hello", "World", "How", "Are", "You"}; // WILL THROW ERROR IF NOT BEING USED!
	// var means thats the variable!
	// mySlice means thats the name of slice!
	anotherSlice := []string{"🍎", "🥭"}
	fmt.Println("anotherSlice: ", anotherSlice)

	// APPENDING NOW!
	anotherSlice = append(anotherSlice, "🍌") // ADDING BANANA TO SLICE!
	fmt.Println("anotherSlice: ", anotherSlice)

	// EXTRACT A RANGE OF SLICES // BREAK SLICES INTO PIECES!
	fmt.Println("anotherSlice[0:2]: ", anotherSlice[0:2]) // WILL RETURN FIRST TWO ELEMENTS

	// Create a slice with make
	dynamicSlice := make([]int, 3, 5) // len=3, cap=5
	fmt.Println("dynamicSlice: ", dynamicSlice)

	copySlice := make([]string, len(anotherSlice))

	copy(copySlice, anotherSlice)

	fmt.Println("copySlice: ", copySlice)
	fmt.Println("OriginalSlice: ", anotherSlice)

	// REMOVING ELEMENT FROM SLICES!
	fmt.Println("New Slice After Removing Banana= ", append(anotherSlice[:2], anotherSlice[3:]...))

	// CHECKING IF A SLICE CONTAINS ANY ELEMENT!
	fmt.Println("Contains Apple: ", contains(anotherSlice, "🍉"))

	test := []string{"🍎", "🍌", "🍊", "🍇", "🍉"}
	// APPENDING TWO SLICES!
	result := append(anotherSlice, test...)

	fmt.Println("Result: ", result)
	fmt.Println("Result Length: ", len(result))

	// MAPS!
	// Create and initialize a map
	var studentGrades map[string]int = map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	// Add a new key-value pair
	studentGrades["David"] = 88

	// Update an existing value
	studentGrades["Bob"] = 87

	fmt.Println("Student Grades: ", studentGrades["Alice"]) // ACCESSING THE KEYS TO GET VALUES!
	fmt.Println("studentsGrades: ", studentGrades)

	// Check if a key exists
	_, exists := studentGrades["Bob"] // VALUE | KEY EXITS RETURN TRUE!
	fmt.Println("Exists: ", exists)

	// Delete a key-value pair
	delete(studentGrades, "Bob")

	// Iterate over the map
	for name, grade := range studentGrades {
		fmt.Printf("%s's grade: %d\n", name, grade)
	}

	// Create a map of slices
	studentsCourses := map[string][]string{
		"Alice":   {"Math", "History"},
		"Bob":     {"English", "Physics"},
		"Charlie": {"Chemistry", "Biology"},
	}

	// Print the map
	fmt.Println(studentsCourses)

	 // Create a map with function values
	 operations := map[string]func(int, int) int{
        "add":    func(a, b int) int { return a + b },
        "subtract": func(a, b int) int { return a - b },
        "multiply": func(a, b int) int { return a * b },
    }

    // Use the function from the map
    result2 := operations["add"](3, 5)
	fmt.Println(result2);


	// STRINGS IN GO LANG!
	message := "String is here";
	fmt.Println("message: ", message);
	// STRING LENGTH!
	fmt.Println("message length: ", len(message));
	// STRING CONCATENATION!
	fmt.Println("message + concatenated: ", message + " concatenated");
	// STRING INDEXING!
	fmt.Println("message[0]: ", message[0]);
	// STRING UPPERCASE!
	fmt.Println("message to uppercase: ", strings.ToUpper(message))

}

// FUNCTION CHECKING!
func contains(slice []string, elem string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

// FUNCTION THAT WILL RETURN STRING!
func thisWillReturnString() string {
	return "Hello"
}

// FUNCTION THAT WILL RETURN INTEGER!
func thisWillReturnInteger() int {
	return 6
}

// FUNCTION THAT WONT RETURN ANY THING!
func thisWillReturnAnything() string {
	return "Learning go is Awesome!"
}

// FUNCTION RETURN TWO VALUES!
func thisWillReturnTwoValues() (int, int) {
	return 56, 56
}

func cannotDivideByZero(top int, bottom int) (int, error) {
	var err error
	if bottom == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, err
	}

	return top / bottom, err
}
