package main

import (
        "sync"
        "bufio"
        "net/http"
        "fmt"
        "os"
        "strings"
        "io/ioutil"
        "net/url"
        "crypto/x509"
        "crypto/tls"
        "net/http/httputil"
        "time"
)

func main(){
        fmt.Println("frequester tool By tojojo !!")
        fmt.Println("\\__(-_-)__/")

        colorReset := "\033[0m"
        colorRed := "\033[31m"
        colorGreen := "\033[32m"

        sc := bufio.NewScanner(os.Stdin)

        proxyUrl, err := url.Parse("http://127.0.0.1:8080") // Update with your Burp Suite proxy configuration
        if err != nil {
             fmt.Println("Failed to parse proxy URL:", err)
             return
        }

        transport := &http.Transport{
             Proxy: http.ProxyURL(proxyUrl),
        }

        // Load the certificate file
        certFile := "burp_certificate.pem" // Update with the path to your certificate file
        caCert, err := ioutil.ReadFile(certFile)
        if err != nil {
             fmt.Println("Failed to load certificate file:", err)
             return
        }

        // Create a certificate pool and add the certificate
        certPool := x509.NewCertPool()
        certPool.AppendCertsFromPEM(caCert)

        // Set the certificate pool in the transport
        transport.TLSClientConfig = &tls.Config{
             RootCAs: certPool,
        }

        client := &http.Client{
             Transport: transport,
             Timeout:   time.Second * 10, // Set a timeout for the HTTP requests
        }

        // Set the HTTP_PROXY environment variable to Burp Suite's proxy address
        err = os.Setenv("HTTP_PROXY", "http://localhost:8080") // Update with your Burp Suite proxy address
        if err != nil {
             fmt.Println("Failed to set HTTP_PROXY environment variable:", err)
             return
        }

        jobs := make(chan string)
        var wg sync.WaitGroup

        for i := 0; i < 20; i++ {
             wg.Add(1)
             go func() {
                defer wg.Done()
                for domain := range jobs {
                        req, err := http.NewRequest("GET", domain, nil)
                        if err != nil {
                          fmt.Println(string(colorRed), "Failed to create request:", err, string(colorReset))
                          continue
                        }

                        resp, err := client.Do(req)
                        if err != nil {
                          fmt.Println(string(colorRed), "Error:", err, string(colorReset))
                          continue
                        }
                        defer resp.Body.Close()

                        requestDump, err := httputil.DumpRequest(req, true)
                        if err != nil {
                          fmt.Println(string(colorRed), "Failed to dump request:", err, string(colorReset))
                          continue
                        }
                        fmt.Println(string(colorGreen), "Request:", string(requestDump), string(colorReset))

                        responseDump, err := httputil.DumpResponse(resp, true)
                        if err != nil {
                          fmt.Println(string(colorRed), "Failed to dump response:", err, string(colorReset))
                          continue
                        }
                        fmt.Println(string(colorGreen), "Response:", string(responseDump), string(colorReset))

                        body, err := ioutil.ReadAll(resp.Body)
                        if err != nil {
                          fmt.Println(string(colorRed), "Error:", err, string(colorReset))
                          continue
                        }

                        sb := string(body)
                        check_result := strings.Contains(sb , "alert(1)")

                        if check_result {
                          fmt.Println(string(colorRed), "XSS FOUND:", domain, string(colorReset))
                        } else {
                          fmt.Println(string(colorGreen), "Not Vulnerable:", domain, string(colorReset))
                        }
                }
             }()
        }

        for sc.Scan() {
             domain := sc.Text()
             jobs <- domain
        }

        close(jobs)
        wg.Wait()
}
 
