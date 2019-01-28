package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func checkFile(path string, fi os.FileInfo, err error) error {
	readonly:=false;
	
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
			if(readonly){
				fmt.Printf(path,"\n");
			}else{
				fmt.Printf(path,"\n");
				var re = regexp.MustCompile(`var _0x(.*)if\(n==!!\[\]\){a\(\);}`);
				newContents := re.ReplaceAllString(string(read), ``);
				fmt.Printf(newContents,"\n")
				ioutil.WriteFile(path+".bak", []byte(string(read)), 0)
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
		if(re.MatchString(string(read))){
			if(readonly){
				fmt.Printf(path,"\n");
			}else{			
				var re = regexp.MustCompile(`<script type='text\/javascript'>var _0x(.*)<\/script>`);
				newContents := re.ReplaceAllString(string(read), ``);
				ioutil.WriteFile(path+".bak", []byte(string(read)), 0)
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
	err := filepath.Walk(".", checkFile)
	if err != nil {
		panic(err)
	}
}