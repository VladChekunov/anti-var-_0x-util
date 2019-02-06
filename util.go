package main

//go build -ldflags "-s -w" util.go 

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"flag"
	"strconv"
)
/*
	Забавно, если кто-то это читает.
	В общем, по-умолчанию, Wishmaster работает только для чтения. Для работы с перезапиьсью вызываем так:
	util -w
	Для поиска новых ширусов на PHP, ибо обычно они за обусфицированны по разному и все паттерны добавить просто нереально, есть флаг h, он найдёт все файлы с не очень стандартным eval, кто вообще его использует? inb4: smartly и прочие индусские поделия, лол.
	util -h
	Эти файлы уже можно будет изучить, написать под них регулярочку и отправить антивирус на сервер.
*/
	var readonly bool;
	var hardmode bool;
	var smallmode bool;

func checkFile(path string, fi os.FileInfo, err error) error {
	//readonly:=true;
	if err != nil {
		return err
	}
	if !!fi.IsDir() {
		return nil
	}

	JSmatched, err := filepath.Match("*.js", fi.Name());
	if err != nil {
		panic(err)
		return err
	}
	PHPmatched, err := filepath.Match("*.php", fi.Name());
	if err != nil {
		panic(err)
		return err
	}

	if JSmatched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var re = regexp.MustCompile(`var _0x`)//1e35
		if(re.MatchString(string(read))){
			if(readonly==true){
				fmt.Printf(path,"\n");
			}else{
				fmt.Printf(path,"\n");
				var re = regexp.MustCompile(`var _0x(.*)if\(n==!!\[\]\){a\(\);}`);
				newContents := re.ReplaceAllString(string(read), ``);
				var re2 = regexp.MustCompile(`var _0x(.*)(\}\}|\)\)\;)`);
				newContents = re2.ReplaceAllString(string(newContents), ``);

				//fmt.Printf(newContents,"\n")
				ioutil.WriteFile(path+".bak", []byte(string(read)), 0)
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile(path, []byte(newContents), 0)
				if err != nil {
					panic(err)
				}
            }
        }
	}else if PHPmatched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var re = regexp.MustCompile(`var _0x`)
		var re2 = regexp.MustCompile(`eval\(String\.fromCharCode\((.*)\);`)
		var re3 = regexp.MustCompile(`<\?php \$(.*)@array_diff_ukey\(@array\(\(string\)\(\$a\)(.*)\?>`)//<\?php \$(.*)\.chr\([0-9]{1,3}\)(.*)@array_diff_ukey\(@array\(\(string\)\(\$a\)(.*)\?>
		var re4 = regexp.MustCompile(`<\?php(.*)\;@include\(\$([a-zA-Z]{1,100})\)\;(.*)>`)

		var re5 = regexp.MustCompile(`(.*)eval\(\$([a-zA-Z]{1,100})(.*)`)
		var re6 = regexp.MustCompile(`(.*)eval\(\$([a-zA-Z]{5,6})\[1\]\(\$([a-zA-Z]{5,6})\[2\]\)\);exit\(\);(.*)`)
		var re7 = regexp.MustCompile(`(.*)(\"\)\;|\;\@|\)\s{|\'\);)eval(.*)`)
		if(re.MatchString(string(read)) || re2.MatchString(string(read)) || re7.MatchString(string(read)) || re6.MatchString(string(read)) || re3.MatchString(string(read)) || re4.MatchString(string(read)) || (hardmode && re5.MatchString(string(read)))){
			if(readonly==true){
				if(smallmode){
					fmt.Printf(strconv.Itoa(len(regexp.MustCompile("\n").FindAllStringIndex(string(read), -1)))," # ",path,"\n");
				}else{
					fmt.Printf(path,"\n");
				}
			}else{
				fmt.Printf(path,"\n");
				var re = regexp.MustCompile(`<script type='text\/javascript'>var _0x(.*)<\/script>`);
				newContents := re.ReplaceAllString(string(read), ``);
				re = regexp.MustCompile(`<script language=javascript>var _0x(.*)<\/script>`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`<script type="text\/javascript">var a1=function\(\)\{(.*)<\/script>`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`<script type="text\/javascript" async> _0x(.*)<\/script>`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`<script language=javascript>eval\(String\.fromCharCode\((.*)<\/script>`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`(.*)eval\(\$([a-zA-Z]{5,6})\[1\]\(\$([a-zA-Z]{5,6})\[2\]\)\);exit\(\);(.*)`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`<\?php(.*)\;@include\(\$([a-zA-Z]{1,100})\)\;(.*)>`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`(.*)(\"\)\;|\;\@|\)\s{|\'\);)eval(.*)`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`(.*)\(base64_decode\(\"(.*)`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`var _0x(.*)(\}\}|\)\)\;)`);
				newContents = re.ReplaceAllString(string(newContents), ``);
				re = regexp.MustCompile(`<\?php echo '<script'; \?>\n\slanguage=javascript>eval\(String\.fromCharCode\((.*)\?>`);
				newContents = re.ReplaceAllString(string(newContents), ``);


				if(smallmode && len(regexp.MustCompile("\n").FindAllStringIndex(string(newContents), -1))<5){
					fmt.Printf(path," ПОПАЛСЯ \n");
					newContents = re5.ReplaceAllString(string(newContents), ``);
				}

				
				

				err = ioutil.WriteFile(path+".bak", []byte(string(read)), 0)
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile(path, []byte(newContents), 0)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	
	return nil
}

func main() {


	fmt.Printf("Wishmaster запущен.\n")

	var writeflag bool;
	flag.BoolVar(&writeflag, "w", false, "режим записи")

	var hardflag bool;
	flag.BoolVar(&hardflag, "h", false, "режим расширенного поиска для php")

	var smallflag bool;
	flag.BoolVar(&smallflag, "s", false, "режим удаления маленьких файлов из режима расширенного поиска")
	
	flag.Parse()

	readonly = true;
	if writeflag==true{
		readonly = false;
	}

	hardmode = false;
	if hardflag==true{
		hardmode = true;
	}

	smallmode = false;
	if smallflag==true{
		smallmode = true;
	}

	if(smallmode==true){
		fmt.Printf("Убираем меленькие с eval.\n")
	}

	if(hardmode==true){
		fmt.Printf("Проверяем все eval.\n")
	}

	if(readonly==true){
		fmt.Printf("Только чтение.\n")
	}else{
		fmt.Printf("Чтение и запись.\n")
	}

	err := filepath.Walk(".", checkFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wishmaster отработал.\n")
}