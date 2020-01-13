package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {

	var client *http.Client
	var remoteURL string
	{
		//setup a mocked http client.
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := httputil.DumpRequest(r, false)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", b)
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `hihi`)
		}))
		defer ts.Close()
		client = ts.Client()
		remoteURL = ts.URL
	}

	//prepare the reader instances to encode
	values := map[string]io.Reader{
		"file":  mustOpen("go.mod"), // lets assume its this file
		"other": strings.NewReader("hello world!"),
	}
	err := Upload(client, remoteURL, values)
	if err != nil {
		panic(err)
	}
}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// 添加图片文件
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	} else {
		data, err := httputil.DumpResponse(res, true)
		if err == nil {
			fmt.Printf("%s", data)
		}
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
