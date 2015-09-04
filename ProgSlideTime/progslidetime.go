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
	w.SetGeometryv(10, 10, 500, 200)

	chkbox := qt5.NewCheckBoxWithText("Таймер вкл/выкл")
	chkbox.SetSizePolicyHV(0, 0)
	timer := qt5.NewTimer()
	timer.SetInterval(100)
	timelab := qt5.NewLabelWithText("0")
	timelab.SetFrameStyle(1)

	timelab.SetSizePolicyHV(0, 0)

	progress := qt5.NewScrollBar()
	progress.SetRange(0, 100)
	progress.SetEnabled(false)

	slide1 := qt5.NewSlider()
	slide1.SetRange(0, 100)

	label1 := qt5.NewLabelWithText("0    ")
	label1.SetSizePolicyHV(0, 0)
	//chkbox.SetSizePolicyHV(0, 0)
	//var sz qt5.Size 	sz.Height = 300 	sz.Width = 300
	hbox := qt5.NewHBoxLayout()
	vbox := qt5.NewVBoxLayout()

	hbox.AddWidget(progress)
	hbox.AddWidget(chkbox)
	hbox.AddWidget(timelab)

	vbox.AddLayout(hbox)
	vbox.AddWidget(slide1)
	vbox.AddWidget(label1)
	w.SetLayout(vbox)
	num := 0

	go func() {

		chkbox.OnStateChanged(func(st int) {
			fmt.Printf("State changed on %d \n", st)
			if st == 2 {
				timer.Start()
			}
			if st == 0 {
				timer.Stop()
			}

		})
		timer.OnTimeout(func() {
			fmt.Printf("Time=%d \n", num)
			timelab.SetText(strconv.Itoa(num))
			num = num + 1
			progress.SetSliderPosition(num)

			if num == 100 {
				num = 0
			}
		})
		slide1.OnValueChanged(func(val int) {
			label1.SetText(strconv.Itoa(val))
		})

	}()
	w.Show()

	qt5.Run()
}
