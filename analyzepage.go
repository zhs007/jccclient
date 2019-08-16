package jccclient

// Viewport - viewport
type Viewport struct {
	Width             int     // viewport width
	Height            int     // viewport height
	DeviceScaleFactor float32 // device scale factor
	IsMobile          bool    // this is a mobile
	IsLandscape       bool    // is landscape
}

// AnalyzePageOptions - AnalyzePage Options
type AnalyzePageOptions struct {
	NeedScreenshots  bool // I want screenshots
	NeedLogs         bool // I want full logs
	Timeout          int  // timeout, default is 3 * 60 second
	ScreenshotsDelay int  // delay second before screenshot
}
