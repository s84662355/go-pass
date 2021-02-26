package process

import (
      "fmt"
)

type ProcessInterface interface{
	Run() 
}

var Register = new(register)

type register struct{
	processContainer  map[string]*processComposition

}

func (l *register ) init(){
	 l.processContainer  = make(map[string]*processComposition)
	  fmt.Printf("11111111运84398438489389s")

}

func (l *register ) Add(name string,process ProcessInterface,totalNumber uint16){
	  processContainer := new(processComposition)

	  processContainer.name = name
	  processContainer.process = process
	  processContainer.totalNumber = totalNumber

	  l.processContainer[name] = processContainer
}

func (l *register ) Start(){
	for k, v := range l.processContainer {
          fmt.Printf("运行%s功能",k)
          go v.run()
    }
    
}