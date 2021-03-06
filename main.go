package main

import (
  "log"
  "net/http"
  "sync"
  "text/template"
  "path/filepath"
)

// templは１つのテンプレートを表します
type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // このonce.Doについてはこのページ参照のこと
  // http://golang.jp/pkg/sync
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, nil)
}

func main() {
  // ルート
  http.Handle("/", &templateHandler{filename: "chat.html"})

  // Webサーバーを開始します
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
