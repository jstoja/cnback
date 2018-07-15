package scheduler

import (
	//"fmt"
  //"io"
	//"os/exec"
	//"time"

	//"github.com/Sirupsen/logrus"
	"github.com/jstoja/cnback/config"
	"github.com/jstoja/cnback/db"
	"github.com/jstoja/cnback/store"
)

type Result struct {}

func backup(plan config.Plan) (*string, error) {

	backupStream, _ := db.FetchBackup(plan)
  err := store.SendBackup(backupStream)
  if err != nil {
  }

	// TODO: Implement timeouts
	// 1. run the backup
	// 2. get the stream
	// 3. add stream encrytion if configured
	// 4. send to storages

	//if err != nil {
	//return Result{}, err
	//}

	//res := Result{}
	//tmpFile := ""
	//if plan.Local != nil {
	//dumpTo := fmt.Sprintf("%v/%v", plan.Local.StoragePath, plan.Name)
	//res, err := dump(archiveStream, plan, dumpTo)
	//if err != nil {
	//return res, err
	//}
	//tmpFile = dumpTo
	//}

	//// TODO: Use io readers/writers to stream directly to APIs
	//// For now this is not possible, so we have to use a temporary file first
	//if plan.SFTP != nil || plan.S3 != nil || plan.GCloud != nil {
	//if tmpFile == "" {
	//tmpFile = fmt.Sprintf("%v/%v", tmpPath, plan.Name)
	//res, err := dump(archiveStream, plan, tmpFile)
	//if err != nil {
	//return res, err
	//}
	//}
	//}

	//if plan.SFTP != nil {
	//sftpOutput, err := sftpUpload(tmpFile, plan)
	//if err != nil {
	//return res, err
	//} else {
	//logrus.WithField("plan", plan.Name).Info(sftpOutput)
	//}
	//}

	//if plan.S3 != nil {
	//s3Output, err := s3Upload(tmpFile, plan)
	//if err != nil {
	//return res, err
	//} else {
	//logrus.WithField("plan", plan.Name).Infof("S3 upload finished %v", s3Output)
	//}
	//}

	//if plan.GCloud != nil {
	//gCloudOutput, err := gCloudUpload(tmpFile, plan)
	//if err != nil {
	//return res, err
	//} else {
	//logrus.WithField("plan", plan.Name).Infof("GCloud upload finished %v", gCloudOutput)
	//}
	//}

	//t2 := time.Now()
	//res.Status = 200
	//res.Duration = t2.Sub(ts)
	return nil, nil
}
