package model
import (
	"fmt"
	"errors"
	"github.com/boltdb/bolt"
	"strconv"
)

func GetUserByName(username string) (bool, User){
	user := User{}
	db, err := bolt.Open("user.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return false, user
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(username))
		if b == nil {
			return errors.New("不存在当前用户")
		}
		user = User{
			Password: string(b.Get([]byte("password"))),
			Username: string(b.Get([]byte("username"))),
		}
		return nil
	}); err != nil {
		return false, user
	}
	return true, user
}

func AddUser(user User) {
	db, err := bolt.Open("user.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	fmt.Println("is going to insert ", user)
	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(user.Username)); err != nil {
			return err
		}
		userBucket := tx.Bucket([]byte(user.Username))
		if err := userBucket.Put([]byte("username"), []byte(user.Username)); err != nil {
			return err
		}
		if err := userBucket.Put([]byte("password"), []byte(user.Password)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("update failed!")
	}
}
func CreateArticle(username string, article Article) {
	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(article.Id)); err != nil {
			return err
		}
		art := tx.Bucket([]byte(article.Id))
		if err := art.Put([]byte("Id"), []byte(article.Id)); err != nil {
			return err
		}
		if err := art.Put([]byte("Name"), []byte(article.Name)); err != nil {
			return err
		}
		if err := art.Put([]byte("Title"), []byte(article.Title)); err != nil {
			return err
		}
		if err := art.Put([]byte("Date"), []byte(article.Date)); err != nil {
			return err
		}
		if err := art.Put([]byte("Content"), []byte(article.Content)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("update failed!")
	}
}
func GetArticleList(username string) (bool, []ArticleList){
	var artlist []ArticleList
	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return false, artlist
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error 	{
		c := tx.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			artlist = append(artlist,
				ArticleList{
					Id: 	string(tx.Bucket([]byte(k)).Get([]byte("id"))),
					Name: 	string(tx.Bucket([]byte(k)).Get([]byte("Name"))),
					Title: 	string(tx.Bucket([]byte(k)).Get([]byte("Title"))),
				})
			return nil
		}
		return nil
	})
	if err != nil {
		return false, artlist
	}
	return true, artlist
}
func GetNextId(username string) string{
	var num string
	var id int = 0
	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return num
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		for k,_ := c.First(); k!= nil; k,_=c.Next() {
			id = id + 1
		}
		return nil
	})
	if err != nil {
		return num
	}
	return strconv.Itoa(id)
}
func GetAllArticleById(username string, id string) (bool, Article) {
	art := Article{}
	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		fmt.Println("open failed")
		return false, art
	} else {
		fmt.Println("open succeed!")
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
	        if id == string(tx.Bucket([]byte(k)).Get([]byte("id"))) {
	        	art = Article{
				Id: id,
				Name: 	string(tx.Bucket([]byte(k)).Get([]byte("Name"))),
				Title: 	string(tx.Bucket([]byte(k)).Get([]byte("Title"))),
				Date:	string(tx.Bucket([]byte(k)).Get([]byte("Date"))),
				Content:string(tx.Bucket([]byte(k)).Get([]byte("Content"))),
				}
			}
			return nil
		}
		return nil
    })
    if err != nil {
    	return false, art
    }
	return true, art
}