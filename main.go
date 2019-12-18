package main

import (
	"fmt"
	"github.com/alexstov/chronos/rtm"
	"github.com/sirupsen/logrus"
)

func main() {
	// Start RTM
	var cap *rtm.Capture
	var capErr error
	if cap, capErr = rtm.BeginCapture("TestApp", rtm.ConfigOption{Option: rtm.Units, Value: rtm.Nanoseconds}); capErr != nil {
		logrus.Error("")
	} else {
		defer cap.Finish()
	}

	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	t1 := cap.Start("Loop")
	for _, v := range xs {
		total += v
	}
	t1.Elapsed()

	t2 := cap.Start("IO")
	fmt.Println(total / float64(len(xs)))
	t2.Elapsed()
}
