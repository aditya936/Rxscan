# Rxscan
this tool help to find reflected xss when you have huge llist of domains .This is a simple tool for sending HTTP requests to domains and checking for potential Cross-Site Scripting (XSS) vulnerabilities. It utilizes goroutines for concurrency and handles request and response details.
connect with burpsuite proxy at 127.0.0.1:8080 to check resquest and response in logger or http-history.
 
* in simple word we can say that it track our paylaod is reflecting in response
 
ex: xss-payload-embeded-urls.txt | freq.go 


-----------DESCRIPTION---------------
The code you provided is a Go programming language code that implements a tool called "frequester." Here's an overview of what the code does:

1. Imports necessary packages for HTTP communication, concurrency, file handling, and TLS.
2. Sets up the configuration for the Burp Suite proxy, including parsing the proxy URL and loading a certificate file.
3. Creates an HTTP client with the Burp Suite proxy configuration and timeout settings.
4. Sets the HTTP_PROXY environment variable to point to the Burp Suite proxy address.
5. Defines a number of worker goroutines to handle concurrent HTTP requests.
6. Reads input from stdin (standard input) and sends each line (domain) to a jobs channel.
7. Each worker goroutine reads from the jobs channel, creates an HTTP GET request for the domain, sends the request using the client, and processes the response.
8. The response is dumped (request and response details are printed) and checked for the presence of the string "alert(1)" in the response body.
9. If the string is found, it prints "XSS FOUND" for the domain; otherwise, it prints "Not Vulnerable."
10. The main goroutine reads input from stdin and sends it to the jobs channel until the input ends.
11. Once all the jobs are sent to the channel, it is closed, and the main goroutine waits for all worker goroutines to finish using the sync.WaitGroup.



