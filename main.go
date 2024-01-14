


package main;

import "fmt";


func main(){
	// DATA TYPES IN GO LANG!
	variableDaNa := 5;
	var floatVariableDaNa float64 = 56.9
	var stringVariableDaNa string = "Babu Ji";
	var booleanVariableDaNa bool = true;

	// ADDITION!
	fmt.Println("Addition: ", variableDaNa + int(floatVariableDaNa))
	// SUBTRACTION!
	fmt.Println("Subtraction: ", variableDaNa - int(floatVariableDaNa))
	// MULTIPLICATION!
	fmt.Println("Multiplication: ", variableDaNa * int(floatVariableDaNa))
	// DIVISION!
	fmt.Println("Division: ", floatVariableDaNa / float64(variableDaNa))
	fmt.Println("variables Di Value = ", variableDaNa, floatVariableDaNa);
	fmt.Println("variables Di String = ", stringVariableDaNa);
	fmt.Println("variables Di String = ", stringVariableDaNa + " is a good boy");
	fmt.Println("Length Of Variable String: ", len(stringVariableDaNa))
	fmt.Println("Boolean Variable: ", booleanVariableDaNa);

	var somethingGotOne, somethingGotTwo int = thisWillReturnTwoValues();
	thisWillReturnInteger();
	thisWillReturnString()

	fmt.Printf("Something Got: %d and %d", somethingGotOne, somethingGotTwo)

	// fmt.Println("Something Got: ", somethingGot);

}


// FUNCTION THAT WILL RETURN STRING!
func thisWillReturnString() string{
	return "Hello";
}

// FUNCTION THAT WILL RETURN INTEGER!
func thisWillReturnInteger() int{
	return 6;
}

// FUNCTION THAT WONT RETURN ANY THING!
func thisWillReturnAnything() string{
	return "Learning go is Awesome!";
}

// FUNCTION RETURN TWO VALUES!
func thisWillReturnTwoValues() (int,int){
	return 56,56;
}