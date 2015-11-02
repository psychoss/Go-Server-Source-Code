package server
const (
	KEY_COLLECTOR = "-keycollotor"
	// For index.html carousel
	PREFIX_INDEX = "index-"
	// For index.hmtl Script
	PREFIX_SCRIPT           = "index-script-"
	DBPATH                  = "/home/psycho/db"
	CACHE_PATH              = "/home/psycho/public/cache/"
	LAYOUT_TEMPLATE_FILE    = "/templates/layout.html"
	ARTICAL_TEMPLATE_FILE   = "/templates/doc.html"
	PUBLIC_PATH             = "/home/psycho/public"
	STATIC_PATH             = "/home/psycho/public/dist"
	INDEX_TEMPLATE_FILE     = "/templates/index.html"
	TEMPLATE_FILE_NOT_FOUND = "the target template file not found!"
	ERROR_NOT_PERMISSION    = "You don't currently have permission to post in this article."
	ERROR_BAD_REQUEST       = "You must provide the certification."
	ERROR_SERVER_INTERNAL   = "Sorry, An internal Server Error occurred."
	BASE_URL                = "http://localhost:9091/"
	//BASE_URL = "http://www.mean101.com/"
	TITLE       = "炫酷的网站技术 - GO+"
	KEYWORD     = "Go语言，SVG 动画，Canvas 动画，MongoDB，RethinkDB，Node.js"
	DESCRIPTION = "Go语言，炫酷动画，建站技术，尽在Go+。"
	// the delimter for the content
	DELIMITER = "^^^"
	KEY_ADMIN_CERTIFICATION_TOKEN = "admin-account"
)
// go run /home/psycho/go/src/web/server/const.go
