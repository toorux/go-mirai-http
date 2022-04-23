package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/kgysf/go-mirai-http/model"
	"io/ioutil"

	//. "github.com/kgysf/go-mirai-http/pkg/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var paramsValue = "测试参数"
var bodyValue = "POST"

func mockServer(t *testing.T) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		var value string
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("mock server 数据获取失败")
				return
			}
			var tmp map[string]string
			err2 := json.Unmarshal(body, &tmp)
			if err2 != nil {
				t.Errorf("mock server 数据获取失败")
				return
			}
			value = tmp["body"]
		} else {
			value = r.URL.Query().Get("value")
		}
		//fmt.Println("params: ", value)
		if value != paramsValue && value != bodyValue {
			t.Errorf("传入的参数值不符合预设值")
		}
		_, err := fmt.Fprintln(w, "{\"Code\":0,\"Msg\":\""+value+"\",\"Data\":\"\"}")
		if err != nil {
			t.Errorf("启动mock server失败")
		}
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestHttpGet(t *testing.T) {

	server := mockServer(t)
	defer server.Close()

	params := HttpParams{"value": paramsValue}
	header := HttpHeader{"Content-Type": "application/json"}
	result, err := HttpGet[string](server.URL, header, params, true)
	fmt.Println(result, err)
	if err != nil || result.Msg != paramsValue {
		t.Errorf("获取到的参数值不符合预设值")
	}
}

func TestHttpPost(t *testing.T) {

	server := mockServer(t)
	defer server.Close()

	body := map[string]string{"body": bodyValue}
	result, err := HttpPost[model.HttpResult[string]](server.URL, &body, true)
	fmt.Println(result, err)
	if err != nil || result.Msg != bodyValue {
		t.Errorf("获取到的参数值不符合预设值")
	}
}
