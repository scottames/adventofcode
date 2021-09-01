package bus

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

var (
	testExample = pie.Strings{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	testExampleSchedule = Schedule{
		earliestDeparture: 939,
		buses: &Buses{
			list: []Bus{
				{id: 7, offset: 0},
				{id: 13, offset: 1},
				{id: 59, offset: 4},
				{id: 31, offset: 6},
				{id: 19, offset: 7},
			},
			min: 7,
			max: 59,
		},
	}
	testActual = pie.Strings{
		"1000053",
		"19,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,23,x,x,x,x,x,29,x,547,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17",
	}
)

// Part 1
func Test_ReadScheduleBuses(t *testing.T) {
	expected := testExampleSchedule.buses
	sched, err := ReadSchedule(testExample)
	if err != nil {
		assert.Error(t, err)
	}
	actual := sched.buses
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ReadScheduleEarliestDeparture(t *testing.T) {
	expected := testExampleSchedule.earliestDeparture
	sched, err := ReadSchedule(testExample)
	if err != nil {
		assert.Error(t, err)
	}
	actual := sched.earliestDeparture
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ScheduleEarliestDeparture(t *testing.T) {
	expected := 939
	actual := testExampleSchedule.EarliestDeparture()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ScheduleNextDeparture(t *testing.T) {
	expected := 944
	actual := testExampleSchedule.NextDeparture()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ScheduleNextDepartureID(t *testing.T) {
	expected := 59
	actual := testExampleSchedule.NextDepartureID()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Part1(t *testing.T) {
	expected := 102
	sched, err := ReadSchedule(testActual)
	if err != nil {
		assert.Error(t, err)
	}
	actual := (sched.NextDeparture() - sched.EarliestDeparture()) * sched.NextDepartureID()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

// Part 2
func Test_SchedEarlistTimestampOffsetMatchExample(t *testing.T) {
	expected := 1068781
	sched, err := ReadSchedule(testExample)
	if err != nil {
		assert.Error(t, err)
	}
	actual := sched.EarliestTimestampOffsetAlignment()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_SchedEarlistTimestampOffsetMatchActual(t *testing.T) {
	expected := 327300950120029
	sched, err := ReadSchedule(testActual)
	if err != nil {
		assert.Error(t, err)
	}
	actual := sched.EarliestTimestampOffsetAlignment()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_EarlistTimestampOffsetMatch(t *testing.T) {
	expected := 1068781
	actual := EarliestTimestampOffsetAlignment(testExample[1])
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_EarlistTimestampOffsetActual(t *testing.T) {
	expected := 327300950120029
	actual := EarliestTimestampOffsetAlignment(testActual[1])
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
