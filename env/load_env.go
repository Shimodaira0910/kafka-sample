package env

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Env struct{}

// 環境変数を読み込むだけ

func (e *Env) LoadEnv(envName string) (string){
	err := godotenv.Load("../.env")
	if err != nil{
		fmt.Println("読み込み失敗")
		return ""
	}

	env := os.Getenv(envName)
	if env == ""{
		fmt.Println(envName + "が設定されていません")
		return ""
	} 
	
	return env
}