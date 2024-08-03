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
	return theme.DefaultTheme().Color(n, v)
   }

func (m *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	if n == theme.IconNameHome{
		return resourceIconJpeg
	}
	return theme.DefaultTheme().Icon(n)
}

func (m *myTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText{
		return 114
	}
	
	return theme.DefaultTheme().Size(n)
}


