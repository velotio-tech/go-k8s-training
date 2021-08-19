package shell

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

const histBufSize int = 1024

//	 A data-structure for history command.
type History struct {
	cmdBuffer  []string
	bufCounter int
	fileName   string
}

//	Initializes the history cmd data-structure
func (hist *History) Init() {
	hist.cmdBuffer = make([]string, histBufSize)
	hist.bufCounter = 0
	wd, _ := os.Getwd()
	(*hist).fileName = wd + "/history.txt"
}

//	Adds a command to the buffer
func (hist *History) addCommand(cmd string) {
	if (*hist).bufCounter == histBufSize-1 {
		(*hist).writeBufferToFile()
		(*hist).resetBuffer()
	}

	(*hist).cmdBuffer[(*hist).bufCounter] = cmd
	(*hist).bufCounter++
}

//	writes the whole buffer to file.
func (hist *History) writeBufferToFile() {
	file, err := os.OpenFile((*hist).fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i := 0; i < (*hist).bufCounter; i++ {
		file.WriteString((*hist).cmdBuffer[i] + "\n")
	}
}

//	resets the buffer.
func (hist *History) resetBuffer() {
	(*hist).bufCounter = 0
}

//	Shows the complete history.
func (hist *History) ShowHistory() {
	tw := tabwriter.NewWriter(os.Stdout, 1, 2, 1, '\t', tabwriter.AlignRight)
	file, _ := os.Open((*hist).fileName)
	fscanner := bufio.NewScanner(file)
	cmdCnt := 1
	for fscanner.Scan() {
		cmd := fscanner.Text()
		fmt.Fprintf(tw, strconv.Itoa(cmdCnt)+" "+cmd+"\n")
		cmdCnt++
	}
	tw.Flush()
}

//	Buffer data is written to file.
func (hist *History) Close() {
	(*hist).writeBufferToFile()
}
