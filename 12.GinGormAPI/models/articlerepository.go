package models

import (
	"github.com/Lexa-san/spc-go2/12.GinGormAPI/storage"
	_ "github.com/lib/pq"
)

func GetAllArticles(a *[]Article) error {
	if err := storage.DB.Find(a).Error; err != nil {
		return err
	}
	return nil
}

func AddNewArticle(a *Article) error {
	if err := storage.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func GetArticleById(a *Article, id string) error {
	if err := storage.DB.Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticleById(a *Article, id string) error {
	if err := storage.DB.Where("id = ?", id).Delete(a).Error; err != nil {
		return err
	}
	return nil
}

func UpdateArticleById(a *Article, id string) error {
	if err := storage.DB.Update(a).Error; err != nil {
		return err
	}
	return nil
}
