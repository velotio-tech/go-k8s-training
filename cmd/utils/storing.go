package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var allvalue string

func PreLoadValues() {
	DecryptMyData()
	data, _ := ioutil.ReadFile("data.txt")
	Fdata := string(data)

	if len(data) == 0 {
		return
	}

	arr := strings.Split(Fdata, "~")

	//fmt.Println(len(arr), "all data len wiht data", arr)

	for _, val := range arr {

		if len(val) < 1 {
			break
		}

		//fmt.Println("values jhkfdk")
		arr2 := strings.Split(val, "#")

		key := arr2[0]

		var obj Entry

		My_map[key] = obj
		// for _, rt := range arr2 {
		// 	fmt.Println(rt)
		// }

		//fmt.Println(key, len(arr2))
		var i int
		for i = 1; i < len(arr2)-1; i += 2 {

			AppendValueToMap(key, arr2[i], arr2[i+1])
			//fmt.Println(key, arr2[i], arr2[i+1])
		}

	}

	//fmt.Println(My_map)
	// e := os.Remove("data.txt")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	//SaveDataToFileHelper()
}

func SaveDataToFileHelper() {
	//data, _ := ioutil.ReadFile("data.txt")
	filename, err := os.Create("data.txt")
	//_, err2 := filename.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
		filename.Close()
		return
	}

	for key, value := range My_map {

		allvalue = ""
		// allvalue = "Username: " + key // # for split
		var i int
		for i = 0; i < len(value.Time); i++ {
			allvalue = allvalue + value.Time[i] + "#" + value.Info[i] + "#"
		}

		allvalue2 := key + "#" + allvalue + "~"

		_, err := filename.Write([]byte(allvalue2))
		if err != nil {
			log.Fatal(err)
			filename.Close()
			return
		}

	}

	EncryptMyData()

}

func EncryptMyData() {
	fdata, _ := ioutil.ReadFile("data.txt")
	filedata := string(fdata)
	e := os.Remove("data.txt")
	if e != nil {
		log.Fatal(e)
	}
	filename, _ := os.Create("data.txt")

	encyFiledata := Encrypt(filedata)
	_, err := filename.Write([]byte(encyFiledata))
	if err != nil {
		log.Fatal(err)
		filename.Close()
		return
	}
}

func DecryptMyData() {
	fdata, _ := ioutil.ReadFile("data.txt")
	filedata := string(fdata)
	e := os.Remove("data.txt")
	if e != nil {
		log.Fatal(e)
	}
	filename, _ := os.Create("data.txt")

	encyFiledata := Decrypt(filedata)
	_, err := filename.Write([]byte(encyFiledata))
	if err != nil {
		log.Fatal(err)
		filename.Close()
		return
	}
}
