package main

import (
	usecase "dddgo/application/usecase"
	infra "dddgo/infrastructure/character"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 依存関係を注入（一応DIっぽいことをしてる）
var characterRepository = infra.NewCharacterRepository()
var usecaseCharacter = usecase.NewCreateCharacterUsecase(characterRepository)

// var tr = repository.NewTodoRepository()
// var tc = controller.NewTodoController(tr)
// var ro = controller.NewRouter(tc)

func main() {

	//ルーティングの設定
	// router := httprouter.New()
	// router.GET("/api/v2/get", userHandler.HandleUserGet)
	// router.POST("/api/v2/signup", userHandler.HandleUserSignup)
	// サーバ起動
	fmt.Println("Server Running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
	r := gin.Default()
	r.GET("/go", response)
	r.Run(":1000") // デフォルトが8080ポートなので今回は変更しません
	// server := http.Server{
	// 	Addr: ":8080",
	// }
	// http.HandleFunc("/todos/", ro.HandleTodosRequest)
	// server.ListenAndServe()

}

func response(c *gin.Context) {
	chara, _ := usecaseCharacter.Create("やまだ　たろう")
	findChara, _ := characterRepository.FindById(chara.GetId())
	fmt.Println(findChara.GetId())
	fmt.Println(findChara.GetName())
	c.JSON(200, findChara.GetName())
}
