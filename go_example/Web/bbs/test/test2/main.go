package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	URL := "http://127.0.0.1/admin.php"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	cookies := res.Cookies()
	for _, cookie := range cookies {
		fmt.Println(cookie.String(), "-")

		fmt.Println(cookie.Domain, "-")

		fmt.Println(cookie.Expires, "\n", cookie.HttpOnly, "\n", cookie.MaxAge, "\n", cookie.Name, "\n", cookie.Path, "\n", cookie.Raw, "\n", cookie.RawExpires, "\n", cookie.Secure, "\n", cookie.Unparsed, "\n", cookie.Value)

		fmt.Printf("cookie -> %s\n", cookie.Value)
	}
	// cookies := fmt.Sprint(res.Cookies()[0])
	// cookie := strings.Split(cookies, " ")[0]
	// fmt.Printf("cookies -> %s\n", cookie)
}
