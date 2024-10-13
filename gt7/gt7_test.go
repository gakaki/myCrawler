package gt7

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCrawlGT7AllCarsReal(t *testing.T) {
	crawlGT7AllCars()
}

// TestCrawlGT7AllCars 测试 crawlGT7AllCars 函数
func TestCrawlGT7AllCarsMock(t *testing.T) {
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	fileContent, err := os.ReadFile("car_mock/car_list.html")
	if err != nil {
		fmt.Errorf("error reading mock file: %v", err)
	}
	httpmock.RegisterResponder("GET", carListURL, httpmock.NewStringResponder(200, string(fileContent)))

	fileContent, err = os.ReadFile("car_mock/car_detail.html")
	if err != nil {
		fmt.Errorf("error reading mock file: %v", err)
	}
	detailUrl := "https://gran-turismo.fandom.com/wiki/Abarth_500_%2709"
	httpmock.RegisterResponder("GET", detailUrl, httpmock.NewStringResponder(200, string(fileContent)))

	crawlGT7AllCars()
}

// 初始化 MongoDB 测试客户端
func TestInitMongoDB(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mt.Close()

	mt.Run("connect to MongoDB", func(mt *mtest.T) {
		initMongoDB() // 你可以使用 mongo-go-driver mtest 模拟 MongoDB 连接
		assert.NotNil(t, mongoClient)
		assert.NotNil(t, carCollection)
	})
}

// 初始化 GORM 测试数据库
func TestInitGORM(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize sqlmock: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize GORM: %v", err)
	}

	// Mock GORM auto migration
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS `cars`").WillReturnResult(sqlmock.NewResult(1, 1))

	initGORM()
	assert.NotNil(t, gormDB)
}

//
//// 测试 fetchCarList 函数，使用 resty 的 mock
//func TestFetchCarList(t *testing.T) {
//	client = resty.New()
//	mockResponse := `<table class="article-table">
//	<tr>
//		<td><img src="car-image.jpg" /></td>
//		<td><a href="/wiki/Car_Detail_Page">Car Name</a></td>
//		<td>Car Brand</td>
//		<td>Gr.3</td>
//	</tr></table>`
//
//	// Mock HTTP request
//	resp := &resty.Response{
//		RawResponse: &http.Response{},
//	}
//	resp.SetBody([]byte(mockResponse))
//
//	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
//		response.RawResponse.Body = strings.NewReader(mockResponse)
//		return nil
//	})
//
//	err := fetchCarList("https://mockurl.com")
//	assert.NoError(t, err)
//
//	assert.Equal(t, 1, len(carList))
//	assert.Equal(t, "Car Name", carList[0].Name)
//	assert.Equal(t, "Car Brand", carList[0].Brand)
//	assert.Equal(t, "Gr.3", carList[0].GrLevel)
//	assert.Equal(t, "https://gran-turismo.fandom.com/wiki/Car_Detail_Page", carList[0].DetailURL)
//}

//
//// 测试 fetchCarDetails 函数
//func TestFetchCarDetails(t *testing.T) {
//	client = resty.New()
//	mockDetailResponse := `
//	<div data-source="weight">1230 kg</div>
//	<div data-source="pp">630.5</div>
//	<div data-source="manufacturer">Porsche</div>
//	<div data-source="category">Sports Car</div>
//	<div data-source="engine">3.8L</div>
//	<div data-source="power">385 hp</div>
//	<div data-source="torque">420 Nm</div>
//	<div data-source="drivetrain">RWD</div>
//	<div data-source="length">4456 mm</div>
//	<div data-source="width">1998 mm</div>
//	<div data-source="height">1282 mm</div>
//	`
//
//	resp := &resty.Response{
//		RawResponse: &http.Response{},
//	}
//	resp.SetBody([]byte(mockDetailResponse))
//
//	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
//		response.RawResponse.Body = strings.NewReader(mockDetailResponse)
//		return nil
//	})
//
//	car := &Car{Name: "Test Car", DetailURL: "https://mockurl.com/details"}
//	fetchCarDetails(car)
//
//	assert.Equal(t, "1230 kg", car.Weight)
//	assert.Equal(t, "630.5", car.Performance)
//	assert.Equal(t, "Porsche", car.Manufacturer)
//	assert.Equal(t, "385 hp", car.Power)
//	assert.Equal(t, "420 Nm", car.Torque)
//}
//
//// 测试 saveCarListToMongo 函数
//func TestSaveCarListToMongo(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//	defer mt.Close()
//
//	mt.Run("save car list to MongoDB", func(mt *mtest.T) {
//		initMongoDB()
//
//		mt.AddMockResponses(mtest.CreateSuccessResponse())
//		err := saveCarListToMongo()
//		assert.NoError(t, err)
//	})
//}
//
//// 测试 saveCarListToGORM 函数
//func TestSaveCarListToGORM(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Failed to initialize sqlmock: %v", err)
//	}
//	defer db.Close()
//
//	dialector := mysql.New(mysql.Config{
//		Conn: db,
//	})
//	gormDB, err = gorm.Open(dialector, &gorm.Config{})
//	if err != nil {
//		t.Fatalf("Failed to initialize GORM: %v", err)
//	}
//
//	// Mock GORM insert statement
//	mock.ExpectBegin()
//	mock.ExpectExec("INSERT INTO `cars`").WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//
//	err = saveCarListToGORM()
//	assert.NoError(t, err)
//}

// 测试 saveCarListToJSON 函数
func TestSaveCarListToJSON(t *testing.T) {
	carList = []Car{
		{Name: "Test Car 1", Brand: "Test Brand", GrLevel: "Gr.3"},
	}

	fileName := "car_list_test.json"
	err := saveCarListToJSON(fileName)
	assert.NoError(t, err)

	// 检查文件是否成功创建
	file, err := os.Open(fileName)
	assert.NoError(t, err)
	defer file.Close()

	// 检查 JSON 文件内容
	var savedCars []Car
	json.NewDecoder(file).Decode(&savedCars)
	assert.Equal(t, carList[0].Name, savedCars[0].Name)

	// 清理测试文件
	os.Remove(fileName)
}
