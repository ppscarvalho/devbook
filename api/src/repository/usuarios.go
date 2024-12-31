package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type ContextUsuario struct {
	db *sql.DB
}

func NovoRepositoryUsuario(db *sql.DB) *ContextUsuario {
	return &ContextUsuario{db}
}

func (contextUsuario ContextUsuario) Criar(usuario models.Usuario) (uint64, error) {

	statement, erro := contextUsuario.db.Prepare(
		"insert into usuario (Nome, Nick, Email, Senha) values(? , ?, ?, ?)",
	)

	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()

	if erro != nil {
		return 0, nil
	}

	return uint64(ID), nil
}

func (contextUsuario ContextUsuario) Buscar(params string) ([]models.Usuario, error) {
	params = fmt.Sprintf("%%%s%%", params)
	result, erro := contextUsuario.db.Query("select Id, Nome, Nick, Email, CriadoEm from usuario where nome like ? or nick like ?", params, params)

	if erro != nil {
		return nil, erro
	}

	defer result.Close()

	var usuarios []models.Usuario

	for result.Next() {
		var usuario models.Usuario
		if erro = result.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (contextUsuario ContextUsuario) BuscarPorId(id uint64) (models.Usuario, error) {

	result, erro := contextUsuario.db.Query("select Id, Nome, Nick, Email, CriadoEm from usuario where id = ?", id)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer result.Close()

	var usuario models.Usuario

	if result.Next() {
		if erro = result.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (contextUsuario ContextUsuario) Atualizar(usuarioID uint64, usuario models.Usuario) (models.Usuario, error) {
	statement, erro := contextUsuario.db.Prepare(
		"update usuario set Nome = ?, Nick = ?, Email = ? where Id = ?",
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); erro != nil {
		return models.Usuario{}, erro
	}

	usuario, erro = contextUsuario.BuscarPorId(usuarioID)
	if erro != nil {
		return models.Usuario{}, erro
	}

	return usuario, nil
}

func (contextUsuario ContextUsuario) Excluir(usuarioID uint64) error {
	statement, erro := contextUsuario.db.Prepare("delete from usuario where Id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuarioID); erro != nil {
		return erro
	}

	return nil
}

func (contextUsuario ContextUsuario) BuscarPorEmail(email string) (models.Usuario, error) {
	result, erro := contextUsuario.db.Query("select id, senha from usuario where Email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}

	defer result.Close()

	var usuario models.Usuario
	if result.Next() {
		if erro = result.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (contextUsuario ContextUsuario) SeguirUsuario(idUsuario, idSeguidor uint64) error {

	statement, erro := contextUsuario.db.Prepare("insert ignore into seguidores (IdUsuario, IdSeguidor) values (?, ?)")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(idUsuario, idSeguidor); erro != nil {
		return erro
	}

	return nil
}

func (contextUsuario ContextUsuario) PararDeSeguirUsuario(idUsuario, idSeguidor uint64) error {

	statement, erro := contextUsuario.db.Prepare("delete from seguidores where IdUsuario = ? and IdSeguidor = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(idUsuario, idSeguidor); erro != nil {
		return erro
	}

	return nil
}

func (contextUsuario ContextUsuario) BuscarSeguidores(idUsuario uint64) ([]models.Usuario, error) {
	result, erro := contextUsuario.db.Query(`
		select u.Id, u.Nome, u.Nick, u.Email, u.CriadoEm
		from usuario u inner join seguidores s
		on u.Id = s.IdSeguidor
		where s.IdUsuario = ?`, idUsuario)

	if erro != nil {
		return nil, erro
	}

	defer result.Close()

	var usuarios []models.Usuario

	for result.Next() {
		var usuario models.Usuario
		if erro = result.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (contextUsuario ContextUsuario) BuscarSeguindo(idSeguidor uint64) ([]models.Usuario, error) {
	result, erro := contextUsuario.db.Query(`
		select u.Id, u.Nome, u.Nick, u.Email, u.CriadoEm
		from usuario u inner join seguidores s
		on u.Id = s.IdUsuario
		where s.idSeguidor = ?`, idSeguidor)

	if erro != nil {
		return nil, erro
	}

	defer result.Close()

	var usuarios []models.Usuario

	for result.Next() {
		var usuario models.Usuario
		if erro = result.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (contextUsuario *ContextUsuario) BuscarSenha(idUsuario uint64) (senha string, erro error) {

	result, erro := contextUsuario.db.Query("select Senha from usuario where id = ?", idUsuario)

	if erro != nil {
		return "", erro
	}

	defer result.Close()

	if result.Next() {
		if erro = result.Scan(&senha); erro != nil {
			return "", erro
		}
	}
	return senha, nil
}

func (contextUsuario ContextUsuario) AtualizarSenha(idUsuario uint64, senha string) error {
	statement, erro := contextUsuario.db.Prepare("update usuario set Senha = ?  where Id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(senha, idUsuario); erro != nil {
		return erro
	}

	return nil
}
