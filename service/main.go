package service

import (
	"github.com/liyuliang/queue-services"
	"github.com/liyuliang/sstorage/models"
	"github.com/liyuliang/sstorage/system"
	"github.com/liyuliang/utils/request"
	"github.com/liyuliang/utils/format"
	"net/url"
	"github.com/liyuliang/access"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var ms models.Models

func Start() {

	initQueue()

	services.AddSingleProcessTask("Pull Job", func(workerNum int) (err error) {

		err = execStorageJobs()
		if err != nil {
			return err //TODO
		}
		return err
	})

	services.Service().Start(true)
}

func execStorageJobs() (err error) {

	html, err := getJobsContent()
	if err != nil {
		return err
	}

	model, err := htmlToModel(html)
	if err != nil {
		return err
	}

	save(model.Sqls())
	extend(model.Extends())
	//esIndex(model.)

	return //TODO return
}

func save(sqls []string) {
	db := system.Mysql()

	tx := db.Begin()

	for _, sql := range sqls {
		tx.Exec(sql)
	}

	tx.Commit()
}

func extend(jobs []*models.Job) {
	log.Println(len(jobs))

	for _, job := range jobs {
		extendJob(job)
	}
}

func extendJob(job *models.Job) (html string, err error) {

	gateway := system.Config()[system.SystemGateway]
	queueAddApi := gateway + system.AddApiPath

	vals := url.Values{}
	vals.Add("type", job.Type)
	vals.Add("token", job.Token)

	for _, u := range job.Urls {
		vals.Add("[]url", u)
	}
	return request.HttpPost(queueAddApi, vals)
}

func getJobsContent() (html string, err error) {

	gateway := system.Config()[system.SystemGateway]
	queueGetApi := gateway + system.GetApiPath

	data := format.ToMap(map[string]string{
		"queue": system.QueueName,
		"n":     "1",
	})

	html, err = request.HttpPost(queueGetApi, data.ToUrlVals())
	if html == "" {
		return "", errors.New("Empty Queue ...")
	}
	return
}

func checkModel(model models.Model) (err error) {
	if model == nil {
		return errors.New("Empty storage model by matching queue data")
	}
	return
}

func htmlToModel(html string) (model models.Model, err error) {
	html = strings.Replace(html, "[\"", "", -1)
	html = strings.Replace(html, "\"]", "", -1)
	html = strings.Replace(html, "\\u0026", "&", -1)
	html, _ = url.QueryUnescape(html)

	html = removeLineBreak(html)

	host := "http://127.0.0.1"
	u := host + "?" + html
	p, e := url.Parse(u)
	if e != nil {
		log.Println(e.Error())
		return nil, err
	}

	params := p.Query()

	modelName, err := modelNameFromQueryParam(params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	modelCreator := models.Get(modelName)
	if modelCreator == nil {
		return nil, err
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

	err = checkModel(model)
	if err != nil {
		return nil, err
	}

	return model, nil
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

func removeLineBreak(v string) string {
	v = strings.Replace(v, "\n", "", -1)
	v = strings.Replace(v, "\t", "", -1)
	v = strings.Replace(v, " ", "", -1)
	return v
}
