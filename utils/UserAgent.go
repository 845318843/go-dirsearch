package utils

import (
	"godir/common"
	"math/rand"
	"time"
)

var (
	MyRand = rand.New(rand.NewSource(time.Now().Unix()))
)

func Random_Num() int32 {
	return MyRand.Int31n(int32(len(common.User_Agents)))
}

func Random_UA() string {
	return common.User_Agents[Random_Num()]
}
