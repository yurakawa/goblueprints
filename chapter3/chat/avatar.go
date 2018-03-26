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

// GravatarAvatar はGravatar用の構造体
type GravatarAvatar struct{}

// UseGravatar あとでこの変数をAvatar型として定義されているフィールドにセットする
var UseGravatar GravatarAvatar

// GetAvatarURL Gravatarのガイドラインに則りメールアドレスに含まれる大文字を小文字に変換し、その結果に対してMD5アルゴリズムを適用してハッシュ値を算出してURLに埋め込む
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
