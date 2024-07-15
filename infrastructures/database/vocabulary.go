package database

import (
	"context"
	"log"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Vocabulary struct {
	gorm.Model
	Title        string
	Definition   string
	Example      string
	PartOfSpeech string
	IsMemorized  string
}

// Insert new vocabulary data
func (g *Gorm) InsertNewVocabulary(ctx context.Context, vocabDto *dto.VocabDto) (uint, error) {
	vocabulary := &Vocabulary{
		Title:        vocabDto.Title,
		Definition:   vocabDto.Definition,
		Example:      vocabDto.Example,
		PartOfSpeech: vocabDto.PartOfSpeech,
		IsMemorized:  vocabDto.IsMemorized,
	}

	result := g.Db.Create(vocabulary)
	if result.Error != nil {
		log.Printf("failed to insert new vocabulary: %v", result.Error)
		return 0, result.Error
	} else {
		return vocabulary.ID, nil
	}
}

func (g *Gorm) SelectVocabularyById(ctx context.Context, vocabularyID uint) (*dto.VocabDto, error) {
	selected := Vocabulary{}
	result := g.Db.First(&selected, vocabularyID)
	vocabDto := &dto.VocabDto{}
	if result.Error != nil {
		log.Printf("failed to select a vocabulary by ID: %v", result.Error)
		return vocabDto, result.Error
	}

	vocabDto.Title = selected.Title
	vocabDto.Definition = selected.Definition
	vocabDto.Example = selected.Example
	vocabDto.PartOfSpeech = selected.PartOfSpeech
	vocabDto.IsMemorized = selected.IsMemorized

	return vocabDto, nil
}

func (g *Gorm) UpdateVocabularyById(ctx context.Context, vocabularyID uint, vocabDto *dto.VocabDto) (uint, error) {
	//convert dto into db model
	vocabulary := &Vocabulary{
		Title:        vocabDto.Title,
		Definition:   vocabDto.Definition,
		Example:      vocabDto.Example,
		PartOfSpeech: vocabDto.PartOfSpeech,
		IsMemorized:  vocabDto.IsMemorized,
	}

	selected := Vocabulary{}
	result := g.Db.First(&selected, vocabularyID)
	if result.Error != nil {
		log.Printf("failed to select a record that is going to update: %v", result.Error)
		return 0, result.Error
	}

	selected.Title = vocabulary.Title
	selected.Definition = vocabulary.Definition
	selected.Example = vocabulary.Example
	selected.PartOfSpeech = vocabulary.PartOfSpeech
	selected.IsMemorized = vocabulary.IsMemorized

	result = g.Db.Save(&selected)
	if result.Error != nil {
		log.Printf("failed to update a selected record: %v", result.Error)
		return 0, result.Error
	} else {
		return selected.ID, nil
	}
}

func (g *Gorm) DeleteVocabularyById(ctx context.Context, vocabularyID uint) (uint, error) {
	var deleted Vocabulary
	result := g.Db.Clauses(clause.Returning{}).Where("ID = ?", vocabularyID).Delete(&deleted)
	if result.Error != nil {
		log.Printf("failed to delete a vocabulary by id: %v", result.Error)
		return 0, result.Error
	} else {
		return deleted.ID, nil
	}
}
