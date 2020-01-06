package main

import ("fmt"
	"strconv"
	chain "github.com/zhengjianfeng/go-blockchain/v1/blockchain"
)

func main() {
	blockChain := chain.InitBlockChain()

	blockChain.AddBlock("First Block")
	blockChain.AddBlock("Second Block")
	blockChain.AddBlock("Third Block")
	for _, block := range blockChain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreHash)
		fmt.Printf("Data in block is: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := chain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
