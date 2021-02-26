package process
import (
    "fmt"
    "GoPass/lib/helper"
    
)


type processComposition struct{
    name string
    process ProcessInterface
    totalNumber uint16
}


func (l *processComposition) run(){

	  ch := make(chan int, l.totalNumber)
	  for true{
	  	  ch <- 1
	  	  fmt.Printf("开启%s功能",l.name)

	  	  go func(){
            defer helper.Recover() 
            defer func(){
                <-ch    
                fmt.Printf("关闭%s功能",l.name)
            }()
             

            l.process.Run()
	  	  }()
	  }

}