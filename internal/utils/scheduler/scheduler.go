package scheduler

// StartScheduler method will implement the functionality to start a scheduler with the required business logic.
// The time input for the schedules are defined in cron expression and defined as constants available for the entire package.
// The logic behind is to implement schedulers for required jobs that implement the required business logic, but the time input is derived from the package level constants.
type Scheduler interface {
	StartScheduler() error
	//StopScheduler() error
}
