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

type MapSqlValue struct {
	arr map[string]interface{}
}

func (t MapSqlValue) MarshalJSON() ([]byte, error) {

	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *MapSqlValue) UnmarshalJSON(data []byte) error {
	mmap := make(map[string]interface{})
	err := json.Unmarshal(data, &mmap)
	if err != nil {
		return err
	}
	t.arr = mmap
	return nil
}

func (t MapSqlValue) Value() (driver.Value, error) {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (t MapSqlValue) String() string {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return ""
	}
	return string(data)
}

func (t *MapSqlValue) Scan(v interface{}) error {
	vv, _ := v.([]byte)
	mmap := make(map[string]interface{})
	err := json.Unmarshal(vv, &mmap)
	if err != nil {
		return err
	}
	t.arr = mmap
	return nil
}

func (ttt MapSqlValue) Create(mmap interface{}) MapSqlValue {
	ttttttt := MapSqlValue{}
	ttttttt.arr = make(map[string]interface{})
	t := reflect.ValueOf(mmap)
	switch t.Kind() {
	case reflect.Map:
		//MapRange Key()  Value()  Next()  String() Interface()
		MapRange := t.MapRange()
		//ttttttt.arr[MapRange.Key().String()] = MapRange.Value().Interface()
		for MapRange.Next() {
			ttttttt.arr[MapRange.Key().String()] = MapRange.Value().Interface()
		}

		break

	}
	return ttttttt

}

func (t MapSqlValue) GetArr() map[string]interface{} {
	return t.arr
}

func (ttttt *MapSqlValue) Cover(mmap interface{}) bool {

	t := reflect.ValueOf(mmap)
	switch t.Kind() {
	case reflect.Map:
		ttttt.arr = make(map[string]interface{})

		MapRange := t.MapRange()

		for MapRange.Next() {
			ttttt.arr[MapRange.Key().String()] = MapRange.Value().Interface()
		}

		return true

	}
	return false

}

func (t *MapSqlValue) Add(key string, data interface{}) {
	t.arr[key] = data
}
