package main

import (
"fmt"
"net/http"
"io/ioutil"
"github.com/mgutz/ansi"

)

func main(){
	asw := ansi.Color("Infogather By Fajar Firdaus", "white+b:green")
	p := ansi.Color("[======================]\n","red")
	a := ansi.Color("Coder : Fajar Firdaus\n", "green")
	b := ansi.Color("Fb : Fajar Firdaus\n", "green")
	c := ansi.Color("IG : fajar_firdaus_7\n", "green")
	d := ansi.Color("YT : iTech7732\n", "green")
	f := ansi.Color("[======================]\n", "red")
	fmt.Print(asw)
	fmt.Print("\n")
	fmt.Print(p)
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
	fmt.Print(f)
	var input string
	fmt.Print("Masukan URL : ")
	fmt.Scan(&input)

response, error := http.Get("https://api.hackertarget.com/whois/?q=" + input)

if error != nil{
fmt.Println("[!] 404 URL NOT FOUND")
}else{
data, error := ioutil.ReadAll(response.Body)
if error != nil{
fmt.Println("")
}else{
fmt.Println(string(data))
}
}
}
