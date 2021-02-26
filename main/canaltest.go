// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	cccc "GoPass/canal"
	"GoPass/es/model"
	"GoPass/lib/canal"
	"GoPass/lib/es"
	"GoPass/lib/mysql"
	"fmt"
	//protocol "github.com/withlin/canal-go/protocol"
)

func main() {
	mysql.InitMysqlConnect()
	es.InitEsConnect()
	canal.InitCanal()

	connector := canal.Canal("default")

	fmt.Println(connector)

	cccccc := cccc.InitCanal()

	cccccc.Run(connector, func(ttt cccc.Row) {

		columns := ttt.GetColumns()
		fmt.Println(ttt)
		for _, col := range columns {
			fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
			if col.GetName() == "id" {
				_, eee := model.Article{}.PostDataById(col.GetValue())
				fmt.Println(eee)
			}
		}
		//
	})

	cateccc := cccc.InitCanal()
	connector1 := canal.Canal("cate")

	cateccc.Run(connector1, func(ttt cccc.Row) {

		columns := ttt.GetColumns()
		for _, col := range columns {
			fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
			if col.GetName() == "id" {
				ammmr := model.Article{}.GetByCateId(col.GetValue())
				for _, rrr := range ammmr {
					rrr.Update()
				}
			}

		}
		//
	})

	vvvv := make(chan int)

	<-vvvv

}
