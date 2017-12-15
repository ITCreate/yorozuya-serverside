# yorozuya-serverside
> すごい経営ゲーム作ります

# how to run

GOPATHで設定したディレクトリ内にcloneしてください。

```
dep ensure
go run app.go
```

クライアントは http://github.com/ITCreate/yorozuya-client に保存されています。
開発時はそちらのリポジトリをクローンし、 `npm run dev` してください。

開発時は `localhost:9999` （サーバーサイド（このリポジトリ）） と `localhost:8080` （クライアントサイド（上記リポジトリ））
を同時に立ててください。

# Tools/Libraries

- gorilla/mux
- gorilla/sessions
- goose
- go-sqlite3
