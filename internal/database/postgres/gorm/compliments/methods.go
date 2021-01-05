package compliments

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"telegram-pug/internal/database/errs"
	"telegram-pug/model"
)

func (c *compliments) CreateTableIfNotExists() error {
	if err := c.db.AutoMigrate(model.Compliment{}); err != nil {
		return err
	}

	if _, err := c.SelectRandom(model.Neutral.String()); err != nil && err != errs.RecordNotFound {
		return err
	}

	if err := c.InsertData(); err != nil {
		return err
	}
	return nil
}

func (c *compliments) Insert(compliment model.Compliment) (*model.Compliment, error) {
	err := c.db.Create(&compliment).Error
	if err != nil {
		return nil, err
	}
	return &compliment, nil
}

func (c *compliments) SelectRandom(gender string) (*model.Compliment, error) {
	ret := &model.Compliment{}
	if err := c.db.Raw(`SELECT * FROM compliments WHERE gender=? AND gender=? ORDER BY RANDOM() LIMIT 1`,
		gender, model.Neutral.String()).Scan(ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.RecordNotFound
		}
		return nil, err
	}
	return ret, nil
}

func (c *compliments) InsertData() error {
	file, err := os.Open(c.dataPath)
	if err != nil {
		return err
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var com struct {
		Compliments []model.Compliment `json:"compliments"`
	}

	if err = json.Unmarshal(raw, &com); err != nil {
		return err
	}

	for _, elem := range com.Compliments {
		if _, err = c.Insert(elem); err != nil {
			return err
		}
	}

	return nil
}
