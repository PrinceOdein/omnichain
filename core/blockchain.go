package core

type Blockchain struct {
	Blocks  []Block
	Mempool []Transaction
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock([]Transaction{}, "", 0)
	return &Blockchain{
		Blocks:  []Block{genesisBlock},
		Mempool: []Transaction{},
	}
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.Mempool = append(bc.Mempool, tx)
}

func (bc *Blockchain) MineBlock() Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(bc.Mempool, prevBlock.Hash, len(bc.Blocks))
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []Transaction{}
	return newBlock
}
