package cf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// var cf sync.Mutex

// UserInfo 响应结构定义
type UserInfo struct {
	Status string `json:"status"`
	Result []User `json:"result"`
}

type UserRating struct {
	Status string         `json:"status"`
	Result []RatingChange `json:"result"`
}

// User 结构定义
type User struct {
	Handle                  string `json:"handle" redis:"-"`
	Email                   string `json:"email,omitempty" redis:"email"`    // 可选字段，仅在用户允许分享联系信息时显示
	VkId                    string `json:"vkId,omitempty" redis:"-"`         // 可选字段，仅在用户允许分享联系信息时显示
	OpenId                  string `json:"openId,omitempty" redis:"-"`       // 可选字段，仅在用户允许分享联系信息时显示
	FirstName               string `json:"firstName,omitempty" redis:"-"`    // 可选字段，可以缺失
	LastName                string `json:"lastName,omitempty" redis:"-"`     // 可选字段，可以缺失
	Country                 string `json:"country,omitempty" redis:"-"`      // 可选字段，可以缺失
	City                    string `json:"city,omitempty" redis:"-"`         // 可选字段，可以缺失
	Organization            string `json:"organization,omitempty" redis:"-"` // 可选字段，可以缺失
	Contribution            int    `json:"contribution" redis:"-"`
	Rank                    string `json:"rank" redis:"-"`
	Rating                  int    `json:"rating" redis:"rating"`
	MaxRank                 string `json:"maxRank" redis:"-"`
	MaxRating               int    `json:"maxRating" redis:"maxRating"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds" redis:"-"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds" redis:"-"`
	FriendOfCount           int    `json:"friendOfCount" redis:"-"`
	Avatar                  string `json:"avatar" redis:"avatar"`
	TitlePhoto              string `json:"titlePhoto" redis:"-"`
}

// RatingChange 结构定义
type RatingChange struct {
	ContestId               int    `json:"contestId" redis:"contestId"`
	ContestName             string `json:"contestName" redis:"contestName"`
	Handle                  string `json:"handle" redis:"-"`
	Rank                    int    `json:"rank" redis:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds" redis:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating" redis:"oldRating"`
	NewRating               int    `json:"newRating" redis:"newRating"`
}

// GetUserInfo 通过 Codeforces API 获取用户信息
func GetUserInfo(userHandle string) ([]User, error) {
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%s&checkHistoricHandles=false", userHandle)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	// 解析 JSON
	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return userInfo.Result, nil
}

// MGetUserInfo 通过 Codeforces API 获取多个用户信息
func MGetUserInfo(userHandles []string) ([]User, error) {
	// 将用户 handles 用逗号拼接成一个字符串
	handles := ""
	for i, handle := range userHandles {
		if i > 0 {
			handles += ";"
		}
		handles += handle
	}

	// 构造请求 URL
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%s&checkHistoricHandles=false", handles)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	// 解析 JSON
	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return userInfo.Result, nil
}

// GetUserRating 通过 Codeforces API 获取用户Rating历史
func GetUserRating(userHandle string) ([]RatingChange, error) {
	url := fmt.Sprintf("https://codeforces.com/api/user.rating?handle=%s", userHandle)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}

	// 解析 JSON
	var userRating UserRating
	err = json.Unmarshal(body, &userRating)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return userRating.Result, nil
}
