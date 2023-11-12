# monitor_ram
Go code to monitor RAM memory 

## To run the code
`go run main.go`

# How the code runs

## RAM Memory Monitoring in Go

Let's break down the code step by step:

```go
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
			fmt.Println("Error retrieving memory information:", err)
			return
		}

		fmt.Printf("Memory Usage:\n")
		fmt.Printf("  Total: %v MB\n", memInfo.Total>>20)
		fmt.Printf("  Available: %v MB\n", memInfo.Available>>20)
		fmt.Printf("  Used: %v MB\n", memInfo.Used>>20)

		time.Sleep(1 * time.Second)
	}
}
```

Now, let's break down the code:

1. **`package main`**: Indicates that this file belongs to the main package. In Go, an executable program starts with the `main` package.

2. **Imports**:
    - `fmt`: Provides formatting functions similar to `printf` in C.
    - `time`: Provides functions for measuring and manipulating time.
    - `github.com/shirou/gopsutil/v3/mem`: An external library that provides information about system memory usage.

3. **`func main()`**: The main function that gets executed when the program is started.

4. **`for {...}`**: An infinite loop.

5. **`memInfo, err := mem.VirtualMemory()`**: Retrieves information about memory usage using the `gopsutil` library. `mem.VirtualMemory()` returns a `*mem.VirtualMemoryInfo` object that contains information about the system's virtual memory.

6. **`if err != nil {...}`**: Checks for an error when retrieving memory information. If an error occurs, it prints an error message and exits the program.

7. **Printing Memory Information**:
    - **`fmt.Printf("Memory Usage:\n")`**: Prints a header.
    - **`fmt.Printf("  Total: %v MB\n", memInfo.Total>>20)`**: Prints the total amount of memory in megabytes.
    - **`fmt.Printf("  Available: %v MB\n", memInfo.Available>>20)`**: Prints the amount of available memory in megabytes.
    - **`fmt.Printf("  Used: %v MB\n", memInfo.Used>>20)`**: Prints the amount of used memory in megabytes.

8. **`time.Sleep(1 * time.Second)`**: Makes the program wait for 1 second before continuing to the next iteration of the loop. This sets the monitoring update frequency.

This program essentially retrieves and displays information about system memory usage in an infinite loop with a 1-second update frequency. The output format is quite simple, but you can customize it as per your needs. Furthermore, this is a starting point; you can expand and enhance the program based on what you want to achieve with it.
