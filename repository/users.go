package repository

import (
	"context"
	"go-fiber-auth-api/db"
	"go-fiber-auth-api/models"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Delete(id string) error
}

type usersRepository struct {
	c *qmgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	ctx := context.Background()
	_, err := r.c.InsertOne(ctx, user)
	return err
}

func (r *usersRepository) Update(user *models.User) error {
	ctx := context.Background()
	return r.c.UpdateId(ctx, user.Id, user)
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	ctx := context.Background()
	err = r.c.Find(ctx, bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	ctx := context.Background()
	err = r.c.Find(ctx, bson.M{"email": email}).One(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	ctx := context.Background()
	err = r.c.Find(ctx, bson.M{}).All(&users)
	return users, err
}

func (r *usersRepository) Delete(id string) error {
	ctx := context.Background()
	return r.c.RemoveId(ctx, bson.ObjectIdHex(id))
}
