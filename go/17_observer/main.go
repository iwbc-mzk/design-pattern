package main

import (
	"fmt"
)

// func main() {
// 	pn := NewPCNotifier()
// 	mn := NewMobileNotifier()
// 	sensor := NewEarthquakeSensor()

// 	sensor.Subscribe(pn)
// 	sensor.Subscribe(mn)
// 	sensor.NotifiyAll()

// 	sensor.Unsubscribe(mn)
// 	sensor.NotifiyAll()
	
// }

func main() {
	sensorForFunc := NewEarthquakeSensorForFunction()
	f_pc := func(n Notification) {
		fmt.Printf("[%s] %s (to PC) [Use Function]\n", n.Title, n.Message)
	}
	f_mobile := func(n Notification) {
		fmt.Printf("[%s] %s (to Mobile) [Use Function]\n", n.Title, n.Message)
	}
	pc_id := sensorForFunc.Subscribe(f_pc)
	sensorForFunc.Subscribe(f_mobile)
	sensorForFunc.NotifiyAll()

	sensorForFunc.Unsubscribe(pc_id)
	sensorForFunc.NotifiyAll()
}