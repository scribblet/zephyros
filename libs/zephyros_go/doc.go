/*
	Example script:

		import (
			"fmt"
			. "./zephyros_go"
		)

		func main() {
			API.Bind("d", []string{"cmd", "shift"}, func() {
				API.Alert("LIKE", 1)

				win := API.FocusedWindow()
				fmt.Println(win.Title())

				f := win.TopLeft()
				f.X += 10
				win.SetTopLeft(f)

				API.ChooseFrom([]string{"foo", "bar"}, "title", 20, 20, func(i int) {
					fmt.Println(i)
				})
			})

			API.Listen("app_launched", func(app App) {
				API.Alert(app.Title(), 1)
			})

			ListenForCallbacks()
		}
*/
package zephyros_go
