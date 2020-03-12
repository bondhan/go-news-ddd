CREATE SEQUENCE m_tag_seq;

CREATE TABLE m_tag (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_tag_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  slug varchar(255) NOT NULL,
  version int check (version > 0) DEFAULT 1,
  UNIQUE(slug),
  PRIMARY KEY (id)
);

CREATE INDEX m_tag_deleted_at ON m_tag (deleted_at);
CREATE INDEX m_tag_slug ON m_tag (slug);
CREATE INDEX m_tag_version ON m_tag (version);