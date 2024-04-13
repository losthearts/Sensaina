package main

import (
	"fmt"
)

func main() {

	var initgame string
	var finalgame string
	var initsens float64
	var initdpi int32
	var finaldpi int32

	fmt.Println("Sensaina 繊細な")
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

	fmt.Printf("Sensitivity in %v : %f\n", finalgame, sens(yaw[initgame], yaw[finalgame], initsens, initdpi, finaldpi))
	fmt.Printf("360/cm : %f\n", cmpi(yaw[initgame], initsens, initdpi))

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
