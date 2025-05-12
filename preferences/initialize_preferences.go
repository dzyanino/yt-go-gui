package preferences

import "fyne.io/fyne/v2"

func InitializePreferences(a fyne.App) {
	var preferencesExist bool = a.Preferences().BoolWithFallback("CONFIGURED", false)
	if !preferencesExist {
		a.Preferences().SetBool("CONFIGURED", false)
		a.Preferences().SetString("FYNE_THEME", "light")
		a.Preferences().SetString("MINIMIZE_ON_CLOSE", "ask")
		a.Preferences().SetString("EXTENSION_ID", "...")
		a.Preferences().SetString("EXTENSION_TOKEN", "...")
		a.Preferences().SetInt("MAX_SIMULTANEOUS_DOWNLOADS", 2)
		a.Preferences().SetInt("MINIMUM_VIDEO_SIZE", 1)
		a.Preferences().SetStringList("BLACK_LIST", []string{"test.com", "another-test.com"})
		a.Preferences().SetBool("OVERWRITE_DUPLICATION", false)
		a.Preferences().SetBool("SINGLE_FOLDER", false)
		a.Preferences().SetString("DESTINATION_FOLDER", "...")
		a.Preferences().SetString("CONFIG_FOLDER", "...")
	}
}
