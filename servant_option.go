/**
 * サーバントのオプションファイルを管理するメソッドを
 * 集約するファイル
 */
package gosseract

import (
  "os"
)

func (o *Options) init() *Options {
  o.UseFile   = false
  o.FilePath  = ""
  o.Digest = make(map[string]string)
  return o
}

func (o *Options) WithFile(path string) /* Error TODO#1 */bool {
  // 存在をチェック
  _,e := os.Open(path)
  if e != nil {
    return false// TODO#1: Error.Message
  }
  o.UseFile  = true
  o.FilePath = path
  return true
}

// 全部サーバントに属した方が良い気がする
func (o *Options) Allow(charAllowed string) {
  if charAllowed == "" {
    return
  }
  o.Digest["tessedit_char_whitelist"] = charAllowed
  return
}
