package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var initgame string
var finalgame string
var initsens float64
var initdpi int64
var finaldpi int64

func main() {
	if len(os.Args) != 1 {
		if os.Args[1] == "--gui" || os.Args[1] == "-g" {
			gui()
		} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()
		} else {
			cli()
		}
	} else {
		tui()
	}
}

func help() {
	fmt.Printf("Sensaina 繊細な\n")
	fmt.Printf("A tool to convert sensitivities between games\n")

	fmt.Printf("Currently CS:GO, Overwatch, and Valorant is supported.\n\n")
	fmt.Printf("Use without any flags to run the TUI\n\n")

	fmt.Printf("--gui || -g \n")
	fmt.Printf("\tUse this flag to run the Sensaina GUI\n===\n\n")

	fmt.Printf("sensanai --flag value\n")

	fmt.Printf("--from\n")
	fmt.Printf("\tFrom Game.\n")

	fmt.Printf("--to\n")
	fmt.Printf("\tTo Game.\n")

	fmt.Printf("--sens\n")
	fmt.Printf("\tInitial Sensitivity.\n")

	fmt.Printf("--idpi\n")
	fmt.Printf("\tInitial DPI.\n")

	fmt.Printf("--fdpi\n")
	fmt.Printf("\tFinal DPI.\n")
}

func cli() {
	fromGamePtr := flag.String("from", "CS:GO", "From Game")
	toGamePtr := flag.String("to", "CS:GO", "To Game")
	fromSensPtr := flag.Float64("sens", 100, "From Sensitivity")
	fromDPIPtr := flag.Int64("idpi", 800, "Initial DPI")

	flag.Parse()

	toDPIPtr := flag.Int64("fdpi", *fromDPIPtr, "Final DPI")

	flag.Parse()

	sensitivity := sens(yaw[*fromGamePtr], yaw[*toGamePtr], *fromSensPtr, *fromDPIPtr, *toDPIPtr)
	cm360 := cmpi(yaw[*fromGamePtr], *fromSensPtr, *fromDPIPtr)

	fmt.Printf("Sensitivity: %f\n", sensitivity)
	fmt.Printf("Centimeter/360: %f\n\n", cm360)

	fmt.Println("Initial Game: ", *fromGamePtr)
	fmt.Println("Final Game: ", *toGamePtr)
	fmt.Println("From Sensitivity: ", *fromSensPtr)
	fmt.Println("From DPI: ", *fromDPIPtr)
	fmt.Println("To DPI: ", *toDPIPtr)

}

func tui() {

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

func gui() {
	sensanaiApp := app.New()
	formWindow := sensanaiApp.NewWindow("Sensaina 繊細な")
	formWindow.Resize(fyne.NewSize(600, 500))
	resultWindow := sensanaiApp.NewWindow("Conversion")
	resultWindow.Resize(fyne.NewSize(600, 500))

	initgame := widget.NewEntry()
	initsens := widget.NewEntry()
	initDPI := widget.NewEntry()

	finalgame := widget.NewEntry()
	finalDPI := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "The game you want to convert from: ", Widget: initgame},
			{Text: "Sensitivity in it: ", Widget: initsens},
			{Text: "DPI in it: ", Widget: initDPI},
			{Text: "The game you want to convert to: ", Widget: finalgame},
			{Text: "The DPI you want to convert to: ", Widget: finalDPI}},
		OnSubmit: func() {
			fmt.Println("Initial Game:", initgame.Text)
			fmt.Println("Initial Sensitivity:", initsens.Text)
			fmt.Println("Initial DPI:", initDPI.Text)
			fmt.Println("Final Game:", finalgame.Text)
			fmt.Println("Final DPI:", finalDPI.Text)

			gameI := initgame.Text
			gameF := finalgame.Text
			sensI, err := strconv.ParseFloat(initsens.Text, 64)
			DPIin, err := strconv.ParseInt(initDPI.Text, 10, 32)
			DPIou, err := strconv.ParseInt(finalDPI.Text, 10, 32)

			fmt.Println(err)

			sensitivity := sens(yaw[gameI], yaw[gameF], sensI, DPIin, DPIou)
			cm360 := cmpi(yaw[gameI], sensI, DPIin)

			resultWindow.SetContent(widget.NewLabel("Sensitivity:: " + strconv.FormatFloat(sensitivity, 'g', -1, 64) + "\n360/cm ::" + strconv.FormatFloat(cm360, 'g', -1, 64)))

			resultWindow.Show()
			formWindow.Close()
		},
	}

	formWindow.SetContent(form)
	formWindow.ShowAndRun()
}

func cmpi(iyaw float64, sensitivity float64, dpi int64) float64 {
	return (360.0 * 2.54 / (iyaw * sensitivity * float64(dpi)))
}

func sens(iyaw float64, fyaw float64, sensitivity float64, iDPI int64, fDPI int64) float64 {
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
