package nsq

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
)

func init() {
	content, err := ioutil.ReadFile("env/nsq.json")
	if err != nil {
		//log.Fatal(err)
		return
	}

	ccc := config{}
	json.Unmarshal(content, &ccc)

	if ccc.DefaultConsumer != "" {
		Config.DefaultConsumer = ccc.DefaultConsumer
	}

	if ccc.DefaultProducer != "" {
		Config.DefaultProducer = ccc.DefaultProducer
	}

	for key, v := range ccc.NsqConsumer {
		Config.NsqConsumer[key] = v
	}

	for key, v := range ccc.NsqProducer {
		Config.NsqProducer[key] = v
	}

}
