package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"https://github.com/harbour-tech/go-ethereum-arbitrum/metrics"
	metrics2 "runtime/metrics"
	"testing"
	"time"
)

func TestCollectRuntimeMetrics(t *testing.T) {
	t.Skip("Only used for generating testdata")
	serialize := func(path string, histogram *metrics2.Float64Histogram) {
		f := new(bytes.Buffer)
		if err := gob.NewEncoder(f).Encode(histogram); err != nil {
			panic(err)
		}
		fmt.Printf("var %v = %q\n", path, f.Bytes())
	}
	time.Sleep(2 * time.Second)
	stats := metrics.ReadRuntimeStats()
	serialize("schedlatency", stats.SchedLatency)
	serialize("gcpauses", stats.GCPauses)
}
