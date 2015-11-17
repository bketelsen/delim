package delim

import "os"
import "encoding/csv"
import "unicode/utf8"
import "fmt"
import "io"

type DelimMapReaderConfig struct {
	Delimiter string
	File      *os.File
}

type DelimMapReader struct {
	delimiter rune
	reader    *csv.Reader
	header    []string
}

func (self *DelimMapReaderConfig) NewDelimMapReader() *DelimMapReader {
	mapReader := new(DelimMapReader)
	mapReader.ApplyConfig(self)
	return mapReader
}

func (self *DelimMapReader) ApplyConfig(config *DelimMapReaderConfig) {
	// set delimiter
	// TODO: Should verify that the string length isn't > 1, also, push to shared library
	if config.Delimiter == "" {
		self.delimiter = '|'
	} else {
		r, size := utf8.DecodeRuneInString(config.Delimiter)
		if size == 0 {
			fmt.Println("Invalid delimiter: %s", config.Delimiter)
			os.Exit(1)
		}
		self.delimiter = r
	}

	// set reader
	self.reader = csv.NewReader(config.File)
	self.reader.Comma = self.delimiter
	self.reader.LazyQuotes = true
	self.reader.FieldsPerRecord = 0

	// set header
	header, err := self.reader.Read()
	if err != nil {
		if err == io.EOF {
			fmt.Println("Missing header on file")
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	self.header = header
}

func (self *DelimMapReader) GetHeader() []string {
	return self.header
}

func (self *DelimMapReader) Next() (map[string]string, error) {
	record := make(map[string]string)
	data, err := self.reader.Read()
	if err != nil {
		return nil, err
	}

	for i, field := range self.header {
		record[field] = data[i]
	}

	return record, nil
}
