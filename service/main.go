package service

import (
	"github.com/liyuliang/queue-services"
	"github.com/liyuliang/sstorage/models"
	"github.com/liyuliang/sstorage/system"
	"github.com/liyuliang/utils/request"
	"github.com/liyuliang/utils/format"
	"net/url"
	"access"
	"github.com/pkg/errors"
)

var ms models.Models

func Start() {

	initQueue()

	services.AddSingleProcessTask("Pull Job", func(workerNum int) (err error) {

		return execStorageQueueJobs()

	})

	services.Service().Start(true)
}
func execStorageQueueJobs() (err error) {
	gateway := system.Config()[system.SystemGateway]
	queueGetApi := gateway + system.GetApiPath

	data := format.ToMap(map[string]string{
		"queue": system.QueueName,
		"n":     "1",
	})

	html, err := request.HttpPost(queueGetApi, data.ToUrlVals())

	if err != nil {
		return
	}

	model := formatResponseData(html)

	err = checkModel(model)
	if err != nil {

		return err
		//TODO
	}

	db := system.Mysql()
	tx := db.Begin()

	for _, sql := range model.Sqls() {
		tx.Exec(sql)
	}

	tx.Commit()
	return //TODO return
}

func checkModel(model models.Model) (err error) {
	if model == nil {
		return errors.New("Empty storage model by matching queue data")
	}
	return
}

func formatResponseData(html string) (model models.Model) {
	host := "http://127.0.0.1"
	u := host + "?" + html
	p, e := url.Parse(u)
	if e != nil {
		return
	}

	params := p.Query()

	modelName, err := modelNameFromQueryParam(params)
	if err != nil {
		return
	}

	modelCreator := models.Get(modelName)
	if modelCreator == nil {
		return
	}

	model = modelCreator()
	for k, vs := range params {

		switch len(vs) {
		case 0:
		case 1:
			access.SetField(model, k, vs[0])
		default:
			access.SetField(model, k, vs)
		}
	}
	return model
}

func modelNameFromQueryParam(vs url.Values) (modelName string, err error) {
	params, ok := vs["database"]
	if !ok {
		return "", errors.New("storage data from get api missing required param named 'database'")
	}

	if len(params) == 0 {
		return "", errors.New("storage param from get api missing value which named 'database'")
	}

	return params[0], nil
}

func initQueue() {

	ms = models.List()
}
