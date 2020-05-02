package db

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/YouDad/blockchain/core"
	"github.com/YouDad/blockchain/global"
	"github.com/YouDad/blockchain/types"
	"github.com/YouDad/localhost/controllers"
	"github.com/YouDad/localhost/log"
)

type DBController struct {
	controllers.BaseController
}

const dir = "/home/manjaro/go/src/github.com/YouDad/blockchain/"

// @router /db/list [get]
func (c *DBController) List() {
	dir, err := ioutil.ReadDir(dir)
	c.ReturnErr(err)

	var ret []string

	for i := 0; i < len(dir); i++ {
		name := dir[i].Name()
		if len(name) >= 3 && name[len(name)-3:] == ".db" {
			ret = append(ret, name)
		}
	}

	c.Return(ret)
}

type Block struct {
	*types.Block
	Hash      types.HashValue
	TxnNumber int
	Time      string
	Txns      []*Transaction
}

func GetBlock(b *types.Block) Block {
	txnNumber := len(b.Txns)
	timestamp := time.Unix(b.Timestamp/1e9, 0).Format("2006/01/02 15:04:05")

	var txns []*Transaction
	for _, txn := range b.Txns {
		txns = append(txns, GetTxn(txn))
	}

	return Block{b, b.Hash(), txnNumber, timestamp, txns}
}

type TxnInput struct {
	types.TxnInput
	Value      int64
	PubKeyHash string
}

type Transaction struct {
	Vin  []TxnInput
	Vout []types.TxnOutput
	Hash types.HashValue
}

func GetTxn(txn *types.Transaction) *Transaction {
	var vin []TxnInput
	for _, in := range txn.Vin {
		pubKeyHash := "coinbase"
		if in.VoutIndex != -1 {
			pubKeyHash = in.PubKey.Hash().String()
		}
		vin = append(vin, TxnInput{
			in, in.VoutValue, pubKeyHash,
		})
	}

	return &Transaction{vin, txn.Vout, txn.Hash()}
}

// @router /db/:db [get]
func (c *DBController) GetBlocks() {
	filename := c.ParamStr("db")

	global.SetRootPath(dir)
	global.Port = filename[10:14]
	group, err := strconv.Atoi(filename[15:16])
	global.MaxGroupNum = group + 1
	log.Err(err)
	bc := core.GetBlockchain(group)

	ret := make(map[int32]interface{})
	height := bc.GetHeight()
	var i int32
	for i = 0; i <= height; i++ {
		block := bc.GetBlockByHeight(i)
		if block == nil {
			log.Traceln(i)
		} else {
			ret[i] = GetBlock(block)
			block.Txns = nil
		}
	}

	c.Return(ret)
}

// @router /db/:db/:id [get]
func (c *DBController) GetObject() {
	filename := c.ParamStr("db")

	global.SetRootPath(dir)
	global.Port = filename[10:14]
	group, err := strconv.Atoi(filename[15:16])
	global.MaxGroupNum = group + 1
	log.Err(err)
	bc := core.GetBlockchain(group)

	hash := c.ParamStr("id")
	if len(hash) == 64 {
		var hashValue types.HashValue
		hashValue.UnmarshalJSON([]byte("`" + hash + "`"))

		block := core.BytesToBlock(bc.Get(hashValue))
		if block != nil {
			c.Return(GetBlock(block))
		}

		txn, err := bc.FindTxn(hashValue)
		c.ReturnErr(err)
		c.Return(GetTxn(txn))
	} else {
		height := c.ParamInt("id")
		c.Return(GetBlock(bc.GetBlockByHeight(int32(height))))
	}
}
