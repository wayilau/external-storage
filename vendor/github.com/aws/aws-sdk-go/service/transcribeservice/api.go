// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribeservice

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetTranscriptionJob = "GetTranscriptionJob"

// GetTranscriptionJobRequest generates a "aws/request.Request" representing the
// client's request for the GetTranscriptionJob operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetTranscriptionJob for more information on using the GetTranscriptionJob
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetTranscriptionJobRequest method.
//    req, resp := client.GetTranscriptionJobRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetTranscriptionJob
func (c *TranscribeService) GetTranscriptionJobRequest(input *GetTranscriptionJobInput) (req *request.Request, output *GetTranscriptionJobOutput) {
	op := &request.Operation{
		Name:       opGetTranscriptionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTranscriptionJobInput{}
	}

	output = &GetTranscriptionJobOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetTranscriptionJob API operation for Amazon Transcribe Service.
//
// Returns information about a transcription job. To see the status of the job,
// check the Status field. If the status is COMPLETE, the job is finished and
// you can find the results at the location specified in the TranscriptionFileUri
// field.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Transcribe Service's
// API operation GetTranscriptionJob for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   There is a problem with one of the input fields. Check the S3 bucket name,
//   make sure that the job name is not a duplicate, and confirm that you are
//   using the correct file format. Then resend your request.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   Either you have sent too many requests or your input file is longer than
//   2 hours. Wait before you resend your request, or use a smaller file and resend
//   the request.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   There was an internal error. Check the error message and try your request
//   again.
//
//   * ErrCodeNotFoundException "NotFoundException"
//   We can't find the requested job. Check the job name and try your request
//   again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetTranscriptionJob
func (c *TranscribeService) GetTranscriptionJob(input *GetTranscriptionJobInput) (*GetTranscriptionJobOutput, error) {
	req, out := c.GetTranscriptionJobRequest(input)
	return out, req.Send()
}

// GetTranscriptionJobWithContext is the same as GetTranscriptionJob with the addition of
// the ability to pass a context and additional request options.
//
// See GetTranscriptionJob for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TranscribeService) GetTranscriptionJobWithContext(ctx aws.Context, input *GetTranscriptionJobInput, opts ...request.Option) (*GetTranscriptionJobOutput, error) {
	req, out := c.GetTranscriptionJobRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListTranscriptionJobs = "ListTranscriptionJobs"

// ListTranscriptionJobsRequest generates a "aws/request.Request" representing the
// client's request for the ListTranscriptionJobs operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListTranscriptionJobs for more information on using the ListTranscriptionJobs
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListTranscriptionJobsRequest method.
//    req, resp := client.ListTranscriptionJobsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListTranscriptionJobs
func (c *TranscribeService) ListTranscriptionJobsRequest(input *ListTranscriptionJobsInput) (req *request.Request, output *ListTranscriptionJobsOutput) {
	op := &request.Operation{
		Name:       opListTranscriptionJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTranscriptionJobsInput{}
	}

	output = &ListTranscriptionJobsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListTranscriptionJobs API operation for Amazon Transcribe Service.
//
// Lists transcription jobs with the specified status.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Transcribe Service's
// API operation ListTranscriptionJobs for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   There is a problem with one of the input fields. Check the S3 bucket name,
//   make sure that the job name is not a duplicate, and confirm that you are
//   using the correct file format. Then resend your request.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   Either you have sent too many requests or your input file is longer than
//   2 hours. Wait before you resend your request, or use a smaller file and resend
//   the request.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   There was an internal error. Check the error message and try your request
//   again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListTranscriptionJobs
func (c *TranscribeService) ListTranscriptionJobs(input *ListTranscriptionJobsInput) (*ListTranscriptionJobsOutput, error) {
	req, out := c.ListTranscriptionJobsRequest(input)
	return out, req.Send()
}

// ListTranscriptionJobsWithContext is the same as ListTranscriptionJobs with the addition of
// the ability to pass a context and additional request options.
//
// See ListTranscriptionJobs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TranscribeService) ListTranscriptionJobsWithContext(ctx aws.Context, input *ListTranscriptionJobsInput, opts ...request.Option) (*ListTranscriptionJobsOutput, error) {
	req, out := c.ListTranscriptionJobsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListTranscriptionJobsPages iterates over the pages of a ListTranscriptionJobs operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListTranscriptionJobs method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListTranscriptionJobs operation.
//    pageNum := 0
//    err := client.ListTranscriptionJobsPages(params,
//        func(page *ListTranscriptionJobsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *TranscribeService) ListTranscriptionJobsPages(input *ListTranscriptionJobsInput, fn func(*ListTranscriptionJobsOutput, bool) bool) error {
	return c.ListTranscriptionJobsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListTranscriptionJobsPagesWithContext same as ListTranscriptionJobsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TranscribeService) ListTranscriptionJobsPagesWithContext(ctx aws.Context, input *ListTranscriptionJobsInput, fn func(*ListTranscriptionJobsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListTranscriptionJobsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListTranscriptionJobsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListTranscriptionJobsOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opStartTranscriptionJob = "StartTranscriptionJob"

// StartTranscriptionJobRequest generates a "aws/request.Request" representing the
// client's request for the StartTranscriptionJob operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StartTranscriptionJob for more information on using the StartTranscriptionJob
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the StartTranscriptionJobRequest method.
//    req, resp := client.StartTranscriptionJobRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/StartTranscriptionJob
func (c *TranscribeService) StartTranscriptionJobRequest(input *StartTranscriptionJobInput) (req *request.Request, output *StartTranscriptionJobOutput) {
	op := &request.Operation{
		Name:       opStartTranscriptionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTranscriptionJobInput{}
	}

	output = &StartTranscriptionJobOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StartTranscriptionJob API operation for Amazon Transcribe Service.
//
// Starts an asynchronous job to transcribe speech to text.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Transcribe Service's
// API operation StartTranscriptionJob for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   There is a problem with one of the input fields. Check the S3 bucket name,
//   make sure that the job name is not a duplicate, and confirm that you are
//   using the correct file format. Then resend your request.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   Either you have sent too many requests or your input file is longer than
//   2 hours. Wait before you resend your request, or use a smaller file and resend
//   the request.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   There was an internal error. Check the error message and try your request
//   again.
//
//   * ErrCodeConflictException "ConflictException"
//   The JobName field is a duplicate of a previously entered job name. Resend
//   your request with a different name.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/StartTranscriptionJob
func (c *TranscribeService) StartTranscriptionJob(input *StartTranscriptionJobInput) (*StartTranscriptionJobOutput, error) {
	req, out := c.StartTranscriptionJobRequest(input)
	return out, req.Send()
}

// StartTranscriptionJobWithContext is the same as StartTranscriptionJob with the addition of
// the ability to pass a context and additional request options.
//
// See StartTranscriptionJob for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TranscribeService) StartTranscriptionJobWithContext(ctx aws.Context, input *StartTranscriptionJobInput, opts ...request.Option) (*StartTranscriptionJobOutput, error) {
	req, out := c.StartTranscriptionJobRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetTranscriptionJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the job.
	//
	// TranscriptionJobName is a required field
	TranscriptionJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTranscriptionJobInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTranscriptionJobInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTranscriptionJobInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetTranscriptionJobInput"}
	if s.TranscriptionJobName == nil {
		invalidParams.Add(request.NewErrParamRequired("TranscriptionJobName"))
	}
	if s.TranscriptionJobName != nil && len(*s.TranscriptionJobName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TranscriptionJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTranscriptionJobName sets the TranscriptionJobName field's value.
func (s *GetTranscriptionJobInput) SetTranscriptionJobName(v string) *GetTranscriptionJobInput {
	s.TranscriptionJobName = &v
	return s
}

type GetTranscriptionJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the results of the transcription job.
	TranscriptionJob *TranscriptionJob `type:"structure"`
}

// String returns the string representation
func (s GetTranscriptionJobOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTranscriptionJobOutput) GoString() string {
	return s.String()
}

// SetTranscriptionJob sets the TranscriptionJob field's value.
func (s *GetTranscriptionJobOutput) SetTranscriptionJob(v *TranscriptionJob) *GetTranscriptionJobOutput {
	s.TranscriptionJob = v
	return s
}

type ListTranscriptionJobsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous request to ListTranscriptionJobs was truncated,
	// include the NextToken to fetch the next set of jobs.
	NextToken *string `type:"string"`

	// When specified, returns only transcription jobs with the specified status.
	//
	// Status is a required field
	Status *string `type:"string" required:"true" enum:"TranscriptionJobStatus"`
}

// String returns the string representation
func (s ListTranscriptionJobsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTranscriptionJobsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTranscriptionJobsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListTranscriptionJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Status == nil {
		invalidParams.Add(request.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListTranscriptionJobsInput) SetMaxResults(v int64) *ListTranscriptionJobsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTranscriptionJobsInput) SetNextToken(v string) *ListTranscriptionJobsInput {
	s.NextToken = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListTranscriptionJobsInput) SetStatus(v string) *ListTranscriptionJobsInput {
	s.Status = &v
	return s
}

type ListTranscriptionJobsOutput struct {
	_ struct{} `type:"structure"`

	// The ListTranscriptionJobs operation returns a page of jobs at a time. The
	// size of the page is set by the MaxResults parameter. If there are more jobs
	// in the list than the page size, Amazon Transcribe returns the NextPage token.
	// Include the token in the next request to the ListTranscriptionJobs operation
	// to return in the next page of jobs.
	NextToken *string `type:"string"`

	// The requested status of the jobs returned.
	Status *string `type:"string" enum:"TranscriptionJobStatus"`

	// A list of objects containing summary information for a transcription job.
	TranscriptionJobSummaries []*TranscriptionJobSummary `type:"list"`
}

// String returns the string representation
func (s ListTranscriptionJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTranscriptionJobsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListTranscriptionJobsOutput) SetNextToken(v string) *ListTranscriptionJobsOutput {
	s.NextToken = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListTranscriptionJobsOutput) SetStatus(v string) *ListTranscriptionJobsOutput {
	s.Status = &v
	return s
}

// SetTranscriptionJobSummaries sets the TranscriptionJobSummaries field's value.
func (s *ListTranscriptionJobsOutput) SetTranscriptionJobSummaries(v []*TranscriptionJobSummary) *ListTranscriptionJobsOutput {
	s.TranscriptionJobSummaries = v
	return s
}

// Describes the input media file in a transcription request.
type Media struct {
	_ struct{} `type:"structure"`

	// The S3 location of the input media file. The general form is:
	//
	// https://<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3-us-west-2.amazonaws.com/examplebucket/example.mp4
	//
	// https://s3-us-west-2.amazonaws.com/examplebucket/mediadocs/example.mp4
	MediaFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Media) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Media) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Media) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Media"}
	if s.MediaFileUri != nil && len(*s.MediaFileUri) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MediaFileUri", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMediaFileUri sets the MediaFileUri field's value.
func (s *Media) SetMediaFileUri(v string) *Media {
	s.MediaFileUri = &v
	return s
}

type StartTranscriptionJobInput struct {
	_ struct{} `type:"structure"`

	// The language code for the language used in the input media file.
	//
	// LanguageCode is a required field
	LanguageCode *string `type:"string" required:"true" enum:"LanguageCode"`

	// An object that describes the input media for a transcription job.
	//
	// Media is a required field
	Media *Media `type:"structure" required:"true"`

	// The format of the input media file.
	//
	// MediaFormat is a required field
	MediaFormat *string `type:"string" required:"true" enum:"MediaFormat"`

	// The sample rate, in Hertz, of the audio track in the input media file.
	MediaSampleRateHertz *int64 `min:"8000" type:"integer"`

	// The name of the job. The name must be unique within an AWS account.
	//
	// TranscriptionJobName is a required field
	TranscriptionJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartTranscriptionJobInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTranscriptionJobInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTranscriptionJobInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartTranscriptionJobInput"}
	if s.LanguageCode == nil {
		invalidParams.Add(request.NewErrParamRequired("LanguageCode"))
	}
	if s.Media == nil {
		invalidParams.Add(request.NewErrParamRequired("Media"))
	}
	if s.MediaFormat == nil {
		invalidParams.Add(request.NewErrParamRequired("MediaFormat"))
	}
	if s.MediaSampleRateHertz != nil && *s.MediaSampleRateHertz < 8000 {
		invalidParams.Add(request.NewErrParamMinValue("MediaSampleRateHertz", 8000))
	}
	if s.TranscriptionJobName == nil {
		invalidParams.Add(request.NewErrParamRequired("TranscriptionJobName"))
	}
	if s.TranscriptionJobName != nil && len(*s.TranscriptionJobName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TranscriptionJobName", 1))
	}
	if s.Media != nil {
		if err := s.Media.Validate(); err != nil {
			invalidParams.AddNested("Media", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLanguageCode sets the LanguageCode field's value.
func (s *StartTranscriptionJobInput) SetLanguageCode(v string) *StartTranscriptionJobInput {
	s.LanguageCode = &v
	return s
}

// SetMedia sets the Media field's value.
func (s *StartTranscriptionJobInput) SetMedia(v *Media) *StartTranscriptionJobInput {
	s.Media = v
	return s
}

// SetMediaFormat sets the MediaFormat field's value.
func (s *StartTranscriptionJobInput) SetMediaFormat(v string) *StartTranscriptionJobInput {
	s.MediaFormat = &v
	return s
}

// SetMediaSampleRateHertz sets the MediaSampleRateHertz field's value.
func (s *StartTranscriptionJobInput) SetMediaSampleRateHertz(v int64) *StartTranscriptionJobInput {
	s.MediaSampleRateHertz = &v
	return s
}

// SetTranscriptionJobName sets the TranscriptionJobName field's value.
func (s *StartTranscriptionJobInput) SetTranscriptionJobName(v string) *StartTranscriptionJobInput {
	s.TranscriptionJobName = &v
	return s
}

type StartTranscriptionJobOutput struct {
	_ struct{} `type:"structure"`

	// An object containing details of the asynchronous transcription job.
	TranscriptionJob *TranscriptionJob `type:"structure"`
}

// String returns the string representation
func (s StartTranscriptionJobOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTranscriptionJobOutput) GoString() string {
	return s.String()
}

// SetTranscriptionJob sets the TranscriptionJob field's value.
func (s *StartTranscriptionJobOutput) SetTranscriptionJob(v *TranscriptionJob) *StartTranscriptionJobOutput {
	s.TranscriptionJob = v
	return s
}

// Describes the output of a transcription job.
type Transcript struct {
	_ struct{} `type:"structure"`

	// The S3 location where the transcription result is stored. The general form
	// of this Uri is:
	//
	// https://<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3-us-west-2.amazonaws.com/examplebucket/example.json
	//
	// https://s3-us-west-2.amazonaws.com/examplebucket/mediadocs/example.json
	TranscriptFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Transcript) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Transcript) GoString() string {
	return s.String()
}

// SetTranscriptFileUri sets the TranscriptFileUri field's value.
func (s *Transcript) SetTranscriptFileUri(v string) *Transcript {
	s.TranscriptFileUri = &v
	return s
}

// Describes an asynchronous transcription job that was created with the StartTranscriptionJob
// operation.
type TranscriptionJob struct {
	_ struct{} `type:"structure"`

	// Timestamp of the date and time that the job completed.
	CompletionTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Timestamp of the date and time that the job was created.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// If the TranscriptionJobStatus field is FAILED, this field contains information
	// about why the job failed.
	FailureReason *string `type:"string"`

	// The language code for the input speech.
	LanguageCode *string `type:"string" enum:"LanguageCode"`

	// An object that describes the input media for a transcription job.
	Media *Media `type:"structure"`

	// The format of the input media file.
	MediaFormat *string `type:"string" enum:"MediaFormat"`

	// The sample rate, in Hertz, of the audio track in the input media file.
	MediaSampleRateHertz *int64 `min:"8000" type:"integer"`

	// An object that describes the output of the transcription job.
	Transcript *Transcript `type:"structure"`

	// A name to identify the transcription job.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The identifier assigned to the job when it was created.
	TranscriptionJobStatus *string `type:"string" enum:"TranscriptionJobStatus"`
}

// String returns the string representation
func (s TranscriptionJob) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TranscriptionJob) GoString() string {
	return s.String()
}

// SetCompletionTime sets the CompletionTime field's value.
func (s *TranscriptionJob) SetCompletionTime(v time.Time) *TranscriptionJob {
	s.CompletionTime = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *TranscriptionJob) SetCreationTime(v time.Time) *TranscriptionJob {
	s.CreationTime = &v
	return s
}

// SetFailureReason sets the FailureReason field's value.
func (s *TranscriptionJob) SetFailureReason(v string) *TranscriptionJob {
	s.FailureReason = &v
	return s
}

// SetLanguageCode sets the LanguageCode field's value.
func (s *TranscriptionJob) SetLanguageCode(v string) *TranscriptionJob {
	s.LanguageCode = &v
	return s
}

// SetMedia sets the Media field's value.
func (s *TranscriptionJob) SetMedia(v *Media) *TranscriptionJob {
	s.Media = v
	return s
}

// SetMediaFormat sets the MediaFormat field's value.
func (s *TranscriptionJob) SetMediaFormat(v string) *TranscriptionJob {
	s.MediaFormat = &v
	return s
}

// SetMediaSampleRateHertz sets the MediaSampleRateHertz field's value.
func (s *TranscriptionJob) SetMediaSampleRateHertz(v int64) *TranscriptionJob {
	s.MediaSampleRateHertz = &v
	return s
}

// SetTranscript sets the Transcript field's value.
func (s *TranscriptionJob) SetTranscript(v *Transcript) *TranscriptionJob {
	s.Transcript = v
	return s
}

// SetTranscriptionJobName sets the TranscriptionJobName field's value.
func (s *TranscriptionJob) SetTranscriptionJobName(v string) *TranscriptionJob {
	s.TranscriptionJobName = &v
	return s
}

// SetTranscriptionJobStatus sets the TranscriptionJobStatus field's value.
func (s *TranscriptionJob) SetTranscriptionJobStatus(v string) *TranscriptionJob {
	s.TranscriptionJobStatus = &v
	return s
}

// Provides a summary of information about a transcription job.
type TranscriptionJobSummary struct {
	_ struct{} `type:"structure"`

	// Timestamp of the date and time that the job completed.
	CompletionTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Timestamp of the date and time that the job was created.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// If the TranscriptionJobStatus field is FAILED, this field contains a description
	// of the error.
	FailureReason *string `type:"string"`

	// The language code for the input speech.
	LanguageCode *string `type:"string" enum:"LanguageCode"`

	// The name assigned to the transcription job when it was created.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The status of the transcription job. When the status is COMPLETED, use the
	// GetTranscriptionJob operation to get the results of the transcription.
	TranscriptionJobStatus *string `type:"string" enum:"TranscriptionJobStatus"`
}

// String returns the string representation
func (s TranscriptionJobSummary) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TranscriptionJobSummary) GoString() string {
	return s.String()
}

// SetCompletionTime sets the CompletionTime field's value.
func (s *TranscriptionJobSummary) SetCompletionTime(v time.Time) *TranscriptionJobSummary {
	s.CompletionTime = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *TranscriptionJobSummary) SetCreationTime(v time.Time) *TranscriptionJobSummary {
	s.CreationTime = &v
	return s
}

// SetFailureReason sets the FailureReason field's value.
func (s *TranscriptionJobSummary) SetFailureReason(v string) *TranscriptionJobSummary {
	s.FailureReason = &v
	return s
}

// SetLanguageCode sets the LanguageCode field's value.
func (s *TranscriptionJobSummary) SetLanguageCode(v string) *TranscriptionJobSummary {
	s.LanguageCode = &v
	return s
}

// SetTranscriptionJobName sets the TranscriptionJobName field's value.
func (s *TranscriptionJobSummary) SetTranscriptionJobName(v string) *TranscriptionJobSummary {
	s.TranscriptionJobName = &v
	return s
}

// SetTranscriptionJobStatus sets the TranscriptionJobStatus field's value.
func (s *TranscriptionJobSummary) SetTranscriptionJobStatus(v string) *TranscriptionJobSummary {
	s.TranscriptionJobStatus = &v
	return s
}

const (
	// LanguageCodeEnUs is a LanguageCode enum value
	LanguageCodeEnUs = "en-US"

	// LanguageCodeEsUs is a LanguageCode enum value
	LanguageCodeEsUs = "es-US"
)

const (
	// MediaFormatMp3 is a MediaFormat enum value
	MediaFormatMp3 = "mp3"

	// MediaFormatMp4 is a MediaFormat enum value
	MediaFormatMp4 = "mp4"

	// MediaFormatWav is a MediaFormat enum value
	MediaFormatWav = "wav"

	// MediaFormatFlac is a MediaFormat enum value
	MediaFormatFlac = "flac"
)

const (
	// TranscriptionJobStatusInProgress is a TranscriptionJobStatus enum value
	TranscriptionJobStatusInProgress = "IN_PROGRESS"

	// TranscriptionJobStatusFailed is a TranscriptionJobStatus enum value
	TranscriptionJobStatusFailed = "FAILED"

	// TranscriptionJobStatusCompleted is a TranscriptionJobStatus enum value
	TranscriptionJobStatusCompleted = "COMPLETED"
)