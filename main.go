package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*This is the first set of comments just briefly organizing what should go on in this application

1. Have a functioning UI
2. Be split into components and functions each with their own tests based on what the best practice for Go is
3. The user first of all enters the form which gets encrypted
4. It gets passed to the AI which takes in the form
5. The use then passes in the data which it should expect
6. The AI will encrypt that
7. Give multiple options when the user logs in. They can choose what form they are entering information for
8. They upload the filled in form as a pdf.
9. This again gets encrypted the same way
10. The AI then reads this, processes the form and explains if any fields are missing, if any data is incorrect, or if it requires human review.
11. The AI will be deepseek running locally
12. Connect to postgres to store the data and the form types, train the AI based on the form encryption
13. Use Go for frontent, use Python for ML
*/
func main() {
    a := app.New()
    w := a.NewWindow("GO Form")
    w.Resize(fyne.NewSize(300,300))
    w.SetContent(container.NewVBox(
        widget.NewLabel("Hello, World!"),
        widget.NewButton("Click me", func() {
            println("Button clicked!")
        }),
    ))
    w.ShowAndRun()
}


