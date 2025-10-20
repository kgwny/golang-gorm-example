package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	// 構造体の中に gorm.Model を記述すると、id と CreatedAt と UpdatedAt と DeletedAt が作られる
	gorm.Model

	Name     string
	Age      int
	IsActive bool
}

type Products struct {
	Code    string
	Price   uint
	Deleted gorm.DeletedAt
}

func main() {
	// db を作成する
	db := dbInit()

	// db を migrate する
	db.AutoMigrate(&Users{})
}

func dbInit() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	return db
}

// レコードを1件登録する
func insert(db *gorm.DB) {
	users := Users{
		Name:     "一郎",
		Age:      20,
		IsActive: true,
	}

	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// レコードを複数件登録する
func inserts(db *gorm.DB) {
	users := []Users{
		{
			Name:     "次郎",
			Age:      19,
			IsActive: false,
		},
		{
			Name:     "花子",
			Age:      18,
			IsActive: true,
		},
		{
			Name:     "タマ",
			Age:      3,
			IsActive: true,
		},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 単体取得
func getOne(db *gorm.DB) {
	// 昇順で単体取得
	users1 := Users{}
	result1 := db.First(&users1)

	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println("first:", users1)

	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result1.Error)
	}
	fmt.Println("count:", result1.RowsAffected)

	// 何も設定せずに単体取得
	users2 := Users{}
	result2 := db.Take(&users2)

	// SELECT + FROM users LIMIT 1;
	fmt.Println("take:", users2)
	if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result2.Error)
	}

	// 降順で単体取得
	users3 := Users{}
	result3 := db.Last(&users3)

	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("last:", users3)

	if errors.Is(result3.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result3.Error)
	}

	// 取得方法別
	// First・・・PKの昇順で取得する、PKがない場合は、モデルの最初のフィールドで順序付けされる
	// Last・・・PKの降順で取得する、PKがない場合は、モデルの最初のフィールドで順序付けされる
	// Take・・・特に条件を指定せずに取得

	// 他の取得条件
	// プライマリーキーで取得する場合
	//db.First(&users1, 1)
	//db.First(&users1, "id = ?", 1)
}

// 全件取得
func find(db *gorm.DB) {
	users := []Users{}
	result := db.Find(&users)

	fmt.Println("user:", users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 更新
func save(db *gorm.DB) {
	// 構造体に id が無い場合は insert される
	user1 := Users{}
	user1.Name = "花子"

	result1 := db.Save(&user1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println("count:", result1.RowsAffected)
	fmt.Println("user1", user1)

	user2 := Users{}
	db.First(&user2)

	// 構造体にidがある場合は update される
	user2.Name = "太郎"
	result2 := db.Save(&user2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println("count:", result2.RowsAffected)
	fmt.Println("user2:", user2)
}

// 単一のカラムを更新する場合
func update(db *gorm.DB) {
	result := db.Model(&Users{}).Where("id = 2").Update("name", "サブロー")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	user := Users{}
	db.Where("id = 2").Take(&user)
	fmt.Println("user:", user)
}

// 複数のカラムを更新する場合
func updates(db *gorm.DB) {
	result := db.Model(&Users{}).Where("id = 1").Updates(Users{Name: "Taro", Age: 9, IsActive: true})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	user := Users{}
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)
}

// 一括更新
func updateAll(db *gorm.DB) {
	user := Users{
		Name:     "ボブ",
		Age:      40,
		IsActive: true,
	}
	result := db.Where("name = ?", "花子").Updates(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	users := []Users{}
	db.Find(&users)
	fmt.Println("users:", users)
}

func noUpdates(db *gorm.DB) {
	// 以下の場合 IsActive が更新されない
	// result := db.Model(Users{}).Where("id = 1").Updates(Users{Name: "まさお", IsActive: false})

	// 以下の場合 IsActive が更新される -> Selectを追加して明示的に更新対象カラムを指定することができる
	result := db.Model(Users{}).Where("id = 1").Select("name", "is_active").Updates(Users{Name: "まさお", IsActive: false})

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)

	user := Users{}
	db.Where("id = 1").Take(&user)
	fmt.Println("user:", user)
}

// 論理削除
func logicalDelete(db *gorm.DB) {
	// id = 1 を条件にしたとき
	//db.Where("id = 1").Delete(&Users{})
	// name = "太郎" を条件にしたとき
	db.Where("name = ?", "太郎").Delete(&Users{})

	// ID=1 のユーザーを論理削除
	if err := db.Delete(&Users{}, 1).Error; err != nil {
		log.Println("Soft delete failed:", err)
	}
}

// 物理削除
func delete(db *gorm.DB) {
	db.Unscoped().Where("id = 1").Delete(&Users{})

	// ID=1 のユーザーを物理削除
	if err := db.Unscoped().Delete(&Users{}, 1).Error; err != nil {
		log.Println("Delete failed:", err)
	}
}
