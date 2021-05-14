package function

import (
"context"
"fmt"
"os"
"strings"
"encoding/json"

cloudevents "github.com/cloudevents/sdk-go/v2"
)

/*
Example Event of interest:
☁️  cloudevents.Event
Validation: valid
Context Attributes,
  specversion: 1.0
  type: dev.knative.apiserver.resource.add
  source: https://172.30.0.1:443
  subject: /apis/v1/namespaces/agcoolserve3/events/web-tlznr-1-deployment-5c94c4885c.167ee5c8cfe38f46
  id: ec19161a-9614-44c9-84ef-b18821bbbd6c
  time: 2021-05-14T09:46:06.553418793Z
  datacontenttype: application/json
Extensions,
  kind: Event
  knativearrivaltime: 2021-05-14T09:46:06.553696617Z
  name: web-tlznr-1-deployment-5c94c4885c.167ee5c8cfe38f46
  namespace: agcoolserve3
Data,
  {
    "apiVersion": "v1",
    "count": 1,
    "eventTime": null,
    "firstTimestamp": "2021-05-14T09:46:06Z",
    "involvedObject": {
      "apiVersion": "apps/v1",
      "kind": "ReplicaSet",
      "name": "web-tlznr-1-deployment-5c94c4885c",
      "namespace": "agcoolserve3",
      "resourceVersion": "5738486",
      "uid": "901ae0fa-222d-4c11-89f5-b7d4b0a2a4e6"
    },
    "kind": "Event",
    "lastTimestamp": "2021-05-14T09:46:06Z",
    "message": "Created pod: web-tlznr-1-deployment-5c94c4885c-n5d4d",
    "metadata": {
      "creationTimestamp": "2021-05-14T09:46:06Z",
      "managedFields": [
        {
          "apiVersion": "v1",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:count": {},
            "f:firstTimestamp": {},
            "f:involvedObject": {
              "f:apiVersion": {},
              "f:kind": {},
              "f:name": {},
              "f:namespace": {},
              "f:resourceVersion": {},
              "f:uid": {}
            },
            "f:lastTimestamp": {},
            "f:message": {},
            "f:reason": {},
            "f:source": {
              "f:component": {}
            },
            "f:type": {}
          },
          "manager": "kube-controller-manager",
          "operation": "Update",
          "time": "2021-05-14T09:46:06Z"
        }
      ],
      "name": "web-tlznr-1-deployment-5c94c4885c.167ee5c8cfe38f46",
      "namespace": "agcoolserve3",
      "resourceVersion": "5738505",
      "selfLink": "/api/v1/namespaces/agcoolserve3/events/web-tlznr-1-deployment-5c94c4885c.167ee5c8cfe38f46",
      "uid": "3fc0b759-82d1-48e6-8dc5-a6c34fdb6e42"
    },
    "reason": "SuccessfulCreate",
    "reportingComponent": "",
    "reportingInstance": "",
    "source": {
      "component": "replicaset-controller"
    },
    "type": "Normal"
  }
*/


func Handle(ctx context.Context, event cloudevents.Event) (resp *cloudevents.Event, err error) {

var search string
var item string

// check environment for config

item = os.Getenv("ITEM")

search = os.Getenv("SEARCH")

//  fmt.Printf("Found env SEARCH %s\n", search);
//  fmt.Printf("Found env ITEM %s\n", item);

// we have filter the event on subject to make sure we are talking about the relevant item
// The subject typically describes the pod/deployment name source of the event and we allow 
// partial matches since the name often includes revision specific information which we can't guess
if strings.Contains(event.Subject(), item) {
	// now marshal the data into json and extract the message field and compare that against 
	// our search value again using a partial match
	var dat map[string]interface{};
	if err = json.Unmarshal(event.Data(), &dat); err != nil {
		fmt.Printf("invalid event data %v", err)
	} else {
//		fmt.Printf("Extracted Message: %s\n", dat["message"])

		// so does the message contain our text?
		if strings.Contains(dat["message"].(string), search) {
			fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
			
			// inject a new event back into the broker to notify our components
			response := cloudevents.NewEvent()
			response.SetID("wakeup")
			response.SetSource("web-coolstore-faas")
			response.SetType("web-wakeup")
	
			// Validate the response
			resp = &response
			if err = resp.Validate(); err != nil {
			fmt.Printf("invalid event created. %v", err)
			}
			return
		}
	}
}
// leave this so we know we are recieving events
fmt.Printf(".");
return
}

