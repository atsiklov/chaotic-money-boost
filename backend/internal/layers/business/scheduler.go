package scheduler

import (
	enums "backend/internal/layers"
	instance "backend/internal/layers/database/challenge/instance"
	"backend/internal/layers/database/challenge/template"
	"context"
	"fmt"

	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type ChgeScheduler struct {
	tmplRepo template.Repository
	instRepo instance.Repository
}

func (chgeShdr *ChgeScheduler) Start(ctx context.Context) {
	scheduler := gocron.NewScheduler(time.Local)
	log.Println("🚀 Starting scheduler...")

	_, err := scheduler.Every(15).Seconds().Do(func() {
		chgeTmpl, _ := chgeShdr.tmplRepo.FindByID(ctx, 2)
		log.Printf("Selected random challenge with id = %d", chgeTmpl.ID)

		startedAt := time.Now()
		expiresAt := startedAt.Add(*chgeTmpl.Duration)

		chgeInst, err := chgeShdr.instRepo.Create(ctx, &instance.ChallengeInstance{
			TemplateID: chgeTmpl.ID,
			Status:     enums.INST_ACTIVE,
			StartedAt:  &startedAt,
			ExpiresAt:  &expiresAt,
		})
		if err != nil {
			log.Printf("Failed to start job. Skipping current round")
			return
		}

		log.Printf("Instance id = %d started at %s", chgeInst.ID, startedAt)

		go func(chgeInst *instance.ChallengeInstance, duration *time.Duration) {
			log.Printf("Instance id = %d will be processed in %s. Waiting...", chgeInst.ID, duration)
			<-time.After(*duration)
			log.Printf("Processing instance id = %d", chgeInst.ID)
			// todo - собираем все инстансы
		}(chgeInst, chgeTmpl.Duration)

		fmt.Println("PUSH")
		fmt.Println("PUSH")
		fmt.Println("PUSH")
	})

	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}

	scheduler.StartAsync()
}

func NewChgeScheduler(
	tmplRepo template.Repository,
	instRepo instance.Repository,
) *ChgeScheduler {
	return &ChgeScheduler{
		tmplRepo: tmplRepo,
		instRepo: instRepo,
	}
}
