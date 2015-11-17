package delim

import "os"
import "encoding/csv"

type DelimMapWriterConfig struct {
	Delimiter string
	File      *os.File
	Header    []string
}

type DelimMapWriter struct {
	delimiter rune
	writer    *csv.Writer
	header    []string
}

func (self *DelimMapWriterConfig) NewDelimMapWriter() *DelimMapWriter {
	mapWriter := new(DelimMapWriter)
	mapWriter.ApplyConfig(self)
	return mapWriter
}

func (self *DelimMapWriter) ApplyConfig(config *DelimMapWriterConfig) {
	// set delimiter
	if config.Delimiter == "" {
		self.delimiter = '|'
	} else {
		runes := []rune(config.Delimiter)
		self.delimiter = runes[0]
	}

	// set writer
	self.writer = csv.NewWriter(config.File)
	self.writer.Comma = self.delimiter

	// set header
	self.header = config.Header
}

func (self *DelimMapWriter) GetHeader() []string {
	return self.header
}

func (self *DelimMapWriter) GetDelimiter() rune {
	return self.delimiter
}

func (self *DelimMapWriter) Flush() {
	self.writer.Flush()
}

func (self *DelimMapWriter) Write(record map[string]string) error {
	var output []string

	for _, field := range self.header {
		val, exists := record[field]
		if !exists {
			val = ""
		}
		output = append(output, val)
	}

	return self.writer.Write(output)
}

func (self *DelimMapWriter) WriteHeader() error {
	return self.writer.Write(self.header)
}
