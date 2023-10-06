package endpointcount

type FormatterGetCount struct {
	Endpoint        string `json:"endpoint"`
	Count           int    `json:"count"`
	UniqueUserAgent int    `json:"unique_user_agent"`
}

func FormatterGetCounts(count Statistics) FormatterGetCount {
	formatterCreate := FormatterGetCount{
		Endpoint:        count.Endpoint,
		Count:           count.Count,
		UniqueUserAgent: count.UniqueUserAgent,
	}
	return formatterCreate
}

func FormatterGet(counts []Statistics) []FormatterGetCount {
	countGetFormatter := []FormatterGetCount{}

	for _, count := range counts {
		countFormatter := FormatterGetCounts(count)
		countGetFormatter = append(countGetFormatter, countFormatter)
	}

	return countGetFormatter
}
