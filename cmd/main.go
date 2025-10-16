package main

import (
	"fmt"
	"strings"
	"dz2/internal/task_1"
)

func main() {
	// Check task1

	// Example 1: normal input 
	input := `q--3wer4301-4+r_e**
dog d1
cat c1
cat c2
bird b1
end`

	stream := strings.NewReader(input)
	results := task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}


	// Example 2: input with non unified register
	input = `1234567qwerty-=_+,./\|
dog d1
DoG d2
dOG d3
cat c1
CAT c2
bird b1
end`

	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")

	for _, result := range results {
		fmt.Println(result)
	}

	// Example 3: input with empty rows 
	input = `q--3wer4301-4+r_e**

dog d1

cat c1
cat c2
	
bird b1

end`

	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")
	for _, result := range results {
		fmt.Println(result)
	}


	// Example 4: input with empty lines and non unified register
	input = `1234567qwerty-=_+,./\|

dog d1
DoG d2
dOG d3

cat c1
	
CAT c2
bird b1
end`

	stream = strings.NewReader(input)
	results = task1.AnimalFeeding(stream)

	fmt.Println("Результаты кормления:")

	for _, result := range results {
		fmt.Println(result)
	}
}