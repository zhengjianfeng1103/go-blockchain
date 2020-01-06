package blockchain

import ("bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//type struct Block
type Block struct {
	Data []byte
	Hash []byte
	PreHash []byte
	Nonce int
}

//type struct BlockChain
type BlockChain struct {
	Blocks []*Block
}

//为区块生成Hash值
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
	sum := sha256.Sum256(info)
	b.Hash = sum[:]
}


//添加区块
func (chain *BlockChain) AddBlock(data string) {
	//get pre block
	preBlock := chain.Blocks[len(chain.Blocks) - 1]
	//use pre block to create new block
	newBlock := CreateBlock(data, preBlock.Hash)
	//add new block to blockChain
	chain.Blocks = append(chain.Blocks, newBlock)
}

//1.创建区块
func CreateBlock(data string, preHash []byte) *Block {
	block := &Block{Data: []byte(data), PreHash: preHash, Hash: []byte{}}
	pow := NewProof(block)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//2.创建世纪区块
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//3.初始化区块链
func InitBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{Genesis()}}
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(res)

	if err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}


