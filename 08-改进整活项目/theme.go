package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type myTheme struct {
	font fyne.Resource
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return m.font
}

func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
	return color.RGBA{190, 233, 255, 1}
	case theme.ColorNameButton:
	return color.RGBA{0, 122, 255, 255}
	case theme.ColorNameDisabledButton:
	 return color.RGBA{142, 142, 147, 255}
	case theme.ColorNameHover:
	 return color.RGBA{230, 230, 230, 255}
	case theme.ColorNameFocus:
	 return color.RGBA{255, 165, 0, 255}
	case theme.ColorNameShadow:
	 return color.RGBA{0, 0, 0, 50}
	default:
	 return theme.DefaultTheme().Color(n, v)
	}
   }

func (m *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

