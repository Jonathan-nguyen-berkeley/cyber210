package chains

import (
	"cyber210/final/utils"
	"hash/fnv"

	"go.uber.org/zap"
)

func BuildChainFromTwigs(log *zap.SugaredLogger) *utils.Chain {
	genesis := utils.NewBlock("GENESIS")
	return utils.NewChain(fnv.New32(), genesis, log)
}
