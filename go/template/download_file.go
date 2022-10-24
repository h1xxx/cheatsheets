package main

import (
        "net/http"
        "os"
        "path"
)

func main() {
        ua := "Wget/1.20.3 (linux-gnu)"
        url := "https://sourceforge.net/projects/procps-ng/files/Production/procps-ng-3.3.17.tar.xz"
        client := &http.Client{
                CheckRedirect: func(req *http.Request, via []*http.Request) error {
                        for k, _ := range req.Header {
                                delete(req.Header, k)
                        }
                        req.Header.Set("User-Agent", ua)
                        req.Header.Set("Accept", "*/*")
                        req.Header.Set("Connection", "Keep-Alive")

                        return nil
                },
        }
        req, _ := http.NewRequest("GET", url, nil)
        for k, _ := range req.Header {
                delete(req.Header, k)
        }
        req.Header.Set("User-Agent", ua)
        req.Header.Set("Accept", "*/*")
        req.Header.Set("Connection", "Keep-Alive")

        r, e := client.Do(req)
        if e != nil {
                panic(e)
        }
        defer r.Body.Close()
        f, e := os.Create(path.Base(url))
        if e != nil {
                panic(e)
        }
        defer f.Close()
        f.ReadFrom(r.Body)
}
