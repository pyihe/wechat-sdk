package vars

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewParam(t *testing.T) {
	p := NewKvs()
	data := []byte("{\n  \"code\": \"PARAM_ERROR\",\n  \"message\": \"参数错误\",\n  \"detail\": {\n    \"field\": \"/amount/currency\",\n    \"value\": \"XYZ\",\n    \"issue\": \"Currency code is invalid\",\n    \"location\" :\"body\"\n  }\n}")
	fmt.Println(json.Unmarshal(data, &p))
	fmt.Println(p)
}
