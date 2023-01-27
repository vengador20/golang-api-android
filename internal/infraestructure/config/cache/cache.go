package cache

// import (
// 	"context"
// 	"fiberapi/database/sqlite"
// 	"fmt"
// )

// // var (
// // 	one   sync.Once
// // 	cache *Cache
// // )

// type Cache struct {
// 	//cache map[string]interface{}
// 	Ctx context.Context
// }

// func (c *Cache) GetCache(nombre string) (string, error) {
// 	db := sqlite.Open()
// 	db.Ctx = c.Ctx
// 	quer, err := db.Query(nombre)

// 	if err != nil {
// 		return "", fmt.Errorf(err.Error())
// 	}

// 	res := quer["datos"].(string)

// 	return res, nil
// }

// func (c *Cache) SaveCache(nombre string, datos interface{}) {
// 	db := sqlite.Open()
// 	//db.Ctx = c.Ctx
// 	db.Insert(nombre, datos)
// }

// func (c *Cache) ClearCache() {
// 	db := sqlite.Open()
// 	db.Clear()
// }
