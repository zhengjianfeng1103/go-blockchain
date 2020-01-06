package blockchain

/*
	======Proof of Working 工作量证明=========
	1.保证签发区块的速率不要太快
	
	
*/

//Take the data from the block

//Create a counter which starts at 0s

//Create a hash of the data plus the counter

//Check the hash to see if it meets a set of requirements

//Requirements:

//The First few bytes must contain 0s

import ("fmt"
	"log"
	"math/big"
	"math"
	"bytes"
	"crypto/sha256"
	"encoding/binary"

)
const Defficulty = 18


type ProofWork struct {
	Block *Block
	Target *big.Int //a number that represents the requirements that we described up
}

//创建 工作量证明
func NewProof(b *Block) *ProofWork {
	//计算要求的0s位数
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Defficulty)) // sets z = x << n and returns z

	pow := &ProofWork{Block: b, Target: target}

	return pow
}

//初始化 数据
func (pow *ProofWork) InitData(nonce int) []byte {
	//计算数据: preHash + data + Hex(nonce) + Hex(Defficulty)
	data := bytes.Join(
		[][]byte{
			pow.Block.PreHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Defficulty)),
		},
		[]byte{},
	)

	return data
}

//运行计算
func (pow *ProofWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		//验证hash的有效性- 有效的Hash: 工作量证明的target 大于 计算出的hash
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		}else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

//验证工作量证明(昂贵的Hash计算):重新计算data的Hash值. 在比较
func (pow *ProofWork) Validate() bool {
	var initHash big.Int
	 
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])

	return initHash.Cmp(pow.Target) == -1
}


//计算 16进制(字节表示)
func ToHex(num int64) []byte{
	//https://blog.csdn.net/waitingbb123/article/details/80504093
	//What different between BigEndian and LittleEndian
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}



