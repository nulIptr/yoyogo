package Abstractions

const (
	APPLICATION_LIFE_START = "APPLICATION_LIFE_START"
	APPLICATION_LIFE_STOP  = "APPLICATION_LIFE_STOP"
)

//var ApplicationCycle = NewApplicationLife()

type ApplicationLife struct {
	eventPublisher     *ApplicationEventPublisher
	ApplicationStopped chan ApplicationEvent
	ApplicationStarted chan ApplicationEvent
}

func NewApplicationLife() *ApplicationLife {
	applife := &ApplicationLife{
		eventPublisher:     NewEventPublisher(),
		ApplicationStopped: make(chan ApplicationEvent),
		ApplicationStarted: make(chan ApplicationEvent),
	}
	applife.eventPublisher.Subscribe(APPLICATION_LIFE_START, applife.ApplicationStarted)
	applife.eventPublisher.Subscribe(APPLICATION_LIFE_STOP, applife.ApplicationStopped)
	return applife
}

func (life *ApplicationLife) StartApplication() {
	life.eventPublisher.Publish(APPLICATION_LIFE_START, "Start")
}

func (life *ApplicationLife) StopApplication() {
	life.eventPublisher.Publish(APPLICATION_LIFE_STOP, "Stop")
}
