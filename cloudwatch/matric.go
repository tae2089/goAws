package cloudwatch

import (
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

type CloudWatchMatrix func() (string, float64)

func MemoryPercent() (string, float64) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	stringPercentage := strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64)
	percentage, err := strconv.ParseFloat(stringPercentage, 64)
	if err != nil {
		panic(err)
	}
	return "memory", percentage
}

func DiskUsagedPercent() (string, float64) {
	return "Disk Used", 1
}
