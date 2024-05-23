package cf

import (
	"fmt"
	"sort"
	"time"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/model"
	"github.com/zilanlann/acmer-manage-system/server/pkg/redis"
	"gorm.io/gorm"
)

func GetUserInfo(userHandle string) (user User, err error) {
	key := fmt.Sprintf("cf:user:%s", userHandle)
	err = global.REDIS.HGetAll(redis.Ctx, key).Scan(&user)
	return
}

func GetUserInfos(userHandles []string) (users []User, err error) {
	for _, userHandle := range userHandles {
		key := fmt.Sprintf("cf:user:%s", userHandle)
		user := User{}
		err = global.REDIS.HGetAll(redis.Ctx, key).Scan(&user)
		users = append(users, user)
	}
	return
}

func GetRatingChange(userHandle string) (ratingChange []RatingChange, err error) {
	key := fmt.Sprintf("cf:rating:%s:*", userHandle)
	keys, _, err := global.REDIS.Scan(redis.Ctx, 0, key, 1000).Result()
	for _, key := range keys {
		tmpRatingChange := RatingChange{}
		global.REDIS.HGetAll(redis.Ctx, key).Scan(&tmpRatingChange)
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
		global.REDIS.HSet(redis.Ctx, key, ratingChange)
	}
	return nil
}

// RefreshCFRating refresh users' codeforces rating and save it in database
func RefreshCFRating(userHandles []string) error {
	userInfos, err := apiMGetUserInfo(userHandles)
	if err != nil {
		return err
	}
	users := make([]model.UserStatus, len(userInfos))
	for i, userInfo := range userInfos {
		users[i].CFRating = userInfo.Rating
	}
	for i, userHandle := range userHandles {
		ratingChanges, err := apiGetUserRating(userHandle)
		if err != nil {
			return err
		}
		monthlyChange, weeklyChange := calcRatingChanges(ratingChanges)
		users[i].CFMonthlyRating = monthlyChange
		users[i].CFWeeklyRating = weeklyChange
	}
	for i, user := range users {
		err = user.UpdateByCFHandle(userHandles[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func RefreshUserInfos(userHandles []string) error {
	var userInfos []User
	queryUserHandles := make([]string, 0, 50)
	for _, userHandle := range userHandles {
		key := fmt.Sprintf("cf:user:%s", userHandle)
		if num, _ := global.REDIS.Exists(redis.Ctx, key).Result(); num > 0 {
			continue
		}
		queryUserHandles = append(queryUserHandles, userHandle)
	}
	userInfos, err := apiMGetUserInfo(queryUserHandles)
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}
	for _, userInfo := range userInfos {
		if err := setUserInfo(userInfo); err != nil {
			global.LOG.Error(err.Error())
		}
	}
	return nil
}

func setUserInfo(userInfo User) error {
	key := fmt.Sprintf("cf:user:%s", userInfo.Handle)
	err := global.REDIS.HSet(redis.Ctx, key, userInfo).Err()
	global.REDIS.Expire(redis.Ctx, key, 6*time.Hour)
	return err
}

// RefreshContests 从 Codeforces 获取最新竞赛信息并更新数据库
func RefreshContests() error {
	// 获取 Codeforces 竞赛列表
	contests, err := apiGetContests()
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}

	// 获取数据库中现有的竞赛
	var existingContests []model.OJContest
	err = global.DB.Find(&existingContests).Error
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}

	// 将现有竞赛转换为 map 以便快速查找
	existingContestsMap := make(map[int]model.OJContest)
	for _, contest := range existingContests {
		existingContestsMap[contest.ContestID] = contest
	}

	// 处理每个从 API 获取的竞赛
	for _, contest := range contests {
		startTime := time.Unix(int64(*contest.StartTimeSeconds), 0)
		ojContest := model.OJContest{
			Name:            contest.Name,
			ContestID:       contest.Id,
			OJ:              "Codeforces",
			Type:            contest.Type,
			DurationSeconds: contest.DurationSeconds,
			StartTime:       startTime,
		}

		// 检查是否已有竞赛
		if existingContest, exists := existingContestsMap[contest.Id]; exists {
			// 如果竞赛存在且有更新，则更新数据库记录
			if existingContest.Name != ojContest.Name || existingContest.DurationSeconds != ojContest.DurationSeconds || !existingContest.StartTime.Equal(ojContest.StartTime) {
				err = global.DB.Model(&existingContest).Updates(ojContest).Error
				if err != nil {
					global.LOG.Error(fmt.Sprintf("Failed to update contest %d: %v", contest.Id, err))
				} else {
					global.LOG.Info(fmt.Sprintf("Updated contest %d", contest.Id))
				}
			}
		} else {
			// 如果竞赛不存在，则插入新记录
			err = global.DB.Create(&ojContest).Error
			if err != nil {
				global.LOG.Error(fmt.Sprintf("Failed to create contest %d: %v", contest.Id, err))
			} else {
				global.LOG.Info(fmt.Sprintf("Created contest %d", contest.Id))
			}
		}
	}

	return nil
}

func RefreshSubmissionsByUser(userHandle string, userId uint) error {
	// Fetch submissions from Codeforces API
	submissions, err := apiGetUserSubmissions(userHandle)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("Failed to fetch submissions for user %s: %v", userHandle, err))
		return err
	}

	// Iterate over fetched submissions and update the database
	for _, sub := range submissions {
		tags := []model.ProblemTag{}
		// Convert API tags to ProblemTag models
		for _, tagName := range sub.Problem.Tags {
			tag := model.ProblemTag{Name: tagName}
			// Check if the tag already exists in the database
			err = global.DB.Where(&model.ProblemTag{Name: tagName}).FirstOrCreate(&tag).Error
			if err != nil {
				global.LOG.Error(fmt.Sprintf("Failed to create or get tag %s: %v", tagName, err))
				continue
			}
			tags = append(tags, tag)
		}

		var rating int
		if sub.Problem.Rating != nil {
			rating = *sub.Problem.Rating
		} else {
			rating = 0
		}

		submission := model.OJSubmission{
			UserID:  userId,
			Name:    sub.Problem.Name,
			Rating:  rating, // Assuming Problem Rating refers to submission rating
			Tags:    tags,
			Verdict: *sub.Verdict,
			OJ:      "Codeforces",
			Time:    time.Unix(sub.CreationTimeSeconds, 0),
		}

		// Check if the submission already exists in the database
		existingSubmission := model.OJSubmission{}
		err := global.DB.Where("user_id = ? AND time = ?", submission.UserID, submission.Time).First(&existingSubmission).Error
		if err == gorm.ErrRecordNotFound {
			// If the submission does not exist, create a new record
			err = submission.Create()
			if err != nil {
				global.LOG.Error(fmt.Sprintf("Failed to create submission %d: %v", sub.Id, err))
			} else {
				global.LOG.Info(fmt.Sprintf("Created submission %d", sub.Id))
			}
		} else if err != nil {
			global.LOG.Error(fmt.Sprintf("Failed to query submission %d: %v", sub.Id, err))
		}
		// else {
		// 	// If the submission exists, update it
		// 	err = submission.Update()
		// 	if err != nil {
		// 		global.LOG.Error(fmt.Sprintf("Failed to update submission %d: %v", sub.Id, err))
		// 	} else {
		// 		global.LOG.Info(fmt.Sprintf("Updated submission %d", sub.Id))
		// 	}
		// }
	}

	return nil
}

func RefreshAllUserSubmisions() error {
	users, err := model.GetACMersList()
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}
	for _, user := range users {
		userId, err := model.GetUserIdByCfHandle(user.CFHandle)
		if err != nil {
			global.LOG.Error(err.Error())
			return err
		}
		err = RefreshSubmissionsByUser(user.CFHandle, userId)
		if err != nil {
			global.LOG.Error(err.Error())
			return err
		}
	}

	return nil
}

func calcRatingChanges(ratingChanges []RatingChange) (monthlyChange, weeklyChange int) {
	sort.Sort(byTimeDesc(ratingChanges))
	currentRating := ratingChanges[0].NewRating
	weeklyChange = sort.Search(len(ratingChanges), func(i int) bool {
		return ratingChanges[i].RatingUpdateTimeSeconds <= int(time.Now().AddDate(0, 0, -7).Unix())
	})
	if weeklyChange == len(ratingChanges) {
		weeklyChange = 0
	} else {
		weeklyChange = ratingChanges[weeklyChange].NewRating
	}
	monthlyChange = sort.Search(len(ratingChanges), func(i int) bool {
		return ratingChanges[i].RatingUpdateTimeSeconds <= int(time.Now().AddDate(0, -1, 0).Unix())
	})
	if monthlyChange == len(ratingChanges) {
		monthlyChange = 0
	} else {
		monthlyChange = ratingChanges[monthlyChange].NewRating
	}
	weeklyChange = currentRating - weeklyChange
	monthlyChange = currentRating - monthlyChange
	return
}
