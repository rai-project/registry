package etcd

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/coreos/etcd/client"

	"github.com/rai-project/registry"
	"github.com/rai-project/utils"
)

// Client wraps some information p3sr services use as etcd clients
type Client struct {
	opts   registry.Options
	client client.Client
}

// New creates a new Client
func New(opts ...registry.Option) (registry.Registry, error) {
	cfg := client.Config{
		Endpoints: Config.Endpoints,
	}

	var options registry.Options
	for _, o := range opts {
		o(&options)
	}

	if options.Timeout == 0 {
		cfg.HeaderTimeoutPerRequest = Config.Timeout
	}

	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{cfg: cfg, client: c}, nil
}

func (ec *Client) Configuration() client.Config {
	return ec.cfg
}

// FindService queries etcd to find an address that provides a service
func (ec *Client) FindService(name string) (string, error) {
	kapi := client.NewKeysAPI(ec.client)
	key := cleanup(serviceDesc)
	resp, err := kapi.Get(context.Background(), key, &client.GetOptions{
		Recursive: true,
	})
	if err != nil {
		return "", err
	}
	// print common key info
	//log.Printf("Get is done. Metadata is %q\n", resp)
	// print value
	//log.Printf("%q key has %q nodes\n", resp.Node.Key, resp.Node.Nodes)

	if resp.Node.Value != "" {
		return resp.Node.Value, nil
	}
	if len(resp.Node.Nodes) == 0 {
		return "", errors.New("No available service for " + serviceDesc.ServiceName)
	}
	return resp.Node.Nodes[0].Value, nil
}

// MustFindService creates a fatal error if it fails
func (ec *Client) MustFindService(name string) string {
	s, err := ec.FindService(name)
	if err != nil {
		log.WithError(err).
			WithField("key", Key(serviceDesc)).
			Fatal("Unable to find ", serviceDesc.ServiceName, " key in etcd")
		panic(err)
	}
	return s
}

// Heartbeat registers a service with etcd and sends a heartbeat
func (ec *Client) Register(name string, address string) {
	kapi := client.NewKeysAPI(ec.client)
	key := cleanup(serviceDesc)
	if strings.HasPrefix(address, "localhost:") {
		ip, err := utils.GetLocalIp()
		if err != nil {
			panic(err)
		}
		address = strings.Replace(address, "localhost", ip, -1)
	}
	go heartbeat(10*time.Second, kapi, key, address)
}

func heartbeat(period time.Duration, kapi client.KeysAPI, key, value string) {
	go func() {
		ticker := time.Tick(period)
		// Set val every period with ttl = 2 * period
		ttl := period * 10
		for range ticker {
			//opts := client.SetOptions{Dir: true}
			opts := &client.SetOptions{
				TTL: ttl,
			}
			if _, err := kapi.Set(context.Background(), key, value, opts); err != nil {
				log.WithError(err).Errorf("Cannot set key = %s to value = %s", key, value)
			}
		}
	}()
}
