package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	CreatedAt  time.Time `json:"created_at"`
	Title      string    `json:"title"`
	LikesCount int       `json:"likes_count"`
	Url        string    `json:"url"`
}

// Qiita APIを叩いてQiitaの投稿を取得する
func GetQiitaItems(n int, word string) []Item {
	// per_page 1ページあたりに含まれる要素数
	baseUrl := fmt.Sprintf("https://qiita.com/api/v2/items?page=1&per_page=%s", strconv.Itoa(n))

	if word != "" {
		baseUrl = fmt.Sprintf("%s&query=body:%s", baseUrl, word)
	}

	req, _ := http.NewRequest("GET", baseUrl, nil)
	req.Header.Set("Authorization", "Bearer "+GetApiToken())
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	byteAry, _ := ioutil.ReadAll(resp.Body)

	var items []Item
	json.Unmarshal(byteAry, &items)

	return items
}
