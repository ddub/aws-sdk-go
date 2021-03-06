// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudtrail provides a client for AWS CloudTrail.
package cloudtrail

import (
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateTrailRequest generates a request for the CreateTrail operation.
func (c *CloudTrail) CreateTrailRequest(input *CreateTrailInput) (req *aws.Request, output *CreateTrailOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateTrail == nil {
		opCreateTrail = &aws.Operation{
			Name:       "CreateTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateTrailInput{}
	}

	req = c.newRequest(opCreateTrail, input, output)
	output = &CreateTrailOutput{}
	req.Data = output
	return
}

// From the command line, use create-subscription.
//
// Creates a trail that specifies the settings for delivery of log data to
// an Amazon S3 bucket.
func (c *CloudTrail) CreateTrail(input *CreateTrailInput) (output *CreateTrailOutput, err error) {
	req, out := c.CreateTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateTrail *aws.Operation

// DeleteTrailRequest generates a request for the DeleteTrail operation.
func (c *CloudTrail) DeleteTrailRequest(input *DeleteTrailInput) (req *aws.Request, output *DeleteTrailOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteTrail == nil {
		opDeleteTrail = &aws.Operation{
			Name:       "DeleteTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteTrailInput{}
	}

	req = c.newRequest(opDeleteTrail, input, output)
	output = &DeleteTrailOutput{}
	req.Data = output
	return
}

// Deletes a trail.
func (c *CloudTrail) DeleteTrail(input *DeleteTrailInput) (output *DeleteTrailOutput, err error) {
	req, out := c.DeleteTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTrail *aws.Operation

// DescribeTrailsRequest generates a request for the DescribeTrails operation.
func (c *CloudTrail) DescribeTrailsRequest(input *DescribeTrailsInput) (req *aws.Request, output *DescribeTrailsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeTrails == nil {
		opDescribeTrails = &aws.Operation{
			Name:       "DescribeTrails",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeTrailsInput{}
	}

	req = c.newRequest(opDescribeTrails, input, output)
	output = &DescribeTrailsOutput{}
	req.Data = output
	return
}

// Retrieves settings for the trail associated with the current region for your
// account.
func (c *CloudTrail) DescribeTrails(input *DescribeTrailsInput) (output *DescribeTrailsOutput, err error) {
	req, out := c.DescribeTrailsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrails *aws.Operation

// GetTrailStatusRequest generates a request for the GetTrailStatus operation.
func (c *CloudTrail) GetTrailStatusRequest(input *GetTrailStatusInput) (req *aws.Request, output *GetTrailStatusOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetTrailStatus == nil {
		opGetTrailStatus = &aws.Operation{
			Name:       "GetTrailStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetTrailStatusInput{}
	}

	req = c.newRequest(opGetTrailStatus, input, output)
	output = &GetTrailStatusOutput{}
	req.Data = output
	return
}

// Returns a JSON-formatted list of information about the specified trail. Fields
// include information on delivery errors, Amazon SNS and Amazon S3 errors,
// and start and stop logging times for each trail.
func (c *CloudTrail) GetTrailStatus(input *GetTrailStatusInput) (output *GetTrailStatusOutput, err error) {
	req, out := c.GetTrailStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetTrailStatus *aws.Operation

// LookupEventsRequest generates a request for the LookupEvents operation.
func (c *CloudTrail) LookupEventsRequest(input *LookupEventsInput) (req *aws.Request, output *LookupEventsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opLookupEvents == nil {
		opLookupEvents = &aws.Operation{
			Name:       "LookupEvents",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &LookupEventsInput{}
	}

	req = c.newRequest(opLookupEvents, input, output)
	output = &LookupEventsOutput{}
	req.Data = output
	return
}

// Looks up API activity events captured by CloudTrail that create, update,
// or delete resources in your account. Events for a region can be looked up
// for the times in which you had CloudTrail turned on in that region during
// the last seven days. Lookup supports five different attributes: time range
// (defined by a start time and end time), user name, event name, resource type,
// and resource name. All attributes are optional. The maximum number of attributes
// that can be specified in any one lookup request are time range and one other
// attribute. The default number of results returned is 10, with a maximum of
// 50 possible. The response includes a token that you can use to get the next
// page of results. The rate of lookup requests is limited to one per second
// per account.
//
// Events that occurred during the selected time range will not be available
// for lookup if CloudTrail logging was not enabled when the events occurred.
func (c *CloudTrail) LookupEvents(input *LookupEventsInput) (output *LookupEventsOutput, err error) {
	req, out := c.LookupEventsRequest(input)
	output = out
	err = req.Send()
	return
}

var opLookupEvents *aws.Operation

// StartLoggingRequest generates a request for the StartLogging operation.
func (c *CloudTrail) StartLoggingRequest(input *StartLoggingInput) (req *aws.Request, output *StartLoggingOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opStartLogging == nil {
		opStartLogging = &aws.Operation{
			Name:       "StartLogging",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &StartLoggingInput{}
	}

	req = c.newRequest(opStartLogging, input, output)
	output = &StartLoggingOutput{}
	req.Data = output
	return
}

// Starts the recording of AWS API calls and log file delivery for a trail.
func (c *CloudTrail) StartLogging(input *StartLoggingInput) (output *StartLoggingOutput, err error) {
	req, out := c.StartLoggingRequest(input)
	output = out
	err = req.Send()
	return
}

var opStartLogging *aws.Operation

// StopLoggingRequest generates a request for the StopLogging operation.
func (c *CloudTrail) StopLoggingRequest(input *StopLoggingInput) (req *aws.Request, output *StopLoggingOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opStopLogging == nil {
		opStopLogging = &aws.Operation{
			Name:       "StopLogging",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &StopLoggingInput{}
	}

	req = c.newRequest(opStopLogging, input, output)
	output = &StopLoggingOutput{}
	req.Data = output
	return
}

// Suspends the recording of AWS API calls and log file delivery for the specified
// trail. Under most circumstances, there is no need to use this action. You
// can update a trail without stopping it first. This action is the only way
// to stop recording.
func (c *CloudTrail) StopLogging(input *StopLoggingInput) (output *StopLoggingOutput, err error) {
	req, out := c.StopLoggingRequest(input)
	output = out
	err = req.Send()
	return
}

var opStopLogging *aws.Operation

// UpdateTrailRequest generates a request for the UpdateTrail operation.
func (c *CloudTrail) UpdateTrailRequest(input *UpdateTrailInput) (req *aws.Request, output *UpdateTrailOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUpdateTrail == nil {
		opUpdateTrail = &aws.Operation{
			Name:       "UpdateTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &UpdateTrailInput{}
	}

	req = c.newRequest(opUpdateTrail, input, output)
	output = &UpdateTrailOutput{}
	req.Data = output
	return
}

// From the command line, use update-subscription.
//
// Updates the settings that specify delivery of log files. Changes to a trail
// do not require stopping the CloudTrail service. Use this action to designate
// an existing bucket for log delivery. If the existing bucket has previously
// been a target for CloudTrail log files, an IAM policy exists for the bucket.
func (c *CloudTrail) UpdateTrail(input *UpdateTrailInput) (output *UpdateTrailOutput, err error) {
	req, out := c.UpdateTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateTrail *aws.Operation

// Specifies the settings for each trail.
type CreateTrailInput struct {
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique
	// identifier that represents the log group to which CloudTrail logs will be
	// delivered. Not required unless you specify CloudWatchLogsRoleArn.
	CloudWatchLogsLogGroupARN *string `locationName:"CloudWatchLogsLogGroupArn" type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user’s log group.
	CloudWatchLogsRoleARN *string `locationName:"CloudWatchLogsRoleArn" type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string" required:"true"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string" required:"true"`

	// Specifies the Amazon S3 key prefix that precedes the name of the bucket you
	// have designated for log file delivery.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery.
	SNSTopicName *string `locationName:"SnsTopicName" type:"string"`

	metadataCreateTrailInput `json:"-", xml:"-"`
}

type metadataCreateTrailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type CreateTrailOutput struct {
	// Specifies the Amazon Resource Name (ARN) of the log group to which CloudTrail
	// logs will be delivered.
	CloudWatchLogsLogGroupARN *string `locationName:"CloudWatchLogsLogGroupArn" type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user’s log group.
	CloudWatchLogsRoleARN *string `locationName:"CloudWatchLogsRoleArn" type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that precedes the name of the bucket you
	// have designated for log file delivery.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery.
	SNSTopicName *string `locationName:"SnsTopicName" type:"string"`

	metadataCreateTrailOutput `json:"-", xml:"-"`
}

type metadataCreateTrailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The request that specifies the name of a trail to delete.
type DeleteTrailInput struct {
	// The name of a trail to be deleted.
	Name *string `type:"string" required:"true"`

	metadataDeleteTrailInput `json:"-", xml:"-"`
}

type metadataDeleteTrailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type DeleteTrailOutput struct {
	metadataDeleteTrailOutput `json:"-", xml:"-"`
}

type metadataDeleteTrailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns information about the trail.
type DescribeTrailsInput struct {
	// The trail returned.
	TrailNameList []*string `locationName:"trailNameList" type:"list"`

	metadataDescribeTrailsInput `json:"-", xml:"-"`
}

type metadataDescribeTrailsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type DescribeTrailsOutput struct {
	// The list of trails.
	TrailList []*Trail `locationName:"trailList" type:"list"`

	metadataDescribeTrailsOutput `json:"-", xml:"-"`
}

type metadataDescribeTrailsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains information about an event that was returned by a lookup request.
// The result includes a representation of a CloudTrail event.
type Event struct {
	// A JSON string that contains a representation of the event returned.
	CloudTrailEvent *string `type:"string"`

	// The CloudTrail ID of the event returned.
	EventID *string `locationName:"EventId" type:"string"`

	// The name of the event returned.
	EventName *string `type:"string"`

	// The date and time of the event returned.
	EventTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A list of resources referenced by the event returned.
	Resources []*Resource `type:"list"`

	// A user name or role name of the requester that called the API in the event
	// returned.
	Username *string `type:"string"`

	metadataEvent `json:"-", xml:"-"`
}

type metadataEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

// The name of a trail about which you want the current status.
type GetTrailStatusInput struct {
	// The name of the trail for which you are requesting the current status.
	Name *string `type:"string" required:"true"`

	metadataGetTrailStatusInput `json:"-", xml:"-"`
}

type metadataGetTrailStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type GetTrailStatusOutput struct {
	// Whether the CloudTrail is currently logging AWS API calls.
	IsLogging *bool `type:"boolean"`

	// Displays any CloudWatch Logs error that CloudTrail encountered when attempting
	// to deliver logs to CloudWatch Logs.
	LatestCloudWatchLogsDeliveryError *string `type:"string"`

	// Displays the most recent date and time when CloudTrail delivered logs to
	// CloudWatch Logs.
	LatestCloudWatchLogsDeliveryTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Displays any Amazon S3 error that CloudTrail encountered when attempting
	// to deliver log files to the designated bucket. For more information see the
	// topic Error Responses (http://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html)
	// in the Amazon S3 API Reference.
	LatestDeliveryError *string `type:"string"`

	// Specifies the date and time that CloudTrail last delivered log files to an
	// account's Amazon S3 bucket.
	LatestDeliveryTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Displays any Amazon SNS error that CloudTrail encountered when attempting
	// to send a notification. For more information about Amazon SNS errors, see
	// the Amazon SNS Developer Guide (http://docs.aws.amazon.com/sns/latest/dg/welcome.html).
	LatestNotificationError *string `type:"string"`

	// Specifies the date and time of the most recent Amazon SNS notification that
	// CloudTrail has written a new log file to an account's Amazon S3 bucket.
	LatestNotificationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Specifies the most recent date and time when CloudTrail started recording
	// API calls for an AWS account.
	StartLoggingTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Specifies the most recent date and time when CloudTrail stopped recording
	// API calls for an AWS account.
	StopLoggingTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	metadataGetTrailStatusOutput `json:"-", xml:"-"`
}

type metadataGetTrailStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Specifies an attribute and value that filter the events returned.
type LookupAttribute struct {
	// Specifies an attribute on which to filter the events returned.
	AttributeKey *string `type:"string" required:"true"`

	// Specifies a value for the specified AttributeKey.
	AttributeValue *string `type:"string" required:"true"`

	metadataLookupAttribute `json:"-", xml:"-"`
}

type metadataLookupAttribute struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains a request for LookupEvents.
type LookupEventsInput struct {
	// Specifies that only events that occur before or at the specified time are
	// returned. If the specified end time is before the specified start time, an
	// error is returned.
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Contains a list of lookup attributes. Currently the list can contain only
	// one item.
	LookupAttributes []*LookupAttribute `type:"list"`

	// The number of events to return. Possible values are 1 through 50. The default
	// is 10.
	MaxResults *int64 `type:"integer"`

	// The token to use to get the next page of results after a previous API call.
	// This token must be passed in with the same parameters that were specified
	// in the the original call. For example, if the original call specified an
	// AttributeKey of 'Username' with a value of 'root', the call with NextToken
	// should include those same parameters.
	NextToken *string `type:"string"`

	// Specifies that only events that occur after or at the specified time are
	// returned. If the specified start time is after the specified end time, an
	// error is returned.
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	metadataLookupEventsInput `json:"-", xml:"-"`
}

type metadataLookupEventsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains a response to a LookupEvents action.
type LookupEventsOutput struct {
	// A list of events returned based on the lookup attributes specified and the
	// CloudTrail event. The events list is sorted by time. The most recent event
	// is listed first.
	Events []*Event `type:"list"`

	// The token to use to get the next page of results after a previous API call.
	// If the token does not appear, there are no more results to return. The token
	// must be passed in with the same parameters as the previous call. For example,
	// if the original call specified an AttributeKey of 'Username' with a value
	// of 'root', the call with NextToken should include those same parameters.
	NextToken *string `type:"string"`

	metadataLookupEventsOutput `json:"-", xml:"-"`
}

type metadataLookupEventsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Specifies the type and name of a resource referenced by an event.
type Resource struct {
	// The name of the resource referenced by the event returned. These are user-created
	// names whose values will depend on the environment. For example, the resource
	// name might be "auto-scaling-test-group" for an Auto Scaling Group or "i-1234567"
	// for an EC2 Instance.
	ResourceName *string `type:"string"`

	// The type of a resource referenced by the event returned. When the resource
	// type cannot be determined, null is returned. Some examples of resource types
	// are: Instance for EC2, Trail for CloudTrail, DBInstance for RDS, and AccessKey
	// for IAM. For a list of resource types supported for event lookup, see Resource
	// Types Supported for Event Lookup (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/lookup_supported_resourcetypes.html).
	ResourceType *string `type:"string"`

	metadataResource `json:"-", xml:"-"`
}

type metadataResource struct {
	SDKShapeTraits bool `type:"structure"`
}

// The request to CloudTrail to start logging AWS API calls for an account.
type StartLoggingInput struct {
	// The name of the trail for which CloudTrail logs AWS API calls.
	Name *string `type:"string" required:"true"`

	metadataStartLoggingInput `json:"-", xml:"-"`
}

type metadataStartLoggingInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type StartLoggingOutput struct {
	metadataStartLoggingOutput `json:"-", xml:"-"`
}

type metadataStartLoggingOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Passes the request to CloudTrail to stop logging AWS API calls for the specified
// account.
type StopLoggingInput struct {
	// Communicates to CloudTrail the name of the trail for which to stop logging
	// AWS API calls.
	Name *string `type:"string" required:"true"`

	metadataStopLoggingInput `json:"-", xml:"-"`
}

type metadataStopLoggingInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type StopLoggingOutput struct {
	metadataStopLoggingOutput `json:"-", xml:"-"`
}

type metadataStopLoggingOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The settings for a trail.
type Trail struct {
	// Specifies an Amazon Resource Name (ARN), a unique identifier that represents
	// the log group to which CloudTrail logs will be delivered.
	CloudWatchLogsLogGroupARN *string `locationName:"CloudWatchLogsLogGroupArn" type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user’s log group.
	CloudWatchLogsRoleARN *string `locationName:"CloudWatchLogsRoleArn" type:"string"`

	// Set to True to include AWS API calls from AWS global services such as IAM.
	// Otherwise, False.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Name of the trail set by calling CreateTrail.
	Name *string `type:"string"`

	// Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.
	S3BucketName *string `type:"string"`

	// Value of the Amazon S3 prefix.
	S3KeyPrefix *string `type:"string"`

	// Name of the existing Amazon SNS topic that CloudTrail uses to notify the
	// account owner when new CloudTrail log files have been delivered.
	SNSTopicName *string `locationName:"SnsTopicName" type:"string"`

	metadataTrail `json:"-", xml:"-"`
}

type metadataTrail struct {
	SDKShapeTraits bool `type:"structure"`
}

// Specifies settings to update for the trail.
type UpdateTrailInput struct {
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique
	// identifier that represents the log group to which CloudTrail logs will be
	// delivered. Not required unless you specify CloudWatchLogsRoleArn.
	CloudWatchLogsLogGroupARN *string `locationName:"CloudWatchLogsLogGroupArn" type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user’s log group.
	CloudWatchLogsRoleARN *string `locationName:"CloudWatchLogsRoleArn" type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string" required:"true"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that precedes the name of the bucket you
	// have designated for log file delivery.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery.
	SNSTopicName *string `locationName:"SnsTopicName" type:"string"`

	metadataUpdateTrailInput `json:"-", xml:"-"`
}

type metadataUpdateTrailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type UpdateTrailOutput struct {
	// Specifies the Amazon Resource Name (ARN) of the log group to which CloudTrail
	// logs will be delivered.
	CloudWatchLogsLogGroupARN *string `locationName:"CloudWatchLogsLogGroupArn" type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user’s log group.
	CloudWatchLogsRoleARN *string `locationName:"CloudWatchLogsRoleArn" type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that precedes the name of the bucket you
	// have designated for log file delivery.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery.
	SNSTopicName *string `locationName:"SnsTopicName" type:"string"`

	metadataUpdateTrailOutput `json:"-", xml:"-"`
}

type metadataUpdateTrailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}