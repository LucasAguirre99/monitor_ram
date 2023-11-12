// main.go
package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	for {
		memInfo, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error al obtener información de memoria:", err)
			return
		}

		fmt.Printf("Uso de Memoria:\n")
		fmt.Printf("  Total: %v MB\n", memInfo.Total>>20)
		fmt.Printf("  Disponible: %v MB\n", memInfo.Available>>20)
		fmt.Printf("  Usado: %v MB\n", memInfo.Used>>20)

		time.Sleep(1 * time.Second)
	}
}
