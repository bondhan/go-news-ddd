CREATE SEQUENCE m_topic_seq;

CREATE TABLE m_topic (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_topic_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  slug varchar(255) NOT NULL,
  version int check (version > 0) DEFAULT 1,
  UNIQUE(slug),
  PRIMARY KEY (id)
);

CREATE INDEX m_topic_deleted_at ON m_topic (deleted_at);
CREATE INDEX m_topic_slug ON m_topic (slug);
CREATE INDEX m_topic_version ON m_topic (version);