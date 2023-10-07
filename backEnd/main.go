package main

import (
	"net/http"

	"github.com/labstack/echo/v4" //引入echo框架
)

// echo.Context中存放了request跟response
// Echo的Handler function 預設會回傳一個error變數，存放錯誤訊息
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!") //回傳字串形式的response跟status code給客戶端
}

func main() {
	e := echo.New()  //建立一個Echo的物件
	e.GET("/", home) //設定當客戶端請求的url為"/"且方法為"GET"時的Handler function

	//e.Logger是Echo物件的日誌記錄器
	//Fatal()是該紀錄器中的一個method
	//當括號內的程式發生錯誤時，Fatal() method會記住錯誤訊息並強制終止程式
	e.Logger.Fatal(e.Start(":8000")) //啟動伺服器，並設定在8000阜監聽客戶端請求
}
