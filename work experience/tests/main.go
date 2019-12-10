package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"parspooyesh.com/Rebin/Test-Bed/internal/config"
)

func init() {
	config.InitConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	wg sync.WaitGroup
	// packetChanel     = make(chan string, buffer)
	scenariosNpingCmdMap     = make(map[int][]string)
	treadCount           int = 200
	buffer               int = 2000
)

func main() {

	// go func() {
	// 	for treadNum := 0; treadNum < treadCount; treadNum++ {
	// 		go sendPacketFromChanel()
	// 	}
	// }()

	for senarioNum := range config.Scenarios.Scenario {
		log.Println("Scenario", config.Scenarios.Scenario[senarioNum].Name, "started")

		go sendPacketToChanel(senarioNum)

		// ticker := time.NewTicker(300 * time.Millisecond)
		// finishCounter := 0

		// 	for range ticker.C {
		// 		log.Println("len of chan", senarioNum, len(packetChanel))
		// 		if len(packetChanel) == 0 {
		// 			finishCounter++
		// 			if finishCounter > 1 {
		// 				// close(packetChanel)
		// 				log.Println(config.Scenarios.Scenario[senarioNum].Name, "finished")
		// 				goto end
		// 			}
		// 		}
		// 	}
		// end:
		// 	// fmt.Println(runtime.NumGoroutine())
	}

	sendPacketFromChanel()
}

//SendPacketToChanel func will send a bash command to packetChanel.
func sendPacketToChanel(senarioNum int) {
	srcIps, _ := hosts(config.Scenarios.Scenario[senarioNum].Params.SrcIP)
	// dstIps, _ := hosts(config.Scenarios.Scenario[senarioNum].Params.DestIp)
	dstIPCount, _ := hostsCount(config.Scenarios.Scenario[senarioNum].Params.DestIp)

	npingPossibility := dstIPCount *
		(config.Scenarios.Scenario[senarioNum].Params.DestPortEnd -
			config.Scenarios.Scenario[senarioNum].Params.DestPortStart)

	packetCounter := 0

sendPacketToChanelStart:
	for _, SrcIP := range srcIps {
		for SrcPort := config.Scenarios.Scenario[senarioNum].Params.SrcPortStart; SrcPort <= config.Scenarios.Scenario[senarioNum].Params.SrcPortEnd; SrcPort++ {
			if npingPossibility <= config.Scenarios.Scenario[senarioNum].Params.Size-packetCounter {
				scenariosNpingCmdMap[senarioNum] = append(scenariosNpingCmdMap[senarioNum],
					fmt.Sprintf("nping -c %v -rate %v --%v -p %v-%v --dest-ip %v -g %v -S %v --data-string justForTest &",
						config.Scenarios.Scenario[senarioNum].Params.Size,
						config.Scenarios.Scenario[senarioNum].Params.Rate,
						strings.ToLower(config.Scenarios.Scenario[senarioNum].Params.ProtocolVersion),
						config.Scenarios.Scenario[senarioNum].Params.DestPortStart,
						config.Scenarios.Scenario[senarioNum].Params.DestPortEnd,
						config.Scenarios.Scenario[senarioNum].Params.DestIp, SrcPort, SrcIP))
				packetCounter += config.Scenarios.Scenario[senarioNum].Params.Size
				if packetCounter == config.Scenarios.Scenario[senarioNum].Params.Size {
					goto sendPacketToChanelEnd
				}
			} else {
				scenariosNpingCmdMap[senarioNum] = append(scenariosNpingCmdMap[senarioNum],
					fmt.Sprintf("nping -c %v -rate %v --%v -p %v-%v --dest-ip %v -g %v -S %v --data-string justForTest &",
						config.Scenarios.Scenario[senarioNum].Params.Size,
						config.Scenarios.Scenario[senarioNum].Params.Rate,
						strings.ToLower(config.Scenarios.Scenario[senarioNum].Params.ProtocolVersion),
						config.Scenarios.Scenario[senarioNum].Params.DestPortStart,
						config.Scenarios.Scenario[senarioNum].Params.DestPortEnd,
						config.Scenarios.Scenario[senarioNum].Params.DestIp, SrcPort, SrcIP))
				packetCounter += npingPossibility
				if packetCounter == config.Scenarios.Scenario[senarioNum].Params.Size {
					goto sendPacketToChanelEnd
				}
			}
		}
	}

sendPacketToChanelEnd:
	if packetCounter < config.Scenarios.Scenario[senarioNum].Params.Size {
		goto sendPacketToChanelStart
	}
	// packetCounter := 0

	// sendPacketToChanelStart:
	// 	for _, SrcIP := range srcIps {
	// 		for SrcPort := config.Scenarios.Scenario[senarioNum].Params.SrcPortStart; SrcPort <= config.Scenarios.Scenario[senarioNum].Params.SrcPortEnd; SrcPort++ {
	// 			packetChanel <- fmt.Sprintf("nping -rate %v --%v -p %v-%v --dest-ip %v -g %v -S %v --data-string justForTest &",
	// 				config.Scenarios.Scenario[senarioNum].Params.Rate,
	// 				strings.ToLower(config.Scenarios.Scenario[senarioNum].Params.ProtocolVersion),
	// 				config.Scenarios.Scenario[senarioNum].Params.DestPortStart,
	// 				config.Scenarios.Scenario[senarioNum].Params.DestPortEnd,
	// 				config.Scenarios.Scenario[senarioNum].Params.DestIp, SrcPort, SrcIP)
	// 			packetCounter++
	// 			if packetCounter == config.Scenarios.Scenario[senarioNum].Params.Size {
	// 				goto sendPacketToChanelEnd
	// 			}

	// 			// for _, DstIP := range dstIps {
	// 			// 	for DestPort := config.Scenarios.Scenario[senarioNum].Params.DestPortStart; DestPort <= config.Scenarios.Scenario[senarioNum].Params.DestPortEnd; DestPort++ {

	// 			// 	}
	// 			// }
	// 		}
	// 	}

	// sendPacketToChanelEnd:
	// 	if packetCounter < config.Scenarios.Scenario[senarioNum].Params.Size {
	// 		goto sendPacketToChanelStart
	// 	}
}

func sendPacketFromChanel() {
	// for cmd := range packetChanel {
	// 	runcmd(cmd, true)
	// }

	for _, cmds := range scenariosNpingCmdMap {
		for _, cmd := range cmds {
			runcmd(cmd, true)
		}
	}
}

//runcmd func will take a linux command and run it then will return the result.
func runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic(err)
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		panic(err)
	}
	return out
}

//hosts func will take a ip and range then will return all of the submasks.
func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

//inc fun is used for hosts func.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func hostsCount(cidr string) (int, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return -1, err
	}

	var ips int

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips++
	}
	// remove network address and broadcast address
	return ips - 2, nil
}
