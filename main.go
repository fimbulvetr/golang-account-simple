package main

import "os"
import "fmt"
import "gopkg.in/macaron.v1"
import "golang-account-simple/repository"

func main() {
	fmt.Println("Hello, world!")
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {

		fmt.Println("getting engine...")

		var dbUser string = os.Getenv("ACCOUNT_DB_USER")
		var dbPass string = os.Getenv("ACCOUNT_DB_PASS")
		var dbHost string = os.Getenv("ACCOUNT_DB_HOST")
		var dbName string = os.Getenv("ACCOUNT_DB_NAME")

		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?%scharset=utf8&parseTime=true", dbUser, dbPass, dbHost, dbName, "")

		r, err := repository.GetEngine(connStr)
		if err != nil {
			fmt.Println(err)
		}


		fmt.Println("making accounts")

		accounts := make([]repository.Account, 0)
		err = r.Find(&accounts)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		if len(accounts) > 0 {
			fmt.Println("Congratulations, we found your account.")
			ctx.Data["Id"] = accounts[0].Id
			ctx.Data["Username"] = accounts[0].Username
			ctx.Data["Type"] = "Unknown"
			if accounts[0].Type == 1 {
				ctx.Data["Type"] = "Primary"
			}
			ctx.HTML(200, "hello") // 200 is the response code.
		} else {
			fmt.Println("No account found.")
			ctx.HTML(200, "hello") // 200 is the response code.
		}
	})
	m.Run()
}
