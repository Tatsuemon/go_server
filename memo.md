- "/"は全てにマッチする("/a", "/images/aaa/bbb”にもマッチする)
- 複数マッチする場合は長い方が優先される

- public/index.htmlを配置すると、それを配信("/")

- json.Marchalは []byte, errorを返すため, stringにキャストする必要あり　
- go get -u github/----/----で更新

- json:"-"で出力しない
- json:"age,omitempty"で空だと出力しない
(https://golang.org/pkg/encoding/json/#Marshal)

- [2020-07-14 17:44:48]  unsupported destination, should be slice or struct 
- GetUserができない理由を調べる
