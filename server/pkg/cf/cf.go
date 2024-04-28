package cf

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/zilanlann/acmer-manage-system/server/pkg/redis"
)

func GetWMRating(userHandle string) (weeklyAgo int, monthlyAgo int, err error) {
	key := fmt.Sprintf("cf:rating:%s:*", userHandle)
	keys, _, _ := redis.RDB.Scan(redis.Ctx, 0, key, 1000).Result()
	ratingChanges := make([]RatingChange, 0, len(keys))
	for _, key := range keys {
		ratingChange := RatingChange{}
		redis.RDB.HGetAll(redis.Ctx, key).Scan(&ratingChange)
		ratingChanges = append(ratingChanges, ratingChange)
	}
	sort.Sort(byTimeDesc(ratingChanges))
	weeklyAgo = sort.Search(len(ratingChanges), func(i int) bool {
		return ratingChanges[i].RatingUpdateTimeSeconds <= int(time.Now().AddDate(0, 0, -7).Unix())
	})
	if weeklyAgo == len(ratingChanges){
		weeklyAgo = 0
	} else {
		weeklyAgo = ratingChanges[weeklyAgo].NewRating
	}
	monthlyAgo = sort.Search(len(ratingChanges), func(i int) bool {
		return ratingChanges[i].RatingUpdateTimeSeconds <= int(time.Now().AddDate(0, -1, 0).Unix())
	})
	if monthlyAgo == len(ratingChanges) {
		monthlyAgo = 0
	} else {
		monthlyAgo = ratingChanges[monthlyAgo].NewRating
	}
	return
}

func GetUserInfo(userHandle string) (user User, err error) {
	key := fmt.Sprintf("cf:user:%s", userHandle)
	err = redis.RDB.HGetAll(redis.Ctx, key).Scan(&user)
	return
}

func GetUserInfos(userHandles []string) (users []User, err error) {
	for _, userHandle := range userHandles {
		key := fmt.Sprintf("cf:user:%s", userHandle)
		user := User{}
		err = redis.RDB.HGetAll(redis.Ctx, key).Scan(&user)
		users = append(users, user)
	}
	return
}

func GetRatingChange(userHandle string) (ratingChange []RatingChange, err error) {
	key := fmt.Sprintf("cf:rating:%s:*", userHandle)
	keys, _, _ := redis.RDB.Scan(redis.Ctx, 0, key, 1000).Result()
	for _, key := range keys {
		tmpRatingChange := RatingChange{}
		redis.RDB.HGetAll(redis.Ctx, key).Scan(&tmpRatingChange)
		ratingChange = append(ratingChange, tmpRatingChange)
	}
	return
}

func RefreshRatingChange(userHandle string) error {
	ratingChanges, err := apiGetUserRating(userHandle)
	if err != nil {
		return err
	}
	for _, ratingChange := range ratingChanges {
		key := fmt.Sprintf("cf:rating:%s:%d", userHandle, ratingChange.ContestId)
		redis.RDB.HSet(redis.Ctx, key, ratingChange)
	}
	return nil
}

func RefreshUserInfos(userHandles []string) error {
	var userInfos []User
	queryUserHandles := make([]string, 0, 50)
	for _, userHandle := range userHandles {
		key := fmt.Sprintf("cf:user:%s", userHandle)
		if num, _ := redis.RDB.Exists(redis.Ctx, key).Result(); num > 0 {
			continue
		}
		queryUserHandles = append(queryUserHandles, userHandle)
	}
	userInfos, err := apiMGetUserInfo(queryUserHandles)
	if err != nil {
		return err
	}
	for _, userInfo := range userInfos {
		if err := setUserInfo(userInfo); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func setUserInfo(userInfo User) error {
	key := fmt.Sprintf("cf:user:%s", userInfo.Handle)
	err := redis.RDB.HSet(redis.Ctx, key, userInfo).Err()
	redis.RDB.Expire(redis.Ctx, key, 6*time.Hour)
	return err
}
