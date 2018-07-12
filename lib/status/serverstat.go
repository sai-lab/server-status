package status

import (
	"encoding/json"

	"github.com/shirou/gopsutil/mem"
)

type ServerStat struct {
	// Host
	HostName             string `json:"hostname"`
	HostID               string `json:"hostid"`
	VirtualizationSystem string `json:"virtualizationSystem"`
	// Memory
	MemStat               mem.VirtualMemoryStat `json:"memStat"`
	MemoryAcquisitionTime string                `json:"memoryAcquisitionTime"`
	// DiskIO
	DiskIO              []DiskStat `json:"diskIO"`
	diskAcquisitionTime string     `json:"diskAcquisitionTime"`
	// Cpu
	CpuUsedPercent     []float64 `json:"cpuUsedPercent"`
	CpuAcquisitionTime string    `json:"cpuAcquisitionTime"`
	// Apache
	ApacheStat            float64 `json:"apacheStat"`
	ApacheLog             int64   `json:"apacheLog"`
	ReqPerSec             float64 `json:"reqPerSec"`
	ApacheAcquisitionTime string  `json:"apacheAcquisitionTime"`
	// Dstat
	DstatLog             string `json:"dstatLog"`
	DstatAcquisitionTime string `json:"dstatAcquisitionTime"`
	// Time
	LastTime string `json:"time"`
	// Error
	ErrorInfo []error `json:"errorInfo"`
}

type DiskStat struct {
	Name       string `json:"name"`
	IoTime     uint64 `json:"ioTime"`
	WeightedIO uint64 `json:"weightedIO"`
}

func (d ServerStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d DiskStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
