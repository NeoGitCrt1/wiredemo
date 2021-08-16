//+build wireinject

package main

import "github.com/google/wire"
import "github.com/NeoGitCrt1/wiredemo/model"

func InitMission(name string) model.Mission {
  wire.Build(model.NewMonster, model.NewPlayer, model.NewMission)
  return model.Mission{}
}

