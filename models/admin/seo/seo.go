package seo

type SeoRequest struct {
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
