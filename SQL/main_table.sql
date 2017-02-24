-- 创建数据库
CREATE DATABASE naturl_language_process;

-- 创建单字属性表
CREATE TABLE naturl_language_process.main_single_word
(
    word_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    word_name VARCHAR(20) NOT NULL,
    word_pinyin VARCHAR(125)
);
CREATE UNIQUE INDEX main_single_word_word_id_uindex ON naturl_language_process.main_single_word (word_id);

-- 创建词汇属性表
CREATE TABLE naturl_language_process.main_single_terms
(
    terms_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    terms_name VARCHAR(100) NOT NULL,
    terms_pinyin VARCHAR(255)
);
CREATE UNIQUE INDEX main_single_terms_terms_id_uindex ON naturl_language_process.main_single_terms (terms_id);
CREATE UNIQUE INDEX main_single_terms_terms_name_uindex ON naturl_language_process.main_single_terms (terms_name);
