- "/"は全てにマッチする("/a", "/images/aaa/bbb”にもマッチする)
- 複数マッチする場合は長い方が優先される

- public/index.htmlを配置すると、それを配信("/")

- json.Marchalは []byte, errorを返すため, stringにキャストする必要あり