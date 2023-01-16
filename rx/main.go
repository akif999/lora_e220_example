// This reads ...
// Example of ...
package main

import (
	"machine"
	"time"
)

var (
	uart = machine.UART1
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
	led  = machine.LED
)

func main() {
	go blink()

	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: tx, RX: rx})
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			print(string(data))
		}
		time.Sleep(100 * time.Microsecond)
	}
}

func blink() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)
		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
