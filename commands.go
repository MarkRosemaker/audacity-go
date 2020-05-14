package audacity

// Import imports a file from the given path.
func Import(path string) {
	Command("Import2", "Filename", path)
}

// NoiseReduction applies noise reduction with default settings.
func NoiseReduction() {
	Command("NoiseReduction")
}

// RemoveTracks removes all tracks, allowing you to start over.
func RemoveTracks() {
	Command("RemoveTracks")
}

// SelectAll selects everything, like CTRL + A.
func SelectAll() {
	Command("SelectAll")
}

// Normalize applies the normalize function with standard parameters.
func Normalize() {
	Command("Normalize")
}

// Export exports the Audacity project to the given file path.
func Export(path string) {
	//  NumChannels="1"
	Command("Export2", "Filename", path)
}
