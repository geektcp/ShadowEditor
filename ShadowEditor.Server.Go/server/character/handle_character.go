package character

import (
	"net/http"

	shadow "github.com/tengge1/shadoweditor"
	"github.com/tengge1/shadoweditor/context"
	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/model"
	"github.com/tengge1/shadoweditor/model/category"
	"github.com/tengge1/shadoweditor/model/character"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	audio := Character{}
	context.Mux.UsingContext().Handle(http.MethodGet, "/api/Character/List", audio.List)
	context.Mux.UsingContext().Handle(http.MethodGet, "/api/Character/Get", audio.Get)
	context.Mux.UsingContext().Handle(http.MethodPost, "/api/Character/Edit", audio.Edit)
	context.Mux.UsingContext().Handle(http.MethodPost, "/api/Character/Save", audio.Save)
	context.Mux.UsingContext().Handle(http.MethodPost, "/api/Character/Delete", audio.Delete)
}

// Character 人物控制器
type Character struct {
}

// List 获取列表
func (Character) List(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db, err := context.Mongo()
	if err != nil {
		helper.WriteJSON(w, model.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	// 获取所有类别
	filter := bson.M{
		"Type": "Character",
	}
	categories := []category.Model{}
	db.FindMany(shadow.CategoryCollectionName, filter, &categories)

	docs := bson.A{}

	opts := options.FindOptions{
		Sort: bson.M{
			"_id": -1,
		},
	}

	if context.Config.Authority.Enabled {
		user, _ := context.GetCurrentUser(r)

		if user != nil {
			filter1 := bson.M{
				"UserID": user.ID,
			}

			if user.Name == "Administrator" {
				filter2 := bson.M{
					"UserID": bson.M{
						"$exists": 0,
					},
				}
				filter1 = bson.M{
					"$or": bson.A{
						filter1,
						filter2,
					},
				}
			}
			db.FindMany(shadow.CharacterCollectionName, filter1, &docs, &opts)
		}
	} else {
		db.FindAll(shadow.CharacterCollectionName, &docs, &opts)
	}

	list := []character.Model{}
	for _, i := range docs {
		doc := i.(primitive.D).Map()
		categoryID := ""
		categoryName := ""

		if doc["Category"] != nil {
			for _, category := range categories {
				if category.ID == doc["Category"].(string) {
					categoryID = category.ID
					categoryName = category.Name
					break
				}
			}
		}

		thumbnail, _ := doc["Thumbnail"].(string)

		info := character.Model{
			ID:           doc["_id"].(primitive.ObjectID).Hex(),
			Name:         doc["Name"].(string),
			CategoryID:   categoryID,
			CategoryName: categoryName,
			TotalPinYin:  helper.PinYinToString(doc["TotalPinYin"]),
			FirstPinYin:  helper.PinYinToString(doc["FirstPinYin"]),
			CreateTime:   doc["CreateTime"].(primitive.DateTime).Time(),
			UpdateTime:   doc["UpdateTime"].(primitive.DateTime).Time(),
			Thumbnail:    thumbnail,
		}

		list = append(list, info)
	}

	helper.WriteJSON(w, model.Result{
		Code: 200,
		Msg:  "Get Successfully!",
		Data: list,
	})
}

// Get 获取
func (Character) Get(w http.ResponseWriter, r *http.Request) {

}

// Edit 编辑
func (Character) Edit(w http.ResponseWriter, r *http.Request) {

}

// Save 保存
func (Character) Save(w http.ResponseWriter, r *http.Request) {

}

// Delete 删除
func (Character) Delete(w http.ResponseWriter, r *http.Request) {

}