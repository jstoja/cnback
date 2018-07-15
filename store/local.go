package store

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/jstoja/cnback/config"
	"github.com/pkg/errors"
)

func dump(archive io.ReadCloser, plan config.Plan, planDir string) (string, error) {
	ts := time.Now()

	err := sh.Command("mkdir", "-p", planDir).Run()
	if err != nil {
		return "", errors.Wrapf(err, "creating dir %v in %v failed", plan.Name, plan.Store.Local.StoragePath)
	}

	archivePath := fmt.Sprintf("%v/%v-%v.gz", planDir, plan.Name, ts.Unix())
	f, err := os.Create(archivePath)
	bytesWritten, err := io.Copy(f, archive)
	if err != nil {
		return "", errors.Wrapf(err, "piping to file %v failed", archivePath)
	}

	fi, err := os.Stat(archivePath)
	if err != nil {
		return "", errors.Wrapf(err, "stat file %v failed", archivePath)
	}
	//_, res.Name = filepath.Split(archivePath)
	if bytesWritten != fi.Size() {
		return "", errors.Wrapf(err, "different size read and saved on disk %v vs %v", bytesWritten, fi.Size())
	}
	//res.Size = fi.Size()

	// TODO: Add stderr logging in file + move to planDir

	if plan.Scheduler.Retention > 0 {
		err = applyRetention(planDir, plan.Scheduler.Retention)
		if err != nil {
			return "", errors.Wrap(err, "retention job failed")
		}
	}

	//file := filepath.Join(planDir, res.Name)
	return "", nil
}

func logToFile(file string, data []byte) error {
	if len(data) > 0 {
		err := ioutil.WriteFile(file, data, 0644)
		if err != nil {
			return errors.Wrapf(err, "writing log %v failed", file)
		}
	}

	return nil
}

func applyRetention(path string, retention int) error {
	gz := fmt.Sprintf("cd %v && rm -f $(ls -1t *.gz | tail -n +%v)", path, retention+1)
	err := sh.Command("/bin/sh", "-c", gz).Run()
	if err != nil {
		return errors.Wrapf(err, "removing old gz files from %v failed", path)
	}

	log := fmt.Sprintf("cd %v && rm -f $(ls -1t *.log | tail -n +%v)", path, retention+1)
	err = sh.Command("/bin/sh", "-c", log).Run()
	if err != nil {
		return errors.Wrapf(err, "removing old log files from %v failed", path)
	}

	return nil
}

// TmpCleanup remove files older than one day
func TmpCleanup(path string) error {
	rm := fmt.Sprintf("find %v -not -name \"mgob.db\" -mtime +%v -type f -delete", path, 1)
	err := sh.Command("/bin/sh", "-c", rm).Run()
	if err != nil {
		return errors.Wrapf(err, "%v cleanup failed", path)
	}

	return nil
}
