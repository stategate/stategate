package elastic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type Service struct {
	client *elasticsearch.Client
}

func NewService(client *elasticsearch.Client) (*Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := esapi.InfoRequest{
		Pretty: true,
		Human:  true,
	}
	_, err := req.Do(ctx, client)
	if err != nil {
		return nil, err
	}
	return &Service{client: client}, nil
}

func (s Service) SaveEvent(ctx context.Context, event *eventgate.EventDetail) error {
	req := esapi.IndexRequest{
		Index:   event.GetChannel(),
		Body:    strings.NewReader(protojson.Format(event)),
		Refresh: "true",
	}
	resp, err := req.Do(ctx, s.client)
	if err != nil {
		return err
	}
	if resp.HasWarnings() {
		return errors.New(strings.Join(resp.Warnings(), ","))
	}
	return nil
}

func (s Service) GetEvents(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.EventDetails, error) {
	if opts.Max == nil {
		opts.Max = timestamppb.New(time.Now().Add(5 * time.Minute))
	}
	if opts.Limit == 0 {
		opts.Limit = 1000
	}
	//	query := fmt.Sprintf(`
	//{
	//  "query": {
	//	"match": {
	//		"channel": "%s"
	//	},
	//    "range": {
	//      "time": {
	//        "gte": "%s",
	//		"lte": "%s"
	//      },
	//    }
	//  }
	//}`, opts.GetChannel(), opts.GetMin().AsTime().String(), opts.GetMax().AsTime().String())

	query := fmt.Sprintf(`
{ 
  "from" : %v, 
  "size" : %v,
  "query": { 
    "bool": { 
      "filter": [ 
        { 
          "range": { 
            "time": {
              "gte": %v,
  			  "lte": %v
            } 
          } 
        }, 
        { 
          "bool": { 
            "must": [ 
              { 
                "match": { 
                  "channel": "%s" 
                } 
              }
            ] 
          } 
        } 
      ] 
    } 
  } 
}
`, opts.GetOffset(), opts.GetLimit(), opts.GetMin().GetSeconds()*1000, opts.GetMax().GetSeconds()*1000, opts.GetChannel())
	rsp, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex(opts.GetChannel()),
		s.client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var mapResp = map[string]interface{}{}

	if err := json.NewDecoder(rsp.Body).Decode(&mapResp); err != nil {
		return nil, err
	}
	events := &eventgate.EventDetails{}
	if mapResp["hits"] != nil {
		for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {
			doc := hit.(map[string]interface{})
			source := doc["_source"].(map[string]interface{})
			var e eventgate.EventDetail
			bits, _ := json.Marshal(source)
			if err := protojson.Unmarshal(bits, &e); err != nil {
				return nil, err
			}
			events.Events = append(events.Events, &e)
		}
	}

	return events, nil
}

func (s Service) Close() error {
	return nil
}
