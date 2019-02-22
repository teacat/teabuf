package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		start(c.Args().Get(0), c.Args().Get(1))
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//
func start(in, out string) {
	var isDir bool
	//
	if info, err := os.Stat(in); err == nil && info.IsDir() {
		isDir = true
	}
	//
	if isDir {
		files, err := ioutil.ReadDir(in)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".teabuf3" {
				parse(in + "/" + f.Name())
			}
		}
		return
	}
	parse(in)
}

type typeAlias struct {
	alias string
	typ   string
}

//
func parse(path string) {
	//b, err := ioutil.ReadFile(path)
	//if err != nil {
	//	panic(err)
	//}
	//s := string(b)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//var (
//	typeRegExp = `type\s*?([A-Z]\w*?)\s*?(string|bool|float|double|bytes|(?:s|u|)int(?:32|64|)|(?:s|)fixed(?:32|64|))`
//)

//func parseTypeAliases(content string) []*typeAlias {
//	pat := regexp.MustCompile(`(\w+)=(\w+)`)
//	matches := pat.FindAllStringSubmatch(data, -1)
//}
