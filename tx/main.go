// This reads from UART1 and outputs to default serial, usually UART0 or USB.
// Example of how to work with UARTs other than the default.
package main

import (
	"bytes"
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
	payload := bytes.Join([][]byte{{0x00, 0x00, 0x00}, []byte("Hello\r\n")}, []byte(""))
	for {
		uart.Write(payload)
		time.Sleep(time.Millisecond * 5000)
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
