package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
    // ...
	
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
}
for i, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	if len(actual) != len(c.expected){
         t.Errorf("Test Case %d Failed" , i)
	} else{
		fmt.Printf("Test Case %d Passed" , i)
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if word != expectedWord{
			 t.Errorf("Test Case Failing word did not match actual : %s expected :%s" , word,expectedWord)
		}	else{
fmt.Printf("Test Case Passing word  match actual : %s expected :%s" , word,expectedWord)
		}

	}
}


}




