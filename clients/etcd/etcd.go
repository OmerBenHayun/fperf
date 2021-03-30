package etcd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fperf/fperf"
	"go.etcd.io/etcd/clientv3"

	. "github.com/fperf/etcd/generator"

)

func init() {
	fperf.Register("etcd", New, "etcd benchmark")
}

// Op is the operation type issued to etcd
type Op string

// Operations
const (
	Put    Op = "put"
	Get    Op = "get"
	Range  Op = "range"
	Delete Op = "delete"
)

type client struct {
	etcd  *clientv3.Client
	space *KeySpace
	op    Op
}

// New creates a fperf client
func New(fs *fperf.FlagSet) fperf.Client {
	var keySize int
	var op Op
	fs.IntVar(&keySize, "key-size", 4, "length of the random key")
	fs.Parse()
	args := fs.Args()
	if len(args) == 0 {
		op = Put
	} else {
		op = Op(args[0])
	}
	return &client{
		space: NewKeySpace(keySize),
		op:    op,
	}
}

// Dial to etcd
func (c *client) Dial(addr string) error {
	endpoints := strings.Split(addr, ",")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		return err
	}
	c.etcd = cli
	return nil
}

// Request etcd
func (c *client) Request() error {
	switch c.op {
	case Put:
		return doPut(c)
	case Get:
		return doGet(c)
	case Range:
		return doRange(c)
	case Delete:
		return doDelete(c)
	}
	return fmt.Errorf("unknown op %s", c.op)
}

func doPut(c *client) error {
	key := c.space.RandKey()
	value := key
	_, err := c.etcd.Put(context.Background(), key, value)
	return err
}
func doGet(c *client) error {
	_, err := c.etcd.Get(context.Background(), c.space.RandKey())
	return err
}
func doDelete(c *client) error {
	_, err := c.etcd.Delete(context.Background(), c.space.RandKey())
	return err
}
func doRange(c *client) error {
	start, end := c.space.RandRange()
	_, err := c.etcd.Get(context.Background(), start, clientv3.WithRange(end))
	return err
}
