package transfer_code

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/hnwgij1/file"
	"os"
	"path/filepath"
	"regexp"
)

func GbkToUtf8(fileName string) {
	content := file.ReadFile(fileName)
	enc := mahonia.NewDecoder("gbk")
	content = enc.ConvertString(content)
	file.WriteFile(fileName, content)
	fmt.Println("转换完成!")
}

func Utf8ToGbk(fileName string) {
	content := file.ReadFile(fileName)
	enc := mahonia.NewEncoder("gbk")
	content = enc.ConvertString(content)
	file.WriteFile(fileName, content)
	fmt.Println("转换完成!")
}

func DirectConvert(directName string, convertFunc func(string)){
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
