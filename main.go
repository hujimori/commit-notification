package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Event struct {
	Id         string  `json:"id"`
	Type       string  `json:"string"`
	Actor      Actor   `json:"actor"`
	Repo       Repo    `json:"repo"`
	Payload    Payload `json:"payload"`
	Public     bool    `json:"bool"`
	Created_at string  `json:"created_at"`
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
	PushId       uint64    `json:"push_id"`
	Size         uint64    `json:"size"`
	DistinctSize uint64    `json:"distinct_size"`
	Ref          string    `json:"ref"`
	Head         string    `json:"string"`
	Before       string    `json:"before"`
	Commits      []*Commit `json:"commits"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
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

	// fmt.Print(string(bodyBytes))

	events := make([]*Event, 0)
	err = json.Unmarshal(bodyBytes, &events)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(len(events))
	for _, event := range events {
		fmt.Printf("%s\n", event.Id)

		for _, c := range event.Payload.Commits {
			fmt.Printf("%s\n", c.Sha)
		}
	}

	// 取得したデータを構造体につめかえる
	// 日付単位で集計
	// 実行日のコミット数を数える
}
