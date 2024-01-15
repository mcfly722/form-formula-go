package formFormula

func recombineRequiredXRecursive(currentPos int, input *[]*int, buffer *[]*int, bufferPos int, ocurrencesLeft int, setValue int, ready func(remained *[]*int)) {
	if ocurrencesLeft == 0 {
		if bufferPos < len(*buffer) {
			for i := currentPos; i < len(*input); i++ {
				(*buffer)[bufferPos+i-currentPos] = (*input)[i]
			}
		}
		ready(buffer)
	} else {
		for i := currentPos; i < len(*input)-(ocurrencesLeft-1); i++ {
			*(*input)[i] = setValue
			recombineRequiredXRecursive(i+1, input, buffer, bufferPos+i-currentPos, ocurrencesLeft-1, setValue, ready)
			if bufferPos+i-currentPos < len(*buffer) {
				(*buffer)[bufferPos+i-currentPos] = (*input)[i]
			}
		}
	}
}

func RecombineRequiredX(input *[]*int, maxOccurrences int, setXValue int, ready func(remained *[]*int)) {
	if len(*input) == 0 {
		ready(input)
		return
	}
	for occurencies := 1; occurencies <= maxOccurrences; occurencies++ {
		buffer := make([]*int, len(*input)-occurencies)
		recombineRequiredXRecursive(0, input, &buffer, 0, occurencies, setXValue, ready)
	}
}

func recombineValues(input *[]*int, possibleValues []int, ready func(originalInput *[]*int), currentPos int) {
	for _, value := range possibleValues {
		*(*input)[currentPos] = value
		if currentPos == 0 {
			ready(input)
		} else {
			recombineValues(input, possibleValues, ready, currentPos-1)
		}
	}
}

func RecombineValues(input *[]*int, possibleValues []int, ready func(originalInput *[]*int)) {
	if len(*input) == 0 || len(possibleValues) == 0 {
		ready(input)
		return
	}
	recombineValues(input, possibleValues, ready, len(*input)-1)
}
