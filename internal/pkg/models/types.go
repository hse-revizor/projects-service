package models

type ProjectTemplate string

const (
	ProjectTemplate_StrictEqualityTempl     ProjectTemplate = "StrictEqualityTempl"
	ProjectTemplate_StrictInequalityTempl   ProjectTemplate = "StrictInequalityTempl"
	ProjectTemplate_StrictGreaterTempl      ProjectTemplate = "StrictGreaterTempl"
	ProjectTemplate_StrictShorterThenTempl  ProjectTemplate = "StrictShorterThenTempl"
	ProjectTemplate_StrictMatchesRegexTempl ProjectTemplate = "StrictMatchesRegexTempl"
	ProjectTemplate_NonStrictMathTempl      ProjectTemplate = "NonStrictMathTempl"
)
