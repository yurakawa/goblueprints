package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
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
	if email, ok := c.userData["email"]; ok {
		// TODO: (string)って書き方何??
		// TODO: アバターのURLが必要になるたびにハッシュ値を計算している。
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}
