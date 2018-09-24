package helps

import (
	"encoding/json"
	"testing"
)

type data struct {
	Name string
	Age  int
}

var tempData = data{Name: "dog", Age: 10}
var b, err = json.Marshal(tempData)

func TestString(t *testing.T) {
	var d data
	JsonToObj(`{"Name":"dog", "Age":10}`, &d)
	if d.Name != "dog" && d.Age != 10 {
		t.Errorf(`TestString`)
	}
}

func TestByte(t *testing.T) {
	var d data
	JsonToObj([]byte(`{"Name":"dog", "Age":10}`), &d)
	if d.Name != "dog" && d.Age != 10 {
		t.Errorf(`TestByte`)
	}
}

func TestByteToMap(t *testing.T) {
	var d map[string]interface{}
	JsonToObj([]byte(`{"name":"dog", "age":10}`), &d)
	if d["name"] != "dog" && d["age"] != 10 {
		t.Errorf(`TestByteToMap`)
	}
}

func TestStringToMap(t *testing.T) {
	var d map[string]interface{}
	JsonToObj(`{"name":"dog", "age":10}`, &d)
	if d["name"] != "dog" && d["age"] != 10 {
		t.Errorf(`TestStringToMap`)
	}
}
