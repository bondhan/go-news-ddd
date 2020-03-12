CREATE TABLE m_topic_news (
  news_id int check (news_id > 0) NOT NULL,
  topic_id int check (topic_id > 0) NOT NULL,
  UNIQUE(news_id, topic_id),
  CONSTRAINT fk_m_topic_news_news_id FOREIGN KEY (news_id) REFERENCES m_news (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT fk_m_topic_news_topic_id FOREIGN KEY (topic_id) REFERENCES m_topic (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_topic_news_news_id ON m_topic_news (news_id);
CREATE INDEX m_topic_news_topic_id ON m_topic_news (topic_id);