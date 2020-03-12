insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (1, current_timestamp, current_timestamp, 'Investment', 'investment');
insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (2, current_timestamp, current_timestamp, 'Politics', 'politics');
insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (3, current_timestamp, current_timestamp, 'Religion', 'religion'); 
insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (4, current_timestamp, current_timestamp, 'International', 'international'); 
insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (5, current_timestamp, current_timestamp, 'National', 'national');
	 insert into m_topic (id, created_at ,updated_at , name, slug )
 	values (6, current_timestamp, current_timestamp, 'Other', 'other-topic');
 
insert into m_tag (id, created_at ,updated_at , name, slug)
 	values (1, current_timestamp, current_timestamp, 'Investment Tag', 'investment-tag');
insert into m_tag (id, created_at ,updated_at , name, slug )
 	values (2, current_timestamp, current_timestamp, 'Politics Tag', 'politics-tag');
insert into m_tag (id, created_at ,updated_at , name, slug )
 	values (3, current_timestamp, current_timestamp, 'Religion Tag', 'religion-tag'); 
insert into m_tag (id, created_at ,updated_at , name, slug )
 	values (4, current_timestamp, current_timestamp, 'International Tag', 'international-tag'); 
insert into m_tag (id, created_at ,updated_at , name, slug )
 	values (5, current_timestamp, current_timestamp, 'National Tag', 'national-tag');
insert into m_tag (id, created_at ,updated_at , name, slug )
 	values (6, current_timestamp, current_timestamp, 'Other', 'other-tag');
 
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(1, current_timestamp, current_timestamp, 'Title Investment 1', 'slug-investment-a', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(2, current_timestamp, current_timestamp, 'Title Investment 2', 'slug-investment-b', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(3, current_timestamp, current_timestamp, 'Title Investment 3', 'slug-investment-c', 'deleted', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(4, current_timestamp, current_timestamp, 'Title Investment 4', 'slug-investment-d', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(5, current_timestamp, current_timestamp, 'Title Investment 5', 'slug-investment-e', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(6, current_timestamp, current_timestamp, 'Title Investment 6', 'slug-investment-f', 'deleted', 'content');

insert into m_topic_news(news_id, topic_id) values(1, 1);
insert into m_topic_news(news_id, topic_id) values(2, 1);
insert into m_topic_news(news_id, topic_id) values(3, 1);
insert into m_topic_news(news_id, topic_id) values(4, 1);
insert into m_topic_news(news_id, topic_id) values(5, 1);
insert into m_topic_news(news_id, topic_id) values(6, 1);

insert into m_tag_news(news_id, tag_id ) values(1, 1);
insert into m_tag_news(news_id, tag_id ) values(2, 1);
insert into m_tag_news(news_id, tag_id ) values(3, 1);
insert into m_tag_news(news_id, tag_id ) values(4, 1);
insert into m_tag_news(news_id, tag_id ) values(5, 1);
insert into m_tag_news(news_id, tag_id ) values(6, 1);

insert into m_tag_news(news_id, tag_id ) values(1, 4);
insert into m_tag_news(news_id, tag_id ) values(2, 4);
insert into m_tag_news(news_id, tag_id ) values(3, 4);
insert into m_tag_news(news_id, tag_id ) values(4, 4);
insert into m_tag_news(news_id, tag_id ) values(5, 4);
insert into m_tag_news(news_id, tag_id ) values(6, 4);


--------------------------------------------------------------------------------
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(11, current_timestamp, current_timestamp, 'Title Politics 11', 'slug-politics-a', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(12, current_timestamp, current_timestamp, 'Title Politics 12', 'slug-politics-b', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(13, current_timestamp, current_timestamp, 'Title Politics 13', 'slug-politics-c', 'deleted', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(14, current_timestamp, current_timestamp, 'Title Politics 14', 'slug-politics-d', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(15, current_timestamp, current_timestamp, 'Title Politics 15', 'slug-politics-e', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(16, current_timestamp, current_timestamp, 'Title Politics 16', 'slug-politics-f', 'deleted', 'content');


insert into m_topic_news(news_id, topic_id) values(11, 2);
insert into m_topic_news(news_id, topic_id) values(12, 2);
insert into m_topic_news(news_id, topic_id) values(13, 2);
insert into m_topic_news(news_id, topic_id) values(14, 2);
insert into m_topic_news(news_id, topic_id) values(15, 2);
insert into m_topic_news(news_id, topic_id) values(16, 2);

insert into m_tag_news(news_id, tag_id ) values(11, 2);
insert into m_tag_news(news_id, tag_id ) values(12, 2);
insert into m_tag_news(news_id, tag_id ) values(13, 2);
insert into m_tag_news(news_id, tag_id ) values(14, 2);
insert into m_tag_news(news_id, tag_id ) values(15, 2);
insert into m_tag_news(news_id, tag_id ) values(16, 2);

insert into m_tag_news(news_id, tag_id ) values(11, 5);
insert into m_tag_news(news_id, tag_id ) values(12, 5);
insert into m_tag_news(news_id, tag_id ) values(13, 5);
insert into m_tag_news(news_id, tag_id ) values(14, 5);
insert into m_tag_news(news_id, tag_id ) values(15, 5);
insert into m_tag_news(news_id, tag_id ) values(16, 5);

--------------------------------------------------------------------------------------------------------------------------
insert into m_news(id, created_at, updated_at, title, slug, status, content) 
	values(21, current_timestamp, current_timestamp, 'Title Religion 21', 'slug-religion-a', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content)  
	values(22, current_timestamp, current_timestamp, 'Title Religion 22', 'slug-religion-b', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content)  
	values(23, current_timestamp, current_timestamp, 'Title Religion 23', 'slug-religion-c', 'deleted', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content)  
	values(24, current_timestamp, current_timestamp, 'Title Religion 24', 'slug-religion-d', 'draft', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content)  
	values(25, current_timestamp, current_timestamp, 'Title Religion 5', 'slug-religion-e', 'publish', 'content');
insert into m_news(id, created_at, updated_at, title, slug, status, content)  
	values(26, current_timestamp, current_timestamp, 'Title Religion 26', 'slug-religion-f', 'deleted', 'content');

insert into m_topic_news(news_id, topic_id) values(21, 3);
insert into m_topic_news(news_id, topic_id) values(22, 3);
insert into m_topic_news(news_id, topic_id) values(23, 3);
insert into m_topic_news(news_id, topic_id) values(24, 3);
insert into m_topic_news(news_id, topic_id) values(25, 3);
insert into m_topic_news(news_id, topic_id) values(26, 3);

insert into m_tag_news(news_id, tag_id ) values(21, 6);
insert into m_tag_news(news_id, tag_id ) values(22, 6);
insert into m_tag_news(news_id, tag_id ) values(23, 6);
insert into m_tag_news(news_id, tag_id ) values(24, 6);
insert into m_tag_news(news_id, tag_id ) values(25, 6);
insert into m_tag_news(news_id, tag_id ) values(26, 6);

