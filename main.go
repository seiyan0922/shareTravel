package main

import (
	"net/http"
	"shareTravel/controller"
	"shareTravel/model"
)

func main() {

	//サーバー起動時にDBコネクションの起動
	model.OpenSQL()
	defer model.Db.Close()

	http.HandleFunc("/", controller.MakeHandler(controller.TopHandler))
	http.HandleFunc("/view/", controller.MakeHandler(controller.ViewHandler))
	http.HandleFunc("/event/", controller.MakeHandler(controller.EventHandler))
	http.HandleFunc("/member/", controller.MakeHandler(controller.MemberHandler))
	http.HandleFunc("/expense/", controller.MakeHandler(controller.ExpenseHandler))

	//cssファイルへのハンドラを定義
	http.Handle("/layout/", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))

	http.ListenAndServe(":8080", nil)
}
