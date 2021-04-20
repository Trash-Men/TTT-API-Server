CREATE TABLE IF NOT EXISTS trash (
  photo_url VARCHAR (100) NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  area VARCHAR (10) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (photo_url)
);

CREATE TABLE IF NOT EXISTS trash_can (
  photo_url VARCHAR (100) NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  area VARCHAR (10) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (photo_url)
);

CREATE TYPE E_ROLE AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS user_table (
  id VARCHAR (20) NOT NULL,
  password VARCHAR (80) NOT NULL,
  role E_ROLE NOT NULL,

  PRIMARY KEY (id)
);

-- insert into trash values('upload/trashCan/2021/04/19/3318fe01-05eb-46dc-93fa-d01bda425afe.png', 127.1479532, 35.8242238, '전라북도');

-- postgreSQL에서는 문자열을 입력할때 "" 쌍따옴표를 쓰면 에러가 난다.
-- "" -> 식별자를 나타낼떄 사용, e.g., table name, column key
-- '' -> 토큰(값)을 나타내기 위해 사용, e.g., insert values할때 문자열 값
-- psql -U woochanleee -d development_db -a -f ddl.sql -p 9090 -h 0.0.0.0

-- if not exists 를 붙여주면 -> NOTICE:  relation "trash_can2" already exists, skipping
-- 안붙이면 -> ERROR:  relation "trash_can2" already exists