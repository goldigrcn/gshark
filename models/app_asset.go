package models

import (
	"time"
	"github.com/neal1991/gshark/vars"
	"fmt"
)

type AppAsset struct {
	Id          int64
	Name        *string `json:"name,omitempty"`
	Description *string
	Market      *string `json:"market,omitempty"`
	Developer   *string
	Version     *string
	DeployDate  *string
	Url         *string
	// obtain from virustotal
	Hash      *string
	Status       int
	CreatedTime time.Time `xorm:"created"`
	UpdatedTime time.Time `xorm:"updated"`
}

// sha256 is utilized to detect if the app exists
func Detect(hash string) (bool, int64) {
	appAsset := new(AppAsset)
	var id int64
	has, err := Engine.Table("app_asset").Where("hash=?", hash).Get(appAsset)
	if err != nil {
		fmt.Println(err)
	}
	if !has {
		id = -1
	} else {
		id = appAsset.Id
	}
	return has, id
}

func (r *AppAsset) Insert() (int64, error) {
	return Engine.Insert(r)
}


func ListAppAssets(page int) ([]InputInfo, int, error) {
	inputs := make([]InputInfo, 0)

	totalPages, err := Engine.Table("app_assets").Count()
	var pages int

	if int(totalPages)%vars.PAGE_SIZE == 0 {
		pages = int(totalPages) / vars.PAGE_SIZE
	} else {
		pages = int(totalPages)/vars.PAGE_SIZE + 1
	}

	if page >= pages {
		page = pages
	}

	if page < 1 {
		page = 1
	}

	err = Engine.Table("input_info").Limit(vars.PAGE_SIZE, (page-1)*vars.PAGE_SIZE).Find(&inputs)

	return inputs, pages, err
}
