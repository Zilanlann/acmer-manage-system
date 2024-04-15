package cf

import (
	"fmt"
	"log"
	"time"

	cf "github.com/zilanlann/acmer-manage-system/server/pkg/cf/api"
	"github.com/zilanlann/acmer-manage-system/server/pkg/redis"
)

func GetCurrentRatings(userHandles []string) error {
	var userInfos []cf.User
	queryUserHandles := make([]string, 50)
	for _, userHandle := range userHandles {
		key := fmt.Sprintf("cf:user:%s", userHandle)
		if num, _ := redis.RDB.Exists(redis.Ctx, key).Result(); num > 0 {
			continue
		}
		queryUserHandles = append(queryUserHandles, userHandle)
	}
	userInfos, err := cf.MGetUserInfo(queryUserHandles)
	if err != nil {
		return err
	}
	for _, userInfo := range userInfos {
		fmt.Printf("userInfo: %v\n", userInfo)
		if err := setUserInfo(userInfo); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func RefreshRatingChange(userHandle string) error {
	ratingChanges, err := cf.GetUserRating(userHandle)
	if err != nil {
		log.Fatal(err)
	}
	for _, ratingChange := range ratingChanges {
		key := fmt.Sprintf("cf:rating:%s:%s", userHandle, ratingChange.ContestId)
		if num, _ := redis.RDB.Exists(redis.Ctx, key).Result(); num > 0 {
			continue
		}
		err := redis.RDB.HSet(redis.Ctx, key, ratingChange).Err()
		return err
	}
	return nil
}

func setUserInfo(userInfo cf.User) error {
	key := fmt.Sprintf("cf:user:%s", userInfo.Handle)
	err := redis.RDB.HSet(redis.Ctx, key, userInfo).Err()
	redis.RDB.Expire(redis.Ctx, key, 6*time.Hour).Err()
	return err
}
