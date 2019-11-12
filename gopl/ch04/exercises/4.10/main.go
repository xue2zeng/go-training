// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.10 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、 超过一年。
package main

import (
	"fmt"
	"go-training/gopl/ch04/github"
	"log"
	"os"
	"time"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("Less than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("More than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

//!-

/*
//!+textoutput
$ go run ./go-training/gopl/ch04/exercises/4.10/main.go repo:golang/go is:open json decoder
41 issues:
Less than a month old:
Less than a year old:
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#34647 babolivie encoding/json: fix byte counter increments when using d
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#32779       rsc proposal: encoding/json: memoize strings during decode?
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#28923     mvdan encoding/json: speed up the decoding scanner
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#31701    lr1980 encoding/json: second decode after error impossible
#33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
#31789  mgritter encoding/json: provide a way to limit recursion depth
#33714    flimzy proposal: encoding/json: Opt-in for true streaming supp
#30701 LouAdrien encoding/json: ignore tag "-" not working on embedded s
More than a year old:
#11046     kurin encoding/json: Decoder internally buffers full input
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#16212 josharian encoding/json: do all reflect work before decoding
#5901        rsc encoding/json: allow override type marshaling
#22480     blixt proposal: encoding/json: add omitnil option
#28189     adnsv encoding/json: confusing errors when unmarshaling custo
#14750 cyberphon encoding/json: parser ignores the case of member names
#7872  extempora encoding/json: Encoder internally buffers full output
#28143    arp242 proposal: encoding/json: add "readonly" tag
#22752  buyology proposal: encoding/json: add access to the underlying d
#27179  lavalamp encoding/json: no way to preserve the order of map keys
#20754       rsc encoding/xml: unmarshal only processes first XML elemen
#20528  jvshahid net/http: connection reuse does not work happily with n
#21092  trotha01 encoding/json: unmarshal into slice reuses element data
#22816 ganelon13 encoding/json: include field name in unmarshal error me
//!-textoutput
*/
