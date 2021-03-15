package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressData struct {
	Message interface{} `json:"message"`
	Results []struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		Address3 string `json:"address3"`
		Kana1    string `json:"kana1"`
		Kana2    string `json:"kana2"`
		Kana3    string `json:"kana3"`
		Prefcode string `json:"prefcode"`
		Zipcode  string `json:"zipcode"`
	} `json:"results"`
	Status int `json:"status"`
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "index.html")

	r.GET("/api/search", func(c *gin.Context) {
		zipcode := c.Query("zipcode") // /api/search?zipcode=?が作られる
		log.Println("zipcode: ", zipcode)
		response, err := http.Get("http://zipcloud.ibsnet.co.jp/api/search?zipcode=" + zipcode)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("response.Body: ", response.Body)
		reader := response.Body

		defer response.Body.Close()

		body, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal("ERROR: ", err)
		}

		var address AddressData
		if err := json.Unmarshal(body, &address); err != nil {
			log.Fatal(err)
		}

		fmt.Println("結果: ", address)
		c.JSON(200, address)

	})

	r.Run()
}
