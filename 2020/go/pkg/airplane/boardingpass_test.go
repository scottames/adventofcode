package airplane

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testBoardingPass0       = "FBFBBFFRLR"
	testBoardingPass0Cols   = "RLR"
	testBoardingPass0Rows   = "FBFBBFF"
	testBoardingPass0column = 5
	testBoardingPass0row    = 44
	testBoardingPass0seatID = 357
	testBoardingPass1       = "BFFFBBFRRR"
	testBoardingPass1Cols   = "RRR"
	testBoardingPass1Rows   = "BFFFBBF"
	testBoardingPass1row    = 70
	testBoardingPass1column = 7
	testBoardingPass1seatID = 567
	testBoardingPass2       = "FFFBBBFRRR"
	testBoardingPass2Cols   = "RRR"
	testBoardingPass2Rows   = "FFFBBBF"
	testBoardingPass2row    = 14
	testBoardingPass2column = 7
	testBoardingPass2seatID = 119
	testBoardingPass3       = "BBFFBBFRLL"
	testBoardingPass3Cols   = "RLL"
	testBoardingPass3Rows   = "BBFFBBF"
	testBoardingPass3row    = 102
	testBoardingPass3column = 4
	testBoardingPass3seatID = 820
)

// TODO: Simplify to one test function

func TestBoardingPass_Columns0(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass0}
	expected := testBoardingPass0Cols
	actual := bp.Columns()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Rows0(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass0}
	expected := testBoardingPass0Rows
	actual := bp.Rows()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Columns1(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass1}
	expected := testBoardingPass1Cols
	actual := bp.Columns()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Columns2(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass2}
	expected := testBoardingPass2Cols
	actual := bp.Columns()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Columns3(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass3}
	expected := testBoardingPass3Cols
	actual := bp.Columns()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Rows1(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass1}
	expected := testBoardingPass1Rows
	actual := bp.Rows()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Rows2(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass2}
	expected := testBoardingPass2Rows
	actual := bp.Rows()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_Rows3(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass3}
	expected := testBoardingPass3Rows
	actual := bp.Rows()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s. Got %s.", expected, actual))
}

func TestBoardingPass_CalculateColumn0(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass0}
	expected := testBoardingPass0column
	actual := bp.calculateColumn()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateColumn1(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass1}
	expected := testBoardingPass1column
	actual := bp.calculateColumn()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateColumn2(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass2}
	expected := testBoardingPass2column
	actual := bp.calculateColumn()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateColumn3(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass3}
	expected := testBoardingPass3column
	actual := bp.calculateColumn()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateRow0(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass0}
	expected := testBoardingPass0row
	actual := bp.calculateRow()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateRow1(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass1}
	expected := testBoardingPass1row
	actual := bp.calculateRow()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateRow2(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass2}
	expected := testBoardingPass2row
	actual := bp.calculateRow()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateRow3(t *testing.T) {
	bp := BoardingPass{Pass: testBoardingPass3}
	expected := testBoardingPass3row
	actual := bp.calculateRow()
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateSeatID0(t *testing.T) {
	bp := NewBoardingPass(testBoardingPass0)
	expected := testBoardingPass0seatID
	actual := bp.Seat
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateSeatID1(t *testing.T) {
	bp := NewBoardingPass(testBoardingPass1)
	expected := testBoardingPass1seatID
	actual := bp.Seat
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateSeatID2(t *testing.T) {
	bp := NewBoardingPass(testBoardingPass2)
	expected := testBoardingPass2seatID
	actual := bp.Seat
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}

func TestBoardingPass_CalculateSeatID3(t *testing.T) {
	bp := NewBoardingPass(testBoardingPass3)
	expected := testBoardingPass3seatID
	actual := bp.Seat
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d. Got %d.", expected, actual))
}
