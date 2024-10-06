package tags

import (
	"log"

	tagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
	"github.com/dot-slash-ann/home-services-api/entities/tags"
	"gorm.io/gorm"
)

type TagsService interface {
	Create(tagsDto.CreateTagDto) (tags.Tag, error)
	FindAll() ([]tags.Tag, error)
	FindOne(string) (tags.Tag, error)
	Update(string, tagsDto.UpdateTagDto) (tags.Tag, error)
	Delete(string) (tags.Tag, error)
	FindOneOrCreate(string) (tags.Tag, error)
}

type TagsServiceImpl struct {
	database *gorm.DB
}

func NewTagsService(database *gorm.DB) *TagsServiceImpl {
	return &TagsServiceImpl{
		database: database,
	}
}

func (service *TagsServiceImpl) Create(createTagDto tagsDto.CreateTagDto) (tags.Tag, error) {
	tag := tags.Tag{
		Name: createTagDto.Name,
	}

	if result := service.database.Create(&tag); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) FindAll() ([]tags.Tag, error) {
	var tagList []tags.Tag

	if results := service.database.Find(&tagList); results.Error != nil {
		return []tags.Tag{}, results.Error
	}

	return tagList, nil
}

func (service *TagsServiceImpl) FindOne(id string) (tags.Tag, error) {
	var tag tags.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) FindOneOrCreate(name string) (tags.Tag, error) {
	var tag tags.Tag

	if result := service.database.FirstOrCreate(&tag, tags.Tag{Name: name}); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	log.Default().Println("service FindOneOrCreate - ", tag.Name, tag.ID)

	return tag, nil
}

func (service *TagsServiceImpl) Update(id string, updateTagDto tagsDto.UpdateTagDto) (tags.Tag, error) {
	var tag tags.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	if result := service.database.Model(&tag).Updates(tags.Tag{
		Name: updateTagDto.Name,
	}); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	return tag, nil
}

func (service *TagsServiceImpl) Delete(id string) (tags.Tag, error) {
	var tag tags.Tag

	if result := service.database.First(&tag, id); result.Error != nil {
		return tags.Tag{}, result.Error
	}

	service.database.Delete(&tags.Tag{}, id)

	return tag, nil
}
