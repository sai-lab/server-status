package status

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetServerStat() ServerStat {
	var wg sync.WaitGroup
	var ss ServerStat

	wg.Add(5)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetHostStat()
	}(&ss)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetMemoryStat()
	}(&ss)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetDiskIOStat()
	}(&ss)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetApacheStat()
	}(&ss)
	go func(d *ServerStat) {
		defer wg.Done()
		d.GetCpuStat()
	}(&ss)

	// errDstatLog := d.GetDstatLog()

	wg.Wait()
	ss.GetTime()

	return ss
}

func (s *ServerStat) GetHostStat() {
	h, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}
	s.HostName = h.Hostname
	s.HostID = h.HostID
	s.VirtualizationSystem = h.VirtualizationSystem
}

func (s *ServerStat) GetMemoryStat() {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	s.MemStat = *m
}

func (s *ServerStat) GetDiskIOStat() {
	var ds []DiskStat
	i, err := disk.IOCounters()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range i {
		var d DiskStat
		d.Name = k
		d.IoTime = v.IoTime
		d.WeightedIO = v.WeightedIO
		ds = append(ds, d)
	}
	s.DiskIO = ds
}

func (s *ServerStat) GetApacheStat() {
	var operatingData, accessData string
	out, err := exec.Command("curl", "localhost/server-status?auto").Output()
	if err != nil {
		log.Fatal(err)
	}

	d := string(out)
	lines := strings.Split(strings.TrimRight(d, "\n"), "\n")

	for _, v := range lines {
		if strings.Index(v, "Scoreboard") != -1 {
			operatingData = v
		}
		if strings.Index(v, "Total Accesses") != -1 {
			accessData = v
		}
	}

	board := operatingData[12:]
	totalAccess := accessData[16:]

	all := len(strings.Split(board, ""))
	idles := strings.Count(board, "_") + strings.Count(board, ".")

	r := float64((all - idles)) / float64(all)
	s.ApacheStat = r
	s.ApacheLog, err = strconv.ParseInt(totalAccess, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ServerStat) GetDstatLog() {
	out, err := exec.Command("tail", "-1", "/home/ansible/dstatlog.csv").Output()
	if err != nil {
		log.Fatal(err)
	}
	s.DstatLog = string(out)
}

func (s *ServerStat) GetTime() {
	now := time.Now()
	s.Time = now.String()
}

func (s *ServerStat) GetCpuStat() {
	c, err := cpu.Percent(0.0, false)
	if err != nil {
		log.Fatal(err)
	}
	s.CpuUsedPercent = c
}
