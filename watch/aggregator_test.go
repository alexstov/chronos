package watch

import (
	"testing"

	"gotest.tools/assert"
)

func TestAggregatorNewAggregator(t *testing.T) {
	var a *Aggregator
	var area string
	var target = "NewAggregator"
	t.Log("Given the need to test NewAggregator.")
	{
		t.Logf("\tTest 0:\tWhen calling %v.", target)
		{
			area = generateRandomString(8)
			a = NewAggregator(area)
		}
	}

	assert.Assert(t, a != nil)
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)

	assert.Equal(t, area, a.Area)
	t.Logf("\t%s\tShould have correct Area set in %v.", succeed, target)

	assert.Assert(t, a.Monitors != nil)
	t.Logf("\t%s\tShould have monitors slice created in %v.", succeed, target)
}
