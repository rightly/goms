package core

import (
	"monitoring/internal"
	"github.com/shirou/gopsutil/cpu"
	"time"
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/docker"
	"fmt"
)

type System struct {
	*CPU                 `json:"cpu"`
	*Disk                `json:"disk"`
	*Memory              `json:"memory"`
	*Network             `json:"network"`
	*Host                `json:"host"`
	Timestamp *time.Time `json:"timestamp"`
}

func New() *System {
	system := &System{
		CPU:     &CPU{},
		Memory:  &Memory{},
		Network: &Network{},
		Disk:    &Disk{},
		Host:    &Host{},
	}

	system.initHost()

	return system
}

func (r *System) Polling() {
	internal.CheckErr(r.CPU.PollingUsage(), "couldn't polling cpu metric")
	internal.CheckErr(r.Memory.Polling(), "couldn't polling memory metric")
	internal.CheckErr(r.Network.Polling(), "couldn't polling network metric")
	internal.CheckErr(r.Disk.PollingUsage("/"), "couldn't polling disk metric")

	now := time.Now()
	nanos := now.UnixNano()
	timestamp := time.Unix(0, nanos)
	r.Timestamp = &timestamp
}

func (r *System) String() string {
	indentedString := "Time: " + r.Timestamp.String() + "\n" +
		r.CPU.String() + "\n" +
			r.Memory.String() + "\n" +
			r.Network.String() + "\n" +
			r.Disk.String() + "\n"

	return indentedString
}

type CPU struct {
	Info  []cpu.InfoStat  `json:"info"`
	Time  []cpu.TimesStat `json:"time"`
	Usage []float64       `json:"usage"`
}

func (r *CPU)PollingInfo() error {
	var err error

	r.Info, err = cpu.Info()
	r.Time, err = cpu.Times(true)
	r.Usage, err = cpu.Percent(1*time.Second, true)

	return err
}

func (r *CPU)PollingUsage() error {
	var err error

	r.Info, err = cpu.Info()
	r.Time, err = cpu.Times(true)
	r.Usage, err = cpu.Percent(1*time.Second, true)

	return err
}

func (r *CPU)String() string {
	indentedInfo, _ := json.MarshalIndent(r.Info, "\t\t", "\t")
	indentedTime, _ := json.MarshalIndent(r.Time, "\t\t", "\t")
	indentedUsage, _ := json.MarshalIndent(r.Usage, "\t\t", "\t")

	indentedString := "{\n\tCPU:{\n\t\t" +
		"Info:" + string(indentedInfo) +
		"\n\t\ttime:" + string(indentedTime) +
		"\n\t\tUsage:" + string(indentedUsage) + "\n}"

	return indentedString
}

type Memory struct {
	Swap    *mem.SwapMemoryStat    `json:"swap"`
	Virtual *mem.VirtualMemoryStat `json:"virtual"`
}

func (r *Memory)Polling() error {
	var err error

	r.Swap, err = mem.SwapMemory()
	r.Virtual, err = mem.VirtualMemory()

	return err
}

func (r *Memory)String() string {
	indentedSwap, _ := json.MarshalIndent(r.Swap, "\t\t", "\t")
	indentedVirtual, _ := json.MarshalIndent(r.Virtual, "\t\t", "\t")

	indentedString := "{\n\tMemory:{\n\t\t" +
		"Swap:" + string(indentedSwap) +
		"\n\t\tVirtual:" + string(indentedVirtual) + "\n}"

	return indentedString
}

type Network struct {
	TCPConnection []net.ConnectionStat `json:"tcp_connection"`
	UDPConnection []net.ConnectionStat `json:"udp_connection"`
	Interface     []net.InterfaceStat  `json:"interface"`
}

func (r *Network)Polling() error {
	var err error

	r.Interface, err = net.Interfaces()
	r.TCPConnection, err = net.Connections("tcp")
	r.UDPConnection, err = net.Connections("udp")

	return err
}

func (r *Network)String() string {
	indentedInfo, _ := json.MarshalIndent(r.Interface, "\t\t", "\t")
	indentedTime, _ := json.MarshalIndent(r.TCPConnection, "\t\t", "\t")
	indentedUsage, _ := json.MarshalIndent(r.UDPConnection, "\t\t", "\t")

	indentedString := "{\n\tNetwork:{\n\t\t" +
		"Interface:" + string(indentedInfo) +
		"\n\t\tTCPConnection:" + string(indentedTime) +
		"\n\t\tUDPConnection:" + string(indentedUsage) + "\n}"

	return indentedString
}

type Disk struct {
	Usage      *disk.UsageStat                `json:"usage"`
	Partition  []disk.PartitionStat           `json:"partition"`
	IOCounters map[string]disk.IOCountersStat `json:"io_counters"`
}

func (r *Disk)PollingInfo(path string) error {
	var err error

	r.Partition, err = disk.Partitions(true)

	return err
}

func (r *Disk)PollingUsage(path string) error {
	var err error

	r.Usage, err = disk.Usage(path)
	r.IOCounters, err = disk.IOCounters("")

	return err
}

func (r *Disk)String() string {
	indentedInfo, _ := json.MarshalIndent(r.Usage, "\t\t", "\t")
	indentedTime, _ := json.MarshalIndent(r.Partition, "\t\t", "\t")
	indentedUsage, _ := json.MarshalIndent(r.IOCounters, "\t\t", "\t")

	indentedString := "{\n\tDisk:{\n\t\t" +
		"Usage:" + string(indentedInfo) +
		"\n\t\tPartition:" + string(indentedTime) +
		"\n\t\tIOCounters:" + string(indentedUsage) + "\n}"

	return indentedString
}

type Host struct {
	Info  *host.InfoStat  `json:"info"`
	Users []host.UserStat `json:"users"`
}

func (r *Host) new() error {
	var err error

	r.Info, err = host.Info()
	r.Users, err = host.Users()

	return err
}

func (r *Host)String() string {
	indentedInfo, _ := json.MarshalIndent(r.Info, "\t\t", "\t")
	indentedUsers, _ := json.MarshalIndent(r.Users, "\t\t", "\t")

	indentedString := "{\n\tHost:{\n\t\t" +
		"Info:" + string(indentedInfo) +
		"\n\t\tUsers:" + string(indentedUsers) + "\n}"

	return indentedString
}

func Tests()  {
	iload, _ := load.Avg()
	idocker, _ := docker.GetDockerIDList()

	il, _ := json.MarshalIndent(iload, "", "\t")
	idoc, _ := json.MarshalIndent(idocker,  "", "\t")

	fmt.Printf("Load Info: %s\n", string(il))
	fmt.Printf("Docker Info: %s\n", string(idoc))
}
