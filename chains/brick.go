package chains

import (
	"crypto/sha256"
	"cyber210/final/utils"
	"fmt"
	"math/big"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func buildChainFromBricks(log *zap.SugaredLogger) *utils.Chain {
	genesis := utils.NewBlock("GENESIS", log)
	return utils.NewChain(sha256.New(), genesis, log)
}

func RunBrickcoinBenchmark(log *zap.SugaredLogger) {
	color.New(color.FgYellow).Print("========================= Testing Brickcoin =========================\n")

	jonathan := utils.User{Name: "jonathan"}
	coleman := utils.User{Name: "coleman"}
	richard := utils.User{Name: "richard"}
	satoshi := utils.User{Name: "satoshi"}

	jonathan.GeneratePrivateKey(big.NewInt(461))
	coleman.GeneratePrivateKey(big.NewInt(641))
	richard.GeneratePrivateKey(big.NewInt(853))
	satoshi.GeneratePrivateKey(big.NewInt(997))

	brickcoin := buildChainFromBricks(log)
	log.Info("Brickcoin Created")

	block1 := utils.NewBlock(brickcoin.Curr.GetHash(brickcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block1, jonathan, coleman, 100)
	utils.AddTransactionHelper(block1, coleman, richard, 20)
	utils.AddTransactionHelper(block1, coleman, satoshi, 34)
	utils.AddTransactionHelper(block1, satoshi, jonathan, 12)
	time := block1.ComputeWork(brickcoin.Hash)
	log.Infof("Time to compute proof of work for new brickcoin block: %s\n", time)
	brickcoin.AddBlock(block1)
	log.Info("Block Added")
	block2 := utils.NewBlock(brickcoin.Curr.GetHash(brickcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block2, richard, coleman, 100)
	utils.AddTransactionHelper(block2, richard, jonathan, 20)
	utils.AddTransactionHelper(block2, richard, satoshi, 34)
	utils.AddTransactionHelper(block2, richard, jonathan, 12)
	time = block2.ComputeWork(brickcoin.Hash)
	log.Infof("Time to compute proof of work for new brickcoin block: %s\n", time)
	brickcoin.AddBlock(block2)
	log.Info("Block Added")
	fmt.Printf(brickcoin.String())
	log.Debugf("Block 1 hash: %x", block1.GetHash(brickcoin.Hash))
}
