/**
 * サーバントのオプションファイルを管理するメソッドを
 * 集約するファイル
 */
package gosseract

import (
  "os"
  "errors"
)

func (o *Options) init() *Options {
  o.UseFile   = false
  o.FilePath  = ""
  o.Digest = make(map[string]string)
  return o
}

func (o *Options) WithFile(path string) error {
  // 存在をチェック
  _,e := os.Open(path)
  if e != nil {
    return errors.New("このメッセージどうする")
  }
  o.UseFile  = true
  o.FilePath = path
  return nil
}

// 全部サーバントに属した方が良い気がする
func (o *Options) Allow(charAllowed string) {
  if charAllowed == "" {
    return
  }
  o.Digest["tessedit_char_whitelist"] = charAllowed
  return
}
