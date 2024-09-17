package mysql

import (
	"fmt"
	"github.com/amirazad1/gin-crud/models"
	"github.com/amirazad1/gin-crud/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type BookGRepository struct {
	db *gorm.DB
}

func NewBookGRepository() *BookGRepository {
	mdb, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return &BookGRepository{
		db: mdb,
	}
}

func (repo *BookGRepository) GetAll() (*[]models.Book, error) {
	var items []models.Book
	err := repo.db.Find(&items).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &items, nil
}

func (repo *BookGRepository) GetById(id int64) (*models.Book, error) {
	var item models.Book
	err := repo.db.Where("id=?", id).First(&item).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func (repo *BookGRepository) GetByName(name string) (*[]models.Book, error) {
	var items []models.Book
	err := repo.db.Where("name=?", name).Find(&items).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &items, nil
}

func (repo *BookGRepository) Create(book *models.Book) error {
	err := repo.db.Create(&book).Error
	return err
}

func (repo *BookGRepository) Update(id int64, book *models.Book) error {
	err := repo.db.Model(&models.Book{}).Where("id=?", id).Updates(book).Error
	if err != nil {
		return err
	}
	return err
}

func (repo *BookGRepository) Delete(id int64) error {
	err := repo.db.Where("id=?", id).Delete(&models.Book{}).Error
	if err != nil {
		return err
	}

	return err
}
