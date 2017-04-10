package consumenodeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"net/http"
)

func (this consumeNodeApi) KillOp(
	req model.KillOpReq,
) (
	err error,
) {

	reqBytes, err := json.Marshal(req)
	if nil != err {
		return
	}

	httpReq, err := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%v/ops/kills", "localhost:42224"),
		bytes.NewBuffer(reqBytes),
	)
	if nil != err {
		return
	}

	_, err = this.httpClient.Do(httpReq)
	return

}
