package formFormula_test

import (
	"encoding/json"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_BracketsToExpressionTree(t *testing.T) {
	bracketSequence := "()((())())"
	expression, err := formFormula.BracketsToExpressionTree(bracketSequence)
	if err != nil {
		t.Errorf("Cant build expression tree for %v. Error: %v", bracketSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", bracketSequence, err)
	}
	t.Log(string(bytes))
}

func testBracketsToExpressionTreeSuccess(t *testing.T, testSequence string) {
	expression, err := formFormula.BracketsToExpressionTree(testSequence)

	if err != nil {
		t.Errorf("error for %v :%v", testSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", testSequence, err)
	}
	t.Log(string(bytes))
}

func testBracketsToExpressionTreeError(t *testing.T, testSequence string) {
	_, err := formFormula.BracketsToExpressionTree(testSequence)

	if err == nil {
		t.Errorf("error for '%v' sequence does not catched", testSequence)
	} else {
		t.Logf("error for '%v' successfully catched.\nerror description:%v", testSequence, err)
	}
}

func testGetNextBracketsSequenceError(t *testing.T, sequence string, maxChilds uint) {
	_, err := formFormula.GetNextBracketsSequence(sequence, maxChilds)
	if err == nil {
		t.Fatal("not catched error")
	}
	t.Logf("successfully catched error: %v", err)
}

func Test_BracketsToExpressionTree_FirstBracket(t *testing.T) {
	testBracketsToExpressionTreeSuccess(t, "()")
}

func Test_BracketsToExpressionTree_IncorrectSymbolFromStart(t *testing.T) {
	testBracketsToExpressionTreeError(t, "a(()(()))")
}

func Test_BracketsToExpressionTree_IncorrectSymbolAtTheEnd(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()(()))b")
}

func Test_BracketsToExpressionTree_IncorrectSymbolInTheMiddle(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()((c)))")
}

func Test_BracketsToExpressionTree_LostOpeningBracket(t *testing.T) {
	testBracketsToExpressionTreeError(t, "())")
}

func Test_BracketsToExpressionTree_LostClosingBracket(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()")
}

func Test_BracketsToExpressionTree_WrongBracketsSequence(t *testing.T) {
	testBracketsToExpressionTreeError(t, "())(")
}

func Test_GetNextBracketsSequence_IncorrectSymbolFromStart(t *testing.T) {
	testGetNextBracketsSequenceError(t, "a(()(()))", 2)
}

func Test_GetNextBracketsSequence_IncorrectSymbolAtTheEnd(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()(()))b", 2)
}

func Test_GetNextBracketsSequence_IncorrectSymbolInTheMiddle(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()((c)))", 2)
}

func Test_GetNextBracketsSequence_LostOpeningBracket(t *testing.T) {
	testGetNextBracketsSequenceError(t, "())", 2)
}

func Test_GetNextBracketsSequence_LostClosingBracket(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()", 2)
}

func Test_GetNextBracketsSequence_WrongBracketsSequence(t *testing.T) {
	testGetNextBracketsSequenceError(t, "())(", 2)
}

func Test_GetNextBracketsSequence_ForSeveralIterations(t *testing.T) {
	current := "()"

	for i := 0; i < 100; i++ {
		c, err := formFormula.GetNextBracketsSequence(current, 3)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%3v %v -> %v", i, current, c)
		current = c
	}

	assert_string(t, current, "(()((()())))")
}
