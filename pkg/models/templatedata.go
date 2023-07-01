package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // for any data type
	CSRFToken string                 // cross site request forgery token
	Flash     string                 // flash message
	Warning   string                 // warning message
	Error     string                 // error message

}
