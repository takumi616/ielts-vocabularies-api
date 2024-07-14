package dto

import "github.com/takumi616/ielts-vocabularies-api/domains"

type VocabDto struct {
	Title        string `json:"title"`
	Definition   string `json:"definition"`
	Example      string `json:"example"`
	PartOfSpeech string `json:"part_of_speech"`
	IsMemorized  string `json:"is_memorized"`
}

func ToDomain(vocabDto *VocabDto) *domains.Vocabulary {
	return &domains.Vocabulary{
		Title:        vocabDto.Title,
		Definition:   vocabDto.Definition,
		Example:      vocabDto.Example,
		PartOfSpeech: vocabDto.PartOfSpeech,
		IsMemorized:  vocabDto.IsMemorized,
	}
}

func FromDomain(vocabulary *domains.Vocabulary) *VocabDto {
	return &VocabDto{
		Title:        vocabulary.Title,
		Definition:   vocabulary.Definition,
		Example:      vocabulary.Example,
		PartOfSpeech: vocabulary.PartOfSpeech,
		IsMemorized:  vocabulary.IsMemorized,
	}
}
