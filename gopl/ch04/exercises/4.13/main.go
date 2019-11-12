// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.13 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对 应的海报。
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const URL = "http://www.omdbapi.com/?t="

type Movie struct {
	Response string
	Error    string
	Poster   string
}

//go run ex4.13/ex4.13.go mission impossible III
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: need to input keywords")
		os.Exit(1)
	}
	keywords := url.QueryEscape(strings.Join(os.Args[1:], " "))
	resp, err := http.Get(URL + keywords)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	if movie.Response != "True" {
		fmt.Println(movie.Error)
		os.Exit(4)
	}
	if movie.Poster == "" {
		fmt.Println("No poster")
		os.Exit(5)
	}
	fmt.Println("Downloading file...")
	rawURL := movie.Poster
	fileURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1] // change the number to accommodate changes to the url.Path position
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp2, err := check.Get(rawURL) // add a filter to check redirect

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp2.Body.Close()
	fmt.Println(resp2.Status)

	size, err := io.Copy(file, resp2.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s with %v bytes downloaded\n", fileName, size)
}
