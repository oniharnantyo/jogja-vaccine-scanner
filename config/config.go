package config

import "time"

type Config struct {
	Service
	Scheduler
	Participant
	SMTP
	Emails
}

type Service struct {
	Slemankab struct {
		BaseUrl string
	}
}

type Scheduler struct {
	Enable bool
	Every  time.Duration
}

type Participant struct {
	Nik string
	Age int
}

type SMTP struct {
	Enable   bool
	Host     string
	Port     int
	Sender   string
	Password string
	Subject  string
}

type Emails []string
