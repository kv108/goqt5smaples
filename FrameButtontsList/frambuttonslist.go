package main

import (
	"fmt"
	"go-qt5/qt5"
)

func main() {
	qt5.Main(ui_main)
}

func ui_main() {
	w := qt5.NewWidget()

	defer w.Close()
	list := qt5.NewListWidget()
	item := qt5.NewListWidgetItem()
	item.SetText("Item1")
	list.AddItem(item)
	list.AddItem(qt5.NewListWidgetItemWithText("Item2"))

	button1 := qt5.NewButton()
	button1.SetText("Нажми меня")
	button2 := qt5.NewButtonWithText("Изменить стиль фрейма")

	var sz qt5.Size
	sz.Height = 35
	sz.Width = 300
	frame := qt5.NewFrame()
	frame.SetMinimumSize(sz)
	button1.SetParent(frame)
	button2.SetParent(frame)
	button1.SetPosv(5, 5)
	button2.SetPosv(100, 5)

	vbox := qt5.NewVBoxLayout()
	vbox.AddWidget(frame)
	vbox.AddWidget(list)

	w.SetLayout(vbox)
	style := 1
	go func() {
		button1.OnClicked(func() {
			fmt.Println("Clicked")
			button1.SetText("Hello")
			list.AddItem(qt5.NewListWidgetItemWithText("Привет"))
		})
		button2.OnClicked(func() {
			fmt.Println(style)
			frame.SetFrameStyle(style)
			style++
			if style == 7 {
				style = 0
			}
		})
		list.OnCurrentItemChanged(func(item, old *qt5.ListWidgetItem) {
			//go func() {
			fmt.Println(item.Attr("text"))
			//}()
		})

	}()
	w.Show()

	qt5.Run()
}
