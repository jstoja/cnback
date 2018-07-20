package scheduler

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jstoja/cnback/config"
	"github.com/jstoja/cnback/notifier"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
)

type Scheduler struct {
	Cron   *cron.Cron
	Plans  []config.Plan
	Config *config.AppConfig
	Stats  *StatusStore
}

func New(plans []config.Plan, conf *config.AppConfig, stats *StatusStore) *Scheduler {
	return &Scheduler{
		Cron:   cron.New(),
		Plans:  plans,
		Config: conf,
		Stats:  stats,
	}
}

func (s *Scheduler) Start() (int, error) {
  scheduleCount := 0
	for _, plan := range s.Plans {
    job := backupJob{plan.Name, plan, s.Config, s.Stats, s.Cron}
    if plan.Scheduler.Cron == "" {
      job.Run()
    } else {
      schedule, err := cron.ParseStandard(plan.Scheduler.Cron)
      if err != nil {
        return 0, errors.Wrapf(err, "Invalid cron %v for plan %v", plan.Scheduler.Cron, plan.Name)
      }
      scheduleCount += 1
      s.Cron.Schedule(schedule, job)
      logrus.Infof("Scheduled %v", plan.Name)
    }
	}
  if scheduleCount == 0 {
    return 0, nil
  }

	// TODO: Shouldn't be here
	//s.Cron.AddFunc("0 0 */1 * *", func() {
	//	backup.TmpCleanup(s.Config.TmpPath)
	//})

	s.Cron.Start()
	stats := make([]*Status, 0)
	for _, e := range s.Cron.Entries() {
		switch e.Job.(type) {
		case backupJob:
			status := &Status{
				Plan:    e.Job.(backupJob).name,
				NextRun: e.Next,
			}
			stats = append(stats, status)
		default:
			logrus.Infof("Next tmp cleanup run at %v", e.Next)
		}
	}

	//if err := s.Stats.Sync(stats); err != nil {
	//	logrus.Errorf("Status store sync failed %v", err)
	//}

	return scheduleCount, nil
}

type backupJob struct {
	name  string
	plan  config.Plan
	conf  *config.AppConfig
	stats *StatusStore
	cron  *cron.Cron
}

func (b backupJob) Run() {
	logrus.WithField("plan", b.plan.Name).Info("Backup started")
	log := ""
	//t1 := time.Now()

	_, err := backup(b.plan)
	if err != nil {
		log = fmt.Sprintf("backup failed %v", err)
		logrus.WithField("plan", b.plan.Name).Error(log)

		if err := notifier.SendNotification(fmt.Sprintf("%v backup failed", b.plan.Name), err.Error(), true, b.plan); err != nil {
			logrus.WithField("plan", b.plan.Name).Errorf("notifier failed %v", err)
		}
	}// else {
		// log = fmt.Sprintf("Backup finished in %v archive %v size %v", res.Duration, res.Name, humanize.Bytes(uint64(res.Size)))

		//logrus.WithField("plan", b.plan.Name).Info(log)
		// if err := notifier.SendNotification(fmt.Sprintf("%v backup finished", b.plan.Name),
		// 	fmt.Sprintf("%v backup finished in %v archive size %v", res.Name, res.Duration, humanize.Bytes(uint64(res.Size))), false, b.plan); err != nil {
		// 	logrus.WithField("plan", b.plan.Name).Errorf("Notifier failed %v", err)
		// }
	//}

	//for _, e := range b.cron.Entries() {
	//	switch e.Job.(type) {
	//	case backupJob:
	//		if e.Job.(backupJob).name == b.plan.Name {
	//			s.NextRun = e.Next
	//			break
	//		}
	//	}
	//}

	//logrus.WithField("plan", b.plan.Name).Infof("Next run at %v", s.NextRun)
	//if err := b.stats.Put(s); err != nil {
	//	logrus.WithField("plan", b.plan.Name).Errorf("Status store failed %v", err)
	//}
}
