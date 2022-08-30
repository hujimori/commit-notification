package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// githubからイベントデータを取得
	// httpリクエストを送る
	req, _ := http.NewRequest("GET", "https://api.github.com/users/hujimori/events", nil)

	req.Header.Set("Accept", "application/vnd.github+json")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Print(string(bodyBytes))

	// 取得したデータを構造体につめかえる
	// 日付単位で集計
	// 実行日のコミット数を数える
}
