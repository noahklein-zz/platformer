package main

// Config represents game configuration properties.
type Config struct {
	width     int
	height    int
	framerate int
}

func getConfig() Config {
	return Config{
		width:     800,
		height:    600,
		framerate: 60,
	}
}
