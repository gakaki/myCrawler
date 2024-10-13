package gt7

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"myCrawler/utils"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var carList []Car
var client *resty.Client = resty.New().
	SetRetryCount(3).
	SetRetryWaitTime(2 * time.Second).
	SetRetryMaxWaitTime(10 * time.Second)
var mongoClient *mongo.Client
var carCollection *mongo.Collection
var gormDB *gorm.DB
var mutex sync.Mutex
var baseUrl = "https://gran-turismo.fandom.com"
var carListURL = baseUrl + "/wiki/Gran_Turismo_7/Car_List"

func crawlGT7AllCars() {
	// 初始化 resty.Client, MongoDB 和 GORM 连接

	//initMongoDB()
	//initGORM()

	// Step 1: 爬取首页并保存所有的详情页面信息
	fmt.Println("Step 1: Fetching car list from main page...")
	err := fetchCarList(carListURL)
	if err != nil {
		log.Fatalf("Failed to fetch car list: %v", err)
	}

	// // 保存初步的车列表到 MongoDB 和 GORM 中
	// saveCarListToMongo()
	// saveCarListToGORM()

	// Step 2: 并发爬取详情页并更新详细信息
	//fmt.Println("Step 2: Fetching car details...")
	//err = loadCarListFromMongo()
	//if err != nil {
	//	log.Fatalf("Failed to load car list from MongoDB: %v", err)
	//}

	carList = carList[0:1]
	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second / 5) // 控制每秒 5 个并发
	defer ticker.Stop()

	for i := range carList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-ticker.C
			fetchCarDetails(&carList[i])
		}(i)
	}
	wg.Wait()
	fmt.Println(len(carList))
	// // Step 3: 保存最终的汽车信息到 MongoDB 和 GORM
	// saveCarListToMongo()
	// saveCarListToGORM()

	// fmt.Println("Car details saved to MongoDB and GORM.")
}

// initMongoDB 初始化 MongoDB 连接
func initMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://iloveNewipad3@127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	carCollection = mongoClient.Database("gran_turismo").Collection("cars")
}

// initGORM 初始化 GORM 数据库连接
func initGORM() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/gran_turismo?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to GORM database: %v", err)
	}

	// 自动迁移，确保表结构与模型保持一致
	err = gormDB.AutoMigrate(&Car{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}

// fetchCarList 爬取主页面并提取汽车信息
func fetchCarList(url string) error {
	resp, err := client.R().Get(url)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return err
	}
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		brandName := s.Find("span.mw-headline").Text()
		brandName = strings.TrimSpace(brandName)
		if brandName == "" {
			return // Skip to the next iteration
		}
		fmt.Println("start crawl brand ....", brandName)

		tableDoc := s.Next()
		if tableDoc == nil {
			return
		}
		tableDoc.Find("tr").Each(func(i int, d *goquery.Selection) {

			carName := d.Find("td:nth-child(1)").Text()
			if carName == "" {
				return
			}
			car := Car{}
			car.Name = strings.TrimSpace(carName)
			car.Brand = strings.TrimSpace(brandName)
			detailLink, _ := d.Find("a").Attr("href")
			car.DetailURL = baseUrl + strings.TrimSpace(detailLink)

			car.GrLevel = strings.TrimSpace((d.Find("td:nth-child(2)").Text()))
			car.ImageURL, _ = d.Find("td:nth-child(3) a").Attr("href")
			car.ImageURL = strings.TrimSpace(car.ImageURL)

			mutex.Lock()
			carList = append(carList, car)
			mutex.Unlock()
		})

	})

	utils.WriteJSON(carList, "cars_json/cars.json")
	return nil
}

// fetchCarDetails 并发爬取详情页面的详细信息
func fetchCarDetails(car *Car) {
	resp, err := client.R().Get(car.DetailURL)
	if err != nil {
		fmt.Printf("Failed to fetch car details for %s: %v\n", car.Name, err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		fmt.Printf("Failed to parse car details for %s: %v\n", car.Name, err)
		return
	}

	car.Weight = strings.TrimSpace(doc.Find("div[data-source='weight']").Text())
	car.Performance = strings.TrimSpace(doc.Find("div[data-source='pp']").Text()) // Performance Points
	car.Description = strings.TrimSpace(doc.Find(".mw-parser-output p").First().Text())

	car.Manufacturer = strings.TrimSpace(doc.Find("div[data-source='manufacturer']").Text())
	car.Category = strings.TrimSpace(doc.Find("div[data-source='category']").Text())
	car.Engine = strings.TrimSpace(doc.Find("div[data-source='engine']").Text())
	car.Aspiration = strings.TrimSpace(doc.Find("div[data-source='aspiration']").Text())
	car.Power = strings.TrimSpace(doc.Find("div[data-source='power']").Text())
	car.Torque = strings.TrimSpace(doc.Find("div[data-source='torque']").Text())
	car.Drivetrain = strings.TrimSpace(doc.Find("div[data-source='drivetrain']").Text())
	car.Length = strings.TrimSpace(doc.Find("div[data-source='length']").Text())
	car.Width = strings.TrimSpace(doc.Find("div[data-source='width']").Text())
	car.Height = strings.TrimSpace(doc.Find("div[data-source='height']").Text())

	mutex.Lock()
	defer mutex.Unlock()
	for i, c := range carList {
		if c.Name == car.Name {
			carList[i] = *car
			break
		}
	}
}

// saveCarListToMongo 将汽车列表保存到 MongoDB
func saveCarListToMongo() {
	ctx := context.Background()
	for _, car := range carList {
		filter := bson.M{"name": car.Name}
		update := bson.M{
			"$set": car,
		}
		opts := options.Update().SetUpsert(true)
		_, err := carCollection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Failed to upsert car %s: %v", car.Name, err)
		}
	}
}

// saveCarListToGORM 将汽车列表保存到 GORM 数据库
func saveCarListToGORM() {
	for _, car := range carList {
		err := gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&car).Error
		if err != nil {
			log.Printf("Failed to save car %s to GORM: %v", car.Name, err)
			err := gormDB.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&car).Error
			if err != nil {
				log.Printf("Failed to save car %s to GORM: %v", car.Name, err)
			}
		}
	}
}

// loadCarListFromMongo 从 MongoDB 中加载汽车列表
func loadCarListFromMongo() error {
	ctx := context.Background()
	cursor, err := carCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var car Car
		err := cursor.Decode(&car)
		if err != nil {
			return err
		}
		carList = append(carList, car)
	}

	if err := cursor.Err(); err != nil {
		return err
	}
	return nil
}

// saveCarListToJSON 将汽车数据保存为 JSON 文件
func saveCarListToJSON(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(carList)
	if err != nil {
		return fmt.Errorf("failed to encode car list to JSON: %v", err)
	}
	fmt.Printf("Car list saved to %s\n", fileName)
	return nil
}
