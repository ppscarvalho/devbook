CREATE DATABASE IT NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuario;

CREATE TABLE usuario(
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Nome varchar(60) not null,
    Nick varchar(60) not null unique,
    Email varchar(60) not null unique,
    Senha varchar(100) not null,
    CriadoEm TIMESTAMP DEFAULT current_timestamp()
)ENGINE=INNODB;


CREATE TABLE seguidores(
    IdUsuario INT not null,
    IdSeguidor INT not null,
    FOREIGN KEY (IdUsuario) REFERENCES usuario(Id) ON DELETE CASCADE,
    FOREIGN KEY (IdSeguidor) REFERENCES usuario(Id) ON DELETE CASCADE,
    PRIMARY KEY (IdUsuario, IdSeguidor)
)ENGINE=INNODB;