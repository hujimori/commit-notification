package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Event struct {
	Id uint32 `json:"id"`
}

type Actor struct {
	Id           uint32 `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"octocat"`
	GravatrId    string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	PushId       uint32    `json:"push_id"`
	Size         string    `json:"size"`
	DistinctSize string    `json:"distinct_size"`
	Ref          string    `json:"ref"`
	Head         string    `json:"string"`
	Before       string    `json:"before"`
	Commits      []*Commit `json:"commits"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct string `json:"distinct"`
	Url      string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

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
