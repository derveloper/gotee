package gotee

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Tee(lineFunc func([]string), flushInterval time.Duration) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)

	ticker := time.NewTicker(flushInterval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <- ticker.C:
				lines = flush(lines, lineFunc)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	close(quit)

	flush(lines, lineFunc)
}

func flush(lines []string, lineFunc func([]string)) []string {
	if len(lines) != 0 {
		lineFunc(lines)
		for _, line := range lines {
			fmt.Println(line)
		}
		lines = make([]string, 0)
	}
	return lines
}
