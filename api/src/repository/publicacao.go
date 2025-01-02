package repository

import (
	"api/src/models"
	"database/sql"
)

type ContextPublicacao struct {
	db *sql.DB
}

func PublicacaoRepository(db *sql.DB) *ContextPublicacao {
	return &ContextPublicacao{db}
}

// Implemente aqui a criação da publicação no banco de dados
func (contextPublicacao ContextPublicacao) Criar(publicacao models.Publicacao) (uint64, error) {

	statement, erro := contextPublicacao.db.Prepare(
		"INSERT INTO publicacao (Titulo, Conteudo, AutorId) VALUES (? , ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	idPublicacao, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(idPublicacao), nil
}

// Implemente aqui a busca da publicação no banco de dados por AutorId
func (contextPublicacao ContextPublicacao) BuscarPorId(idPublicacao uint64) (models.Publicacao, error) {
	result, erro := contextPublicacao.db.Query(`
		SELECT p.*, u.Nick 
		FROM publicacao p INNER JOIN usuario u 
		ON p.AutorId = u.Id
		WHERE p.Id = ?`, idPublicacao)
	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer result.Close()

	var publicacao models.Publicacao
	if result.Next() {
		if erro = result.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Implemente aqui a busca da publicações do usuário seguido e do Autor da publicação no banco de dados
func (contextPublicacao ContextPublicacao) Buscar(idUsuario uint64) ([]models.Publicacao, error) {
	result, erro := contextPublicacao.db.Query(`
	SELECT DISTINCT p.*, u.Nick
	FROM publicacao p INNER JOIN usuario u 
	ON p.AutorId = u.Id INNER JOIN seguidores s 
	ON p.AutorId = s.IdUsuario 
	WHERE p.AutorId = ? OR s.IdSeguidor = ?	
	ORDER BY 1 DESC`, idUsuario, idUsuario)

	if erro != nil {
		return []models.Publicacao{}, erro
	}

	defer result.Close()

	var publicacoes []models.Publicacao

	for result.Next() {
		var publicacao models.Publicacao
		if erro = result.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick); erro != nil {
			return []models.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Implemente aqui a atualização da publicação no banco de dados
func (contextPublicacao ContextPublicacao) Atualizar(publicacao models.Publicacao) error {
	statement, erro := contextPublicacao.db.Prepare(
		"UPDATE publicacao SET Titulo = ?, Conteudo = ? WHERE Id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.Id); erro != nil {
		return erro
	}

	return nil
}

// Implemente aqui a exclusão da publicação no banco de dados
func (contextPublicacao ContextPublicacao) Deletar(idPublicacao uint64) error {
	statement, erro := contextPublicacao.db.Prepare("DELETE FROM publicacao WHERE Id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(idPublicacao); erro != nil {
		return erro
	}

	return nil
}

// Implemente aqui a busca todas as publicações de um determinando usuário
func (contextPublicacao ContextPublicacao) BuscarPorUsuario(idUsuario uint64) ([]models.Publicacao, error) {
	result, erro := contextPublicacao.db.Query(`
		SELECT p.*, u.Nick 
		FROM publicacao p INNER JOIN usuario u 
		ON p.AutorId = u.Id
		WHERE p.AutorId = ?`, idUsuario)
	if erro != nil {
		return []models.Publicacao{}, erro
	}

	defer result.Close()

	var publicacoes []models.Publicacao
	for result.Next() {
		var publicacao models.Publicacao
		if erro = result.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick); erro != nil {
			return []models.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Implemente aqui curtir a publicação de um usuario do banco de dados
func (contextPublicacao ContextPublicacao) Curtir(idPublicacao uint64) error {
	statement, erro := contextPublicacao.db.Prepare("UPDATE publicacao SET Curtidas = Curtidas + 1 WHERE Id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(idPublicacao); erro != nil {
		return erro
	}

	return nil
}

// Implemente aqui descutir uma publicação do  banco de dados por AutorId
func (contextPublicacao ContextPublicacao) DesCurtir(idPublicacao uint64) error {
	statement, erro := contextPublicacao.db.Prepare(`
		UPDATE publicacao SET Curtidas = 
		CASE WHEN Curtidas > 0 THEN Curtidas - 1
		ELSE Curtidas END 
		WHERE Id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(idPublicacao); erro != nil {
		return erro
	}

	return nil
}
