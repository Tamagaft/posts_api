package repository

import (
	"posts/internals/entity"

	"github.com/jmoiron/sqlx"
)

type UserPostPSQL struct {
	db *sqlx.DB
}

const postsRange = 10

func NewUserPostPSQL(db *sqlx.DB) *UserPostPSQL {
	return &UserPostPSQL{db: db}
}

func (r UserPostPSQL) CreatePost(post entity.Post) error {
	stmt, err := r.db.Prepare("INSERT INTO posts(text,date,author_pk) VALUES($1,$2,$3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Text, post.Date, post.Author)
	if err != nil {
		return err
	}
	return nil
}

func (r UserPostPSQL) GetPostById(postId int) (*entity.Post, error) {
	var post entity.Post
	stmt, err := r.db.Prepare("SELECT * FROM posts WHERE id = $1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(postId)
	if err = row.Scan(&post.Id, &post.Text, &post.Date, &post.Author); err != nil {
		return nil, err
	}
	return &post, nil
}

func (r UserPostPSQL) GetUserPostsRange(userId, part int) ([]entity.Post, error) {
	var posts []entity.Post
	stmt, err := r.db.Prepare("SELECT * FROM posts WHERE author_pk = $1 offset $2 limit $3")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(userId, postsRange*part, postsRange)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := entity.Post{}
		err := rows.Scan(&post.Id, &post.Text, &post.Date, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
