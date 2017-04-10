package consumenodeapi

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/opspec-io/sdk-golang/model"
	"net/url"
	"strings"
)

func (this consumeNodeApi) GetEventStream(
	req *model.GetEventStreamReq,
) (eventStream chan model.Event, err error) {

	eventStream = make(chan model.Event, 1000)

	// construct query params
	queryParams := []string{}
	if filter := req.Filter; nil != filter {
		var filterBytes []byte
		filterBytes, err = json.Marshal(filter)
		if nil != err {
			return
		}
		queryParams = append(
			queryParams,
			fmt.Sprintf("filter=%v", url.QueryEscape(string(filterBytes))),
		)
	}

	c, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("ws://%v/events/stream?%v", "localhost:42224", strings.Join(queryParams, "&")),
		nil,
	)
	if err != nil {
		return
	}

	go func() {
		// ensure web socket closed on exit
		defer c.Close()

		// ensure channel closed on exit
		defer close(eventStream)

		for {

			_, bytes, err := c.ReadMessage()
			if nil != err {
				return
			}

			var event model.Event
			err = json.Unmarshal(bytes, &event)
			if nil != err {
				return
			}
			eventStream <- event

		}
	}()

	return

}
