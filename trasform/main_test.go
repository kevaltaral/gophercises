package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

var testRoutes = []struct {
	Method    string
	Path      string
	ErrorCode int
}{
	//{Method: "GET", Path: "http://localhost:50000/", ErrorCode: 200},
	{Method: "POST", Path: "http://localhost:50000/upload", ErrorCode: 400},
	{Method: "POST", Path: "http://localhost:50000/modify/827865555.png", ErrorCode: 400},
	{Method: "POST", Path: "http://localhost:50000/modify/testme.png?mode= ", ErrorCode: 400},
	{Method: "POST", Path: "http://localhost:50000/modify/index.jpeg?mode=4", ErrorCode: 200},
	{Method: "POST", Path: "http://localhost:50000/modify/index.jpeg?", ErrorCode: 200},
	{Method: "POST", Path: "http://localhost:50000/modify/index.jpeg?mode=4&n=1", ErrorCode: 302},
}

var testRoutesNegative = []struct {
	Method    string
	Path      string
	ErrorCode int
}{
	//{Method: "POST", Path: "http://localhost:50000/modify/index.jpeg?mode=4", ErrorCode: 500},
	//{Method: "POST", Path: "http://localhost:50000/modify/index.jpeg?", ErrorCode: 500},
}

/*func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}*/
func TestRoutes(t *testing.T) {
	go main()
	time.Sleep(500)
	oldGenImages := myGenImages
	myGenImages = mockGenImages
	for index, tt := range testRoutesNegative {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			executeRequestForRoutes(t, tt.Method, tt.Path, tt.ErrorCode)
		})
	}
	myGenImages = oldGenImages

	for index, tt := range testRoutes {
		t.Run("TC_Negative_"+strconv.Itoa(index), func(t *testing.T) {
			executeRequestForRoutes(t, tt.Method, tt.Path, tt.ErrorCode)
		})
	}

}
func executeRequestForRoutes(t *testing.T, method, path string, errorCode int) {
	resp, err := followURL(method, path)
	if err != nil {
		t.Fatalf("FAIL: Got error %v while hitting %s, %s", err, method, path)
		//continue
	}
	if resp.StatusCode == 200 {
		message, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf(string(message))
		defer resp.Body.Close()
	}
	if resp.StatusCode != errorCode {
		t.Errorf("FAIL: %s returned status %d, expected %d", path, resp.StatusCode, errorCode)
	}
}
func followURL(method, path string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 240,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	var req *http.Request
	var resp *http.Response
	var err error
	req, _ = http.NewRequest(method, path, nil)
	resp, err = client.Do(req)
	//fmt.Printf("Request is %v\n", req)
	//fmt.Printf("Response is %v\n", resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func TestUploadRequest(t *testing.T) {

	filepath := "./index.jpeg"

	req, err := fileUploadRequest("http://localhost:50000/upload", "image", filepath)
	client := &http.Client{
		Timeout: time.Second * 240,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	} else {
		if resp.StatusCode != http.StatusFound {
			t.Errorf("Expected status 302 got %d", resp.StatusCode)
		}
		fmt.Printf("Response is %v\n", resp)
		redirPath, err := resp.Location()
		if err != nil {
			t.Errorf("Got error %v as response redirect\n", err)
		}

		resp, err = followURL("GET", redirPath.String())
		if err != nil {
			t.Errorf("Got error %v on following %s\n", err, redirPath)
		}
		fmt.Printf("Response to %s is %v\n", redirPath, resp)

		if resp.StatusCode != 200 {
			t.Errorf("Got response %d for %s. Expected 200\n", resp.StatusCode, redirPath)
		}

	}

}

func fileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func mockIOCopy(dst io.Writer, src io.Reader) (int64, error) {
	return 0, errors.New("Test Error")
}

func mockTempfile(prefix, ext string) (*os.File, error) {
	return nil, errors.New("Test Error")
}

func mockGenImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	return nil, errors.New("Test Error")
}

func TestUploadFile_Negative(t *testing.T) {
	//Negative TestCase for testing tempfile
	//go main()
	t.Run("UploadFile_TempFile_Negative", func(t *testing.T) {
		oldmytempfilefunc := mytempfile
		mytempfile = mockTempfile
		executeUploadReqForNegativeTestCases(t)
		mytempfile = oldmytempfilefunc
	})

	//Negative testCase for testing io.Copy

	t.Run("UploadFile_TempFile_Negative", func(t *testing.T) {
		oldmyIOCopy := myCopy
		myCopy = mockIOCopy
		executeUploadReqForNegativeTestCases(t)
		myCopy = oldmyIOCopy
	})

}

func executeUploadReqForNegativeTestCases(t *testing.T) {
	//go main()
	filepath := "./index.jpeg"
	req, err := fileUploadRequest("http://localhost:50000/upload", "image", filepath)
	client := &http.Client{
		Timeout: time.Second * 240,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	fmt.Printf("Response is %v\n", resp)
	if err != nil {
		t.Fatal(err)
	} else {
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected status 500 got %d", resp.StatusCode)
		}
	}
}
