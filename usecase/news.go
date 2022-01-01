package usecase

import (
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/bondhan/godddnews/usecase/view"
	"github.com/sirupsen/logrus"
	"net/url"
)

type NewsApp interface {
	GetAllNews(queryStr url.Values) (view.NewsView, error)
	//AddNews(newsData view.NewsData) error
	//GetNewsBySlug(slug string) (domain.News, error)
	//GetNewsByID(id uint) (domain.News, error)
	//UpdateNewsBySlug(newsData view.NewsData, slug string) error
	//DeleteNewsBySlug(slug string) error
	//UpdateNewsByID(newsData view.NewsData, ID uint) error
	//DeleteNewsByID(ID uint) error
	//GetNewsByTopicSlug(topicSlug string) ([]*domain.News, error)
	//GetNewsByTagSlug(tagSlug string) ([]*domain.News, error)
}

type newsApp struct {
	newsRepo  repository.NewsRepository
	topicRepo repository.TopicRepository
	tagRepo   repository.TagRepository
}

func NewNewsApp(newsRepo repository.NewsRepository, topicRepo repository.TopicRepository, tagRepo repository.TagRepository) NewsApp {
	return &newsApp{
		newsRepo:  newsRepo,
		topicRepo: topicRepo,
		tagRepo:   tagRepo,
	}
}

// GetAllNews will get all news. Pagination and filter is also supported.
func (n *newsApp) GetAllNews(queryStr url.Values) (view.NewsView, error) {
	var newsView view.NewsView
	var err error

	logrus.Debug("GetAllNews")
	newsView, err = n.newsRepo.GetAllNews(queryStr)

	return newsView, err
}

//
//// AddNews will add news to database
//func (n *newsApp) AddNews(newsData view.NewsData) error {
//	var err error
//
//	logrus.Debug("AddNews")
//	err = utils.ValidateModels(newsData)
//	if err != nil {
//		return err
//	}
//
//	err = utils.ValidateSlug(newsData.Slug)
//	if err != nil {
//		return err
//	}
//
//	nws, err := n.newsRepo.GetNewsBySlug(newsData.Slug)
//	//if error was not caused by empty select, means something happened during query to db
//	if err != nil && !gorm.IsRecordNotFoundError(err) {
//		return err
//	}
//
//	//if returned non empty data
//	if (nws.Model != gorm.Model{}) {
//		return errors.New("Data already exist")
//	}
//
//	news := domain.News{
//		Title: newsData.Title,
//		Slug:  newsData.Slug,
//	}
//
//	if newsData.ID > 0 {
//		news.ID = newsData.ID
//	}
//
//	if strings.TrimSpace(newsData.Content) != "" {
//		news.Content = newsData.Content
//	}
//
//	if strings.TrimSpace(newsData.Status) != "" {
//		news.Status = newsData.Status
//	}
//	// when creating a new news version default is 0
//	// we override to 1 as default value
//	if news.Version < 1 {
//		news.Version = 1
//	}
//
//	var topics []domain.Topic
//	for _, slug := range newsData.TopicSlugs {
//		topic, err := n.topicRepo.GetATopicBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Topic")
//		}
//		topics = append(topics, topic)
//	}
//
//	var tags []domain.Tag
//	for _, slug := range newsData.TagSlugs {
//		tag, err := n.tagRepo.GetATagBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Tag")
//		}
//		tags = append(tags, tag)
//	}
//
//	err = n.newsRepo.InsertNews(news, topics, tags)
//
//	return err
//}
//
//// GetNewsBySlug ...
//func (n *newsApp) GetNewsBySlug(slug string) (domain.News, error) {
//	var news domain.News
//	var err error
//
//	logrus.Debug("GetNewsBySlug")
//	news, err = n.newsRepo.GetNewsBySlug(slug)
//	if gorm.IsRecordNotFoundError(err) {
//		err = errors.New("Data not found")
//	}
//
//	return news, err
//}
//
//func (n *newsApp) GetNewsByID(id uint) (domain.News, error) {
//	var news domain.News
//	var err error
//
//	logrus.Debug("GetNewsByID")
//	news, err = n.newsRepo.GetNewsByID(id)
//	if gorm.IsRecordNotFoundError(err) {
//		err = errors.New("Data not found")
//	}
//
//	return news, err
//}
//
//func (n *newsApp) UpdateNewsBySlug(newsData view.NewsData, oldslug string) error {
//
//	var err error
//
//	logrus.Debug("UpdateNewsBySlug")
//
//	err = utils.ValidateModels(newsData)
//	if err != nil {
//		return err
//	}
//
//	err = utils.ValidateSlug(newsData.Slug)
//	if err != nil {
//		return err
//	}
//
//	nws, err := n.newsRepo.GetNewsBySlug(oldslug)
//	//if error was not caused by empty select, means something happened during query to db
//	if gorm.IsRecordNotFoundError(err) {
//		return errors.New("Data to be updated not exist")
//	}
//
//	//if returned non empty data
//	if (nws.Model == gorm.Model{}) {
//		return errors.New("Data to be updated not exist")
//	}
//
//	nws.Title = newsData.Title
//	nws.Slug = newsData.Slug
//
//	if strings.TrimSpace(newsData.Content) != "" {
//		nws.Content = newsData.Content
//	}
//
//	if strings.TrimSpace(newsData.Status) != "" {
//		nws.Status = newsData.Status
//	}
//
//	var newTopics []domain.Topic
//	for _, slug := range newsData.TopicSlugs {
//		topic, err := n.topicRepo.GetATopicBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Topic")
//		}
//		newTopics = append(newTopics, topic)
//	}
//
//	var newTags []domain.Tag
//	for _, slug := range newsData.TagSlugs {
//		tag, err := n.tagRepo.GetATagBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Tag")
//		}
//		newTags = append(newTags, tag)
//	}
//
//	err = n.newsRepo.UpdateNews(nws, newTopics, newTags)
//
//	return err
//}
//
//func (n *newsApp) DeleteNewsBySlug(slug string) error {
//
//	var err error
//
//	logrus.Debug("DeleteNewsBySlug")
//
//	nws, err := n.newsRepo.GetNewsBySlug(slug)
//	//if error was not caused by empty select, means something happened during query to db
//	if gorm.IsRecordNotFoundError(err) {
//		return errors.New("Data to be deleted not exist")
//	}
//
//	//if returned non empty data
//	if (nws.Model == gorm.Model{}) {
//		return errors.New("Data to be deleted not exist")
//	}
//
//	var oldTopics []domain.Topic
//	for _, topic := range nws.Topics {
//		topic, err := n.topicRepo.GetATopicBySlug(topic.Slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Topic")
//		}
//		oldTopics = append(oldTopics, topic)
//	}
//
//	var oldTags []domain.Tag
//	for _, tag := range nws.Tags {
//		tag, err := n.tagRepo.GetATagBySlug(tag.Slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Tag")
//		}
//		oldTags = append(oldTags, tag)
//	}
//
//	err = n.newsRepo.DeleteNewsBySlug(nws, oldTopics, oldTags)
//
//	return err
//}
//
//func (n *newsApp) UpdateNewsByID(newsData view.NewsData, ID uint) error {
//
//	var err error
//
//	logrus.Debug("UpdateNewsByID")
//
//	err = utils.ValidateModels(newsData)
//	if err != nil {
//		return err
//	}
//
//	err = utils.ValidateSlug(newsData.Slug)
//	if err != nil {
//		return err
//	}
//
//	nws, err := n.newsRepo.GetNewsByID(ID)
//	//if error was not caused by empty select, means something happened during query to db
//	if gorm.IsRecordNotFoundError(err) {
//		return errors.New("Data to be updated not exist")
//	}
//
//	//if returned non empty data
//	if (nws.Model == gorm.Model{}) {
//		return errors.New("Data to be updated not exist")
//	}
//
//	nws.Title = newsData.Title
//	nws.Slug = newsData.Slug
//
//	if strings.TrimSpace(newsData.Content) != "" {
//		nws.Content = newsData.Content
//	}
//
//	if strings.TrimSpace(newsData.Status) != "" {
//		nws.Status = newsData.Status
//	}
//
//	var newTopics []domain.Topic
//	for _, slug := range newsData.TopicSlugs {
//		topic, err := n.topicRepo.GetATopicBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Topic")
//		}
//		newTopics = append(newTopics, topic)
//	}
//
//	var newTags []domain.Tag
//	for _, slug := range newsData.TagSlugs {
//		tag, err := n.tagRepo.GetATagBySlug(slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Tag")
//		}
//		newTags = append(newTags, tag)
//	}
//
//	err = n.newsRepo.UpdateNews(nws, newTopics, newTags)
//
//	return err
//}
//
//func (n *newsApp) DeleteNewsByID(ID uint) error {
//
//	var err error
//
//	logrus.Debug("DeleteNewsByID")
//
//	nws, err := n.newsRepo.GetNewsByID(ID)
//	//if error was not caused by empty select, means something happened during query to db
//	if gorm.IsRecordNotFoundError(err) {
//		return errors.New("Data to be deleted not exist")
//	}
//
//	//if returned non empty data
//	if (nws.Model == gorm.Model{}) {
//		return errors.New("Data to be deleted not exist")
//	}
//
//	var oldTopics []domain.Topic
//	for _, topic := range nws.Topics {
//		topic, err := n.topicRepo.GetATopicBySlug(topic.Slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Topic")
//		}
//		oldTopics = append(oldTopics, topic)
//	}
//
//	var oldTags []domain.Tag
//	for _, tag := range nws.Tags {
//		tag, err := n.tagRepo.GetATagBySlug(tag.Slug)
//		if gorm.IsRecordNotFoundError(err) {
//			return errors.New("Unknown Tag")
//		}
//		oldTags = append(oldTags, tag)
//	}
//
//	err = n.newsRepo.DeleteNewsBySlug(nws, oldTopics, oldTags)
//
//	return err
//}
//
//// GetNewsByTopicSlug ...
//func (n *newsApp) GetNewsByTopicSlug(topicSlug string) ([]*domain.News, error) {
//
//	logrus.Debug("GetNewsByTopicSlug")
//	news, err := n.newsRepo.GetNewsByTopicSlug(topicSlug)
//	if gorm.IsRecordNotFoundError(err) {
//		err = errors.New("Data not found")
//	}
//
//	return news, err
//}
//
//// GetNewsBySlug ...
//func (n *newsApp) GetNewsByTagSlug(tagSlug string) ([]*domain.News, error) {
//	logrus.Debug("GetNewsByTagSlug")
//	news, err := n.newsRepo.GetNewsByTagSlug(tagSlug)
//	if gorm.IsRecordNotFoundError(err) {
//		err = errors.New("Data not found")
//	}
//
//	return news, err
//}
