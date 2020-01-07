package main

import (
	"fmt"
	"time"

	"bitbucket.org/wmsight/flourish-api/flrr"
	"github.com/pieterclaerhout/go-waitgroup"
	"github.com/sirupsen/logrus"
)

// JobInput represents an input type to a job
type JobInput struct {
	ID   int
	Rand int
}

// JobInputError represents an error which ocurrs while putting things into a job queue
type JobInputError struct {
	ID    int
	Error error
}

// JobInputDoneResult represents the type of value returned when all input to the job has ceased
type JobInputDoneResult struct {
	Successes []*JobInput
	Errors    []*JobInputError
}

// JobOutput represents a completed JobInput
type JobOutput struct {
	Input JobInput
}

// JobOutputError represents an error while processing a JobInput
type JobOutputError struct {
	Input JobInput
	Error error
}

// JobDoneResult represents the type of value returned when the job completes
type JobDoneResult struct {
	Successes []*JobOutput
	Errors    []*JobOutputError
}

func main() {

	// create a channel to place inputs into a job queue
	// hopefully, no errors ;)
	jobInputChan := make(chan JobInput, 5)

	// create a channel to listen for when all inputs have completed
	// the result returns an accumulation of data sent to the job
	jobInputDoneChan := make(chan JobInputDoneResult)

	// create a channel to listen for when all jobs have finished
	// the result returns an accumulation of job data
	jobDoneChan := make(chan JobDoneResult)

	// create an allDoneChan to listen key process markers
	// and to orchestrate the notification on the allDone channel
	allDoneChan := make(chan bool)

	// spawn a goroutine (or pool of routines) to process inputs and enumerate results
	// TODO: add cancellation chan?
	go RunJobInputQueue(jobInputChan, jobDoneChan)

	// spawn a goroutine to listen to the input and processing threads
	go ListenForResultsAndNotifyAllDone(jobInputDoneChan, jobDoneChan, allDoneChan)

	// spawn a goroutine to add inputs to the queue
	go AddInputsToJobQueue(jobInputChan, jobInputDoneChan)

	// main thread can work while the above threads are working

	// block until all done notified
	<-allDoneChan
	logrus.Infof("ALL DONE")
}

// AddInputsToJobQueue adds inputs to the job inputs channel
func AddInputsToJobQueue(jobInput chan JobInput, jobInputDone chan JobInputDoneResult) {
	result := JobInputDoneResult{}
	for i := 0; i < 100; i++ {

		// simulate an intermittent error
		if i == 3 {
			// add the error and continue
			errInput := JobInputError{
				ID:    i,
				Error: fmt.Errorf("simulated error while generating inputs to queue"),
			}
			result.Errors = append(result.Errors, &errInput)
			logrus.Warnf("ERROR: (ID) %d --> (%s)", errInput.ID, errInput.Error)
			continue
		}

		// simulate a wait
		<-time.After(time.Millisecond * 300)

		// successful add to queue
		input := JobInput{ID: i}
		result.Successes = append(result.Successes, &input)
		jobInput <- input
		logrus.Infof("ADDED TO QUEUE: (ID) %d", input.ID)
	}

	// close input channel & notify done
	close(jobInput)
	jobInputDone <- result
	close(jobInputDone)
}

// RunJobInputQueue runs the job queue to process inputs
func RunJobInputQueue(inputs chan JobInput, jobDone chan JobDoneResult) {
	// need a new result struct
	result := &JobDoneResult{}

	// get a new wait group
	// This is a special flavor of sync.WaitGroup based on channels
	// Need to run: `go get github.com/pieterclaerhout/go-waitgroup`
	// It will always keep X number of worker threads awaited
	// Pass 0 or -1 in to use it in the same way as sync.WaitGroup
	wg := waitgroup.NewWaitGroup(5)

	// or be able to batch? or be able to sense when one is added?
	for processingInput := range inputs {

		// need to declare this in the local scope
		input := processingInput

		// add a new wait group
		wg.BlockAdd()

		// spawn new goroutine for each input in channel as it comes
		go func() {
			funcTag := "processInputWorker"
			defer wg.Done()

			logrus.Infof("PROCESSING: (ID) %d", input.ID)

			// process with simulated wait
			<-time.After(time.Second * 5)

			// mock up an intermittent error
			var err error
			if input.ID == 1 {
				err = fmt.Errorf("simulated processing error")
			}

			// if error, append to errors result
			if err != nil {
				errInput := &JobOutputError{
					Input: input,
					Error: flrr.Errorf(err, funcTag, "failed to complete job"),
				}
				result.Errors = append(result.Errors, errInput)
				logrus.Warnf("ERROR: (ID) %d --> (%s)", errInput.Input.ID, errInput.Error)
				return
			}

			// append to success result
			result.Successes = append(result.Successes, &JobOutput{
				Input: input,
			})
			logrus.Infof("DONE: (ID) %d", input.ID)
			return
		}()
	}

	// block until all processed in parallel
	logrus.Infof("WAIT ON QUEUE TO FINISH")
	wg.Wait()

	// done and close
	logrus.Infof("DONE PROCESSING")
	jobDone <- *result
	close(jobDone)
}

// ListenForResultsAndNotifyAllDone orchestrates the notification of allDoneChan
func ListenForResultsAndNotifyAllDone(inputDone chan JobInputDoneResult, jobDone chan JobDoneResult, allDone chan bool) {

	// block until inputDone received
	inputResult := <-inputDone
	logrus.Infof("====================")
	logrus.Infof("INPUTS TO QUEUE: %d", len(inputResult.Successes))

	// display errors if any
	if len(inputResult.Errors) > 0 {
		// report the errors
		logrus.Infof("INPUT ERRORS: %d", len(inputResult.Errors))
		for _, ie := range inputResult.Errors {
			logrus.Warnf("FAILED: (ID) %d --> (%s)", ie.ID, ie.Error)
		}
	}

	// block until processingDone received
	processingResult := <-jobDone
	logrus.Infof("====================")
	logrus.Infof("PROCESSED BY JOB: %d", len(processingResult.Successes))

	// display errors if any
	if len(processingResult.Errors) > 0 {
		// report the errors
		logrus.Warnf("JOB ERRORS: %d", len(processingResult.Errors))
		for _, pe := range processingResult.Errors {
			logrus.Warnf("FAILED: (ID) %d --> (%s)", pe.Input.ID, pe.Error)
		}
	}

	// notify all done
	allDone <- true
}
