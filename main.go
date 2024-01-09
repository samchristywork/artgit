package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/radovskyb/watcher"
	"time"
)

func perform_git_init() {
	_, err := git.PlainInit("./test", false)
	if err != nil {
		panic(err)
	}
}

func perform_git_add_all() {
	r, err := git.PlainOpen("./test")
	if err != nil {
		panic(err)
	}
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	_, err = w.Add(".")
	if err != nil {
		panic(err)
	}
}

func perform_git_status() {
	r, err := git.PlainOpen("./test")
	if err != nil {
		panic(err)
	}
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	status, err := w.Status()
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}

func perform_git_commit(message string) {
	r, err := git.PlainOpen("./test")
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	author := &object.Signature{
		Name:  "ArtGit",
		Email: "ArtGit@example.com",
		When:  time.Now(),
	}
	_, err = w.Commit(message, &git.CommitOptions{
		Author: author,
	})
}

func get_git_log() object.CommitIter {
	r, err := git.PlainOpen("./test")
	if err != nil {
		panic(err)
	}

	ref, err := r.Head()
	if err != nil {
		panic(err)
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	return cIter
}

func watch_directory(callback func()) {
	w := watcher.New()
	w.FilterOps(watcher.Write, watcher.Create, watcher.Remove, watcher.Rename)
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
				callback()
			case err := <-w.Error:
				panic(err)
			case <-w.Closed:
				return
			}
		}
	}()
	err := w.Add("./test")
	if err != nil {
		panic(err)
	}
	err = w.Start(time.Millisecond * 100)
	if err != nil {
		panic(err)
	}
}

func main() {
	app := app.New()
	window := app.NewWindow("GUI Application")

	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {}),
		fyne.NewMenuItem("Open", func() {}),
		fyne.NewMenuItem("Save", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(fileMenu)
	window.SetMainMenu(mainMenu)

	dirtyIndicator := widget.NewLabel("Clean")
	dirtyIndicator.TextStyle.Bold = true
	dirtyIndicator.TextStyle.Italic = true
	dirtyIndicator.TextStyle.Monospace = true

	textBox := widget.NewEntry()

	button := widget.NewButton("Click Me", func() {
		fmt.Println("Button Clicked")
	})

	stringList := make([]string, 0)
	for i := 0; i < 100; i++ {
		stringList = append(stringList, fmt.Sprintf("Item %d", i))
	}

	list := widget.NewList(
		func() int {
			return len(stringList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(stringList[i])
		},
	)

	listScrollContainer := container.NewVScroll(list)

	content := container.NewBorder(
		container.NewVBox(dirtyIndicator, textBox, button),
		nil, nil, nil,
		container.New(layout.NewMaxLayout(), listScrollContainer),
	)
	window.SetContent(content)

	window.ShowAndRun()
}
