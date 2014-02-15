package operation

import (
	"code.google.com/p/weed-fs/go/glog"
	"code.google.com/p/weed-fs/go/util"
	"encoding/json"
)

type ClusterStatusResult struct {
	IsLeader bool
	Leader   string
	Peers    []string
}

func ListMasters(server string) ([]string, error) {
	jsonBlob, err := util.Get("http://" + server + "/cluster/status")
	glog.V(2).Info("list masters result :", string(jsonBlob))
	if err != nil {
		return nil, err
	}
	var ret ClusterStatusResult
	err = json.Unmarshal(jsonBlob, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Peers, nil
}
