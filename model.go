package main

import (
	"database/sql"
	"errors"
)

type blogs struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PostContent string `json:"post_content"`
	Author      string `json:"author"`
}

func (b *blogs) getBlog(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *blogs) updateBlog(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *blogs) deleteBlog(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *blogs) createBlog(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getBlogs(db *sql.DB, start int, count int) ([]blogs, error) {
	return nil, errors.New("Not implemented")
}
