CREATE TABLE m_tag_news (
  news_id int check (news_id > 0) NOT NULL,
  tag_id int check (tag_id > 0) NOT NULL,
  UNIQUE(news_id, tag_id),
  CONSTRAINT fk_m_tag_news_news_id FOREIGN KEY (news_id) REFERENCES m_news (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT fk_m_tag_news_tag_id FOREIGN KEY (tag_id) REFERENCES m_tag (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_tag_news_news_id ON m_tag_news (news_id);
CREATE INDEX m_tag_news_tag_id ON m_tag_news (tag_id);