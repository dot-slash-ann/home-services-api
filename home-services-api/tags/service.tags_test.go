package tags_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/dot-slash-ann/home-services-api/tags"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to in-memory SQLite database: %v", err)
	}

	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Tag{})
	db.AutoMigrate(&entities.Transaction{})
	db.AutoMigrate(&entities.User{})

	return db
}

func getService() (tags.TagsService, *gorm.DB) {
	db := setupTestDB()
	tagsService := tags.NewTagsService(db)

	return tagsService, db
}

func create(t *testing.T, service tags.TagsService, names []string) []entities.Tag {
	tagsList := make([]entities.Tag, len(names))

	for _, name := range names {
		createTagDto := tags.CreateTagDto{
			Name: name,
		}

		tag, err := service.Create(createTagDto)

		if err != nil {
			t.Errorf("failed creating tags: %v", err)
		}

		tagsList = append(tagsList, tag)
	}

	return tagsList
}

func TestTagsServiceCreate(t *testing.T) {
	tagsService, db := getService()

	createTagDto := tags.CreateTagDto{
		Name: "mock tag",
	}

	tag, err := tagsService.Create(createTagDto)

	var count int64
	db.Model(&entities.Tag{}).Count(&count)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, "mock tag", tag.Name)
	assert.Equal(t, int64(1), count)
}

func TestTagsServiceFindAll(t *testing.T) {
	tagsService, _ := getService()

	names := make([]string, 0, 5)

	for i := 0; i < 5; i++ {
		names = append(names, fmt.Sprintf("mock tag %v", i))
	}

	create(t, tagsService, names)

	tagsList, err := tagsService.FindAll()

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, 5, len(tagsList))

	for i, tag := range tagsList {
		assert.Equal(t, fmt.Sprintf("mock tag %v", i), tag.Name)
	}
}

func TestTagsServiceFindOne(t *testing.T) {
	tagsService, _ := getService()

	create(t, tagsService, []string{"mock tag"})

	tag, err := tagsService.FindOne("1")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)

		assert.Equal(t, uint(1), tag.ID)
		assert.Equal(t, "mock tag", tag.Name)
	}
}

func TestTagsServiceFindOneOrCreate(t *testing.T) {
	tagsService, db := getService()

	create(t, tagsService, []string{"mock tag"})

	tag, err := tagsService.FindOne("1")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, uint(1), tag.ID)
	assert.Equal(t, "mock tag", tag.Name)

	tag2, err := tagsService.FindOneOrCreate("missing tag")

	log.Default().Println(tag2.Name)

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	assert.Equal(t, uint(2), tag2.ID)
	assert.Equal(t, "missing tag", tag2.Name)

	var count int64

	db.Model(&entities.Tag{}).Count(&count)

	assert.Equal(t, int64(2), count)
}

func TestTagsServiceUpdate(t *testing.T) {
	tagsService, _ := getService()

	create(t, tagsService, []string{"old name"})

	tag, _ := tagsService.FindOne("1")

	assert.Equal(t, "old name", tag.Name)

	tag, err := tagsService.Update("1", tags.UpdateTagDto{
		Name: "new name",
	})

	assert.Nil(t, err)

	tag, _ = tagsService.FindOne("1")

	assert.Equal(t, "new name", tag.Name)
}

func TestTagsServiceDelete(t *testing.T) {
	tagsService, _ := getService()

	create(t, tagsService, []string{"to be deleted"})

	tag, err := tagsService.Delete("1")

	assert.Nil(t, err)
	assert.Equal(t, "to be deleted", tag.Name)

	tag, err = tagsService.FindOne("1")

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}
