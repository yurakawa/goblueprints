package thesaurus

type Thesaurus interface {
	// term: 検索語
	Synonyms(term string) ([]string, error)
}
