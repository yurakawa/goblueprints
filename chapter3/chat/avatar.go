package main

import (
	"errors"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("caht: アバターのURLを取得できません。")

// Avatar はユーザのプロフィール画像を表す型
type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar を空の構造体として定義
type AuthAvatar struct{}

// UseAuthAvatar あとでこの変数をAvatar型として定義されているフィールドにセットする
var UseAuthAvatar AuthAvatar

// GetAvatarURL は指定されたクライアントのアバターのURLを返す
// 問題が発生した場合にはエラーを返す。特に、URLを取得できなかった場合にはErrNoAvatarURLを返す
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
