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

type Contest struct {
	Id                  int     `json:"id"`
	Name                string  `json:"name"`
	Type                string  `json:"type"`
	Phase               string  `json:"phase"`
	Frozen              bool    `json:"frozen"`
	DurationSeconds     int     `json:"durationSeconds"`
	StartTimeSeconds    *int    `json:"startTimeSeconds,omitempty"`
	RelativeTimeSeconds *int    `json:"relativeTimeSeconds,omitempty"`
	PreparedBy          *string `json:"preparedBy,omitempty"`
	WebsiteURL          *string `json:"websiteUrl,omitempty"`
	Description         *string `json:"description,omitempty"`
	Difficulty          *int    `json:"difficulty,omitempty"`
	Kind                *string `json:"kind,omitempty"`
	ICPCRegion          *string `json:"icpcRegion,omitempty"`
	Country             *string `json:"country,omitempty"`
	City                *string `json:"city,omitempty"`
	Season              *string `json:"season,omitempty"`
}

// Problem represents a problem.
type Problem struct {
	ContestId      *int     `json:"contestId,omitempty"`
	ProblemsetName *string  `json:"problemsetName,omitempty"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Type           string   `json:"type"` // PROGRAMMING or QUESTION
	Points         *float64 `json:"points,omitempty"`
	Rating         *int     `json:"rating,omitempty"`
	Tags           []string `json:"tags"`
}

// Party represents a party, participating in a contest.
type Party struct {
	ContestId        *int     `json:"contestId,omitempty"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"` // CONTESTANT, PRACTICE, VIRTUAL, MANAGER, OUT_OF_COMPETITION
	TeamId           *int     `json:"teamId,omitempty"`
	TeamName         *string  `json:"teamName,omitempty"`
	Ghost            bool     `json:"ghost"`
	Room             *int     `json:"room,omitempty"`
	StartTimeSeconds *int64   `json:"startTimeSeconds,omitempty"`
}

// Member represents a member of a party.
type Member struct {
	Handle string  `json:"handle"`
	Name   *string `json:"name,omitempty"`
}

// Submission represents a submission.
type Submission struct {
	Id                  int      `json:"id"`
	ContestId           *int     `json:"contestId,omitempty"`
	CreationTimeSeconds int64    `json:"creationTimeSeconds"`
	RelativeTimeSeconds int64    `json:"relativeTimeSeconds"`
	Problem             Problem  `json:"problem"`
	Author              Party    `json:"author"`
	ProgrammingLanguage string   `json:"programmingLanguage"`
	Verdict             *string  `json:"verdict,omitempty"` // FAILED, OK, PARTIAL, COMPILATION_ERROR, etc.
	Testset             string   `json:"testset"`           // SAMPLES, PRETESTS, TESTS, CHALLENGES, etc.
	PassedTestCount     int      `json:"passedTestCount"`
	TimeConsumedMillis  int      `json:"timeConsumedMillis"`
	MemoryConsumedBytes int64    `json:"memoryConsumedBytes"`
	Points              *float64 `json:"points,omitempty"`
}

type byTimeDesc []RatingChange

func (b byTimeDesc) Len() int      { return len(b) }
func (b byTimeDesc) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byTimeDesc) Less(i, j int) bool {
	return b[i].RatingUpdateTimeSeconds > b[j].RatingUpdateTimeSeconds
}

// apiGetUserInfo 通过 Codeforces API 获取用户信息
func apiGetUserInfo(userHandle string) ([]User, error) {
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

// apiMGetUserInfo 通过 Codeforces API 获取多个用户信息
func apiMGetUserInfo(userHandles []string) ([]User, error) {
	// 将用户 handles 用分号拼接成一个字符串
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

// apiGetUserRating 通过 Codeforces API 获取用户Rating历史
func apiGetUserRating(userHandle string) ([]RatingChange, error) {
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

// apiGetContests 通过 Codeforces API 获取竞赛列表
func apiGetContests() ([]Contest, error) {
	url := "https://codeforces.com/api/contest.list?gym=false"

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
	var contests struct {
		Status string    `json:"status"`
		Result []Contest `json:"result"`
	}
	err = json.Unmarshal(body, &contests)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return contests.Result, nil
}

// apiGetUserSubmissions 通过 Codeforces API 获取用户的提交记录
func apiGetUserSubmissions(userHandle string) ([]Submission, error) {
	url := fmt.Sprintf("https://codeforces.com/api/user.status?handle=%s", userHandle)

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
	var submissions struct {
		Status string       `json:"status"`
		Result []Submission `json:"result"`
	}
	err = json.Unmarshal(body, &submissions)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return submissions.Result, nil
}
