package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var file_name string
	fmt.Println("Please type in the file name :")
	fmt.Scanf("%s", &file_name)

	generateFile(file_name, "Controller")
	generateFile(file_name, "Service")
	generateFile(file_name, "Repository")
	generateModel(file_name)
}

func generateFile(file_name, file_type string) {
	// 讀模板檔，存在content
	template, read_err := ioutil.ReadFile("template/" + file_type + "Template.txt")
	if read_err != nil {
		log.Fatal(read_err)
	}

	// 在模板中將檔名取代進去
	file_content := strings.Replace(string(template), "{file_name}", file_name, -1)

	// 設定輸出的檔案路徑並寫檔
	file := filepath.Join("Output", file_name+file_type+".php")
	os.WriteFile(file, []byte(file_content), 0755)
}

func generateModel(file_name string) {
	// 讀模板檔，存在content
	template, read_err := ioutil.ReadFile("template/modelTemplate.txt")
	if read_err != nil {
		log.Fatal(read_err)
	}

	// 在模板中將檔名取代進去
	file_content := strings.Replace(string(template), "{file_name}", file_name, -1)

	// 設定輸出的檔案路徑並寫檔
	file := filepath.Join("Output", file_name+".php")
	os.WriteFile(file, []byte(file_content), 0755)
}
