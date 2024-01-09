package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/", Hello())
    e.GET("/api/", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
<<<<<<< HEAD
        return c.String(http.StatusOK, "4423: こんにちは これはCI/CDのテストです")
=======
        return c.String(http.StatusOK, "4423: 全部自動になったCDCIのテスト2")
>>>>>>> a8a3fc6df83d5678a018d62eafc3fa388a07a840
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {
<<<<<<< HEAD
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4423", "message": "こんにちは これはCI/CDのテストです"})
=======
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4423", "message": "全部自動になったCDCIのテスト2"})
>>>>>>> a8a3fc6df83d5678a018d62eafc3fa388a07a840
    }
}