package core

import (
	"monitoring/internal"
	"fmt"
	"os"
	"time"
)

func (r *System)initHost() int {
	internal.CheckErr(r.Host.new(),"couldn't load host info")
	internal.CheckErr(r.CPU.PollingInfo(), "couldn't load cpu info")
	internal.CheckErr(r.Disk.PollingInfo("/"), "couldn't load disk info")

	fmt.Fprintf(os.Stdout,"System Initialization\n" +
		"Host\t\t%v\n" +
		"Uptime\t\t%v\n" +
		"BootTime\t%v\n" +
		"OS/Platform\t%v/%v\n" +
		"Kernal\t\t%v\n" +
		"CPU Vendor\t%v\n" +
		"Core\t\t%v\n" +
		"Model\t\t%v\n" +
		"Disk\n%v\n",r.Host.Info.Hostname, r.Host.Info.Uptime, r.Host.Info.BootTime, r.Host.Info.OS, r.Host.Info.Platform,
		r.Host.Info.KernelVersion, r.CPU.Info[0].VendorID, r.CPU.Info[0].Cores, r.CPU.Info[0].ModelName, r.Disk)

	return 0
}

func (r *System)Collect(c chan bool) {
	for {
		r.Polling()
		//fmt.Fprintf(os.Stdout, "%d.polling : %v\n  CPU Usage: %v\n", cnt, r.Timestamp, r.CPU.Usage)
		c <- true
		time.Sleep(5 * time.Second)
	}
}