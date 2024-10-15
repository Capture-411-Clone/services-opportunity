package openapi

/*
 * @apiEnum: GenderEnum
 */
type GenderEnum struct {
	Select string `json:"select" openapi:"enumValue:select"`
	Male   string `json:"male" openapi:"enumValue:male"`
	Female string `json:"female" openapi:"enumValue:female"`
}

/*
 * @apiEnum: FilterOperatorsEnum
 */
type FilterOperatorsEnum struct {
	Equals            string `json:"equals" openapi:"enumValue:equals"`
	Contains          string `json:"contains" openapi:"enumValue:contains"`
	StartsWith        string `json:"startsWith" openapi:"enumValue:startsWith"`
	EndsWith          string `json:"endsWith" openapi:"enumValue:endsWith"`
	IsEmpty           string `json:"isEmpty" openapi:"enumValue:isEmpty"`
	IsNotEmpty        string `json:"isNotEmpty" openapi:"enumValue:isNotEmpty"`
	IsAnyOf           string `json:"isAnyOf" openapi:"enumValue:isAnyOf"`
	NumberEquals      string `json:"=" openapi:"enumValue:="`
	NumberNotEquals   string `json:"!=" openapi:"enumValue:!="`
	GreaterThan       string `json:">" openapi:"enumValue:>"`
	GreaterThanEquals string `json:">=" openapi:"enumValue:>="`
	LessThan          string `json:"<" openapi:"enumValue:<"`
	LessThanEquals    string `json:"<=" openapi:"enumValue:<="`
	DateIs            string `json:"is" openapi:"enumValue:is"`
	DateIsNot         string `json:"isNot" openapi:"enumValue:isNot"`
}

/*
 * @apiEnum: ProjectStatusesEnum
 */
type ProjectStatusesEnum struct {
	SubttileProcessing string `json:"subtitle_processing" openapi:"enumValue:subtitle_processing"`
	SubtitleDone       string `json:"subtitle_done" openapi:"enumValue:subtitle_done"`
}
