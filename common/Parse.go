package common

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	Domain     string       //域名
	Pwdpath, _ = os.Getwd() //当前工作路径
)

func Parse() {
	ParseUrl()       //解析URL
	ParseThreadNum() //解析线程数量
	ParseOut()       //解析输出文件
	ParseExtenlen()  //解析扩展名
	ParseLine()      //解析字典行数
	ParseTimeout()
	OutTerm() //显示终端配置
}

func ParseTimeout() {
	if Timeout <= 0 {
		Colerr.Printf("The number of timeout must be greater than 0\n")
		Colban.Println("use --help or -h for help")
		os.Exit(1)
	}
}

func ParseLine() {
	file, err := os.Open(WordList)
	if err != nil {
		Colerr.Printf("Open wordlist file fail!")
		os.Exit(1)
	}
	defer file.Close()
	count := 0
	scan := bufio.NewScanner(file)
	elen := int(ParseExtenlen())
	for scan.Scan() {
		if strings.Contains(scan.Text(), "%EXT%") {
			count += elen
		} else {
			count++
		}
	}
	FileLine = uint32(count)
}

func OutTerm() {
	yellow := color.New(color.FgYellow)
	blue := color.New(color.FgBlue)
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	now := time.Now()
	yellow.Printf("Extensions: ")
	blue.Printf("%s", Extention)
	red.Printf(" | ")
	yellow.Printf("Threads: ")
	blue.Printf(" %d ", ThreadNum)
	red.Printf(" | ")
	yellow.Printf("Wordlist size: ")
	blue.Printf("%d\n\n", FileLine)
	green.Printf("Output File: ")
	fmt.Printf("%s\n\n", OutFile)
	yellow.Printf("Target: ")
	blue.Printf("%s\n\n", Url)
	yellow.Printf("[%v:%v:%v] Starting:\n\n", now.Hour(), now.Minute(), now.Second())
}

func ParseExtenlen() uint8 {
	ExtGroup = strings.Split(Extention, ",")
	return uint8(len(ExtGroup))
}

func ParseThreadNum() {
	if ThreadNum < 1 {
		Colerr.Println("The number of threads must be greater than 0")
		Colban.Println("use --help or -h for help")
		os.Exit(1)
	}
}

func ParseUrlFile() {
	file, err := os.Open(UrlFile)
	if err != nil {
		Colerr.Printf("Open %s Fail!", UrlFile)
		os.Exit(1)
	}
	UrlNum = 0
	r := bufio.NewScanner(file)

	for r.Scan() {
		if r.Text() != "" {
			UrlNum++
			Urls = append(Urls, r.Text())
		}
	}

}

func ParseUrl() {
	//解析URL和相应的域名或IP
	if Url == "" && UrlFile == "" {
		Colerr.Println("mgtj_dirsearch.exe -u http://127.0.0.1:9090")
		Colban.Println("use ./dict/dict.txt")
		os.Exit(1)
	} else if Url != "" {
		Urls = strings.Split(Url, ",")
		for key, value := range Urls {
			UrlNum++
			firin := strings.Index(value, "://")
			if !strings.HasSuffix(value, "/") {
				value += "/"
			}
			if firin == -1 {
				Domain = value
				value = "http://" + value
			} else {
				Domain = value[firin+3:]
			}
			Urls[key] = value
		}
	} else {
		ParseUrlFile()
	}
}

func ParseOut() {
	now := time.Now()
	os.Mkdir("report", 0755)
	if OutFile == "" && UrlFile == "" {
		Domain = strings.ReplaceAll(Domain, ":", "_")
		Domain = Domain[:strings.Index(Domain, "/")]
		OutFile = fmt.Sprintf("%s/_%d_%02d_%02d_%02d-%02d-%02d.result.txt", Domain, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		OutFile = filepath.Join(Pwdpath, "report", OutFile)
		os.Mkdir("report/"+Domain, 0755) //创建对应文件夹
	} else if OutFile == "" {
		OutFile = filepath.Join(Pwdpath, "report", fmt.Sprintf("%s.result.txt", UrlFile))
	}
}
