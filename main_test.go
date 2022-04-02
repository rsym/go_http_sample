package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"./handler"
)

func TestTop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.TopHandler))
	defer ts.Close()

	// test1
	u := ts.URL
	fmt.Printf("[TEST] GET %s\n", u)

	res, err := http.Get(u)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode should be 200, but %d", res.StatusCode)
	}

}

func TestForm(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.FormHandler))
	defer ts.Close()

	// test1
	u := ts.URL + "/form"
	fmt.Printf("[TEST] GET %s\n", u)

	res, err := http.Get(u)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode should be 200, but %d", res.StatusCode)
	}

}

func TestSubmit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.SubmitHandler))
	defer ts.Close()

	// test1
	u := ts.URL + "/submit"
	fmt.Printf("[TEST] GET %s\n", u)

	res, err := http.Get(u)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode should be 200, but %d", res.StatusCode)
	}

	// test2
	u = ts.URL + "/submit?param1=foo&param2=bar"
	fmt.Printf("[TEST] GET %s\n", u)

	res, err = http.Get(u)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode should be 200, but %d", res.StatusCode)
	}
	s := "[ ]+<p>param1 : foo</p>\n[ ]+<p>param2 : bar</p>"
	if !regexp.MustCompile(s).MatchString(string(body)) {
		t.Errorf("responce body should be contain '%s', but doesn't contain.", s)
	}

	// test3
	u = ts.URL + "/submit"
	v := url.Values{}
	v.Add("param1", "foo")
	v.Add("param2", "bar")
	fmt.Printf("[TEST] POST %s\n", u)

	res, err = http.PostForm(u, v)
	body, _ = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode should be 200, but %d", res.StatusCode)
	}
	s = "[ ]+<p>param1 : foo</p>\n[ ]+<p>param2 : bar</p>"
	if !regexp.MustCompile(s).MatchString(string(body)) {
		t.Errorf("responce body should be contain '%s', but doesn't contain.", s)
	}
}
