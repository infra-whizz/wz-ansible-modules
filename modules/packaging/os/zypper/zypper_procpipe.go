package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type TextProcessStream struct {
	filePipe *os.File
}

// NewTextProcessStream creates a TextProcessStream instance. Management of the pipe file is solely on module caller.
func NewTextProcessStream(fname string) *TextProcessStream {
	var err error
	zs := new(TextProcessStream)
	ioutil.WriteFile(fname, []byte(""), 0644)
	zs.filePipe, err = os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zs
}

func (zs *TextProcessStream) parseLine(line string) string {
	line = strings.TrimSpace(line)
	var obj string
	switch lcLine := strings.ReplaceAll(strings.ToLower(line), " ", ""); {
	case strings.HasPrefix(lcLine, "<progress"):
		obj = zs.parseProgress(line)
	case strings.HasPrefix(lcLine, "<download"):
		obj = zs.parseDownload(line)
	case strings.HasPrefix(lcLine, "<message"):
		obj = zs.parseMessage(line)
	}

	return obj
}

func (zs *TextProcessStream) parseMessage(line string) string {
	var message ZypperMessage
	if err := xml.Unmarshal([]byte(line), &message); err != nil {
		panic(err)
	}
	return fmt.Sprintf("LOG\t%s\t%s", message.Type, strings.ReplaceAll(message.Text, "\t", " "))
}

func (zs *TextProcessStream) parseDownload(line string) string {
	var downloadMessage ZypperDownload
	if err := xml.Unmarshal([]byte(line), &downloadMessage); err != nil {
		panic(err)
	}
	return fmt.Sprintf("LOG\tinfo\t%s", downloadMessage.Url)
}

func (zs *TextProcessStream) parseProgress(line string) string {
	var progressMessage ZypperProgress
	if err := xml.Unmarshal([]byte(line), &progressMessage); err != nil {
		panic(err)
	}
	return fmt.Sprintf("PGS\t%s\t%s", progressMessage.Value, progressMessage.Name)
}

// Write data to the underlying pipe file
func (zs *TextProcessStream) Write(data []byte) (n int, err error) {
	line := zs.parseLine(strings.TrimSpace(string(data)))
	if line != "" {
		zs.filePipe.WriteString(line + "\n")
	}
	return len(data), nil
}

// Close stream
func (zs *TextProcessStream) Close() error {
	return zs.filePipe.Close()
}
