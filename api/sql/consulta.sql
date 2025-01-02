SELECT * FROM usuario;

SELECT * FROM seguidores;

select u.Id, u.Nome, u.Nick, u.Email, u.CriadoEm
from usuario u inner join seguidores s
on u.Id = s.IdSeguidor
where s.IdUsuario = 1;

select u.Id, u.Nome, u.Nick, u.Email, u.CriadoEm
from usuario u inner join seguidores s
on u.Id = s.IdUsuario
where s.IdSeguidor = 1;

select * from publicacao;