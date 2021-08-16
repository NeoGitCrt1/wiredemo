//+build wireinject

package main

import (
	"github.com/NeoGitCrt1/wiredemo/model"
	"github.com/google/wire"
)

var monsterPlayerSet = wire.NewSet(model.NewMonster, model.NewPlayer)

func InitEndingA(name string) model.EndingA {
	wire.Build(monsterPlayerSet, model.NewEndingA)
	return model.EndingA{}
}

func InitEndingB(name string) model.EndingB {
	wire.Build(monsterPlayerSet, model.NewEndingB)
	return model.EndingB{}
}
