# go-bookings
Go言語でWebアプリケーションを作成する

## 使用しているライブラリー
### GOライブラリー
- Built in Go version 1.17
- Uses [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)
- Uses [govalidator](github.com/asaskevich/govalidator)
- Uses [pgx](github.com/jackc/pgx/v4)

### CSS Framework
- Uses [bootstrap](https://getbootstrap.com/) V5.1

### Javascript Framework
- Uses [vanillajs-datepicker](https://github.com/mymth/vanillajs-datepicker)
- Uses [notie](https://github.com/jaredreich/notie)
- Uses [sweetalert](https://github.com/sweetalert2/sweetalert2)

## データベース設計
![ER図](doc/ER.png)

### マイグレーションデータベース
- PostgreSQLクライアントツールインストール
```
brew install libpq
echo 'export PATH="/usr/local/opt/libpq/bin:$PATH"' >> ~/.zshrc
```

- sodaコマンドツールインストール
```
go get github.com/gobuffalo/pop/...
go install github.com/gobuffalo/pop/v6/soda@latest
```
