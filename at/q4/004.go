// Demo of API usage:
// GET /admin/teachers?page=1&size=30
// GET /admin/teachers?page=1&size=30&deleted=true
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// ============= DO NOT EDIT =============

var db gorm.DB         // connected db client
var cache redis.Client // connected redis client
var logger zap.Logger  // some useful logger for all developers in company

type Teacher struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func cacheMarshaller(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func cacheUnMarshaller(val string, container interface{}) error {
	return json.Unmarshal([]byte(val), container)
}

type APPError struct {
	Message string `json:"message"`
}

func newAPPError(msg string) *APPError {
	return &APPError{Message: msg}
}

// =========== DO NOT EDIT END ===========
// =======================================

func main() {
	r := gin.Default()
	r.GET("/admin/teachers", GetTeachers)
	r.Run()
}

func GetTeachersFromCache(ctx *gin.Context, key string) ([]Teacher, error) {
	cacheResult, err := cache.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	var teachers []Teacher
	if err := cacheUnMarshaller(cacheResult, &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

func GetTeachers(ctx *gin.Context) {
	withDeleted := false
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	size, _ := strconv.ParseInt(ctx.Query("size"), 10, 64)

	if ctx.Query("deleted") == "true" {
		withDeleted = true
	}

	cacheKey := fmt.Sprintf("%d-%d-%s", page, size, withDeleted)

	var ret []Teacher
	ret, err := GetTeachersFromCache(ctx, cacheKey)
	if err != nil && err != redis.Nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if err == redis.Nil {
		ret, err = GetTeachersBiz(ctx, size, page, withDeleted)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return
		}
	}

	ret, err = GetTeachersFromCache(ctx, cacheKey)
	if err != nil && err != redis.Nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, ret)
}

func GetTeachersBiz(ctx *gin.Context, size int64, page int64, withDeleted bool) ([]Teacher, error) {
	offset := size * (page - 1)

	query := db.Where("").Limit(size).Offset(offset)

	if !withDeleted {
		query.Where("deleted_at IN NOT NULL")
	}

	var teachers []Teacher
	err := query.Find(&teachers).Error

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return nil, err
	}

	if bs, err := cacheMarshaller(teachers); err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return nil, err
	} else {
		if err := cache.Set(ctx, fmt.Sprintf("%d-%d-%s", page, size, withDeleted), bs, time.Minute*3).Err(); err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return nil, err
		}
	}

	return teachers, nil
}
