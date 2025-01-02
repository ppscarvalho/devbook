SELECT 
    p.*,
    u.Nick
FROM publicacao p INNER JOIN usuario u 
ON p.AutorId = u.Id
WHERE p.AutorId = 1;

SELECT DISTINCT 
    p.*,
    u.Nick
FROM publicacao p INNER JOIN usuario u 
ON p.AutorId = u.Id INNER JOIN seguidores s 
ON p.AutorId = s.IdUsuario 
WHERE p.AutorId = 3 OR s.IdSeguidor = 3;