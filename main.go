package main

import (
	"fmt"

	"github.com/sethlessard/goku/roku"
)

func main() {
	devices, errors := roku.DiscoverRokus(5)

	for _, value := range devices {
		fmt.Printf("%v\n", value)
	}

	for _, value := range errors {
		fmt.Printf("%v\n", value.Error)
	}

	// bedroomUSN := "uuid:roku:ecp:2N0031703463"
	livingroomUSN := "uuid:roku:ecp:YN004S825465"

	apps, err := devices[livingroomUSN].Roku.GetApps()
	if err != nil {
		panic(err)
	}
	fmt.Println(apps)

	activeApp, err := devices[livingroomUSN].Roku.GetActiveApp()
	if err != nil {
		panic(err)
	}
	fmt.Println(activeApp)

	livingroom := devices[livingroomUSN].Roku
	info, err := livingroom.GetDeviceInfo()

	if err != nil {
		panic(err)
	}

	fmt.Println(info)
}
