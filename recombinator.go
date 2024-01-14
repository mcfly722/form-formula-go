package formFormula

type Recombiner interface {
	RecombineRequiredX(input []*int, maxOccurrences int, setXValue int, next func(remained *[]*int))
	Recombine(input *[]*int, possibleValues []int, ready func(originalInput *[]*int))
}

type recombiner struct{}

func NewRecombiner() Recombiner {
	return &recombiner{}
}

func recombineRequiredXRecursive(currentPos int, input []*int, buffer *[]*int, bufferPos int, ocurrencesLeft int, setValue int, next func(remained *[]*int)) {
	if ocurrencesLeft == 0 {

		if bufferPos < len(*buffer) {
			for i := currentPos; i < len(input); i++ {
				(*buffer)[bufferPos+i-currentPos] = input[i]
			}
		}
		next(buffer)
	} else {

		for i := currentPos; i < len(input)-(ocurrencesLeft-1); i++ {
			*(input[i]) = setValue
			recombineRequiredXRecursive(i+1, input, buffer, bufferPos+i-currentPos, ocurrencesLeft-1, setValue, next)
			if bufferPos+i-currentPos < len(*buffer) {
				(*buffer)[bufferPos+i-currentPos] = input[i]
			}
		}
	}
}

func (recombiner *recombiner) RecombineRequiredX(input []*int, maxOccurrences int, setXValue int, next func(remained *[]*int)) {
	for occurencies := 1; occurencies <= maxOccurrences; occurencies++ {
		buffer := make([]*int, len(input)-occurencies)
		recombineRequiredXRecursive(0, input, &buffer, 0, occurencies, setXValue, next)
	}
}

func (recombiner *recombiner) recombine(input *[]*int, possibleValues []int, ready func(originalInput *[]*int), currentPos int) {
	for _, value := range possibleValues {
		*(*input)[currentPos] = value
		if currentPos == 0 {
			ready(input)
		} else {
			recombiner.recombine(input, possibleValues, ready, currentPos-1)
		}
	}

}

func (recombiner *recombiner) Recombine(input *[]*int, possibleValues []int, ready func(originalInput *[]*int)) {
	if len(*input) == 0 || len(possibleValues) == 0 {
		ready(input)
		return
	}
	recombiner.recombine(input, possibleValues, ready, len(*input)-1)
}
