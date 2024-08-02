package chains

import (
	"cyber210/final/utils"
	"fmt"
	"math/big"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func buildChainFromStraw(log *zap.SugaredLogger) *utils.Chain {
	genesis := utils.NewBlock("GENESIS", log)
	return utils.NewChain(nil, genesis, log)
}

func RunStrawcoinBenchmark(log *zap.SugaredLogger) {
	color.New(color.FgYellow).Print("========================= Testing Strawcoin =========================\n")

	jonathan := utils.User{Name: "jonathan"}
	coleman := utils.User{Name: "coleman"}
	richard := utils.User{Name: "richard"}
	satoshi := utils.User{Name: "satoshi"}

	jonathan.GeneratePrivateKey(big.NewInt(461))
	coleman.GeneratePrivateKey(big.NewInt(641))
	richard.GeneratePrivateKey(big.NewInt(853))
	satoshi.GeneratePrivateKey(big.NewInt(997))

	strawcoin := buildChainFromStraw(log)
	log.Info("Strawcoin Created")

	block1 := utils.NewBlock(strawcoin.Curr.GetHash(strawcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block1, jonathan, coleman, 100)
	utils.AddTransactionHelper(block1, coleman, richard, 20)
	utils.AddTransactionHelper(block1, coleman, satoshi, 34)
	utils.AddTransactionHelper(block1, satoshi, jonathan, 12)
	time := block1.ComputeWork(strawcoin.Hash)
	color.New(color.FgCyan).Printf("Time to compute proof of work for new strawcoin block: %s\n", time)
	strawcoin.AddBlock(block1)
	log.Info("Block Added")

	block2 := utils.NewBlock(strawcoin.Curr.GetHash(strawcoin.Hash), log)
	log.Info("Block Created")
	utils.AddTransactionHelper(block2, richard, coleman, 100)
	utils.AddTransactionHelper(block2, richard, jonathan, 20)
	utils.AddTransactionHelper(block2, richard, satoshi, 34)
	utils.AddTransactionHelper(block2, richard, jonathan, 12)
	time = block2.ComputeWork(strawcoin.Hash)
	color.New(color.FgCyan).Printf("Time to compute proof of work for new strawcoin block: %s\n", time)
	strawcoin.AddBlock(block2)
	log.Info("Block Added")

	fmt.Printf(strawcoin.String())

}
