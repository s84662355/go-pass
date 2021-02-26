package canal

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
	"time"
)

type CanalListener struct {
	ccc chan int
}

func InitCanal() CanalListener {
	can := CanalListener{}
	can.ccc = make(chan int, 10)
	return can
}

func (l *CanalListener) Run(connector *client.SimpleCanalConnector, handler func(rr Row)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("  Recovered in PanicError", r)
			}
			fmt.Println("canal 出错")

		}()
		for {
			message, err := connector.Get(100, nil, nil)
			if err != nil {
			}
			batchId := message.Id
			if batchId == -1 || len(message.Entries) <= 0 {
				time.Sleep(800 * time.Millisecond)
				fmt.Println(connector.Filter + "===没有数据了===")
				continue
			}
			///connector.Ack(batchId)
			l.goRun(message.Entries, handler)
			//connector.Ack(batchId)
		}
	}()
}

func (l *CanalListener) goRun(entrys []protocol.Entry, handler func(rr Row)) {
	for _, entry := range entrys {
		fmt.Println("-----处理")
		l.ccc <- 1
		go l.runn(entry, handler)
		fmt.Println("-----处理后")
	}
}

func (l *CanalListener) runn(entry protocol.Entry, handler func(rr Row)) {
	defer func() {
		<-l.ccc
		if r := recover(); r != nil {
			fmt.Println("Recovered in PanicError", r)
		}

	}()
	if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
		return
	}
	rowChange := new(protocol.RowChange)

	//err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
	proto.Unmarshal(entry.GetStoreValue(), rowChange)
	//checkError(err)
	if rowChange != nil {
		eventType := rowChange.GetEventType()
		header := entry.GetHeader()
		fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

		for _, rowData := range rowChange.GetRowDatas() {
			///printColumn(rowData.GetBeforeColumns())

			if eventType == protocol.EventType_DELETE {
				handler(initRow(header, rowData.GetBeforeColumns()))
			} else if eventType == protocol.EventType_INSERT {
				handler(initRow(header, rowData.GetAfterColumns()))
			} else {
				handler(initRow(header, rowData.GetAfterColumns()))
				//fmt.Println("-------> before")
				//printColumn(rowData.GetBeforeColumns())
				///	fmt.Println("-------> after")

			}

		}
	}

}
