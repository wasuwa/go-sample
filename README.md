## 概要
Go で作成した Web API のサンプル。

## 技術構成
- Go
- Echo
- Gorm
- PostgreSQL
- Viper
- Gorilla

## API
### Users
|      | メソッド | URI | 権限 |
| ---- | ---- | ---- | ---- |
| 一覧表示 | GET | /users | 不要 |
| 個別表示 | GET | /users/:id | 不要 |
| 作成 | POST | /users | 不要 |
| 更新 | PATCH | /users/:id | login が必要 |
| 削除 | DELETE | /users/:id | login が必要 |

### Sessions
|      | メソッド | URI | 権限 |
| ---- | ---- | ---- | ---- |
| ログイン | POST | /login | 不要 |
| ログアウト | DELETE | /logout | 不要 |

### Tweets
|      | メソッド | URI | 権限 |
| ---- | ---- | ---- | ---- |
| 一覧表示 | GET | /tweets | 不要 |
| 作成 | POST | /tweets | 不要 |
| 削除 | DELETE | /tweets/:id | login が必要 |

