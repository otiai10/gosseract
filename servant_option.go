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
  o.WhiteList = ""
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
