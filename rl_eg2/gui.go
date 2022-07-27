package main

import (
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// 主界面
type UI struct {
	env      *Maze
	myApp    fyne.App
	myWindow fyne.Window
	canGrid  [][]fyne.CanvasObject
	gridSize float32
}

// 添加网格标签
func genCanGrid(env *Maze) [][]fyne.CanvasObject {
	ans := make([][]fyne.CanvasObject, N+2)
	for i := 0; i < N+2; i++ {
		ans[i] = make([]fyne.CanvasObject, N+2)
	}
	for i := 0; i < N+2; i++ {
		for j := 0; j < N+2; j++ {
			switch env.arr[i][j] {
			case 0:
				ans[i][j] = canvas.NewCircle(color.White)
			case 1:
				ans[i][j] = canvas.NewCircle(color.RGBA{255, 0, 0, 255})
			case 2:
				ans[i][j] = canvas.NewRectangle(color.RGBA{0, 0, 0, 255})
			case 3:
				ans[i][j] = canvas.NewRectangle(color.RGBA{0, 0, 0, 255})
			}
		}
	}
	return ans
}

// 生成
func newUI(env *Maze) UI {
	gridSize := float32(10)
	myApp := app.New()
	myWindow := myApp.NewWindow("MageGame")
	CanGrid := genCanGrid(env)
	ctnSlice := make([]fyne.CanvasObject, (N+2)*(N+2))
	k := 0
	for i := 0; i < N+2; i++ {
		for j := 0; j < N+2; j++ {
			ctnSlice[k] = CanGrid[i][j]
			k++
		}
	}
	// grid := container.NewGridWrap(fyne.NewSize(gridSize, gridSize), ctnSlice...)
	grid := container.NewAdaptiveGrid(N+2, ctnSlice...)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(340, 340))
	return UI{env, myApp, myWindow, CanGrid, gridSize}
}

// 刷新
func (this UI) flush() {
	for i := 0; i < N+2; i++ {
		for j := 0; j < N+2; j++ {
			switch this.env.arr[i][j] {
			case 0:
				this.canGrid[i][j].(*canvas.Circle).FillColor = color.White
				this.canGrid[i][j].Refresh()
			case 1:
				this.canGrid[i][j].(*canvas.Circle).FillColor = color.RGBA{255, 0, 0, 255}
				this.canGrid[i][j].Refresh()
			}
		}
	}
}
