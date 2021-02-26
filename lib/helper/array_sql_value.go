package helper

import (
	"database/sql/driver"
	_ "fmt"
	//"strconv"
	"encoding/json"
	_ "strings"
	_ "time"
)

// JSONTime format json time field by myself
type ArraySqlValue struct {
	arr []interface{}
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t ArraySqlValue) MarshalJSON() ([]byte, error) {

	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t *ArraySqlValue) UnmarshalJSON(data []byte) error {

	mSlice := make([]interface{}, 0)
	err := json.Unmarshal(data, &mSlice)
	if err != nil {
		return err
	}
	t.arr = mSlice
	return nil
}

// Value insert timestamp into mysql need this function.
func (t ArraySqlValue) Value() (driver.Value, error) {

	data, err := json.Marshal(t.arr)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (t ArraySqlValue) String() string {
	data, err := json.Marshal(t.arr)
	if err != nil {
		return ""
	}
	return string(data)
}

// Scan valueof time.Time
func (t *ArraySqlValue) Scan(v interface{}) error {
	vv, _ := v.([]byte)
	jsonStr := string(vv)

	mSlice := make([]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &mSlice)
	if err != nil {
		return err
	}
	t.arr = mSlice
	return nil

}

func (t ArraySqlValue) Create(slice interface{}) ArraySqlValue {
	val, ok := isSlice(slice)
	tt := ArraySqlValue{}
	if ok {
		sliceLen := val.Len()
		tt.arr = make([]interface{}, sliceLen)
		for i := 0; i < sliceLen; i++ {
			tt.arr[i] = val.Index(i).Interface()
		}
	} else {
		tt.arr = make([]interface{}, 0)
	}
	return tt
}

func (t ArraySqlValue) GetArr() []interface{} {
	return t.arr
}

func (t *ArraySqlValue) Cover(slice interface{}) bool {

	val, ok := isSlice(slice)

	if ok {
		sliceLen := val.Len()
		t.arr = make([]interface{}, sliceLen)
		for i := 0; i < sliceLen; i++ {
			t.arr[i] = val.Index(i).Interface()
		}
		return true

	}

	return false

}

func (t *ArraySqlValue) Add(data interface{}) {
	t.arr = append(t.arr, data)
}
