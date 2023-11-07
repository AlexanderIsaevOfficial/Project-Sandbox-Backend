package types

//
type Config struct {
	SQLConnection string `json:"SQLConnection"`
	MaxProcs      uint   `json:"MaxProcs"`
	Port          string `json:"Port"`
	ImagePath     string `json:"ImagePath"`
	LocoFilePath  string `json:"LocoFilePath"`
}

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LocationId  int    `json:"locationId"`
	Logo        string `json:"logo"`
	ScreenShots string `json:"screenShots"`
}

type Category struct {
	Id    int
	Title string
}

type CategoryDBRequest struct {
	Id    int
	Title string
	I     Item
}

type CategoryResponse struct {
	Id    int
	Title string
	Items []Item
}

type MainResponse struct {
	Cat    []CategoryResponse `json:"data"`
	Status bool               `json:"status"`
}

type Location struct {
	Id       int
	FilePath string
}

type ItemsResponse struct {
	Status bool   `json:"status"`
	Items  []Item `json:"data"`
}
