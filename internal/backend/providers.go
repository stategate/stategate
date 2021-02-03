package backend

import (
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	nats2 "github.com/autom8ter/eventgate/internal/backend/nats"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"os"
	"strings"
	"time"
)

type Provider string

const (
	NATS Provider = "nats"
)

func GetProvider(provider Provider, providerConfig map[string]string) (eventgate.EventGateServiceServer, func(), error) {
	if providerConfig == nil {
		return nil, nil, errors.Errorf("empty provider config")
	}
	switch provider {
	case NATS:
		lgger := logger.New(cast.ToBool(providerConfig["debug"]))
		natsUrl := providerConfig["url"]
		if natsUrl == "" {
			return nil, nil, errors.New("nats config: empty url")
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, nil, err
		}
		hostname = strings.NewReplacer(".", "_").Replace(hostname)
		var (
			conn *nats.Conn
		)
		for i := 0; i < 10; i++ {
			var (
				clientID = fmt.Sprintf("%s-%s", hostname, nuid.Next())
			)
			conn, err = nats.Connect(
				natsUrl,
				nats.Name(clientID),
			)
			if err == nil && conn != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			return nil, nil, err
		}
		svc, err := nats2.NewService(lgger, conn)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			conn.Drain()
			conn.Close()
		}, nil
	default:
		return nil, nil, errors.Errorf("unsupported provider: %s", provider)
	}
}
