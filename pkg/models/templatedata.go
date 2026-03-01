package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      // This is a map that can hold string key-value pairs. It is useful for passing simple data to templates. For example, you might want to pass a message or a title to the template, and you can store that in the StringMap with a descriptive key.
	IntMap    map[string]int         // This is a map that can hold integer key-value pairs. It is useful for passing numeric data to templates. For example, you might want to pass a count of items or a user ID to the template, and you can store that in the IntMap with a descriptive key.
	FloatMap  map[string]float64     // This is a map that can hold any type of data. It is useful for passing complex data structures to templates. For example, you might have a struct that represents a user and you want to pass that struct to the template. You can create a map with a string key and an empty interface value, and then store the user struct in that map. In the template, you can access the fields of the struct using the dot notation. For example, if you have a user struct with a Name field, you can access it in the template using {{index .DataMap "user"}.Name}}.
	DataMap   map[string]interface{} // This is a map that can hold any type of data. It is useful for passing complex data structures to templates. For example, you might have a struct that represents a user and you want to pass that struct
	CSRFToken string                 // This is used to prevent cross-site request forgery attacks. It is a unique token that is generated for each user session and is included in forms that are submitted to the server. The server then checks the token to ensure that the request is coming from a legitimate source.
	Flash     string                 // This is used to display one-time messages to the user. For example, if a user submits a form successfully, you might want to display a success message. You can set the Flash field in the TemplateData struct and then display it in the template. After the message is displayed, it will be cleared from the session.
	Error     string                 // This is used to display error messages to the user. Similar to Flash, you can set the Error field in the TemplateData struct when an error occurs and then display it in the template. This is useful for providing feedback to the user when something goes wrong.
	Warning   string
}
