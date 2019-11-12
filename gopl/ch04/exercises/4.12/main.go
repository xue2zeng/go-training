// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.12 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。
// 下载每 个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打 印和命令行输入的检索词相匹配的漫画的URL。
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

// suppress auto-escaping with template HTML
type Comic struct {
	Title      template.HTML
	Num        int
	Img        string
	Transcript template.HTML
	Year       string
	Month      string
	Day        string
	SafeTitle  string `json:safe_title`
	Alt        string
	Link       string
	News       string
}

const URLPrefix = "https://xkcd.com/"

const templ = `*****************************************************
Title: {{.Title}}
URL: {{.Num | getURL}}
{{.Transcript}}
*****************************************************
`

func getURL(i int) string {
	return URLPrefix + fmt.Sprintf("%d", i)
}

// database
var db = make(map[int]Comic)

func main() {
	report := template.Must(template.New("comic").
		Funcs(template.FuncMap{"getURL": getURL}).
		Parse(templ))

	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("=> ")
		cmd, err := input.ReadString('\n')
		if err != nil {
			log.Fatalf("IO error: %v\n", err)
		}
		// delete tailing '\n'
		cmd = cmd[:len(cmd)-1]
		if cmd == "" {
			continue
		}

		if cmd == "q" || cmd == "exit" || cmd == "quit" {
			os.Exit(0)
		}

		index, err := strconv.ParseInt(cmd, 10, 64)
		if err != nil {
			fmt.Println("=> integer required")
			continue
		}

		if comic, ok := db[int(index)]; !ok {
			url := URLPrefix + cmd + "/info.0.json"
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("xkcd: Web IO error: %v\n", err)
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("xkcd: can't get commic of %d, try another.\n", index)
				resp.Body.Close()
				continue
			}
			if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
				fmt.Printf("xkcd: internal error, invalid json.\n")
				resp.Body.Close()
				continue
			}
			// insert into db
			db[int(index)] = comic

			// don't forget close
			resp.Body.Close()
		}

		comic := db[int(index)]

		// print
		//fmt.Printf("Title: %s\n", comic.Title)
		//fmt.Printf("URL: %s\n\n", )
		//fmt.Printf("%s\n", comic.Transcript)

		// give text/template a try
		report.Execute(os.Stdout, comic)
	}
}
