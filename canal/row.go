package canal

import (
	_ "fmt"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
	_ "time"
)

///header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType(

func initRow(header *protocol.Header, columns []*protocol.Column) Row {
	rrr := Row{}
	rrr.binlog = header.GetLogfileName()
	rrr.logfileOffset = header.GetLogfileOffset()
	rrr.schemaName = header.GetSchemaName()
	rrr.tableName = header.GetTableName()
	rrr.eventType = header.GetEventType()
	rrr.columns = columns
	return rrr
}

type Row struct {
	binlog        string
	logfileOffset int64
	schemaName    string
	tableName     string
	eventType     protocol.EventType
	columns       []*protocol.Column
}

func (l *Row) GetBinlog() string {
	return l.binlog
}

func (l *Row) GetLogfileOffset() int64 {
	return l.logfileOffset
}

func (l *Row) GetSchemaName() string {
	return l.schemaName
}

func (l *Row) GetTableName() string {
	return l.tableName
}

func (l *Row) GeteventType() protocol.EventType {
	return l.eventType
}

func (l *Row) GetColumns() []*protocol.Column {
	return l.columns
}
