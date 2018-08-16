package config

import (
	"testing"
)

func Test_GetConfigItem(t *testing.T){

	config := New(ConfigName)
	if nil == &config {
		t.Error("UNPASS: failed to init a config.")
	}

	parsedStruct := config.GetConfigItem("producer")
	if nil == parsedStruct {
		t.Error("UNPASS: failed to get a config item.")
	}

	pasedItem := config.GetConfigItem("producer.policy")
	if nil != pasedItem && "time" == pasedItem.(string) {
		t.Log("PASS: success to get a exist item in config.")
	}else{
		t.Error("UNPASS: failed to get a exist item in config.")
	}

	pasedItem = config.GetConfigItem("producer.pol")
	if nil != pasedItem {
		t.Error("UNPASS: failed to test getting a not exist item.")
	}else{
		t.Log("PASS: success to test getting a not exist item.")
	}
}