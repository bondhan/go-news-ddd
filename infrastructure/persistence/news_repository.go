package persistence

import (
	"bytes"
	"math"
	"net/url"
	"strconv"

	"github.com/bondhan/godddnews/application/view"
	"github.com/bondhan/godddnews/domain"
	"github.com/bondhan/godddnews/domain/repository"
	"github.com/jinzhu/gorm"
)

const queryPageSizeDefault = 10
const queryPageNumberDefault = 1

type newsRepository struct {
	db        *gorm.DB
	topicRepo repository.TopicRepository
	tagRepo   repository.TagRepository
}

//NewNewsRepository ...
func NewNewsRepository(newDB *gorm.DB, topicRepo repository.TopicRepository, tagRepo repository.TagRepository) repository.NewsRepository {
	return &newsRepository{
		db:        newDB,
		topicRepo: topicRepo,
		tagRepo:   tagRepo,
	}
}

func (n *newsRepository) GetAllNews(queryStr url.Values) (view.NewsView, error) {
	var nv view.NewsView

	status := queryStr.Get("status")
	pageSize := queryStr.Get("page_size")
	pageNumber := queryStr.Get("page_number")

	var bufferWhere bytes.Buffer
	if status == "draft" || status == "deleted" || status == "publish" {
		bufferWhere.WriteString("status = " + "'" + status + "'")
	}
	// if status is not in the list, means query all

	pgSize, err := strconv.Atoi(pageSize)
	if err != nil || pgSize < 1 {
		pgSize = queryPageSizeDefault
	}

	pgNumber, err := strconv.Atoi(pageNumber)
	if err != nil || pgNumber < 1 {
		pgNumber = queryPageNumberDefault
	}
	pgNumberOut := pgNumber
	if pgNumber > 1 {
		pgNumber--
	}

	/////// get data as per filter ///////////
	nx := []domain.News{}
	err = n.db.Where(bufferWhere.String()).Preload("Topics").
		Preload("Tags").
		Offset((pgNumber - 1) * pgSize).
		Limit(pgSize).
		Find(&nx).Error

	/////// get total data as per filter ///////////
	var totalData view.TotalData
	var totalPage int
	err = n.db.Table("m_news").Select("count(*)").
		Where(bufferWhere.String()).
		Scan(&totalData).Error

	if err != nil {
		return nv, err
	}
	totalPage = int(math.Ceil(float64(totalData.Count) / float64(pgSize)))

	nv = view.NewsView{
		Total:      totalData.Count,
		TotalPage:  totalPage,
		PageNumber: pgNumberOut,
		PageSize:   pgSize,
		News:       nx,
	}

	return nv, err
}

func (n *newsRepository) GetNewsByID(id uint) (domain.News, error) {
	var News domain.News
	err := n.db.Where("id = ?", id).Preload("Topics").Preload("Tags").First(&News).Error

	return News, err
}

func (n *newsRepository) GetNewsBySlug(slug string) (domain.News, error) {
	var News domain.News
	err := n.db.Where("slug = ?", slug).Preload("Topics").Preload("Tags").First(&News).Error

	return News, err
}

func (n *newsRepository) InsertNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error {
	var err error

	tx := n.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	//update and create new many2many relation
	err = tx.Create(&News).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, topic := range NewTopics {
		tpn := domain.TopicNews{
			NewsID:  News.ID,
			TopicID: topic.ID,
		}
		err = tx.Create(&tpn).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, tag := range NewTags {
		tg := domain.TagNews{
			NewsID: News.ID,
			TagID:  tag.ID,
		}
		err = tx.Create(&tg).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (n *newsRepository) UpdateNews(News domain.News, NewTopics []domain.Topic, NewTags []domain.Tag) error {
	tx := n.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	//delete the old topic_news relations
	for _, topic := range News.Topics {
		tn := domain.TopicNews{}

		err := tx.Unscoped().Where("news_id = ? and topic_id= ?", News.ID, topic.ID).Delete(&tn).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//delete the old tag_news relations
	for _, tag := range News.Tags {
		tg := domain.TagNews{}
		err := tx.Unscoped().Where("news_id = ? and tag_id= ?", News.ID, tag.ID).Delete(&tg).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//update and create new many2many relation
	ver := News.Version
	News.Version = News.Version + 1
	err := tx.Model(&News).Where("version = ?", ver).UpdateColumns(News).Error //optimistic lock
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, topic := range NewTopics {
		tpn := domain.TopicNews{
			NewsID:  News.ID,
			TopicID: topic.ID,
		}
		err = tx.Create(&tpn).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, tag := range NewTags {
		tg := domain.TagNews{
			NewsID: News.ID,
			TagID:  tag.ID,
		}
		err = tx.Create(&tg).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (n *newsRepository) DeleteNewsBySlug(News domain.News, OldTopics []domain.Topic, OldTags []domain.Tag) error {
	tx := n.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	//delete the old topic_news relations
	for _, topic := range News.Topics {
		tn := domain.TopicNews{}

		err := tx.Unscoped().Where("news_id = ? and topic_id= ?", News.ID, topic.ID).Delete(&tn).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//delete the old tag_news relations
	for _, tag := range News.Tags {
		tg := domain.TagNews{}
		err := tx.Unscoped().Where("news_id = ? and tag_id= ?", News.ID, tag.ID).Delete(&tg).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//update and create new many2many relation
	err := tx.Unscoped().Delete(&News).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (n *newsRepository) GetNewsByTopicSlug(slug string) ([]*domain.News, error) {

	us := make([]*domain.News, 0)
	err := n.db.Raw("select m_news.*  from  m_topic_news left join m_news on m_topic_news.news_id = m_news.id where m_topic_news.topic_id =( select id as topic_id from m_topic  where slug = ?)", slug).Scan(&us).Error

	return us, err
}

func (n *newsRepository) GetNewsByTagSlug(slug string) ([]*domain.News, error) {
	us := make([]*domain.News, 0)

	err := n.db.Raw("select m_news.* from m_tag_news left join m_news on m_tag_news.news_id = m_news.id where m_tag_news.tag_id =(select id as tag_id from m_tag where slug = ?)", slug).Scan(&us).Error

	return us, err
}
