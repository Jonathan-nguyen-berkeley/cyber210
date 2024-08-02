package utils

import (
	"hash"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

type Chain struct {
	Hash   hash.Hash
	Head   *Block
	Curr   *Block
	size   int
	logger *zap.SugaredLogger
}

func NewChain(Hash hash.Hash, start *Block, logger *zap.SugaredLogger) *Chain {
	return &Chain{Hash, start, start, 1, logger}
}

func (c *Chain) AddBlock(block *Block) bool {
	if !block.verified(c.Hash, c.Curr.GetHash(c.Hash)) {
		c.logger.Warn("Invalid Block")
		return false
	}
	c.Curr.next = block
	c.size++
	c.Curr = block
	return true
}

func (c *Chain) String() string {
	res := color.New(color.FgMagenta).Sprintf("Chain has %d block(s)\n\n", c.size)
	tail := c.Head
	for i := 0; i < c.size; i++ {
		res += color.New(color.FgBlue).Sprintf("BLOCK # %d:\n%s\n", i, tail.String())
		tail = tail.next
	}
	return res
}
