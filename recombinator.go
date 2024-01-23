package formFormula

func recombineRequiredXRecursive(currentPos uint, input *[]*uint, buffer *[]*uint, bufferPos uint, occurrencesLeft uint, setValue uint, ready func(remained *[]*uint)) {
	if occurrencesLeft == 0 {
		if bufferPos < uint(len(*buffer)) {
			for i := currentPos; i < uint(len(*input)); i++ {
				(*buffer)[bufferPos+i-currentPos] = (*input)[i]
			}
		}
		ready(buffer)
	} else {
		for i := currentPos; i < uint(len(*input))-occurrencesLeft+1; i++ {
			*(*input)[i] = setValue
			recombineRequiredXRecursive(i+1, input, buffer, bufferPos+i-currentPos, occurrencesLeft-1, setValue, ready)
			if bufferPos+i-currentPos < uint(len(*buffer)) {
				(*buffer)[bufferPos+i-currentPos] = (*input)[i]
			}
		}
	}
}

func RecombineRequiredX(input *[]*uint, maxOccurrences uint, setXValue uint, ready func(remained *[]*uint)) {
	if len(*input) == 0 {
		ready(input)
		return
	}
	for occurrences := uint(1); occurrences <= maxOccurrences; occurrences++ {
		buffer := make([]*uint, len(*input)-int(occurrences))
		recombineRequiredXRecursive(0, input, &buffer, 0, occurrences, setXValue, ready)
	}
}

func recombineValues(input *[]*uint, possibleValues *[]uint, ready func(), currentPos uint) {
	for _, value := range *possibleValues {
		*(*input)[currentPos] = value
		if currentPos == 0 {
			ready()
		} else {
			recombineValues(input, possibleValues, ready, currentPos-1)
		}
	}
}

func RecombineValues(input *[]*uint, possibleValues *[]uint, ready func()) {
	if len(*input) == 0 || len(*possibleValues) == 0 {
		ready()
		return
	}
	recombineValues(input, possibleValues, ready, uint(len(*input)-1))
}
