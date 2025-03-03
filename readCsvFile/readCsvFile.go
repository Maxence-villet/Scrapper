package readcsvfile

import (
	"encoding/csv"
	"log"
	"os"
)

type readCsvFileHandler interface {
	ReadCsvFile()
	GetBlackList() []string
	GetWhiteList() []string
	SplitCsvData()
	NewReadCsvFile()
}

type readCsvFile struct {
	filename  string
	whitelist []string
	blacklist []string
}

func (c *readCsvFile) GetBlackList() []string {
	return c.blacklist
}

func (c *readCsvFile) GetWhiteList() []string {
	return c.whitelist
}

func (c *readCsvFile) ReadCsvFile() [][]string {
	f, err := os.Open(c.filename)
	if err != nil {
		log.Fatal("Unable to read input file "+c.filename, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+c.filename, err)
	}

	return records
}

func (c *readCsvFile) SplitCsvData() {
	records := c.ReadCsvFile()
	for i, _ := range records {
		c.blacklist = append(c.blacklist, records[i][1])
		c.whitelist = append(c.whitelist, records[i][0])
	}
}

func NewReadCsvFile(filename string) *readCsvFile {
	c := &readCsvFile{
		filename:  filename,
		blacklist: []string{},
		whitelist: []string{},
	}

	return c
}
