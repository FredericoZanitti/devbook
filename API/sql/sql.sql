CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(200) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';