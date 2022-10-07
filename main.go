package main

import (
	"dddgo/domain/model"
	"fmt"
)

func main() {
	// 依存関係を注入（一応DIっぽいことをしてる）
	//ルーティングの設定
	// router := httprouter.New()
	// router.GET("/api/v2/get", userHandler.HandleUserGet)
	// router.POST("/api/v2/signup", userHandler.HandleUserSignup)
	// サーバ起動
	fmt.Println("Server Running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
	test := model.NewCharacter("たろう")
	fmt.Println(test.GetId())
	fmt.Println(test.GetName())

}
