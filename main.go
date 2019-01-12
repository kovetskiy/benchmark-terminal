package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/docopt/docopt-go"
)

var (
	version = "[manual build]"
	usage   = "benchmark-terminal " + version + `

Usage:
  benchmark-terminal [options] -t <seconds>
  benchmark-terminal [options] -i <iterations>
  benchmark-terminal -h | --help
  benchmark-terminal --version

Options:
  -t --time <seconds>           Limit benchmark by time.
  -i --iterations <iterations>  Limit benchmark by number of iterations.  
  -c --char <char>              Print specified symbol [default: X].
  -o --output <path>            Use specified output. [default: /dev/stdout]
  -w --width <n>                Force to use specified width instead of automatically obtained one.
  -h --help                     Show this screen.
  --version                     Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile(
		args["--output"].(string),
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var width int
	if argWidth, ok := args["--width"].(string); ok {
		width, err = strconv.Atoi(argWidth)
		if err != nil {
			log.Fatalf("unable to parse --width flag: %s", err)
		}
	} else {
		width = getWidth()
	}

	char := args["--char"].(string)
	line := strings.Repeat(char, width) + "\n"

	argTime, withTime := args["--time"].(string)
	if withTime {
		seconds, err := strconv.Atoi(argTime)
		if err != nil {
			log.Fatalln(err)
		}

		benchmarkSeconds(output, line, seconds)
	}

	argIterations, withIterations := args["--iterations"].(string)
	if withIterations {
		iterations, err := strconv.Atoi(argIterations)
		if err != nil {
			log.Fatalln(err)
		}

		benchmarkIterations(output, line, iterations)
	}
}

func benchmarkSeconds(output *os.File, line string, seconds int) {
	buffer := []byte(line)
	done := make(chan struct{})

	start := time.Now()
	go func() {
		<-time.After(time.Duration(seconds) * time.Second)
		close(done)
	}()

	i := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			output.Write(buffer)
			i++
		}
	}

	finish := time.Since(start)

	fmt.Printf("line length: %d\n", len(line))
	fmt.Printf(
		"%d lines per %v (%.2f seconds)\n",
		i,
		finish.String(),
		finish.Seconds(),
	)
	fmt.Printf("speed: %.2f l/s\n", float64(i)/finish.Seconds())
}

func benchmarkIterations(output *os.File, line string, iterations int) {
	buffer := []byte(line)

	start := time.Now()

	for i := 0; i < iterations; i++ {
		_, err := output.Write(buffer)
		if err != nil {
			panic(err)
		}
	}

	finish := time.Since(start)

	fmt.Printf("line length: %d\n", len(line))
	fmt.Printf(
		"%d lines per %v (%.2f seconds)\n",
		iterations,
		finish.String(),
		finish.Seconds(),
	)
	fmt.Printf("speed: %.2f l/s\n", float64(iterations)/finish.Seconds())
}

func getWidth() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col)
}
