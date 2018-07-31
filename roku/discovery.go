package roku

import (
	"net"
	"strings"
	"sync"
	"time"
)

type DiscoverResult struct {
	GoroutineId int
	Roku        Roku
}

type DiscoverError struct {
	GoroutineId int
	Error       error
}

func discoverRoku(id int, timeoutSeconds time.Duration) (chan DiscoverResult, chan DiscoverError) {
	discoverChan := make(chan DiscoverResult)
	errorChan := make(chan DiscoverError)

	go func(id int, timeout time.Duration, discoverChan chan DiscoverResult, errorChan chan DiscoverError) {
		ssdp, err := net.ResolveUDPAddr("udp", "239.255.255.250:1900")
		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		address, err := net.ResolveUDPAddr("udp", ":0")
		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		socket, err := net.ListenUDP("udp", address)
		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		// broadcast search signal
		_, err = socket.WriteToUDP([]byte("M-SEARCH * HTTP/1.1\r\n"+
			"HOST: 239.255.255.250:1900\r\n"+
			"MAN: \"ssdp:discover\"\r\n"+
			"ST: roku:ecp\r\n"+
			"MX: 3 \r\n\r\n"), ssdp)

		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		answerBytes := make([]byte, 1024)
		err = socket.SetReadDeadline(time.Now().Add(timeoutSeconds * time.Second))
		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		_, _, err = socket.ReadFromUDP(answerBytes[:])
		if err != nil {
			se := DiscoverError{
				GoroutineId: id,
				Error:       err,
			}
			errorChan <- se
		}

		// fmt.Printf("%s\n", string(answerBytes))
		ret := strings.Split(string(answerBytes), "\r\n")
		usn := strings.TrimPrefix(ret[3], "USN: ")
		location := strings.TrimPrefix(ret[6], "LOCATION: ")

		discoverChan <- DiscoverResult{
			GoroutineId: id,
			Roku: Roku{
				APIAddress: location,
				USN:        usn,
			},
		}
	}(id, timeoutSeconds, discoverChan, errorChan)

	return discoverChan, errorChan
}

// Discover Rokus on the network
func DiscoverRokus(timeoutSeconds time.Duration) (map[string]DiscoverResult, map[int]DiscoverError) {
	discoverChan := make(chan DiscoverResult)
	errorChan := make(chan DiscoverError)

	for i := 0; i < 40; i++ {
		go func(searchId int, discoverChan chan DiscoverResult, errorChan chan DiscoverError) {
			_roku, _errors := discoverRoku(searchId, 10)

		loop:
			for {
				select {
				case discoverResult := <-_roku:
					discoverChan <- discoverResult
					break loop
				case discoverError := <-_errors:
					errorChan <- discoverError
					break loop
				}
			}

		}(i, discoverChan, errorChan)
	}

	var mux sync.Mutex
	var errMux sync.Mutex
	discovered := make(map[string]DiscoverResult)
	errors := make(map[int]DiscoverError)

	for i := 0; i < 10; i++ {
		select {
		case discoverResult := <-discoverChan:
			mux.Lock()
			discovered[discoverResult.Roku.USN] = discoverResult
			mux.Unlock()
		case discoverError := <-errorChan:
			errMux.Lock()
			errors[discoverError.GoroutineId] = discoverError
			errMux.Unlock()
		}
	}

	return discovered, errors
}
