package ovs

import (
	//"context"
	"fmt"
	"errors"

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
	/*
		_, err := c.etcd.Get(context.Background(), c.space.RandKey())
		return err
	*/

	/*
	//TODO this is only the get schema for now ... and not normal get that we with to implement
	//r,err:=c.ovs.GetSchema("ovnsb_db.db")
	//r,err:=c.ovs.GetSchema("ovsdb_server")
	//r,err:=c.ovs.GetSchema("/opt/ovn/ovnsb_db.db")
	r,err:=c.ovs.GetSchema("_Server")
	fmt.Print(r)
	if err != nil {
		return err
	}
	*/
	//uuid := "some-uuid"
	//uuid := "2f77b348-9768-4866-b761-89d5177ecdab"

	////condition := libovsdb.NewCondition("_uuid","!=",libovsdb.UUID{uuid})
	condition := libovsdb.NewCondition("ip","!=","check")

	getOp := libovsdb.Operation{
			Op:"select",
			Table: "MAC_Binding",
			Where: []interface{}{condition},
	}

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

func doPut(c *client) error {
	/*
		key := c.space.RandKey()
		value := key
		_, err := c.etcd.Put(context.Background(), key, value)
		return err
	*/
	/*
	SetConfig()
	if testing.Short() {
		t.Skip()
	}
	*/

	/*
	// NamedUUID is used to add multiple related Operations in a single Transact operation

	var bridgeName = "gopher-br7"
	namedUUID := "gopher"

	externalIds := make(map[string]string)
	externalIds["go"] = "awesome"
	externalIds["docker"] = "made-for-each-other"
	oMap, err := libovsdb.NewOvsMap(externalIds)
	if err != nil {
		return err
	}
	// bridge row to insert
	bridge := make(map[string]interface{})
	bridge["name"] = bridgeName
	bridge["external_ids"] = oMap

	// simple insert operation
	insertOp := libovsdb.Operation{
		Op:       "insert",
		//Table:    "Bridge",
		Table:    "SB_global",
		//Row:      bridge,
		Row:      "nb_cfg",
		UUIDName: namedUUID,
	}

	// Inserting a Bridge row in Bridge table requires mutating the open_vswitch table.
	mutateUUID := []libovsdb.UUID{{namedUUID}}
	mutateSet, _ := libovsdb.NewOvsSet(mutateUUID)
	mutation := libovsdb.NewMutation("bridges", "insert", mutateSet)
	// hacked Condition till we get Monitor / Select working
	condition := libovsdb.NewCondition("_uuid", "!=", libovsdb.UUID{"2f77b348-9768-4866-b761-89d5177ecdab"})

	// simple mutate operation
	mutateOp := libovsdb.Operation{
		Op:        "mutate",
		Table:     "Open_vSwitch",
		Mutations: []interface{}{mutation},
		Where:     []interface{}{condition},
	}

	operations := []libovsdb.Operation{insertOp, mutateOp}
	reply, err := c.ovs.Transact("_Server", operations...)
	if err!=nil{
			return err
	}

	if len(reply) < len(operations) {
		fmt.Printf("number of replies is %v and nomber of operations is %v",len(reply),len(operations))
		return errors.New("Number of Replies should be atleast equal to number of Operations")
	}
	//ok := true
	for i, o := range reply {
		if o.Error != "" && i < len(operations) {
			return errors.New(fmt.Sprintf("Transaction Failed due to an error :", o.Error, " details:", o.Details, " in ", operations[i]))
			//ok = false
		} else if o.Error != "" {
			return errors.New(fmt.Sprintf("Transaction Failed due to an error :", o.Error))
			//ok = false
		}
	}
	*/
	/*
	if ok {
		fmt.Println("Bridge Addition Successful : ", reply[0].UUID.GoUUID)
		bridgeUUID = reply[0].UUID.GoUUID
	}
	 */
	// simple insert operation

	bridge := make(map[string]interface{})
	bridge["ip"] = "omer"
	//bridge["external_ids"] = "yay"

	insertOp := libovsdb.Operation{
		Op:       "insert",
		Table:    "MAC_Binding",
		Row : bridge,
		//Table:    "SHOKO",
		//Row:      bridge,
		//Row:      "nb_cfg",
		//Row : map[string]string{"a":"b"},
		//UUIDName: "omer",
	}
	operations := []libovsdb.Operation{insertOp}
	//reply, err := c.ovs.Transact("_Server", operations...)
	reply, err := c.ovs.Transact("OVN_Southbound", operations...)
	if err!=nil{
			return err
	}

	if len(reply) < len(operations) {
		fmt.Printf("number of replies is %v and nomber of operations is %v",len(reply),len(operations))
		return errors.New("Number of Replies should be atleast equal to number of Operations")
	}
	//ok := true


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
	/*
	start, end := c.space.RandRange()
	_, err := c.etcd.Get(context.Background(), start, clientv3.WithRange(end))
	return err
	*/
	return nil
}
