package agent

import (
	"time"

	"github.com/dcos/dcos-metrics/producers"
)

func (c *Collector) createDatapoints() []producers.Datapoint {
	ts := thisTime()
	dps := []producers.Datapoint{}
	for _, container := range c.containerMetrics {
		addDps := []producers.Datapoint{
			producers.Datapoint{
				Name:      "cpus.user.time",
				Value:     container.Statistics.CpusUserTimeSecs,
				Unit:      SECONDS,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "cpus.system.time",
				Value:     container.Statistics.CpusSystemTimeSecs,
				Unit:      SECONDS,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "cpus.limit",
				Value:     container.Statistics.CpusLimit,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "cpus.throttled.time",
				Value:     container.Statistics.CpusThrottledTimeSecs,
				Unit:      SECONDS,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "mem.total",
				Value:     container.Statistics.MemTotalBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "mem.limit",
				Value:     container.Statistics.MemLimitBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "disk.limit",
				Value:     container.Statistics.DiskLimitBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "disk.used",
				Value:     container.Statistics.DiskUsedBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.rx.packets",
				Value:     container.Statistics.NetRxPackets,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.rx.bytes",
				Value:     container.Statistics.NetRxBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.rx.errors",
				Value:     container.Statistics.NetRxErrors,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.rx.dropped",
				Value:     container.Statistics.NetRxDropped,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.tx.packets",
				Value:     container.Statistics.NetRxPackets,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.tx.bytes",
				Value:     container.Statistics.NetRxBytes,
				Unit:      BYTES,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.tx.errors",
				Value:     container.Statistics.NetRxErrors,
				Unit:      COUNT,
				Timestamp: ts,
			},
			producers.Datapoint{
				Name:      "net.tx.dropped",
				Value:     container.Statistics.NetRxDropped,
				Unit:      COUNT,
				Timestamp: ts,
			},
		}

		for _, dp := range addDps {
			dps = append(dps, dp)
		}
	}

	return dps
}

// -- helpers

func thisTime() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
