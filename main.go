package main

import (
	"cyber210/final/chains"
	"cyber210/final/utils"
	"fmt"
	"log"
	"math/big"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panicf("Failed to create logger %s\n", err.Error())
	}
	log := logger.Sugar()

	public := big.NewInt(461)
	private := utils.GeneratePrivateKey("jonathan", public)

	strawcoin := chains.BuildChainFromStraw(log)
	newBlock := utils.NewBlock(strawcoin.Curr.GetHash(strawcoin.Hash))
	transaction1 := utils.NewTransaction("jonathan", "coleman", 100, newBlock.GetTransactionCount()+1)
	transaction1.Sign(private)
	added := newBlock.AddTransaction(transaction1)
	if !added {
		log.Warn(color.New(color.FgYellow).Sprint("Could not add transaction\n"))
	}
	transaction2 := utils.NewTransaction("jonathan", "richard", 20, newBlock.GetTransactionCount()+1)
	transaction2.Sign(private)
	added = newBlock.AddTransaction(transaction2)
	if !added {
		log.Warn(color.New(color.FgYellow).Sprint("Could not add transaction\n"))
	}
	color.New(color.FgCyan).Printf("Time to compute proof of work for new strawcoin block: %s\n", newBlock.ComputeWork(strawcoin.Hash))
	strawcoin.AddBlock(newBlock)
	fmt.Printf(strawcoin.String())

	/* ==============================
	 */
	twigcoin := chains.BuildChainFromTwigs(log)
	newBlock = utils.NewBlock(twigcoin.Curr.GetHash(twigcoin.Hash))
	transaction1 = utils.NewTransaction("jonathan", "coleman", 100, newBlock.GetTransactionCount()+1)
	transaction1.Sign(private)
	added = newBlock.AddTransaction(transaction1)
	if !added {
		log.Warn(color.New(color.FgYellow).Sprint("Could not add transaction\n"))
	}
	transaction2 = utils.NewTransaction("jonathan", "richard", 20, newBlock.GetTransactionCount()+1)
	transaction2.Sign(private)
	added = newBlock.AddTransaction(transaction2)
	if !added {
		log.Warn(color.New(color.FgYellow).Sprint("Could not add transaction\n"))
	}
	color.New(color.FgCyan).Printf("Time to compute proof of work for new twigcoin block: %s\n", newBlock.ComputeWork(twigcoin.Hash))
	twigcoin.AddBlock(newBlock)
	fmt.Printf(twigcoin.String())
}
