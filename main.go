package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DNSRecords struct {
	DNS []DNSRecord `json:"records"`
}

type DNSRecord struct {
	Domain  string `json:"name"`
	Type    string `json:"type"`
	Records string `json:"value"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var allRecords []DNSRecords
	var dmnRecords []DNSRecord
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allDmnRecords := DNSRecords{DNS: dmnRecords}
			allRecords = append(allRecords, allDmnRecords)
			dmnRecords = nil
		}

		fields := strings.Fields(line)
		if len(fields) >= 3 {
			newRecord := DNSRecord{
				Domain:  strings.TrimSuffix(fields[0], "."),
				Type:    fields[1],
				Records: strings.TrimSuffix(fields[2], "."),
			}
			dmnRecords = append(dmnRecords, newRecord)
		}
	}
	
	RecordsObj, _ := json.Marshal(allRecords)
	fmt.Println(string(RecordsObj))
}
