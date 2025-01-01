INSERT INTO usuario(Nome, Nick, Email, Senha) 
VALUES
('Nazareno Caio Geraldo Campos', 'caio', 'nazareno.caio@gmail.com', '$2a$10$Sx28.w1nkgXbETmryrH0FOMyN6poQZrgMhjWcFlbHkM121/6rUkEK'),
('Bento Thiago Gabriel da Cruz', 'thiago', 'bento.thiago@gmail.com', '$2a$10$Sx28.w1nkgXbETmryrH0FOMyN6poQZrgMhjWcFlbHkM121/6rUkEK'),
('Gael Diego Pietro Jesus', 'gael', 'gael.diego@gmail.com', '$2a$10$Sx28.w1nkgXbETmryrH0FOMyN6poQZrgMhjWcFlbHkM121/6rUkEK'),
('Murilo Carlos Enzo Duarte', 'murilo', 'murilo.carlos@gmail.com', '$2a$10$Sx28.w1nkgXbETmryrH0FOMyN6poQZrgMhjWcFlbHkM121/6rUkEK'),
('Oliver Lucca Lopes', 'oliver', 'oliver.lucca@gmail.com', '$2a$10$Sx28.w1nkgXbETmryrH0FOMyN6poQZrgMhjWcFlbHkM121/6rUkEK');

INSERT INTO seguidores(IdUsuario, IdSeguidor) 
VALUES
(1, 2),
(1, 3),
(1, 4),
(2, 1),
(4, 1),
(5, 1);