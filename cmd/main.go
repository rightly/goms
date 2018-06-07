package main

import (
	"monitoring/internal"
	"monitoring/core"
	"monitoring/collector"
	"monitoring/manager"
	"sync"
)

func main() {
	// WaitGroup 초기화
	var wg sync.WaitGroup
	wg.Add(1)

	config := internal.SetConfigFile()
	//cfg := readCommand()

	system := core.New()
	client := collector.New(system, config)

	go client.Start(&wg)
	if config.Role == "manager" || config.Role =="dev" {
		server := manager.New(config)
		wg.Add(1)
		go server.Start(&wg)
	}

	wg.Wait()
}

/*func readCommand() {
	application := new(string)
	flag.StringVar(application, "app","system","-app is flags that purpose of this server\n")

	flag.Parse()

	// command line argument 의 갯수가 0개 이거나 설정하지 않은 남은 argument 가 있다면 return
	if flag.NFlag() == 0 || flag.NArg() != 0{
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(*application)
}*/