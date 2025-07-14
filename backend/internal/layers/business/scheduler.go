package scheduler

import (
	enums "backend/internal/layers"
	instance "backend/internal/layers/database/challenge/instance"
	submissn "backend/internal/layers/database/challenge/submission"
	"backend/internal/layers/database/challenge/template"
	"context"

	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type ChgeScheduler struct {
	tmplRepo template.Repository
	instRepo instance.Repository
	sbmnRepo submissn.Repository
}

func (chgeShdr *ChgeScheduler) Start(ctx context.Context) {
	scheduler := gocron.NewScheduler(time.Local)
	log.Println("🚀 Starting scheduler...")

	_, err := scheduler.Every(60).Seconds().Do(func() {
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
		log.Println("📨 Sending new challenge notifications... 📢")
		log.Println()

		go func(chgeInst *instance.ChallengeInstance, duration *time.Duration) {
			log.Printf("Challenge id = %d will expire in %s. Waiting for submissions...", chgeInst.ID, duration)
			log.Println()

			<-time.After(*duration)
			log.Printf("Getting submissions for challenge id = %d", chgeInst.ID)

			// todo - сделать обработку инстансов с прошлого круга (если вдруг там была ошибка)
			submissions, err := chgeShdr.sbmnRepo.FindAllByInstanceID(ctx, chgeInst.ID)
			if err != nil {
				log.Printf("Failed to get submissions for instance id = %d", chgeInst.ID)
				return
			}
			log.Printf("📂✅ Found %d submissions for challenge id = %d", len(submissions), chgeInst.ID)
			log.Println()
			log.Println("🎯 Starting to determine a winner...")
			time.Sleep(3 * time.Second)
			log.Println("👑🏆 WINNER FOUND! 🎉")
			log.Println()
			log.Println("🧮 Calculating rewards...")
			time.Sleep(3 * time.Second)
			log.Println("💸💰 REWARDS CALCULATED!")
			log.Println()
			log.Println("🔄 Updating wallets...")
			time.Sleep(3 * time.Second)
			log.Println("💳✅ WALLETS UPDATED!")
			log.Println()
			log.Println("📨 Sending winners notifications... 📢")
		}(chgeInst, chgeTmpl.Duration)
	})

	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}

	scheduler.StartAsync()
}

func NewChgeScheduler(
	tmplRepo template.Repository,
	instRepo instance.Repository,
	sbmnRepo submissn.Repository,
) *ChgeScheduler {
	return &ChgeScheduler{
		tmplRepo: tmplRepo,
		instRepo: instRepo,
		sbmnRepo: sbmnRepo,
	}
}
