-- PROJECT TABLE
CREATE TABLE IF NOT EXISTS projects (
  id              SERIAL PRIMARY KEY,
  project         VARCHAR(128) NOT NULL,
  targets         JSON,
  path            VARCHAR(1024) NOT NULL
);

-- BUILD TABLE
CREATE TABLE IF NOT EXISTS builds (
  id              SERIAL PRIMARY KEY,
  title           VARCHAR(128) NOT NULL,
  target          VARCHAR(128) NOT NULL,
  manifest_url    VARCHAR(256) NOT NULL,
  path            VARCHAR(1024) NOT NULL -- path to ipa directory
);
