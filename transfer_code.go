package transfer_code 

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func writeFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()        //关闭文件
	file.WriteString(content) //写入文件
}

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileSize,_ := file.Seek(0, io.SeekEnd)
	content := make([]byte, fileSize)
	file.Seek(0, io.SeekStart)
	sumReadNum := 0
	for {
		readNum, err := file.Read(content)
		if err != nil && err != io.EOF {
			panic(err) //有错误抛出异常
		}
		if 0 == readNum {
			break //当读取完毕时候退出循环
		}
		sumReadNum += readNum
	}
	return string(content[:sumReadNum])
}

func gbk_to_utf8(fileName string) {
	content := readFile(fileName)
	enc := mahonia.NewDecoder("gbk")
	content = enc.ConvertString(content)
	writeFile(fileName, content)
	fmt.Println("转换完成!")
}


func utf8_to_gbk(fileName string) {
	content := readFile(fileName)
	enc := mahonia.NewEncoder("gbk")
	content = enc.ConvertString(content)
	writeFile(fileName, content)
	fmt.Println("转换完成!")
}

func direct_convert(directName string, convertFunc func(string)){
	    err := filepath.Walk(directName, func(path string, f os.FileInfo, err error) error {
        if f == nil {
			return err
		}
        fileName := f.Name()
        if f.IsDir() {
        	return nil
		}
        reg := regexp.MustCompile("\\.[ch]")
		res := reg.FindAllString(fileName, -1)
		if res == nil {
			return nil
		}
			println(path)
			convertFunc(path)
			return nil
        })
    if err != nil {
        fmt.Printf("filepath.Walk() returned %v\n", err)
    }
}



//func main() {
//	direct_gbk_to_utf8(fileName, gbk_to_utf8)
	//direct_convert(fileName, utf8_to_gbk)
//}
