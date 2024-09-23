package TagsService

import (
	"github.com/dot-slash-ann/home-services-api/database"
	TagsDto "github.com/dot-slash-ann/home-services-api/dtos/tags"
	TagsEntity "github.com/dot-slash-ann/home-services-api/entities/tags"
)

func Create(createTagDto TagsDto.CreateTagDto) (TagsEntity.Tag, error) {
	tag := TagsEntity.Tag{
		Name: createTagDto.Name,
	}

	if result := database.Connection.Create(&tag); result.Error != nil {
		return TagsEntity.Tag{}, result.Error
	}

	return tag, nil
}

func FindAll() ([]TagsEntity.Tag, error) {
	var tags []TagsEntity.Tag

	if results := database.Connection.Find(&tags); results.Error != nil {
		return []TagsEntity.Tag{}, results.Error
	}

	return tags, nil
}

func FindOne(id string) (TagsEntity.Tag, error) {
	var tag TagsEntity.Tag

	if result := database.Connection.First(&tag, id); result.Error != nil {
		return TagsEntity.Tag{}, result.Error
	}

	return tag, nil
}

func Update(id string, updateTagDto TagsDto.UpdateTagDto) (TagsEntity.Tag, error) {
	var tag TagsEntity.Tag

	if result := database.Connection.First(&tag, id); result.Error != nil {
		return TagsEntity.Tag{}, result.Error
	}

	if result := database.Connection.Model(&tag).Updates(TagsEntity.Tag{
		Name: updateTagDto.Name,
	}); result.Error != nil {
		return TagsEntity.Tag{}, result.Error
	}

	return tag, nil
}

func Delete(id string) (TagsEntity.Tag, error) {
	var tag TagsEntity.Tag

	if result := database.Connection.First(&tag, id); result.Error != nil {
		return TagsEntity.Tag{}, result.Error
	}

	database.Connection.Delete(&TagsEntity.Tag{}, id)

	return tag, nil
}
