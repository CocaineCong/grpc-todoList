package repository

import (
	"favorites/internal/service"
	"fmt"
	"testing"
)

func TestCheckExist(t *testing.T) {
	InitDB()
	f := new(Favorites)
	req := new(service.FavoritesRequest)
	exist:=f.CheckExist(req)
	fmt.Println(exist)
}

func TestFavorites_Create(t *testing.T) {
	InitDB()
	f := new(Favorites)
	req := new(service.FavoritesRequest)
	req.FavoriteName="38324"
	req.UserID=4
	err := f.Create(req)
	fmt.Println(err)
}

func TestFavorites_Show(t *testing.T) {
	InitDB()
	f := new(Favorites)
	req := new(service.FavoritesRequest)
	req.UserID=4
	res,err := f.Show(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestFavorites_Update(t *testing.T) {
	InitDB()
	f := new(Favorites)
	req := new(service.FavoritesRequest)
	req.FavoriteID=1
	req.FavoriteName="knowledge again"
	err := f.Update(req)
	fmt.Println(err)
}

func TestFavorites_Delete(t *testing.T) {
	InitDB()
	f := new(Favorites)
	req := new(service.FavoritesRequest)
	req.FavoriteID=1
	err := f.Delete(req)
	fmt.Println(err)
}
