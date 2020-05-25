package database

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Description   string
	Image         string
	DistributorID uint
	Weight        uint
	Price         uint
	PackageID     uint
	TypeID        uint
}

type MoreProduct struct {
	Product
	Distributor
	Package
	Type
}

func GetProducts(typeID, distributorID uint) ([]MoreProduct, error) {
	products := []MoreProduct{}
	var err error
	if typeID == 0 && distributorID != 0 {
		err = gormDB.Raw(
			"SELECT * FROM products p INNER JOIN distributors d ON p.distributor_id=d.id INNER JOIN packages k ON p.package_id=k.id INNER JOIN types t ON p.type_id=t.id WHERE p.deleted_at IS NULL AND d.deleted_at IS NULL AND k.deleted_at IS NULL AND t.deleted_at IS NULL AND distributor_id = ?",
			distributorID).Scan(&products).Error
	} else if distributorID == 0 && typeID != 0 {
		err = gormDB.Raw(
			"SELECT * FROM products p INNER JOIN distributors d ON p.distributor_id=d.id INNER JOIN packages k ON p.package_id=k.id INNER JOIN types t ON p.type_id=t.id WHERE p.deleted_at IS NULL AND d.deleted_at IS NULL AND k.deleted_at IS NULL AND t.deleted_at IS NULL AND type_id = ?",
			typeID).Scan(&products).Error
	} else if distributorID == 0 && typeID == 0 {
		err = gormDB.Raw(
			"SELECT * FROM products p INNER JOIN distributors d ON p.distributor_id=d.id INNER JOIN packages k ON p.package_id=k.id INNER JOIN types t ON p.type_id=t.id WHERE p.deleted_at IS NULL AND d.deleted_at IS NULL AND k.deleted_at IS NULL AND t.deleted_at IS NULL",
		).Scan(&products).Error
	} else {
		err = gormDB.Raw(
			"SELECT * FROM products p INNER JOIN distributors d ON p.distributor_id=d.id INNER JOIN packages k ON p.package_id=k.id INNER JOIN types t ON p.type_id=t.id WHERE p.deleted_at IS NULL AND d.deleted_at IS NULL AND t.deleted_at IS NULL AND k.deleted_at IS NULL AND type_id = ? AND distributor_id = ?",
			typeID, distributorID).Scan(&products).Error
	}
	return products, err
}

func GetProductsOffsetLimit(offset, limit uint) ([]MoreProduct, error) {
	products := []MoreProduct{}
	err := gormDB.Raw(
		"SELECT * FROM products p INNER JOIN distributors d ON p.distributor_id=d.id INNER JOIN packages k ON p.package_id=k.id INNER JOIN types t ON p.type_id=t.id WHERE p.deleted_at IS NULL AND d.deleted_at IS NULL AND t.deleted_at IS NULL AND k.deleted_at IS NULL LIMIT ? OFFSET ?",
		limit, offset).Scan(&products).Error
	return products, err
}

func GetProductsCount() uint {
	count := 0
	gormDB.Model(&Product{}).Count(&count)
	return uint(count)
}

func GetProductByID(id uint) (Product, error) {
	products := Product{}
	err := gormDB.Where("id = ?", id).First(&products).Error
	return products, err
}

func DeleteProducts(id uint) error {
	pr := Product{}
	gormDB.Where("id = ?", id).First(&pr)
	os.Remove("../web/res/" + pr.Image)
	return gormDB.Where("id = ?", id).Delete(&Product{}).Error
}

func CreateProduct(r *http.Request) error {
	image, err := SaveImage(r.PostFormValue("image"))
	if err != nil {
		return err
	}
	disr, _ := strconv.Atoi(r.PostFormValue("distributor"))
	weig, _ := strconv.Atoi(r.PostFormValue("weight"))
	pric, _ := strconv.Atoi(r.PostFormValue("price"))
	typs, _ := strconv.Atoi(r.PostFormValue("type"))
	pkg, _ := strconv.Atoi(r.PostFormValue("package"))
	return gormDB.Create(&Product{
		Name:          r.PostFormValue("name"),
		Description:   r.PostFormValue("description"),
		Image:         image,
		DistributorID: uint(disr),
		Weight:        uint(weig),
		Price:         uint(pric),
		TypeID:        uint(typs),
		PackageID:     uint(pkg),
	}).Error
}

func UpdateProduct(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	pr, err := GetProductByID(uint(id))
	image := pr.Image
	if r.PostFormValue("image") != "" {
		os.Remove("../web/res/" + pr.Image)
		image, err = SaveImage(r.PostFormValue("image"))
		if err != nil {
			return err
		}
	}
	disr, _ := strconv.Atoi(r.PostFormValue("distributor"))
	weig, _ := strconv.Atoi(r.PostFormValue("weight"))
	pric, _ := strconv.Atoi(r.PostFormValue("price"))
	typs, _ := strconv.Atoi(r.PostFormValue("type"))
	pkg, _ := strconv.Atoi(r.PostFormValue("package"))
	pr = Product{
		Name:          r.PostFormValue("name"),
		Description:   r.PostFormValue("description"),
		Image:         image,
		DistributorID: uint(disr),
		Weight:        uint(weig),
		Price:         uint(pric),
		TypeID:        uint(typs),
		PackageID:     uint(pkg),
	}
	pr.ID = uint(id)
	return gormDB.Save(&pr).Error
}

func SaveImage(url string) (string, error) {
	name := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < 32; i++ {
		name += string(97 + rand.Intn(26))
	}
	response, e := http.Get(url)
	if e != nil {
		return name, e
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("../web/res/" + name)
	if err != nil {
		return name, err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return name, err
	}
	return name, nil
}
