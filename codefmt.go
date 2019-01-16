package main

import ( // {{{
	"bufio"
	"flag"
	"fmt"
	//	"io"
	"log"
	"os"
	//	"regexp"
	"strings"
	"unsafe"
) // }}}

func openWriteFile(fn string) *os.File { // {{{
	// ファイル開く
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("openWriteFile")
		log.Fatal(err)
	}
	return fp
} // }}}

func openReadFile(fn string) *os.File { // {{{

	// ファイルを読み込みモードでオープン
	fp, err := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println("openReadFile")
		log.Fatal(err)
	}
	return fp
} // }}}

// str, err = Readln(reader)
func Readln(r *bufio.Reader) (string, error) { // {{{
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
} // }}}

func bytesToString(b []byte) string { // {{{
	return *(*string)(unsafe.Pointer(&b))
} //}}}

func stringToBytes(s string) []byte { //{{{
	return *(*[]byte)(unsafe.Pointer(&s))
} // }}}

func openAndReadFile(inputfile string, input_lines *[]string) { //{{{
	fp, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("readln")
		log.Fatalln(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		*input_lines = append(*input_lines, scanner.Text())
	}
} // }}}

func replace(line string, output_lines *[]string) {

	//rep := []string{"/*", "*/", "//", "\""}
	//n_rep := len(rep)

	//fmt.Println(line)
	// 編集する
	// 置換するかどうかチェック
	var flag bool
	flag = false
	// コメントやダブルクォート、切り捨て除算があったら置換しない
	//flag = flag || regexp.MustCompile(`[ ]*/\*[ ]*`).Match(stringToBytes(line))
	//flag = flag || regexp.MustCompile(`[ ]*\*/[ ]*`).Match(stringToBytes(line))
	//flag = flag || regexp.MustCompile(`"`).Match(stringToBytes(line))
	//flag = flag || regexp.MustCompile(`//`).Match(stringToBytes(line))
	// 短い正規表現はこっちが早い
	flag = flag || strings.Contains(line, "/*")
	flag = flag || strings.Contains(line, "*/")
	flag = flag || strings.Contains(line, "//")
	flag = flag || strings.Contains(line, "\"")
	// 思ったより速度早くなかった書き方
	// flag = strings.Contains(line, "/*") || strings.Contains(line, "*/") || strings.Contains(line, "//") || strings.Contains(line, "\"")
	//for j := 0; j < len(rep); j++ {
	/* j = 0
	for j < n_rep && flag != true {
		flag = strings.Contains(line, rep[j])
		j = j + 1
	} */

	if flag == false {
		for j := 0; j < len(line); j++ {
			//output_line = fmt.Sprint(i) + ":" + fmt.Sprint(len(line)) + line + "\n"
			//fmt.Println(line)
			//fmt.Println(output_line)
			//output_line = regexp.MustCompile(`[ ]*/[ ]*`).ReplaceAllString(line, " / ")
			// 短い正規表現はこっちが早い
			line = strings.Replace(line, `[ ]*/[ ]*`, " / ", -1)
		}
	}
	line = line + "\n"
	*output_lines = append(*output_lines, line)
}

func main() {
	// コマンドライン引数の読み込み
	flag.Parse()
	//args := flag.Args()
	//fmt.Println("input args :", args)

	input_lines := []string{}
	openAndReadFile(flag.Arg(0), &input_lines)

	fp := openWriteFile(flag.Arg(1))
	defer fp.Close()

	//var output_line string
	var output_lines []string

	//for _, line := range input_lines {
	num := len(input_lines)

	for i := 0; i < num; i++ {
		replace(input_lines[i], &output_lines)
	}
	//fmt.Println(output_lines)

	var output string
	for _, l := range output_lines {
		output = output + l
	}

	var writer *bufio.Writer
	writer = bufio.NewWriter(fp)
	writer.WriteString(output)
	writer.Flush()
}
