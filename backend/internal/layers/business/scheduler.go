package scheduler

import (
	enums "backend/internal/layers"
	instance "backend/internal/layers/database/challenge/instance"
	"backend/internal/layers/database/challenge/template"
	"context"

	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type ChgeScheduler struct {
	tmplRepo template.Repository
	instRepo instance.Repository
}

func (chgeShdr *ChgeScheduler) Start(ctx context.Context) {
	s := gocron.NewScheduler(time.Local)
	log.Println("🚀 Starting scheduler...")

	_, err := s.Every(1).Hour().Every(1).Minute().Every(10).Seconds().Do(func() {
		chgeTmpl, _ := chgeShdr.tmplRepo.FindByID(ctx, 1)
		log.Printf("Selected random challenge with id = %d", chgeTmpl.ID)

		_, err := chgeShdr.instRepo.Create(ctx, &instance.ChallengeInstance{TemplateID: chgeTmpl.ID, Status: enums.INST_ACTIVE})
		if err != nil {
			log.Printf("Failed to start job. Skipping current round")
			return
		}
	})

	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}
	s.StartAsync()
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
