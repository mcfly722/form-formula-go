package formFormula

import "fmt"

const max_supported_diagonals = 32

type Expression struct {
	Arguments []*Expression
}

type bracketStep struct {
	Opens  int
	Closes int
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

func brackets2points(brackets string) ([]bracketStep, int, error) {
	totalOpens := 0
	totalCloses := 0

	points := []bracketStep{}

	for i := 0; i < len(brackets); {

		opens := 0
		for ; i+opens < len(brackets) && brackets[i+opens] == '('; opens++ {
		}
		i = i + opens
		totalOpens = totalOpens + opens

		if i < len(brackets) && brackets[i] != ')' {
			return nil, 0, fmt.Errorf("unexpected symbol %v <- expecting '(' or ')'", brackets[:i+1])
		}

		closes := 0
		for ; i+closes < len(brackets) && brackets[i+closes] == ')'; closes++ {
		}
		i = i + closes
		totalCloses = totalCloses + closes

		if i < len(brackets) && brackets[i] != '(' {
			return nil, 0, fmt.Errorf("unexpected symbol %v <- expecting '(' or ')'", brackets[:i+1])
		}

		points = append(points, bracketStep{Opens: opens, Closes: closes})

		if totalOpens < totalCloses {
			return nil, 0, fmt.Errorf("%v <- total closes=%v are greater than opens=%v", brackets[:i], totalCloses, totalOpens)
		}
	}

	if totalOpens != totalCloses {
		return nil, 0, fmt.Errorf("opened brackets=%v closed brackets=%v should be equal", totalOpens, totalCloses)
	}

	return points, (totalOpens + totalCloses) / 2, nil
}

func recursionNext(srcBracketsStack []bracketStep, dstBracketsStack []bracketStep, _x int, _y int, maxBracketPairs int, maxChilds uint, diagonal [max_supported_diagonals]uint, currentRecursionStep int, previousSolutionAlreadyReached bool) ([]bracketStep, bool, bool) {
	if _x == maxBracketPairs && _y == maxBracketPairs {
		if !previousSolutionAlreadyReached && len(srcBracketsStack) > 0 {
			return []bracketStep{}, true, false
		}
		return dstBracketsStack, true, true
	}

	if diagonal[_x-_y] < maxChilds {
		for x := _x + 1; x <= maxBracketPairs; x++ {
			diagonal[x-_y-1] = diagonal[x-_y-1] + 1

			if previousSolutionAlreadyReached || len(srcBracketsStack) == 0 || (x-_x) >= srcBracketsStack[currentRecursionStep].Opens {

				for y := _y + 1; y <= maxBracketPairs; y++ {

					if previousSolutionAlreadyReached || len(srcBracketsStack) == 0 || (y-_y) >= srcBracketsStack[currentRecursionStep].Closes {

						if y <= x {
							if diagonal[x-y] < maxChilds+1 {

								newBracketsStack := append(dstBracketsStack, bracketStep{Opens: x - _x, Closes: y - _y})
								tail, reached, solutionFound := recursionNext(srcBracketsStack, newBracketsStack, x, y, maxBracketPairs, maxChilds, diagonal, currentRecursionStep+1, previousSolutionAlreadyReached)

								previousSolutionAlreadyReached = reached

								if solutionFound {
									return tail, previousSolutionAlreadyReached, solutionFound
								}

							}
						}
					}
				}

			}

		}
	}

	return []bracketStep{}, previousSolutionAlreadyReached, false
}

// bracketsStepsToString serialize bracketSteps to String
func bracketsStepsToString(tail []bracketStep) string {
	output := ""
	for _, step := range tail {
		for i := 0; i < step.Opens; i++ {
			output += "("
		}
		for i := 0; i < step.Closes; i++ {
			output += ")"
		}
	}
	return output
}

// GetNextBracketsSequence get current brackets representation of tree and return next one tree in brackets representation
func GetNextBracketsSequence(brackets string, maxChilds uint) (string, error) {

	bracketsStack, maxBracketPairs, err := brackets2points(brackets)
	if err != nil {
		return "", err
	}

	diagonal := [32]uint{}

	if len(bracketsStack) == 1 {
		bracketsStack = []bracketStep{}
		maxBracketPairs++
	}

	nextBracketCombination, _, _ := recursionNext(bracketsStack, []bracketStep{}, 0, 0, maxBracketPairs, maxChilds, diagonal, 0, false)

	return bracketsStepsToString(nextBracketCombination), nil
}
