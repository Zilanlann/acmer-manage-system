package cf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var cf sync.Mutex

// UserInfo 响应结构定义
type UserInfo struct {
	Status string `json:"status"`
	Result []User `json:"result"`
}

type UserRating struct {
	Status string `json:"status"`
	Result []RatingChange `json:"result"`
}

// User 结构定义
type User struct {
	Handle                 string `json:"handle"`
	Email                  string `json:"email,omitempty"` // 可选字段，仅在用户允许分享联系信息时显示
	VkId                   string `json:"vkId,omitempty"`  // 可选字段，仅在用户允许分享联系信息时显示
	OpenId                 string `json:"openId,omitempty"` // 可选字段，仅在用户允许分享联系信息时显示
	FirstName              string `json:"firstName,omitempty"`  // 可选字段，可以缺失
	LastName               string `json:"lastName,omitempty"`   // 可选字段，可以缺失
	Country                string `json:"country,omitempty"`    // 可选字段，可以缺失
	City                   string `json:"city,omitempty"`       // 可选字段，可以缺失
	Organization           string `json:"organization,omitempty"` // 可选字段，可以缺失
	Contribution           int    `json:"contribution"`
	Rank                   string `json:"rank"`
	Rating                 int    `json:"rating"`
	MaxRank                string `json:"maxRank"`
	MaxRating              int    `json:"maxRating"`
	LastOnlineTimeSeconds  int    `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int   `json:"registrationTimeSeconds"`
	FriendOfCount          int    `json:"friendOfCount"`
	Avatar                 string `json:"avatar"`
	TitlePhoto             string `json:"titlePhoto"`
}

// RatingChange 结构定义
type RatingChange struct {
    ContestId              int    `json:"contestId"`
    ContestName            string `json:"contestName"`
    Handle                 string `json:"handle"`
    Rank                   int    `json:"rank"`
    RatingUpdateTimeSeconds int   `json:"ratingUpdateTimeSeconds"`
    OldRating              int    `json:"oldRating"`
    NewRating              int    `json:"newRating"`
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



func main() {
	// userHandle := []string{"Chrisann", "HUFUAI", "Dayyun"}
	// users, err := MGetUserInfo(userHandle)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// for _, user := range users {
	// 	fmt.Printf("%+v\n", user)
	// }
	ratings, err := GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	GetUserRating("Chrisann")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, rating := range ratings {
		fmt.Printf("%+v\n", rating)
	}
}
