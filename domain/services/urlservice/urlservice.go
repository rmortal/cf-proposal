package urlservice

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/model"
	"context"
	"fmt"
	"strconv"
)

type UrlRepo interface {
	Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error)
	DeleteUrl(ctx context.Context, id int32) error
	GetUrl(ctx context.Context, id int32) (model.UrlDto, error)
	GetLongUrl(ctx context.Context, shortUrl string) (model.LongUrlDto, error)
	GetShortUrlByLongUrl(ctx context.Context, longUrl string) (model.UrlDto, error)
}

type Url struct {
	repo UrlRepo
}

func Init(repo UrlRepo) *Url {
	return &Url{
		repo: repo,
	}
}

func (u *Url) Create(ctx context.Context, urlDto model.UrlDto) (model.UrlDto, error) {
	urlDto, err := u.repo.Create(ctx, urlDto)
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("%w", err)
	}
	return urlDto, nil
}

func (u *Url) DeleteUrl(ctx context.Context, id string) error {

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	err = u.repo.DeleteUrl(ctx, int32(convertedId))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (u *Url) GetUrl(ctx context.Context, id string) (model.UrlDto, error) {

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return model.UrlDto{}, &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	urlDto, err := u.repo.GetUrl(ctx, int32(convertedId))
	if err != nil {
		return model.UrlDto{}, err
	}
	return urlDto, nil
}

func (u *Url) GetLongUrl(ctx context.Context, shortUrl string) (model.LongUrlDto, error) {
	longUrl, err := u.repo.GetLongUrl(ctx, shortUrl)
	if err != nil {
		return model.LongUrlDto{}, fmt.Errorf("%w", err)
	}
	return longUrl, nil
}

func (u *Url) GetShortUrlByLongUrl(ctx context.Context, longUrl string) (model.UrlDto, error) {
	urlDto, err := u.repo.GetShortUrlByLongUrl(ctx, longUrl)
	if err != nil {
		return model.UrlDto{}, fmt.Errorf("%w", err)
	}
	return urlDto, nil
}
