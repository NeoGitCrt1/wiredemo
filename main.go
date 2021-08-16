package main

import "github.com/NeoGitCrt1/wiredemo/model"

func main() {
	monster := model.NewMonster()
	player := model.NewPlayer("dj")
	mission := model.NewMission(player, monster)

	mission.Start()

//	InitMission("injectedMisson").Start()

//	InitEndingA("injA").Appear()
}
