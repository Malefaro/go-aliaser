package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
	"strings"	
)

const mapfile string = ".go-aliaser-map"
const file string = ".go-aliaser"
var HOME string = os.Getenv("HOME")

func ReadJSONFile() (map[string]string, error) {
	filepath := HOME + "/" + mapfile
	f, _ := ioutil.ReadFile(filepath)

	v := make(map[string]string)
	err := json.Unmarshal(f, &v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return v, nil
}

func WriteJSONFile(m map[string]string) error {
	filepath := HOME + "/" + mapfile
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(filepath, data, 0700)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func AddProfileFile() error {

	data , err := ioutil.ReadFile(HOME+"/.bashrc")
	if err != nil {
		fmt.Println(err)
		return err
	}
	str := string(data)
	const line string =`
if [ -f "$HOME/.go-aliaser" ]; then
	. "$HOME/.go-aliaser"
fi
`
	if strings.Contains(str, line){
		return nil
	}
	
	f, err := os.OpenFile(HOME+"/.bashrc", os.O_APPEND | os.O_WRONLY, 0700) 
	if err != nil {
		fmt.Println(err)
		return err
	}	

	f.WriteString(line)
	f.Close()
	return nil
}

func Writefile(m map[string]string) error {
	var s string
	for k,v := range(m) {
		s += "alias "+k + "=" + "\"" + v + "\"\n"
	}
	err := ioutil.WriteFile(HOME+"/"+file, []byte(s), 0700)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func CreateAlias(str string) error{
	var name string
	var command string
	res := strings.Split(str,"=")
	name = res[0]
	command = res[1]
	exist, err := exists(HOME+"/"+mapfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !exist {
		filename := HOME + "/" + mapfile
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer f.Close()
		err = AddProfileFile() 
		if err != nil {
			fmt.Println(err)
			return err
		}
		m := make(map[string]string)
		m[name] = command
		WriteJSONFile(m)
		Writefile(m)
	} else {
		//filename := HOME + "/" + mapfile
		m, err := ReadJSONFile()
		if err != nil {
			fmt.Println(err)
			return err
		}
		m[name] = command
		WriteJSONFile(m)
		Writefile(m)
	}

	return nil
}

func RemoveAlias(name string) error {
	m, err := ReadJSONFile()
	if err != nil {
			fmt.Println(err)
			return err
	}
	delete(m,name)
	WriteJSONFile(m)
	Writefile(m)
	return nil
}

func AddFunction(str string) error{
	name, fun := strings.SplitN(str, " ", 2)
	
}

// func ExecCmd() {
// 	// path := HOME + "/.profile"
// 	// c := "source " + path
// 	cmd := exec.Command("bash","-c",HOME+"/"+".go-aliaser.sh")
// 	cmd.Stdout = os.Stdout
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage: go-aliaser <create|rm> <name>[=<command>]")
		return 
	}
	switch(os.Args[1]){
	case "create":
		err := CreateAlias(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "rm":
		err := RemoveAlias(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "function":
		err := AddFunction(os.Args[2])
		if err != nil{
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("usage: go-aliaser <create|rm> <name>[=<command>]")
		return 
	}

	// filename := HOME + "/" + mapfile
	// m, err := ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(m)
	// m["lol"] = "check"
	// res, _ := json.Marshal(&m)
	// ioutil.WriteFile(filename, res, 0700)
}