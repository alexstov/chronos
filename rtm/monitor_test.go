package rtm

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2714"

func TestMonitorNilPtr(t *testing.T) {
	t.Log("Given the need to call Start when nil Monitor poiter is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v uning nil Monitor poiter.", "monitor.Start")
		{
			var m *Monitor
			m.Start(OpsTotal)
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}

// func BenchmarkSanitizeAWSClientIDMatched(b *testing.B) {
// 	var s string

// 	for i := 0; i < b.N; i++ {
// 		s = awsClient.sanitizeAWSClientID(regexMatch)
// 	}

// 	gc = s
// }
