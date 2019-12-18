package rtm

import (
	"testing"

	"gotest.tools/assert"
)

const succeed = "\u2713"
const failed = "\u2714"

func TestCaptureNilPtr(t *testing.T) {
	t.Log("Given the need to test Start when nil Capture pointer is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v using nil Capture poiter.", "capture.Start")
		{
			var c *Capture
			c.Start("TestCaptureArea")
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}

func TestCaptureBeginCapture(t *testing.T) {
	var err error
	var c *Capture
	var target = "BeginCapture"
	var area string

	t.Logf("Given the need to test %v.", target)
	{
		area = generateRandomString(8)
		t.Logf("\tTest 0:\tWhen calling %v for area <%v>.", target, area)
		{
			c, err = BeginCapture(area)
		}
	}

	assert.NilError(t, err)
	t.Logf("\t%s\tShould not return error.", succeed)

	assert.Assert(t, c.elapsed != nil) // NotNil
	t.Logf("\t%s\tShould have elapsed monitor.", succeed)

	assert.Assert(t, c.elapsed.running) // NotNil
	t.Logf("\t%s\tShould have elapsed monitor running.", succeed)

	assert.Equal(t, totalArea, c.elapsed.Area) // NotNil
	t.Logf("\t%s\tShould have elapsed monitor area set to %v.", succeed, totalArea)
}

func TestCaptureStart(t *testing.T) {
	var c *Capture
	var target = "Start"
	var area string
	var sector string
	var m *Monitor

	t.Logf("Given the need to test %v of the Capture.", target)
	{
		area = generateRandomString(8)
		sector = generateRandomString(32)
		t.Logf("\tTest 0:\tWhen calling %v of the Capture area <%v>, monitor sector <%v>.", target, area, sector)
		{
			c, _ = BeginCapture(area)
			m = c.Start(sector)
		}
	}

	assert.Assert(t, m != nil) // NotNil
	t.Logf("\t%s\tSector monitor should not be nil.", succeed)

	assert.Assert(t, m.running) // NotNil
	t.Logf("\t%s\tSector monitor should be running.", succeed)

	assert.Assert(t, c.aggregators != nil) // NotNil
	t.Logf("\t%s\tCapture aggregators should not be nil.", succeed)

	assert.Assert(t, c.aggregators[sector] != nil) // NotNil
	t.Logf("\t%s\tCapture aggregators should contain Monitor sector.", succeed)

	// assert.Assert(t, c.aggregators[sector] != nil) // NotNil
	// t.Logf("\t%s\tCapture aggregators should contain Monitor sector.", succeed)

}
