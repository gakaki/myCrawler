package juejinBook

import (
	"fmt"
	"log"
	"myCrawler/utils"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// DefaultSaveDir default save dir.
	DefaultSaveDir = "book"
	// GetSectionURL get section url.
	GetSectionURL = "https://api.juejin.cn/booklet_api/v1/section/get"
	// GetBookInfoURL get book info url.
	GetBookInfoURL = "https://api.juejin.cn/booklet_api/v1/booklet/get"
)

type JuejinListRequest struct {
	CategoryID int
	Cursor     string
	Sort       int
	IsVIP      int
	Limit      int
}

func GetAllBookListSortLatestSaveToJSON() []JuejinResponseBook {
	// 先从chrome 链接里爬取 , 然后找到 body里的参数进行 修改参数
	url := "https://api.juejin.cn/booklet_api/v1/booklet/listbycategory"

	juejinListRequest := JuejinListRequest{
		CategoryID: 0,
		Cursor:     "0",
		Sort:       10,
		IsVIP:      0,
		Limit:      1000,
	}

	response, err := utils.PostToStructInputStruct[JuejinResponse](url, juejinListRequest, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("掘金一共有", len(response.Data), "本册子")
	utils.WriteJSON(response, "juejin_book.json")
	return response.Data
}

func (jm *Juejinxiaoce2Markdown) InitStep() (err error) {
	if jm.Sessionid == "" {
		return fmt.Errorf("sessionid is empty")
	}
	if len(jm.BookIDs) == 0 {
		return fmt.Errorf("bookIDs is empty")
	}
	pwd, _ := os.Getwd()
	if pwd == "" {
		return fmt.Errorf("PWD is empty")
	}

	jm.ImgPattern = regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
	jm.RequestHeaders = map[string]string{"cookie": fmt.Sprintf("sessionid=%s;", jm.Sessionid)}
	jm.MarkdownSavePaths = make(map[string]string)

	if jm.SaveDir == "" {
		jm.SaveDir = filepath.Join(pwd, DefaultSaveDir)
	}
	if err := os.MkdirAll(jm.SaveDir, os.ModePerm); err != nil {
		return fmt.Errorf("create save dir failed: %v", err)
	}
	return nil
}

func (j *Juejinxiaoce2Markdown) GetSectionRes(sectionID string) (JuejinSectionContent, error) {
	data := map[string]string{
		//"section_id": strconv.FormatInt(sectionID, 10),
		"section_id": sectionID,
	}
	//if sectionID == "6982024263184154661" {
	//	fmt.Println(sectionID)
	//}
	return utils.PostToStructInputStruct[JuejinSectionContent](GetSectionURL, data, j.Sessionid)
}

func (j *Juejinxiaoce2Markdown) GetBookInfoRes(bookID string) (JuejinSection, error) {
	data := map[string]string{
		"booklet_id": bookID,
	}
	return utils.PostToStructInputStruct[JuejinSection](GetBookInfoURL, data, j.Sessionid)
}

func dealBookAndSectionTitle(s string) string {
	tmp := strings.ReplaceAll(s, "\\", "")
	tmp = strings.ReplaceAll(tmp, "/", "")
	tmp = strings.ReplaceAll(tmp, "|", "")
	return tmp
}

func (j *Juejinxiaoce2Markdown) Download() {
	// 并发 下载
	fmt.Println("books need to download count ", len(j.BookIDs))

	maxWorkerCount := 10
	queue := make(chan string, maxWorkerCount)
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	for i := 0; i < maxWorkerCount; i++ {
		time.Sleep(time.Microsecond * 1000)

		go func() {
			defer wg.Done()
			wg.Add(1)
			for bookId := range queue {
				err := j.DownloadOneBook(bookId)
				if err != nil {
					fmt.Println("Error", err)
				}
			}
		}()
	}
	for _, bookID := range j.BookIDs {
		queue <- bookID
	}
	close(queue)
	wg.Wait()
	//utils.WriteJSON(salttigerItems, "salttigerItems.json")
}

func (j *Juejinxiaoce2Markdown) DownloadOneBook(bookID string) error {
	log.Printf("开始处理小册")

	juejinSection, err := j.GetBookInfoRes(bookID)
	if err != nil {
		return fmt.Errorf("GetBookInfoRes failed: %v", err)
	}

	bookTitle := dealBookAndSectionTitle(juejinSection.Data.Booklet.BaseInfo.Title)
	bookSavePath := filepath.Join(j.SaveDir, bookTitle)
	log.Printf("book_title: %s %s", bookTitle, bookSavePath)

	// 创建目录在线
	bookSavePathOnline := bookSavePath + "_online"
	if err := os.MkdirAll(bookSavePathOnline, os.ModePerm); err != nil {
		return fmt.Errorf("create book onlone save path failed: %v", err)
	}
	//if err := os.MkdirAll(bookSavePath, os.ModePerm); err != nil {
	//	return fmt.Errorf("create book save path failed: %v", err)
	//}
	// 创建目录离线并且带图片
	imgDir := filepath.Join(bookSavePath, "img")
	if err := os.MkdirAll(imgDir, os.ModePerm); err != nil {
		return fmt.Errorf("create img dir failed: %v", err)
	}

	sectionIDList := make([]string, 0, len(juejinSection.Data.Sections))
	for _, section := range juejinSection.Data.Sections {
		sectionIDList = append(sectionIDList, section.SectionID)
	}

	sectionTotalLength := len(sectionIDList)
	for sectionIndex, sectionID := range sectionIDList {

		sectionOrder := sectionIndex + 1

		juejinSectionContent, err := j.GetSectionRes(sectionID)

		if err != nil {
			return fmt.Errorf("GetSectionRes failed: %v", err)
		}

		sectionTitle := dealBookAndSectionTitle(juejinSectionContent.Data.Section.Title)
		markdownStr := juejinSectionContent.Data.Section.MarkdownShow
		if markdownStr == "" || len(markdownStr) <= 0 {
			log.Printf("warning: >>>> markdown为空 检查session section %s >> %s", sectionID, sectionTitle)
		}

		markdownFilePath := filepath.Join(bookSavePath, fmt.Sprintf("%d-%s.md", sectionOrder, sectionTitle))

		sectionImgDir := filepath.Join(imgDir, strconv.Itoa(sectionOrder))

		log.Printf("进度: %d/%d, 处理 section >> %s", sectionOrder, sectionTotalLength, sectionTitle)

		// build online file
		markdownFilePathOnline := filepath.Join(bookSavePathOnline, fmt.Sprintf("%d-%s.md", sectionOrder, sectionTitle))

		err = os.WriteFile(markdownFilePathOnline, []byte(markdownStr), 0644)
		if err != nil {
			return fmt.Errorf("download markdown failed: %v", err)
		}

		if err := os.MkdirAll(sectionImgDir, os.ModePerm); err != nil {
			return fmt.Errorf("create section img dir failed: %v", err)
		}

		markdownRelativeImgDir := filepath.Join("img", strconv.Itoa(sectionOrder))
		j.MarkdownSavePaths[sectionID] = markdownFilePath

		if j.DownloadImage == true {
			j.SaveMarkdownOffline(sectionIndex, markdownFilePath, sectionImgDir, markdownRelativeImgDir, markdownStr)
		}
	}
	log.Printf("处理完成")
	return nil
}

func GetMarkDownImageUrl(markDown string) []string {
	images := make([]string, 0)
	// 定义正则表达式，匹配Markdown中的图片URL
	re := regexp.MustCompile(`!\[.*\]\((https?://[^\s]+)\)`)
	// 使用FindAllStringSubmatch函数查找所有匹配项
	matches := re.FindAllStringSubmatch(markDown, -1)
	// 遍历匹配项，提取图片URL
	for _, match := range matches {
		if len(match) > 1 {
			fmt.Println(match[1]) // 打印图片URL
			images = append(images, match[1])
		}
	}
	return images
}
func FindImageUrls(sectionIndex int, htmls string) []string {
	if sectionIndex == 4 {
		fmt.Println(sectionIndex)
	}

	imgRE := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(htmls, -1)
	out := make([]string, 0)
	for _, img := range imgs {
		if strings.Contains(img[1], "http") {
			out = append(out, strings.Replace(img[1], "\\", "", -1))
		}
	}
	if len(out) <= 0 {
		imgRE = regexp.MustCompile(`https://.*?\.(jpg|jpeg|gif|image|awebp|webp)`)
		imgs := imgRE.FindAllStringSubmatch(htmls, -1)
		out = make([]string, 0)
		for _, img := range imgs {
			if strings.Contains(img[0], "http") {
				out = append(out, strings.Replace(img[0], "\\", "", -1))

			}
		}
	}

	out = GetMarkDownImageUrl(htmls)

	return out
}

func (j *Juejinxiaoce2Markdown) SaveMarkdownOffline(sectionIndex int, markdownFilePath string, sectionImgDir string, markdownRelativeImgDir string, markdownStr string) {
	//j.Down //是否需要下载图片 分为下载和不下载2种

	imgUrls := FindImageUrls(sectionIndex, markdownStr)

	// 并发 下载
	fmt.Println("sectionIndex images download count ", sectionIndex, len(imgUrls))

	type Image struct {
		imgUrl        string
		saveImagePath string
	}
	maxWorkerCount := 8
	queue := make(chan *Image, maxWorkerCount)
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	for i := 0; i < maxWorkerCount; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			for image := range queue {
				// Download image
				err := utils.RequestThanSaveImage(image.imgUrl, image.saveImagePath)
				if err != nil {
					fmt.Println("Error downloading image:", err)
				}
				//time.Sleep(time.Second * 1)
			}
		}()
	}

	for imgIndex, imgUrl := range imgUrls {
		newImgUrl := strings.TrimSpace(imgUrl) // Remove newlines and extra spaces
		if strings.HasPrefix(newImgUrl, "//") {
			newImgUrl = "https:" + newImgUrl // Add https:// if missing
		}

		suffix := filepath.Ext(newImgUrl)
		suffix = ".png"                                                         // Get file extension
		imgFileName := fmt.Sprintf("%d%s", imgIndex+1, suffix)                  // Generate filename
		mdRelativeImgPath := filepath.Join(markdownRelativeImgDir, imgFileName) // Relative path for Markdown
		imgSavePath := filepath.Join(sectionImgDir, imgFileName)                // Full path to save image
		// Replace URL in Markdown string with relative path
		markdownStr = strings.ReplaceAll(markdownStr, imgUrl, mdRelativeImgPath)
		queue <- &Image{
			imgUrl:        newImgUrl,
			saveImagePath: imgSavePath,
		}
	}
	close(queue)
	wg.Wait()

	err := os.WriteFile(markdownFilePath, []byte(markdownStr), 0644)
	if err != nil {
		fmt.Println("Error saving Markdown file:", err)
	}
}
