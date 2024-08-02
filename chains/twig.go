package chains

import (
	"cyber210/final/utils"
	"fmt"
	"hash/fnv"
	"math/big"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func buildChainFromTwigs(log *zap.SugaredLogger) *utils.Chain {
	genesis := utils.NewBlock("GENESIS", log)
	return utils.NewChain(fnv.New32(), genesis, log)
}

func RunTwigcoinBenchmark(log *zap.SugaredLogger) {
	color.New(color.FgYellow).Print("========================= Testing Twigcoin =========================\n")

	jonathan := utils.User{Name: "jonathan"}
	coleman := utils.User{Name: "coleman"}
	richard := utils.User{Name: "richard"}
	satoshi := utils.User{Name: "satoshi"}

	jonathan.GeneratePrivateKey(big.NewInt(461))
	coleman.GeneratePrivateKey(big.NewInt(641))
	richard.GeneratePrivateKey(big.NewInt(853))
	satoshi.GeneratePrivateKey(big.NewInt(997))

	twigcoin := buildChainFromTwigs(log)
	log.Info("Twigcoin Created")

	block1 := utils.NewBlock(twigcoin.Curr.GetHash(twigcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block1, jonathan, coleman, 100)
	utils.AddTransactionHelper(block1, coleman, richard, 20)
	utils.AddTransactionHelper(block1, coleman, satoshi, 34)
	utils.AddTransactionHelper(block1, satoshi, jonathan, 12)
	time := block1.ComputeWork(twigcoin.Hash)
	log.Infof("Time to compute proof of work for new twigcoin block: %s\n", time)
	twigcoin.AddBlock(block1)
	log.Info("Block Added")
	block2 := utils.NewBlock(twigcoin.Curr.GetHash(twigcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block2, richard, coleman, 100)
	utils.AddTransactionHelper(block2, richard, jonathan, 20)
	utils.AddTransactionHelper(block2, richard, satoshi, 34)
	utils.AddTransactionHelper(block2, richard, jonathan, 12)
	time = block2.ComputeWork(twigcoin.Hash)
	log.Infof("Time to compute proof of work for new twigcoin block: %s\n", time)
	twigcoin.AddBlock(block2)
	log.Info("Block Added")
	fmt.Printf(twigcoin.String())

	log.Infof("Block 0 hash: %x", twigcoin.Head.GetHash(twigcoin.Hash))
	log.Infof("Block 1 header: %x", block1.GetHeader())

	log.Infof("Block 1 hash: %x", block1.GetHash(twigcoin.Hash))
	log.Infof("Block 2 header: %x", block2.GetHeader())

	log.Infof("Block 2 hash: %x", block2.GetHash(twigcoin.Hash))

}
