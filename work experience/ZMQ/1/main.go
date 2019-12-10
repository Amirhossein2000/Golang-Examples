package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strconv"
	"sync/atomic"
)

var counter int64

func main() {
	go router()

	for id := 0; id < 20; id++ {
		go client(id)
		go worker(id)
	}

	worker(21)
}

func router() {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	frontend.Bind("tcp://*:5559")
	backend.Bind("tcp://*:5560")

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	//  Switch messages between sockets
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				for {
					msg, _ := s.Recv(0)
					if more, _ := s.GetRcvmore(); more {
						backend.Send(msg, zmq.SNDMORE)
					} else {
						backend.Send(msg, 0)
						break
					}
				}
			case backend:
				for {
					msg, _ := s.Recv(0)
					if more, _ := s.GetRcvmore(); more {
						frontend.Send(msg, zmq.SNDMORE)
					} else {
						frontend.Send(msg, 0)
						break
					}
				}
			}
		}
	}
}

func worker(id int) {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect("tcp://localhost:5560")

	for ; ; atomic.AddInt64(&counter, 1) {
		//  Wait for next request from client
		request, _ := responder.Recv(0)
		//  Do some 'work'

		//  Send reply back to client
		responder.Send(request, 0)
		//fmt.Println(id, counter)
	}
}

func client(clientID int) {
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5559")

	for request := 0; request < 10; request++ {
		requester.Send(strconv.Itoa(clientID), 0)
		reply, _ := requester.Recv(0)
		if strconv.Itoa(clientID) == reply {
			fmt.Printf("%d = %s\n", clientID, reply)
		}
	}
}
