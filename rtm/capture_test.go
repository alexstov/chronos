package rtm

import "testing"

func TestCaptureNilPtr(t *testing.T) {
	t.Log("Given the need to call Start when nil Capture poiter is used.")
	{
		t.Logf("\tTest 0:\tWhen calling %v uning nil Capture poiter.", "capture.Start")
		{
			var c *Capture
			c.Start(OpsTotal)
		}
	}
	t.Logf("\t%s\tShould be able to make the Start call and not crash.", succeed)
}

// func TestSanitizeAWSClientID(t *testing.T) {
// 	var s string

// 	s = awsClient.sanitizeAWSClientID(regexMatch)
// 	assert.Equal(t, sanitized, s)
// }

// func BenchmarkSanitizeAWSClientIDMatched(b *testing.B) {
// 	var s string

// 	for i := 0; i < b.N; i++ {
// 		s = awsClient.sanitizeAWSClientID(regexMatch)
// 	}

// 	gc = s
// }
