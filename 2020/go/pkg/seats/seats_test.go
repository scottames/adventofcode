package seats

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

var (
	start = pie.Strings{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	round0 = []pie.Ints{
		{1, 0, 1, 1, 0, 1, 1, 0, 1, 1}, // L.LL.LL.LL
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 1}, // LLLLLLL.LL
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 0}, // L.L.L..L..
		{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}, // LLLL.LL.LL
		{1, 0, 1, 1, 0, 1, 1, 0, 1, 1}, // L.LL.LL.LL
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 1}, // L.LLLLL.LL
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // LLLLLLLLLL
		{1, 0, 1, 1, 1, 1, 1, 1, 0, 1}, // L.LLLLLL.L
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 1}, // L.LLLLL.LL
	}
	part1Round1 = []pie.Ints{
		{2, 0, 2, 2, 0, 2, 2, 0, 2, 2}, // #.##.##.##
		{2, 2, 2, 2, 2, 2, 2, 0, 2, 2}, // #######.##
		{2, 0, 2, 0, 2, 0, 0, 2, 0, 0}, // #.#.#..#..
		{2, 2, 2, 2, 0, 2, 2, 0, 2, 2}, // ####.##.##
		{2, 0, 2, 2, 0, 2, 2, 0, 2, 2}, // #.##.##.##
		{2, 0, 2, 2, 2, 2, 2, 0, 2, 2}, // #.#####.##
		{0, 0, 2, 0, 2, 0, 0, 0, 0, 0}, // ..#.#.....
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, // ##########
		{2, 0, 2, 2, 2, 2, 2, 2, 0, 2}, // #.######.#
		{2, 0, 2, 2, 2, 2, 2, 0, 2, 2}, // #.#####.##
	}
	part1Round2 = []pie.Ints{
		{2, 0, 1, 1, 0, 1, 2, 0, 2, 2}, // #.LL.L#.##
		{2, 1, 1, 1, 1, 1, 1, 0, 1, 2}, // #LLLLLL.L#
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 0}, // L.L.L..L..
		{2, 1, 1, 1, 0, 1, 1, 0, 1, 2}, // #LLL.LL.L#
		{2, 0, 1, 1, 0, 1, 1, 0, 1, 1}, // #.LL.LL.LL
		{2, 0, 1, 1, 1, 1, 2, 0, 2, 2}, // #.LLLL#.##
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 2}, // #LLLLLLLL#
		{2, 0, 1, 1, 1, 1, 1, 1, 0, 1}, // #.LLLLLL.L
		{2, 0, 2, 1, 1, 1, 1, 0, 2, 2}, // #.#LLLL.##
	}
	part1Round3 = []pie.Ints{
		{2, 0, 2, 2, 0, 1, 2, 0, 2, 2}, // #.##.L#.##
		{2, 1, 2, 2, 2, 1, 1, 0, 1, 2}, // #L###LL.L#
		{1, 0, 2, 0, 2, 0, 0, 2, 0, 0}, // L.#.#..#..
		{2, 1, 2, 2, 0, 2, 2, 0, 1, 2}, // #L##.##.L#
		{2, 0, 2, 2, 0, 1, 1, 0, 1, 1}, // #.##.LL.LL
		{2, 0, 2, 2, 2, 1, 2, 0, 2, 2}, // #.###L#.##
		{0, 0, 2, 0, 2, 0, 0, 0, 0, 0}, // ..#.#.....
		{2, 1, 2, 2, 2, 2, 2, 2, 1, 2}, // #L######L#
		{2, 0, 1, 1, 2, 2, 2, 1, 0, 1}, // #.LL###L.L
		{2, 0, 2, 1, 2, 2, 2, 0, 2, 2}, // #.#L###.##
	}
	part1Round4 = []pie.Ints{
		{2, 0, 2, 1, 0, 1, 2, 0, 2, 2}, // #.#L.L#.##
		{2, 1, 1, 1, 2, 1, 1, 0, 1, 2}, // #LLL#LL.L#
		{1, 0, 1, 0, 1, 0, 0, 2, 0, 0}, // L.L.L..#..
		{2, 1, 1, 1, 0, 2, 2, 0, 1, 2}, // #LLL.##.L#
		{2, 0, 1, 1, 0, 1, 1, 0, 1, 1}, // #.LL.LL.LL
		{2, 0, 1, 1, 2, 1, 2, 0, 2, 2}, // #.LL#L#.##
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{2, 1, 2, 1, 1, 1, 1, 2, 1, 2}, // #L#LLLL#L#
		{2, 0, 1, 1, 1, 1, 1, 1, 0, 1}, // #.LLLLLL.L
		{2, 0, 2, 1, 2, 1, 2, 0, 2, 2}, // #.#L#L#.##
	}
	part1Round5 = []pie.Ints{
		{2, 0, 2, 1, 0, 1, 2, 0, 2, 2}, // #.#L.L#.##
		{2, 1, 1, 1, 2, 1, 1, 0, 1, 2}, // #LLL#LL.L#
		{1, 0, 2, 0, 1, 0, 0, 2, 0, 0}, // L.#.L..#..
		{2, 1, 2, 2, 0, 2, 2, 0, 1, 2}, // #L##.##.L#
		{2, 0, 2, 1, 0, 1, 1, 0, 1, 1}, // #.#L.LL.LL
		{2, 0, 2, 1, 2, 1, 2, 0, 2, 2}, // #.#L#L#.##
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{2, 1, 2, 1, 2, 2, 1, 2, 1, 2}, // #L#L##L#L#
		{2, 0, 1, 1, 1, 1, 1, 1, 0, 1}, // #.LLLLLL.L
		{2, 0, 2, 1, 2, 1, 2, 0, 2, 2}, // #.#L#L#.##
	}
	part2Round1 = []pie.Ints{
		{2, 0, 2, 2, 0, 2, 2, 0, 2, 2}, // #.##.##.##
		{2, 2, 2, 2, 2, 2, 2, 0, 2, 2}, // #######.##
		{2, 0, 2, 0, 2, 0, 0, 2, 0, 0}, // #.#.#..#..
		{2, 2, 2, 2, 0, 2, 2, 0, 2, 2}, // ####.##.##
		{2, 0, 2, 2, 0, 2, 2, 0, 2, 2}, // #.##.##.##
		{2, 0, 2, 2, 2, 2, 2, 0, 2, 2}, // #.#####.##
		{0, 0, 2, 0, 2, 0, 0, 0, 0, 0}, // ..#.#.....
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, // ##########
		{2, 0, 2, 2, 2, 2, 2, 2, 0, 2}, // #.######.#
		{2, 0, 2, 2, 2, 2, 2, 0, 2, 2}, // #.#####.##
	}
	part2Round2 = []pie.Ints{
		{2, 0, 1, 1, 0, 1, 1, 0, 1, 2}, // #.LL.LL.L#
		{2, 1, 1, 1, 1, 1, 1, 0, 1, 1}, // #LLLLLL.LL
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 0}, // L.L.L..L..
		{1, 1, 1, 1, 0, 1, 1, 0, 1, 1}, // LLLL.LL.LL
		{1, 0, 1, 1, 0, 1, 1, 0, 1, 1}, // L.LL.LL.LL
		{1, 0, 1, 1, 1, 1, 1, 0, 1, 1}, // L.LLLLL.LL
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, // LLLLLLLLL#
		{2, 0, 1, 1, 1, 1, 1, 1, 0, 1}, // #.LLLLLL.L
		{2, 0, 1, 1, 1, 1, 1, 0, 1, 2}, // #.LLLLL.L#
	}
	part2Round3 = []pie.Ints{
		{2, 0, 1, 2, 0, 2, 2, 0, 1, 2}, // #.L#.##.L#
		{2, 1, 2, 2, 2, 2, 2, 0, 1, 1}, // #L#####.LL
		{1, 0, 2, 0, 2, 0, 0, 2, 0, 0}, // L.#.#..#..
		{2, 2, 1, 2, 0, 2, 2, 0, 2, 2}, // ##L#.##.##
		{2, 0, 2, 2, 0, 2, 1, 0, 2, 2}, // #.##.#L.##
		{2, 0, 2, 2, 2, 2, 2, 0, 2, 1}, // #.#####.#L
		{0, 0, 2, 0, 2, 0, 0, 0, 0, 0}, // ..#.#.....
		{1, 1, 1, 2, 2, 2, 2, 1, 1, 2}, // LLL####LL#
		{2, 0, 1, 2, 2, 2, 2, 2, 0, 1}, // #.L#####.L
		{2, 0, 1, 2, 2, 2, 2, 0, 1, 2}, // #.L####.L#
	}
	part2Round4 = []pie.Ints{
		{2, 0, 1, 2, 0, 1, 2, 0, 1, 2}, // #.L#.L#.L#
		{2, 1, 1, 1, 1, 1, 1, 0, 1, 1}, // #LLLLLL.LL
		{1, 0, 1, 0, 1, 0, 0, 2, 0, 0}, // L.L.L..#..
		{2, 2, 1, 1, 0, 1, 1, 0, 1, 2}, // ##LL.LL.L#
		{1, 0, 1, 1, 0, 1, 1, 0, 1, 2}, // L.LL.LL.L#
		{2, 0, 1, 1, 1, 1, 1, 0, 1, 1}, // #.LLLLL.LL
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // ..L.L.....
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, // LLLLLLLLL#
		{2, 0, 1, 1, 1, 1, 1, 2, 0, 1}, // #.LLLLL#.L
		{2, 0, 1, 2, 1, 1, 2, 0, 1, 2}, // #.L#LL#.L#
	}
	part2Round5 = []pie.Ints{
		{2, 0, 1, 2, 0, 1, 2, 0, 1, 2}, // #.L#.L#.L#
		{2, 1, 1, 1, 1, 1, 1, 0, 1, 1}, // #LLLLLL.LL
		{1, 0, 1, 0, 1, 0, 0, 2, 0, 0}, // L.L.L..#..
		{2, 2, 1, 2, 0, 2, 1, 0, 1, 2}, // ##L#.#L.L#
		{1, 0, 1, 2, 0, 2, 1, 0, 1, 2}, // L.L#.#L.L#
		{2, 0, 1, 2, 2, 2, 2, 0, 1, 1}, // #.L####.LL
		{0, 0, 2, 0, 2, 0, 0, 0, 0, 0}, // ..#.#.....
		{1, 1, 1, 2, 2, 2, 1, 1, 1, 2}, // LLL###LLL#
		{2, 0, 1, 1, 1, 1, 1, 2, 0, 1}, // #.LLLLL#.L
		{2, 0, 1, 2, 1, 1, 2, 0, 1, 2}, // #.L#LL#.L#
	}
	part2Round6 = []pie.Ints{
		{2, 0, 1, 2, 0, 1, 2, 0, 1, 2}, // #.L#.L#.L#
		{2, 1, 1, 1, 1, 1, 1, 0, 1, 1}, // #LLLLLL.LL
		{1, 0, 1, 0, 1, 0, 0, 2, 0, 0}, // L.L.L..#..
		{2, 2, 1, 2, 0, 2, 1, 0, 1, 2}, // ##L#.#L.L#
		{1, 0, 1, 2, 0, 1, 1, 0, 1, 2}, // L.L#.LL.L#
		{2, 0, 1, 1, 1, 1, 2, 0, 1, 1}, // #.LLLL#.LL
		{0, 0, 2, 0, 1, 0, 0, 0, 0, 0}, // ..#.L.....
		{1, 1, 1, 2, 2, 2, 1, 1, 1, 2}, // LLL###LLL#
		{2, 0, 1, 1, 1, 1, 1, 2, 0, 1}, // #.LLLLL#.L
		{2, 0, 1, 2, 1, 1, 2, 0, 1, 2}, // #.L#LL#.L#
	}

	arrangement1      = Arrangement{set: round0}
	arrangementPart12 = Arrangement{set: part1Round1}
	arrangementPart13 = Arrangement{set: part1Round2}
	arrangementPart14 = Arrangement{set: part1Round3}
	arrangementPart15 = Arrangement{set: part1Round4}

	arrangementPart22 = Arrangement{set: part2Round1}
	arrangementPart23 = Arrangement{set: part2Round2}
	arrangementPart24 = Arrangement{set: part2Round3}
	arrangementPart25 = Arrangement{set: part2Round4}
	arrangementPart26 = Arrangement{set: part2Round5}
)

func Test_NewArrangementFromStart(t *testing.T) {
	expected := &arrangement1
	actual := NewArrangement(start)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart1Round1(t *testing.T) {
	expected := part1Round1
	actual := arrangement1.Next(seatingLogicAdjacent).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart1Round2(t *testing.T) {
	expected := part1Round2
	actual := arrangementPart12.Next(seatingLogicAdjacent).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart1Round3(t *testing.T) {
	expected := part1Round3
	actual := arrangementPart13.Next(seatingLogicAdjacent).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart1Round4(t *testing.T) {
	expected := part1Round4
	actual := arrangementPart14.Next(seatingLogicAdjacent).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart1Round5(t *testing.T) {
	expected := part1Round5
	actual := arrangementPart15.Next(seatingLogicAdjacent).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextUntilMatchingAdjacent(t *testing.T) {
	expected := part1Round5
	actual := NewArrangement(start).NextUntilMatchingAdjacent().set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextUntilMatchingAdjacentOccupied(t *testing.T) {
	expected := 37
	actual := NewArrangement(start).NextUntilMatchingAdjacent().OccupiedSeats()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round1(t *testing.T) {
	expected := part2Round1
	actual := arrangement1.Next(seatingLogicEightDirections).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round2(t *testing.T) {
	expected := part2Round2
	actual := arrangementPart22.Next(seatingLogicEightDirections).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round3(t *testing.T) {
	expected := part2Round3
	actual := arrangementPart23.Next(seatingLogicEightDirections).set
	fmt.Println("----")
	for i := range part2Round3 {
		fmt.Println(part2Round3[i])
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round4(t *testing.T) {
	expected := part2Round4
	actual := arrangementPart24.Next(seatingLogicEightDirections).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round5(t *testing.T) {
	expected := part2Round5
	actual := arrangementPart25.Next(seatingLogicEightDirections).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextPart2Round6(t *testing.T) {
	expected := part2Round6
	actual := arrangementPart26.Next(seatingLogicEightDirections).set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextUntilMatchingEightDirections(t *testing.T) {
	expected := part2Round6
	actual := NewArrangement(start).NextUntilMatchingEightDirections().set
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NextUntilMatchingEightDirectionsOccupied(t *testing.T) {
	expected := 26
	actual := NewArrangement(start).NextUntilMatchingEightDirections().OccupiedSeats()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
