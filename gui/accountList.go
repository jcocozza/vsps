package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeAccountList(appState *VspsAppState) *widget.List {
  list := widget.NewListWithData(appState.BoundAccountList,
    func() fyne.CanvasObject {
      return widget.NewLabel("template")
    },
    func(i binding.DataItem, o fyne.CanvasObject) {
      o.(*widget.Label).Bind(i.(binding.String))
    })
  return list
}

func makeAccountsViewer(appState *VspsAppState, window fyne.Window) *container.Split {
  var split *container.Split
  
  acctsList := makeAccountList(appState)
  acctsList.OnSelected = func(id widget.ListItemID) {
    name, _ := appState.BoundAccountList.GetValue(id)
    currentAcct, err := appState.Accounts.Get(name)
    if err != nil {
      return
    }
    split.Trailing = makeAccountViewer(currentAcct, appState, window)
    split.Refresh()
  }

  defal := widget.NewLabel("Select Account")
  new := widget.NewButton("new", func() {
    dial := makeNewAccountDialog(window, appState)
    dial.Show()
  })

  split = container.NewHSplit(
    container.NewBorder(
      widget.NewLabel("My Accounts"),
      new,
      nil,
      nil,
      acctsList,
    ),
    defal,
  )
  return split
}
