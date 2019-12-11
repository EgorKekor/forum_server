package database

import (
	"database/sql"

	"db-forum/models"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	pg *sql.DB

	CreateUserStmt        *sql.Stmt
	GetUserStmt           *sql.Stmt
	GetUserByUsernameStmt *sql.Stmt
	UpdateUserStmt        *sql.Stmt

	CreateForumStmt             *sql.Stmt
	GetForumStmt                *sql.Stmt
	GetForumThreadsStmt         *sql.Stmt
	GetForumThreadsWithTimeStmt *sql.Stmt

	CreateThreadStmt    *sql.Stmt
	GetThreadStmt       *sql.Stmt
	GetThreadByIDStmt   *sql.Stmt
	GetThreadBySlugStmt *sql.Stmt

	CreatePostStmt  *sql.Stmt
	GetPostByIDStmt *sql.Stmt

	GetPrevVoteThreadStmt *sql.Stmt
	CreatVoteThreadStmt   *sql.Stmt
	UpdateVoteThreadStmt  *sql.Stmt
	BigInsert *sql.Stmt
}

var (
	db           *DB
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate")
)

func InitDB(DSN string) error {
	var err error
	var nDB DB
	nDB.pg, err = sql.Open("postgres", DSN)
	if err != nil {
		return errors.Wrap(err, "can't open database")
	}
	if err = nDB.pg.Ping(); err != nil {
		return errors.Wrap(err, "can't connect to database")
	}
	db = &nDB
	if err = initStmts(); err != nil {
		return errors.Wrap(err, "can't prepare statements")
	}
	if _, err = db.pg.Exec(clearDB); err != nil {
		return errors.Wrap(err, "can't clear db")
	}
	return nil
}


func initStmts() error {
	db.CreateUserStmt, _ = db.pg.Prepare(createUser)
	db.GetUserStmt, _ = db.pg.Prepare(getUser)
	db.UpdateUserStmt, _ = db.pg.Prepare(updateUser)
	db.GetUserByUsernameStmt, _ = db.pg.Prepare(getUserByUsername)


	db.CreateForumStmt, _ = db.pg.Prepare(createForum)
	db.GetForumStmt, _ = db.pg.Prepare(getForum)
	db.GetForumThreadsWithTimeStmt, _ = db.pg.Prepare(getForumThreadsWithTime)
	db.GetForumThreadsStmt, _ = db.pg.Prepare(getForumThreads)


	db.CreateThreadStmt, _ = db.pg.Prepare(createThread)
	db.GetThreadStmt, _ = db.pg.Prepare(getThread)
	db.GetThreadByIDStmt, _ = db.pg.Prepare(getThreadByID)
	db.GetThreadBySlugStmt, _ = db.pg.Prepare(getThreadBySlug)


	db.CreatePostStmt, _ = db.pg.Prepare(createPost)
	db.GetPostByIDStmt, _ = db.pg.Prepare(getPostByID)


	db.BigInsert, _ = db.pg.Prepare(bigInsert)
	db.UpdateVoteThreadStmt, _ = db.pg.Prepare(updateVoteThread)
	db.CreatVoteThreadStmt, _ = db.pg.Prepare(createVoteThread)
	return nil
}

func ClearTable() {
	db.pg.Exec(clearDB)
}

var clearDB = `DELETE FROM users; DELETE FROM forum; DELETE FROM thread; DELETE FROM post; DELETE FROM voice;`

func GetStatus() *models.Status {
	var status models.Status
	db.pg.QueryRow(`SELECT count(*) FROM users;`).Scan(&status.User)
	db.pg.QueryRow(`SELECT count(*) FROM thread;`).Scan(&status.Thread)
	db.pg.QueryRow(`SELECT count(*) FROM post;`).Scan(&status.Post)
	db.pg.QueryRow(`SELECT count(*) FROM forum;`).Scan(&status.Forum)
	return &status
}
