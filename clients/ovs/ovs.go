package ovs

import (
	//"context"
	"errors"
	"fmt"
	"strconv"

	//"strings"
	//"time"

	"github.com/ebay/libovsdb"
	"github.com/fperf/fperf"
	//. "github.com/fperf/etcd/generator"
)

func init() {
	fperf.Register("ovs", New, "ovs benchmark")
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
	ovs  *libovsdb.OvsdbClient
	//space *KeySpace
	op    Op
}

// New creates a fperf client
func New(fs *fperf.FlagSet) fperf.Client {
	var op Op
	/*
	var keySize int
	fs.IntVar(&keySize, "key-size", 4, "length of the random key")
	*/
	fs.Parse()
	args := fs.Args()
	if len(args) == 0 {
		op = Put
	} else {
		op = Op(args[0])
	}
	return &client{
		//space: NewKeySpace(keySize),
		op:    op,
	}
}

// Dial to ovs
func (c *client) Dial(addr string) error {
	// TODO: mange configs
	cli,err := libovsdb.Connect(addr,nil)
	if err != nil {
		fmt.Printf("Dial error:") //FIXME change this in the future
		return err
	}
	c.ovs=cli
	/*
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
	 */
	return nil
}

// Request ovs
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

func doGet(c *client) error {
	condition := libovsdb.NewCondition("type","!=",77)

	getOp := libovsdb.Operation{
			Op:"select",
			Table: "nb_cfg",
			Where: []interface{}{condition},
	}
	/*
	getOp := libovsdb.Operation{
			Op:"select",
			Table: "MAC_Binding",
			Where: []interface{}{condition},
	}
	*/

	operations := []libovsdb.Operation{getOp}

	reply, err := c.ovs.Transact("OVN_Southbound", operations...)
	if err!=nil{
			return err
	}

	if len(reply) < len(operations) {
		fmt.Printf("number of replies is %v and nomber of operations is %v",len(reply),len(operations))
		return errors.New("Number of Replies should be atleast equal to number of Operations")
	}
	fmt.Print("\n111:\n")
	fmt.Print(reply)


	return nil
}

var i=0

func doPut(c *client) error {
	/*

	bridge := make(map[string]interface{})
	bridge["ip"] = "omer"

	insertOp := libovsdb.Operation{
		Op:       "insert",
		Table:    "MAC_Binding",
		Row : bridge,
	}
	operations := []libovsdb.Operation{insertOp}
	reply, err := c.ovs.Transact("OVN_Southbound", operations...)
	if err!=nil{
			return err
	}

	if len(reply) < len(operations) {
		fmt.Printf("number of replies is %v and nomber of operations is %v",len(reply),len(operations))
		return errors.New("Number of Replies should be atleast equal to number of Operations")
	}
	//ok := true
	*/
	r := map[string]interface{}{}
	i++
	s:=strconv.Itoa(i)
	r["name"]="yaya123"+s
	r["hostname"]="yaya123"+s
	//r["data"]="gaga"
	insertOp := libovsdb.Operation{
		Op:       "insert",
		//Op:       "update",
		Table:    "Chassis",
		Row : r,
	}
	operations := []libovsdb.Operation{insertOp}
	//reply, err := c.ovs.Transact("OVN_Southbound", operations...)
	reply, err := c.ovs.Transact("OVN_Southbound", operations...)
	if err!=nil{
			return err
	}else{
			fmt.Print(reply)
	}
	fmt.Print("\n")



	return nil
}
func doDelete(c *client) error {
	/*
	_, err := c.etcd.Delete(context.Background(), c.space.RandKey())
	return err
	*/
	return nil
}
func doRange(c *client) error {
	//TODO metod for debug delete in the future
	//TODO this is only the get schema for now ... and not normal get that we with to implement
	//r,err:=c.ovs.GetSchema("ovnsb_db.db")
	//r,err:=c.ovs.GetSchema("ovsdb_server")
	//r,err:=c.ovs.GetSchema("/opt/ovn/ovnsb_db.db")
	r,err:=c.ovs.GetSchema("_Server")
	if err != nil {
		return err
	}
	fmt.Print(r)
	/*
	start, end := c.space.RandRange()
	_, err := c.etcd.Get(context.Background(), start, clientv3.WithRange(end))
	return err
	*/
	return nil
}
