package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ファイル操作
	// OpenFile
	// os.O_RDONLY 読み込み専用
	// os.O_WRONLY　書き込み専用
	// os.O_RDWR　読み書き可能
	// os.O_APPEND　ファイル末尾に追記
	// os.O_CREATE　ファイルなければ作成
	// os.O_TRUNC　可能であればファイルの内容をオープン時にからにする
	// 0666 パーミッション データ記録など読み書き可能なファイル
	// ファイルがなければ読み書き可能なファイルを作成して、ログが生じたら末尾に追加していく
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
