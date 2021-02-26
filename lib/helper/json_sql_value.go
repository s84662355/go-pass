package helper

import (
	"database/sql/driver"
	"encoding/json"
	_ "fmt"
	"reflect"
	_ "strconv"
	_ "strings"
	_ "time"
)

type JsonSqlValue struct {
	arr interface{}
}

func (t JsonSqlValue) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *JsonSqlValue) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &t.arr)
	if err != nil {
		return err
	}
	return nil
}

func (t JsonSqlValue) Value() (driver.Value, error) {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (t JsonSqlValue) String() string {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return ""
	}
	return string(data)
}

func (t *JsonSqlValue) Scan(v interface{}) error {
	vv, _ := v.([]byte)
	err := json.Unmarshal(vv, &t.arr)
	if err != nil {
		return err
	}
	return nil
}

func (ttt JsonSqlValue) Get() interface{} {
	return ttt.arr
}

func (ttt JsonSqlValue) Create(mmap interface{}) JsonSqlValue {
	ttttttt := JsonSqlValue{}
	switch mmap.(type) {
	case string:
		sss, _ := mmap.(string)
		json.Unmarshal([]byte(sss), &ttttttt.arr)
		return ttttttt
		break
	}
	t := reflect.ValueOf(mmap)
	switch t.Kind() {
	case reflect.Map:
		arr := make(map[string]interface{})
		//MapRange Key()  Value()  Next()  String() Interface()
		MapRange := t.MapRange()
		//ttttttt.arr[MapRange.Key().String()] = MapRange.Value().Interface()
		for MapRange.Next() {
			arr[MapRange.Key().String()] = MapRange.Value().Interface()
		}
		ttttttt.arr = arr
		break
	case reflect.Slice:
		sliceLen := t.Len()
		arr := make([]interface{}, sliceLen)
		for i := 0; i < sliceLen; i++ {
			arr[i] = t.Index(i).Interface()
		}
		ttttttt.arr = arr
		break
	}
	return ttttttt
}

func (ttttt *JsonSqlValue) Cover(mmap interface{}) bool {
	switch mmap.(type) {
	case string:
		sss := mmap.(string)
		err := json.Unmarshal([]byte(sss), &ttttt.arr)
		if err != nil {
			return false
		}
		return true
		break
	}

	t := reflect.ValueOf(mmap)
	switch t.Kind() {
	case reflect.Map:
		arr := make(map[string]interface{})
		//MapRange Key()  Value()  Next()  String() Interface()
		MapRange := t.MapRange()
		//ttttttt.arr[MapRange.Key().String()] = MapRange.Value().Interface()
		for MapRange.Next() {
			arr[MapRange.Key().String()] = MapRange.Value().Interface()
		}
		ttttt.arr = arr
		return true
	case reflect.Slice:
		sliceLen := t.Len()
		arr := make([]interface{}, sliceLen)
		for i := 0; i < sliceLen; i++ {
			arr[i] = t.Index(i).Interface()
		}

		ttttt.arr = arr
		return true
	}
	return false

}
