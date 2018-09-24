package helps

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

// M - json
type M map[string]interface{}

// JSONToObj -
func JSONToObj(data interface{}, obj interface{}) error {
	if v, ok := data.(io.ReadCloser); ok {
		body, err := ioutil.ReadAll(v)
		if err != nil {
			return err
		}
		return json.Unmarshal(body, obj)
	} else if v, ok := data.([]byte); ok {
		return json.Unmarshal(v, obj)
	} else if v, ok := data.(string); ok {
		return json.Unmarshal([]byte(v), obj)
	}
	return errors.New("error data type")
}
