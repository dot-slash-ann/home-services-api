package tags

import (
	"log"

	tagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
	"github.com/dot-slash-ann/home-services-api/entities"
	"gorm.io/gorm"
)

type TagsService interface {
	Create(tagsDto.CreateTagDto) (entities.Tag, error)
	FindAll() ([]entities.Tag, error)
	FindOne(string) (entities.Tag, error)
	Update(string, tagsDto.UpdateTagDto) (entities.Tag, error)
	Delete(string) (entities.Tag, error)
	FindOneOrCreate(string) (entities.Tag, error)
}

type TagsServiceImpl struct {
	database *gorm.DB
}

func NewTagsService(database *gorm.DB) *TagsServiceImpl {
	return &TagsServiceImpl{
		database: database,
	}
}

func (service *TagsServiceImpl) Create(createTagDto tagsDto.CreateTagDto) (entities.Tag, error) {
	tag := entities.Tag{
		Name: createTagDto.Name,
	}

	if result := service.database.Create(&tag); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) FindAll() ([]entities.Tag, error) {
	var tagList []entities.Tag

	if results := service.database.Find(&tagList); results.Error != nil {
		return []entities.Tag{}, results.Error
	}

	return tagList, nil
}

func (service *TagsServiceImpl) FindOne(id string) (entities.Tag, error) {
	var tag entities.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) FindOneOrCreate(name string) (entities.Tag, error) {
	var tag entities.Tag

	if result := service.database.FirstOrCreate(&tag, entities.Tag{Name: name}); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	log.Default().Println("service FindOneOrCreate - ", tag.Name, tag.ID)

	return tag, nil
}

func (service *TagsServiceImpl) Update(id string, updateTagDto tagsDto.UpdateTagDto) (entities.Tag, error) {
	var tag entities.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	if result := service.database.Model(&tag).Updates(entities.Tag{
		Name: updateTagDto.Name,
	}); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) Delete(id string) (entities.Tag, error) {
	var tag entities.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return entities.Tag{}, result.Error
	}

	service.database.Delete(&entities.Tag{}, id)

	return tag, nil
}
