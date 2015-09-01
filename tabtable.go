package main

import (
	"fmt"
	"go-qt5/qt5"
	"strconv"
)

func main() {
	qt5.Main(ui_main)
}

func ui_main() {
	w := qt5.NewWidget()
	defer w.Close()
	w.SetGeometryv(10, 10, 500, 500)

	button2 := qt5.NewButtonWithText("Добавить вкладку")
	button2.SetSizePolicyHV(0, 0)
	var sz qt5.Size
	sz.Height = 300
	sz.Width = 300

	vbox := qt5.NewVBoxLayout()
	vbox.AddWidget(button2)

	table := qt5.NewTableWidget()
	table.SetRowCount(10)
	table.SetColumnCount(10)

	tab := qt5.NewTabWidget()
	var icon qt5.Icon
	tab.AddTab(table, "вкладка", &icon)
	vbox.AddWidget(tab)
	w.SetLayout(vbox)
	num := 1
	//var ti *qt5.TableWidgetItem

	go func() {

		button2.OnClicked(func() {
			lab := strconv.Itoa(num)
			tab.AddTab(qt5.NewFrame(), "вкладка "+lab, &icon)
			num++

		})

		table.OnCellActivated(func(row int, column int) {
			fmt.Printf("row=%d column=%d \n", row, column)
			curit := table.Item(row, column)
			if curit == nil {
				ti := qt5.NewTableWidgetItem()
				ti.SetText("NewItem")
				table.SetItem(row, column, ti)
			} else {
				curit.SetText("ItemExist")
			}
		})

	}()
	w.Show()

	qt5.Run()
}
