package output

import (
	"bytes"
	"context"
	"net/http"
)

const (
	contentType = "text/plain"
)

// TODO: 引数 `out chan []byte` で文字列を受信した際に、その内容 Body として引数 `url string` への HTTP#POST リクエストを行う
// TODO: `Content-Type: plain/text` を Header に添えて送信を行う
// TODO: ctx context.Context がキャンセルされた場合には速やかに関数を終了する
// TODO: エラーが発生した際には errc chan error へエラーを送信する
func Forward(ctx context.Context, out chan []byte, errc chan error, url string) {
	client := &http.Client{}

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-out:
			if !ok {
				return
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(msg))
			if err != nil {
				errc <- err
				continue
			}
			req.Header.Set("Content-Type", "plain/text")

			res, err := client.Do(req)
			if err != nil {
				errc <- err
				continue
			}
			res.Body.Close()
		}
	}
}
