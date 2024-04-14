package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func main() {

	var initgame string
	var finalgame string
	var initsens float64
	var initdpi int32
	var finaldpi int32

	b := lipgloss.NewStyle().Background(lipgloss.Color("#000")).Bold(true).Italic(true).Underline(true).Foreground(lipgloss.Color("#FDB0C0")).Render
	v := lipgloss.NewStyle().Foreground(lipgloss.Color("#FDB0C0")).Render
	k := lipgloss.NewStyle().Foreground(lipgloss.Color("#DFCDD0")).Render
	t := table.New()

	fmt.Println(b("Sensaina (繊細な)"))
	fmt.Print("What game are you converting from: ")
	fmt.Scanln(&initgame)
	fmt.Print("Sensitivity in that game: ")
	fmt.Scanln(&initsens)
	fmt.Print("DPI in that game: ")
	fmt.Scanln(&initdpi)

	fmt.Print("What game do you want to convert to: ")
	fmt.Scanln(&finalgame)
	fmt.Print("DPI in that game: ")
	fmt.Scanln(&finaldpi)

	sensitivity := sens(yaw[initgame], yaw[finalgame], initsens, initdpi, finaldpi)
	cm360 := cmpi(yaw[initgame], initsens, initdpi)

	t.Row(k("Game: "), v(finalgame))
	t.Row(k("Sensitivity: "), v(fmt.Sprintf("%f", sensitivity)))
	t.Row(k("Centimeter/360: "), v(fmt.Sprintf("%f", cm360)))
	fmt.Println(t.Render())
}

func cmpi(iyaw float64, sensitivity float64, dpi int32) float64 {
	return (360.0 * 2.54 / (iyaw * sensitivity * float64(dpi)))
}

func sens(iyaw float64, fyaw float64, sensitivity float64, iDPI int32, fDPI int32) float64 {
	return (sensitivity * iyaw * float64(iDPI) / (fyaw * float64(fDPI)))
}

var yaw = map[string]float64{
	"Overwatch": 0.0066,
	"OW":        0.0066,
	"OW2":       0.0066,
	"CS:GO":     0.022,
	"CSGO":      0.022,
	"CS":        0.022,
	"Valorant":  0.07,
	"Valo":      0.07,
	"Val":       0.07,
}
