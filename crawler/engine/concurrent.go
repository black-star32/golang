package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WrokerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){

	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WrokerCount; i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items{
			go func() {e.ItemChan <- item}()
		}

		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier){
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
