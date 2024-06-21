CREATE TABLE IF NOT EXISTS users (
  id         VARCHAR(20)  NOT NULL,
  name       VARCHAR(100) NOT NULL,
  project_v2 VARCHAR(100),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS repositories (
  id         VARCHAR(20)  NOT NULL,
  owner      VARCHAR(20)  NOT NULL,
  name       VARCHAR(100) NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT repositories_owner_fk FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS issues (
  id         VARCHAR(20)  NOT NULL,
  url        VARCHAR(200) NOT NULL,
  title      VARCHAR(200) NOT NULL,
  closed     BOOLEAN      NOT NULL DEFAULT FALSE,
  number     INT          NOT NULL,
  author     VARCHAR(20)  NOT NULL,
  repository VARCHAR(20)  NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT issues_repository_fk FOREIGN KEY (repository) REFERENCES repositories (id)
);

CREATE TABLE IF NOT EXISTS projects (
  id     VARCHAR(20)  NOT NULL,
  title  VARCHAR(200) NOT NULL,
  url    VARCHAR(200) NOT NULL,
  number INT          NOT NULL,
  owner  VARCHAR(20)  NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT projects_owner_fk FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS pullrequests (
  id            VARCHAR(20)  NOT NULL,
  base_ref_name VARCHAR(200) NOT NULL,
  closed        BOOLEAN      NOT NULL DEFAULT FALSE,
  head_ref_name VARCHAR(200) NOT NULL,
  url           VARCHAR(200) NOT NULL,
  number        INT          NOT NULL,
  repository    VARCHAR(20)  NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT pullrequests_repository_fk FOREIGN KEY (repository) REFERENCES repositories (id)
);

CREATE TABLE IF NOT EXISTS projectcards (
  id          VARCHAR(20) NOT NULL,
  project     VARCHAR(20) NOT NULL,
  issue       VARCHAR(20),
  pullrequest VARCHAR(20),
  PRIMARY KEY (id),
  CONSTRAINT projectcards_project_fk FOREIGN KEY (project) REFERENCES projects (id),
  CONSTRAINT projectcards_check CHECK ( issue IS NOT NULL OR pullrequest IS NOT NULL )
);
