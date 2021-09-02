package cmd

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// some functions: parse bytes

// iCIBA response struct
type IcibaResp struct {
	Word        string   `xml:"word_name"`
	Ps          []string `xml:"ps"`
	Pos         []string `xml:"pos"`
	Acceptation []string `xml:"acceptation"`
	Sent        []sent   `xml:"sent"`
}

type sent struct {
	Orig  string `xml:"orig"`
	Trans string `xml:"trans"`
}

// receive bytes
func IcibaParse(data []byte) {
	var resp = IcibaResp{}
	// parse XML
	// nice function!
	var err = xml.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	var phonetic string
	for index, value := range resp.Ps {
		if index == 0 {
			phonetic += "en-UK[ " + value + " ] "
			continue
		}
		phonetic += "en-US[ " + value + " ] "
	}
	fmt.Printf(" %s %s %s \n\n", resp.Word, phonetic, " ~ from iciba.com")
	if !isChinese(resp.Word) {
		for i := 0; i < len(resp.Pos); i = i + 1 {
			fmt.Printf(" %s %s %s", "-", resp.Pos[i], resp.Acceptation[i])
		}
	}
	fmt.Print("\n")
	for i := 0; i < len(resp.Sent); i = i + 1 {
		fmt.Printf("%s %s\n", strconv.Itoa(i+1)+".", del(resp.Sent[i].Orig))
		fmt.Printf("    %s\n", del(resp.Sent[i].Trans))
	}
	fmt.Println()
	fmt.Println("   --------")
	fmt.Println()
}
