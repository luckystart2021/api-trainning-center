package seo

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type Seo struct {
	Id                 int    `json:"id"`
	Description        string `json:"description"`
	Keywords           string `json:"keywords"`
	FbAppId            string `json:"fb_app_id"`
	OgTitle            string `json:"og_title"`
	OgUrl              string `json:"og_url"`
	OgImage            string `json:"og_image"`
	OgDescription      string `json:"og_description"`
	OgSiteName         string `json:"og_site_name"`
	OgSeeAlso          string `json:"og_see_also"`
	OgLocale           string `json:"og_locale"`
	ArticleAuthor      string `json:"article_author"`
	TwitterCard        string `json:"twitter_card"`
	TwitterUrl         string `json:"twitter_url"`
	TwitterTitle       string `json:"twitter_title"`
	TwitterDescription string `json:"twitter_description"`
	TwitterImage       string `json:"twitter_image"`
	Author             string `json:"author"`
	Generator          string `json:"generator"`
	Copyright          string `json:"copyright"`
}

type SeoResponse struct {
	Id                 int                   `json:"id"`
	Description        string                `json:"description"`
	Keywords           string                `json:"keywords"`
	FbAppId            string                `json:"fb_app_id"`
	OgTitle            string                `json:"og_title"`
	OgUrl              string                `json:"og_url"`
	OgImage            string                `json:"og_image"`
	OgDescription      string                `json:"og_description"`
	OgSiteName         string                `json:"og_site_name"`
	OgSeeAlso          string                `json:"og_see_also"`
	OgLocale           string                `json:"og_locale"`
	ArticleAuthor      string                `json:"article_author"`
	TwitterCard        string                `json:"twitter_card"`
	TwitterUrl         string                `json:"twitter_url"`
	TwitterTitle       string                `json:"twitter_title"`
	TwitterDescription string                `json:"twitter_description"`
	TwitterImage       string                `json:"twitter_image"`
	Author             string                `json:"author"`
	Generator          string                `json:"generator"`
	Copyright          string                `json:"copyright"`
	Articles           []ArticlesTagResponse `json:"articles_tag"`
}

type ArticlesTag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ArticlesTagResponse struct {
	Name string `json:"name"`
}

func (st StoreSeo) ShowSeos() (SeoResponse, error) {
	seoResponse := SeoResponse{}
	seo, err := FindOneSeo(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowSeos] error : ", err)
		return SeoResponse{}, err
	}

	articlesTag, err := FindArticlesTag(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[articlesTag] error : ", err)
		return SeoResponse{}, err
	}
	articlesTagResponses := []ArticlesTagResponse{}
	for _, data := range articlesTag {
		articlesTagResponse := ArticlesTagResponse{}
		articlesTagResponse.Name = data.Name
		articlesTagResponses = append(articlesTagResponses, articlesTagResponse)
	}

	seoResponse.Id = seo.Id
	seoResponse.Description = seo.Description
	seoResponse.Keywords = seo.Keywords
	seoResponse.FbAppId = seo.FbAppId
	seoResponse.OgTitle = seo.OgTitle
	seoResponse.OgUrl = seo.OgUrl
	seoResponse.OgImage = seo.OgImage
	seoResponse.OgDescription = seo.OgDescription
	seoResponse.OgSiteName = seo.OgSiteName
	seoResponse.OgSeeAlso = seo.OgSeeAlso
	seoResponse.OgLocale = seo.OgLocale
	seoResponse.ArticleAuthor = seo.ArticleAuthor
	seoResponse.TwitterCard = seo.TwitterCard
	seoResponse.TwitterUrl = seo.TwitterUrl
	seoResponse.TwitterTitle = seo.TwitterTitle
	seoResponse.TwitterDescription = seo.TwitterDescription
	seoResponse.TwitterImage = seo.TwitterImage
	seoResponse.Author = seo.Author
	seoResponse.Generator = seo.Generator
	seoResponse.Copyright = seo.Copyright
	seoResponse.Articles = articlesTagResponses
	return seoResponse, nil
}

func FindArticlesTag(db *sql.DB) ([]ArticlesTag, error) {
	articlesTag := []ArticlesTag{}
	query := `
	SELECT
		id,
		name
	FROM
		article_tag;
	`
	rows, err := db.Query(query)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindArticlesTag] query error  %v", err)
		return articlesTag, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		articleTag := ArticlesTag{}
		err = rows.Scan(&articleTag.Id, &articleTag.Name)
		articlesTag = append(articlesTag, articleTag)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RetrieveAccounts] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(articlesTag) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindArticlesTag] No Data  %v", err)
		return articlesTag, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articlesTag, nil

}

func FindOneSeo(db *sql.DB) (Seo, error) {
	seo := Seo{}
	query := `
	SELECT
		id,
		description,
		keywords,
		fb_app_id,
		og_title,
		og_url,
		og_image,
		og_description,
		og_site_name,
		og_see_also,
		og_locale,
		article_author,
		twitter_card,
		twitter_url,
		twitter_title,
		twitter_description,
		twitter_image,
		author,
		generator,
		copyright
	FROM
		seo
	`
	rows := db.QueryRow(query)
	err := rows.Scan(&seo.Id, &seo.Description, &seo.Keywords, &seo.FbAppId, &seo.OgTitle, &seo.OgUrl, &seo.OgImage, &seo.OgDescription, &seo.OgSiteName,
		&seo.OgSeeAlso, &seo.OgLocale, &seo.ArticleAuthor, &seo.TwitterCard, &seo.TwitterUrl, &seo.TwitterTitle, &seo.TwitterDescription,
		&seo.TwitterImage, &seo.Author, &seo.Generator, &seo.Copyright)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneSeo] No Data  %v", err)
		return seo, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneSeo] Scan error  %v", err)
	}
	return seo, nil
}
