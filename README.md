## Laravel Template Generator
---
### 介紹
開發時常需要針對某功能一次建立Controller、Service、Repository、Model，但原生php artisan並沒有針對Service、Repository做設計，剛好最近在練習Golang變自己十分土炮的寫個產生程式。

---
### Note
此僅為自己練習Golang時為了有個具體目標而做的練習，相同問題及功能已有許多人開發類似更齊全功能並發布在composer上，建議使用這些較完善的產品。

---
### 準備
於Template資料夾中，依照個人習慣修改template，包含其固定路徑、繼承、namespace等<br>
1. controllerTemplate.txt
2. serviceTemplate.txt
3. repositoryTemplate.txt
4. modelTemplate.txt

---
### 執行
```
$ go run makeFile.go
$ {type in your file name}
```
或者自行修改完template後編譯該makeFile.go
<br>
```
$ go build .
$ ./LaravelTemplateGenerator.exe
$ {type in your file name}
```
<br>
於Output資料夾中找到產出的檔案