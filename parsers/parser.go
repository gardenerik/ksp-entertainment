package parsers

type URLParser interface {
	CanParse(url string) bool
	GetStreamURL(url string) string
	GetName(url string) string
}

var parsers = []URLParser{
	YoutubeDLParser{},
}

func ParseURL(url string) string {
	for _, parser := range parsers {
		if parser.CanParse(url) {
			return parser.GetStreamURL(url)
		}
	}

	return url
}

func GetName(url string) string {
	for _, parser := range parsers {
		if parser.CanParse(url) {
			return parser.GetName(url)
		}
	}

	return "(unknown)"
}
