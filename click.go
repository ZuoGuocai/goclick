/* ************************************* *\
|     Auto Click  Version 1.0             |
|                                         |
|     Author: zuoguocai@126.com           |
|                                         |
|     Last_change: 2017年4月26日          |
|                                         |
|     Apache Licene 2.0                   |
*\ ************************************* */

package main

import (
"github.com/toqueteos/webbrowser"
"fmt"
"time"
"os/exec"
)

func main() {
	 c := exec.Command("cls")
	 c.Run()
     fmt.Println("")
     fmt.Println(("== ^_^ Auto Click the WebPage for JianShu Blog... =="))
	 fmt.Println("")
	 for a := 0; a < 10 ;a++ {
	  webbrowser.Open("http://www.jianshu.com/p/XXXXXXXX")
      fmt.Printf("The %d time click done...\n", a)
      
      time.Sleep(time.Second * 2)  
     
      cmd := exec.Command("taskkill",  "/f", "/t", "/im",  "chrome.exe")
      cmd.Run()
      
    } 
    fmt.Println("") 
    fmt.Println("(*_*) Auto Click WebPage Done...")
}





	


  
   
