package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type CollectionSqliteRepo struct {
	DB *sqlx.DB
}

func (r *CollectionSqliteRepo) Migrate() error {
	q := `
	CREATE TABLE IF NOT EXISTS collections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT NOT NULL,
		user_id INT NOT NULL,
		CONSTRAINT fk_users
		FOREIGN KEY (user_id)
		REFERENCES users (id)
		ON DELETE CASCADE
	)
	`
	fmt.Println("Migrating collections table...")
	_, err := r.DB.Exec(q)
	if err != nil {
		fmt.Println("error migrating collections: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating collections")
	}

	qa := `
	CREATE TABLE IF NOT EXISTS collection_articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		article_id INTEGER NOT NULL,
		collection_id INTEGER NOT NULL,
		CONSTRAINT fk_articles
			FOREIGN KEY (article_id)
				REFERENCES articles(id)
				ON DELETE CASCADE,
		CONSTRAINT fk_collections
			FOREIGN KEY(collection_id)
				REFERENCES collections(id)
				ON DELETE CASCADE
	);
	`
	fmt.Println("Migrating collection_articles table...")
	_, err = r.DB.Exec(qa)
	if err != nil {
		fmt.Println("error migrating collections_articles: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating collection_articles")
	}
	return err
}

func NewCollectionSqliteRepo(db *sqlx.DB) *CollectionSqliteRepo {
	return &CollectionSqliteRepo{DB: db}
}

func (r *CollectionSqliteRepo) Get(id int64) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) Create(c *model.Collection) (*model.Collection, error) {
	q := `
	INSERT INTO collections(title, slug, user_id)
		VALUES(?, ?, ?);
	`
	res, err := r.DB.Exec(q, c.Title, c.Slug, c.UserId)

	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	c.Id = id
	return c, nil
}

func (r *CollectionSqliteRepo) All(userId int64) ([]*model.Collection, error) {
	ss := []*model.Collection{}
	err := r.DB.Select(&ss, "SELECT * FROM collections WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (r *CollectionSqliteRepo) Update(s *model.Collection) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) Delete(id int64) error {
	return nil
}

func (r *CollectionSqliteRepo) FindByTitle(title string) (*model.Collection, error) {
	coll := &model.Collection{}
	err := r.DB.Get(coll, "SELECT * FROM collections WHERE title=?", title)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

func (r *CollectionSqliteRepo) FindBySlug(slug string) (*model.Collection, error) {
	return nil, nil
}

func (r *CollectionSqliteRepo) InsertCollectionItem(itemId int64, collectionId int64) error {
	q := `
	INSERT INTO collection_articles(article_id, collection_id)
		VALUES(?, ?)
	`
	_, err := r.DB.Exec(q, itemId, collectionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionSqliteRepo) InsertManyCollectionItems(aa []*model.Article, cId int64) error {
	for _, a := range aa {
		err := r.InsertCollectionItem(a.ID, cId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CollectionSqliteRepo) GetArticlesByCollectionName(collectionName string) []*model.Article {
	// q := `
	// SELECT * FROM articles
	// 	INNER JOIN
	// `
	return nil
}
