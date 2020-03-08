# media_platform_sunday
## 概要
- メディアプラットフォームのAPIを作成する.
- 使用言語はGo.
- ORM以外のライブラリの使用については追って検討.

## 目的
- webアプリケーションを作ることにより、Go言語の概要を学ぶ.
- 極力ライブラリを用いないことで、webに関する本質的な理解を深める

## 背景
- webサービスを作った経験がpythonのみであることに不安
- python以外のモダンな言語で書いてみよう => Go

## エンドポイント
### `users/signup`
#### Request Method
Post
#### Request Body
```json
{
	"user_id": "hogehoge",
	"password": "hogehoge",
	"name": "ほげほげ"
}
```
---

### `users/signin`
#### Request Method
Post
#### Request Body
```json
{
	"user_id": "hogehoge",
	"password": "hogehoge"
}
```
---

### `users/signout`
#### Request Method
Get

---

### `users/update`
#### Request Method
Post
#### Request Body
```json
{
	"name": "ほげほげさん",
	"age": 2
}
```
---

### `post/new`
#### Request Method
Post
#### Request Body
```json
{
	"title": "hogeのためのhoge",
	"text": "hogeとはhogeである"
}
```
---

### `post/index`
#### Request Method
Get

---

### `post/mine`
#### Request Method
GET

---


### `post/update`
#### Request Method
POST
#### Request Body
```json
{
	"id": 4,
	"title": "やっぱりhogeかもしれん"
}
```
---

### `post/delete`
#### Request Method
POST
#### Request Body
```json
{
	"id": 4
}
```
---

## 作業履歴
### 2020/2/6(木)
#### やったこと
- signup, signinの完成
- passwordのhash化
- signupでuser_idがかぶっているときに異常系でreturn
