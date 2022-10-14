package watch

import (
	"math/rand"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestMonitorNewMonitor(t *testing.T) {
	var m *Monitor
	var area string
	target := "NewMonitor"
	t.Log("Given the need to test NewMonitor.")
	{
		t.Logf("\tTest 0:\tWhen calling %v.", target)
		{
			area = generateRandomString(8)
			m, _ = NewMonitor(area, false)
		}
	}
	assert.Assert(t, m != nil)
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)

	assert.Equal(t, area, m.Area)
	t.Logf("\t%s\tShould have correct Area set in %v.", succeed, target)

	assert.Equal(t, false, m.running)
	t.Logf("\t%s\tShould not be started in %v.", succeed, target)

	assert.Equal(t, int64(0), m.Nanoseconds())
	t.Logf("\t%s\tShould have elapsed set to 0 in %v.", succeed, target)
}

func TestMonitorNilPtr(t *testing.T) {
	t.Log("Given the need to test Start when nil Monitor pointer is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v using nil Capture poiter.", "m.Start")
		{
			var m *Monitor
			_ = m.Start()
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}

func TestMonitorElapsed(t *testing.T) {
	var m *Monitor
	var area string
	target := "NewMonitor"
	t.Log("Given the need to test NewMonitor.")
	{
		t.Logf("\tTest 0:\tWhen calling %v.", target)
		{
			area = generateRandomString(8)
			m, _ = NewMonitor(area, false)
		}

		assert.Assert(t, m != nil)
		t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)

		assert.Equal(t, area, m.Area)
		t.Logf("\t%s\tShould have correct Area set in %v.", succeed, target)

		assert.Equal(t, false, m.running)
		t.Logf("\t%s\tShould not be started in %v.", succeed, target)

		assert.Equal(t, int64(0), m.Nanoseconds())
		t.Logf("\t%s\tShould have elapsed set to 0 in %v.", succeed, target)
	}
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials

	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf) // E.g. "3i[g0|)z"

	return str
}
