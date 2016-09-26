package main 

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main(){
    hqyzm()
    fmt.Println()
    cyfp()

}

func hqyzm() {

    targetUrl := "https://inv-veri.chinatax.gov.cn/WebQuery/do/1100/yzmQuery"
    
    var resp *http.Response
    var err error
    var data []byte
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //忽略证书签名
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}

    resp, err = client.PostForm(targetUrl, nil)
    defer  resp.Body.Close()

    if data, err = ioutil.ReadAll(resp.Body); err == nil {
        fmt.Printf("%s\n", data)
    }

}

func cyfp() {

    targetUrl := "https://inv-veri.chinatax.gov.cn/WebQuery/do/1100/query"

    var resp *http.Response
    var err error
    var data []byte
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //忽略证书签名
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}

/*
param = {
    'fpdm': fpdm,
    'fphm': fphm,
    'kprq': kprq.replace(new RegExp('-', 'gm'), ''),
    'fpje': kjje,
    'fplx': fplx,
    'yzm': yzm,
    'yzmSj': yzmSj,
    'loginSj': loginSj,
    'token': token,
    'username': username,
    'index': jmmy
};
*/

    param := url.Values{//map[string][]string
            "fpdm": {"1200143320"},
            "fphm": {"00724"},
            "kprq": {"20141001"},
            "fpje": {"16.00"},
            "fplx": {"04"},
            "yzm": {"1111"},
            "yzmSj": {"11"},
            "loginSj": {"22"},
            "token": {"33"},
            "username": {"44"},
            "index": {"55"},
    };

    resp, err = client.PostForm(targetUrl,param)
    defer  resp.Body.Close()

    if data, err = ioutil.ReadAll(resp.Body); err == nil {
        fmt.Printf("%s\n", data)
    }

}