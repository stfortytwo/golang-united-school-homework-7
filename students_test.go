package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var (
	testPersonA1 = Person{
		firstName: "AaaFirstName",
		lastName:  "AaaLastName",
		birthDay:  time.Date(2012, 1, 1, 2, 3, 4, 5, time.UTC),
	}
	testPersonA2 = Person{
		firstName: "AaaFirstName",
		lastName:  "AaaLastName2",
		birthDay:  time.Date(2012, 1, 1, 2, 3, 4, 5, time.UTC),
	}
	testPersonB1 = Person{
		firstName: "BbbFirstName",
		lastName:  "BbbLastName",
		birthDay:  time.Date(2010, 12, 30, 2, 4, 1, 0, time.UTC),
	}
	testPersonB2 = Person{
		firstName: "BbbFirstName2",
		lastName:  "BbbLastName",
		birthDay:  time.Date(2010, 12, 30, 2, 4, 1, 0, time.UTC),
	}
	testPeople = People{testPersonA1, testPersonA2, testPersonB1, testPersonB2}
)

func TestPeople_Len(t *testing.T) {
	var testLenData = map[string]struct {
		people      People
		expectedLen int
	}{
		"empty slice": {people: People{}, expectedLen: 0},
		"two persons": {people: []Person{testPersonA1, testPersonB1}, expectedLen: 2},
	}
	for name, v := range testLenData {
		tCase := v
		t.Run(name, func(t *testing.T) {
			result := tCase.people.Len()
			if result != tCase.expectedLen {
				t.Errorf("[%s] expected: %d, got %d", name, tCase.expectedLen, result)
			}
		})
	}
}

func TestPeople_Less(t *testing.T) {
	var testLessData = map[string]struct {
		i        int
		j        int
		expected bool
	}{
		"diff last names false":  {i: 1, j: 0, expected: false},
		"diff last names true":   {i: 0, j: 1, expected: true},
		"diff first names false": {i: 3, j: 2, expected: false},
		"diff first names true":  {i: 2, j: 3, expected: true},
		"diff dates false":       {i: 2, j: 1, expected: false},
		"diff dates true":        {i: 1, j: 2, expected: true},
	}
	for name, v := range testLessData {
		tCase := v
		t.Run(name, func(t *testing.T) {
			result := testPeople.Less(tCase.i, tCase.j)
			if result != tCase.expected {
				t.Errorf("[%s] expected: %v, got %v", name, tCase.expected, result)
			}
		})
	}
}

func TestPeople_Swap(t *testing.T) {
	testSwapPeople := make(People, 0)
	testSwapPeople = append(testSwapPeople, testPeople...)
	testSwapPeople.Swap(0, 3)
	if testSwapPeople[0] != testPeople[3] || testSwapPeople[3] != testPeople[0] {
		t.Errorf("swapping didn't work")
	}
	if testSwapPeople[1] != testPeople[1] || testSwapPeople[2] != testPeople[2] {
		t.Errorf("swapped wrong elements")
	}
}

func TestNew(t *testing.T) {
	expMatrix := Matrix{rows: 3, cols: 3, data: []int{0, 1, 2, -3, 4, 5, 6, 7, 8}}
	if m, err := New("0 1 2 \n -3 4 5 \n 6 7 8"); err != nil {
		t.Errorf("not expected error occured: %s", err)
	} else if !reflect.DeepEqual(*m, expMatrix) {
		t.Errorf("matrix is wrong")
	}
	if m, _ := New(""); m != nil {
		t.Errorf("got %v while expected nil", m)
	}
	if _, err := New("a"); err == nil {
		t.Errorf("didn't get expected string to int conversion error")
	}
}

func TestMatrix_Cols(t *testing.T) {
	var testData = map[string]struct {
		input    Matrix
		expected [][]int
	}{
		"1 element": {input: Matrix{rows: 1, cols: 1, data: []int{1}}, expected: [][]int{{1}}},
		"1r x 2c":   {input: Matrix{rows: 1, cols: 2, data: []int{1, 2}}, expected: [][]int{{1}, {2}}},
		"2r x 1c":   {input: Matrix{rows: 2, cols: 1, data: []int{1, 2}}, expected: [][]int{{1, 2}}},
	}
	for name, v := range testData {
		tCase := v
		t.Run(name, func(t *testing.T) {
			result := tCase.input.Cols()
			if !reflect.DeepEqual(result, tCase.expected) {
				t.Errorf("wrong output: expected %v, got %v", tCase.expected, result)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	var testData = map[string]struct {
		input    Matrix
		expected [][]int
	}{
		"1 element": {input: Matrix{rows: 1, cols: 1, data: []int{1}}, expected: [][]int{{1}}},
		"1r x 2c":   {input: Matrix{rows: 1, cols: 2, data: []int{1, 2}}, expected: [][]int{{1, 2}}},
		"2r x 1c":   {input: Matrix{rows: 2, cols: 1, data: []int{1, 2}}, expected: [][]int{{1}, {2}}},
	}
	for name, v := range testData {
		tCase := v
		t.Run(name, func(t *testing.T) {
			result := tCase.input.Rows()
			if !reflect.DeepEqual(result, tCase.expected) {
				t.Errorf("wrong output: expected %v, got %v", tCase.expected, result)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	m := Matrix{rows: 3, cols: 4, data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}
	var testData = map[string]struct {
		row int
		col int
		val int
		exp bool
	}{
		"row < 0":   {row: -1, col: 2, val: 111, exp: false},
		"row > len": {row: 3, col: 0, val: 111, exp: false},
		"col < 0":   {row: 1, col: -1, val: 111, exp: false},
		"col > len": {row: 3, col: 4, val: 111, exp: false},
		"good":      {row: 2, col: 2, val: 111, exp: true},
	}
	mTest := m
	mTest.Set(2, 2, 111)
	if mTest.data[10] != 111 {
		t.Errorf("Wrong value after changing, expected %d, got %d", 111, mTest.data[10])
	}
	for name, v := range testData {
		tCase := v
		mTest := m
		t.Run(name, func(t *testing.T) {
			result := mTest.Set(tCase.row, tCase.col, tCase.val)
			if result != tCase.exp {
				t.Errorf("got %t while expected %t", result, tCase.exp)
			}
		})
	}
}
