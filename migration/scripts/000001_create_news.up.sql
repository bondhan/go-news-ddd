CREATE SEQUENCE m_news_seq;
CREATE TYPE status AS ENUM('draft','deleted','publish');

CREATE TABLE m_news (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_news_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  title varchar(255) NOT NULL,
  slug varchar(255) NOT NULL,
  content text DEFAULT NULL,
  status status DEFAULT 'draft',
  version int check (version > 0) DEFAULT 1,
  UNIQUE(slug),
  PRIMARY KEY (id)
);

CREATE INDEX m_news_deleted_at ON m_news (deleted_at);
CREATE INDEX m_news_slug ON m_news (slug);
CREATE INDEX m_news_version ON m_news (version);