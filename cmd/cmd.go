package cmd

import (
	"commandline-translation/config"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func Run() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s word \n\n\n", os.Args[0])
		flag.PrintDefaults()
		eg := `Examples:
  			$ tr word
  			$ tr world peace
  			$ tr 中文`
		fmt.Println(eg)
	}
	flag.Parse()
	var query string
	if len(os.Args[1:]) == 0 {
		text, err := clipboard.ReadAll()
		if err != nil || text == "" {
			// nothing in clipboard
			flag.Usage()
			return
		}
		fmt.Printf(" \n read clipboard by default: %s\n", text)
		query = text
	} else {
		query = strings.Join(flag.Args(), " ")
	}
	query = url.QueryEscape(query)
	fmt.Println()
	ch := make(chan string)
	go youdaoRequest(query, ch)
	go icibaRequest(query, ch)
	// TODO: Parse Google API
	for i := 0; i < 2; i++ {
		<-ch
	}
}

func youdaoRequest(query string, ch chan<- string) {
	api := config.API
	youdaoUrl := strings.Replace(api.Youdao, "${word}", query, 1)
	resp, err := http.Get(youdaoUrl)
	if err != nil {
		log.Println("error with Youdao API")
		ch <- "fanyi.youdao.com error"
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	YoudaoParse(data)
	ch <- "youdao done"
}

func icibaRequest(query string, ch chan<- string) {
	api := config.API
	icibaUrl := strings.Replace(api.Iciba, "${word}", query, 1)
	resp, err := http.Get(icibaUrl)
	if err != nil {
		log.Println("error with iCIBA API")
		ch <- "iciba.com error"
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	IcibaParse(data)
	ch <- "iciba done"
}
