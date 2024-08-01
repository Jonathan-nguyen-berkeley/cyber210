package chains

import (
	"cyber210/final/utils"

	"go.uber.org/zap"
)

func BuildChainFromStraw(log *zap.SugaredLogger) *utils.Chain {
	genesis := utils.NewBlock("GENESIS")
	return utils.NewChain(nil, genesis, log)
}
