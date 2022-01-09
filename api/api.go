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
	Url        string    `json:"url"`
	Tags       []*Tag    `json:"tags"`
	LikesCount int       `json:"likes_count"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Tag struct {
	Name string `json:"name"`
}

// Qiita APIを叩いてQiitaの投稿を取得する
func GetQiitaItems(n, p int, w string) []Item {
	// per_page 1ページあたりに含まれる要素数
	url := fmt.Sprintf("https://qiita.com/api/v2/items?page=1&per_page=%s", strconv.Itoa(n))
	baseUrl := baseUrlSelector(n, p, w, url)

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

// APIのエンドポイントをCLIの引数によって切り替える
func baseUrlSelector(n, p int, w, url string) string {
	//検索された日付の1ヶ月前の時間を取得
	now := time.Now()
	date := now.Add(time.Hour * 24 * -30)
	monthAgo := date.Format("2006-01-02")

	if w != "" && p > 0 {
		url = fmt.Sprintf("%s&query=body%s&created:>%s+stocks:>%s", url, w, monthAgo, strconv.Itoa(p))
	} else if p > 0 {
		// LGTM数で記事をソート
		url = fmt.Sprintf("%s&query=created:>%s+stocks:>%s", url, monthAgo, strconv.Itoa(p))
	} else if w != "" {
		url = fmt.Sprintf("%s&query=body:%s", url, w)
	}

	return url
}
