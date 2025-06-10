// Package client @author: Violet-Eva @date  : 2025/6/10 @notes :
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/violet-eva-01/one-stop/client/types"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client[t types.T] struct {
	Method       string
	Url          string
	Path         string
	Proxy        *url.URL
	Size         int
	Headers      map[string]string
	RequestBody  types.ReqBody
	ResponseBody []t
}

func NewClient[t types.T]() *Client[t] {
	return &Client[t]{}
}

func (c *Client[t]) SetMethod(method string) *Client[t] {
	c.Method = method
	return c
}

func (c *Client[t]) SetUrl(url string) *Client[t] {
	c.Url = url
	return c
}

func (c *Client[t]) SetPath(path string) *Client[t] {
	c.Path = path
	return c
}

func (c *Client[t]) SetProxy(proxy string) *Client[t] {
	if parse, err := url.Parse(proxy); err != nil {
		log.Fatal(err)
	} else {
		c.Proxy = parse
	}
	return c
}

func (c *Client[t]) SetSize(size int) *Client[t] {
	c.Size = size
	return c
}

func (c *Client[t]) SetHeaders(headers map[string]string) *Client[t] {
	c.Headers = headers
	return c
}

func (c *Client[t]) SetRequestBody(body types.ReqBody) *Client[t] {
	c.RequestBody = body
	return c
}

func (c *Client[t]) SetResponseBody(body []t) *Client[t] {
	c.ResponseBody = body
	return c
}

func (c *Client[t]) NewRequest() (err error) {
	var (
		reqBody *bytes.Buffer
	)

	if c.RequestBody.WorksheetId != "" {
		var marshal []byte
		if marshal, err = json.Marshal(c.RequestBody); err != nil {
			return
		}
		reqBody = bytes.NewBuffer(marshal)
	}

	if c.Size != 0 {
		c.RequestBody.SetPageSize(c.Size)
		for i := 1; true; i++ {
			c.RequestBody.SetPageIndex(i)
			if err = c.GetResponseBody(reqBody); err != nil {
				return
			}
			if len(c.ResponseBody)%c.Size != 0 {
				break
			}
		}
	} else {
		if err = c.GetResponseBody(reqBody); err != nil {
			return
		}
	}

	return
}

func (c *Client[t]) GetResponseBody(reqBody *bytes.Buffer) (err error) {
	defer func() {
		if err != nil {
			err = errors.New(fmt.Sprintf("GetResponseBody error: %s", err))
		}
	}()
	var (
		req      *http.Request
		resp     *http.Response
		respData []byte
	)

	if req, err = http.NewRequest(c.Method, c.Url, reqBody); err != nil {
		return
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	if resp, err = (&http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(c.Proxy)}}).Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	if respData, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	var respBody types.RespBody[t]
	if err = json.Unmarshal(respData, &respBody); err != nil {
		return
	}

	c.ResponseBody = append(c.ResponseBody, respBody.Data.Rows...)
	return
}
