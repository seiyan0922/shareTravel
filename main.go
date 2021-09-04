package main

import (
	"net/http"
	"os"
	"shareTravel/controller"
	"shareTravel/model"
)

func main() {
	//環境変数の設定
	//DB設定
	os.Setenv("Driver", "mysql")
	os.Setenv("User", "root")
	os.Setenv("Pass", "")
	os.Setenv("Host", "127.0.0.1")
	os.Setenv("Port", "3306")
	os.Setenv("DataBase", "share_travel")
	//サーバー起動時にDBコネクションの起動
	model.Connect()
	//処理が完了した際コネクションをクローズする
	defer model.Db.Close()

	http.HandleFunc("/", controller.MakeHandler(controller.TopHandler))
	http.HandleFunc("/event/", controller.MakeHandler(controller.EventHandler))
	http.HandleFunc("/member/", controller.MakeHandler(controller.MemberHandler))
	http.HandleFunc("/expense/", controller.MakeHandler(controller.ExpenseHandler))

	//cssファイルへのハンドラを定義
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))

	http.ListenAndServe(":8080", nil)
}
