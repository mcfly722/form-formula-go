package formFormula

import "fmt"

type Expression struct {
	Arguments []*Expression
}

// BracketsToExpressionTree generates expression tree based on string of brackets
func BracketsToExpressionTree(input string) (*Expression, error) {
	root := Expression{Arguments: []*Expression{}}

	if input == "" {
		return &root, nil
	}

	counter := 0
	from := 0

	for i := 0; i < len(input); i++ {

		if input[i] == '(' {
			if counter == 0 {
				from = i
			}
			counter++
		}

		if input[i] == ')' {
			counter--

			if counter < 0 {
				return nil, fmt.Errorf("%v<- incorrect brackets balance, could not close not opened bracket", input[0:i+1])
			}

			if counter == 0 {
				argument, _ := BracketsToExpressionTree(input[from+1 : i])
				if argument != nil {
					root.Arguments = append(root.Arguments, argument)
				}
			}

		}

		if input[i] != '(' && input[i] != ')' {
			return nil, fmt.Errorf("%v<- unknown symbol", input[0:i+1])
		}
	}

	if counter != 0 {
		return nil, fmt.Errorf("number of opened brackets are not equal to closed (difference=%v)", counter)
	}

	return &root, nil
}
