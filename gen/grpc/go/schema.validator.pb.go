// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package eventgate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *HistoryOpts) Validate() error {
	if this.Channel == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must not be an empty string`, this.Channel))
	}
	if this.Min != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Min); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Min", err)
		}
	}
	if this.Max != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Max); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Max", err)
		}
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}
func (this *ReceiveOpts) Validate() error {
	if this.Channel == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must not be an empty string`, this.Channel))
	}
	return nil
}
func (this *Event) Validate() error {
	if this.Channel == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must not be an empty string`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if nil == this.Claims {
		return github_com_mwitkow_go_proto_validators.FieldError("Claims", fmt.Errorf("message must exist"))
	}
	if this.Claims != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Claims); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Claims", err)
		}
	}
	if this.Time != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Time); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Time", err)
		}
	}
	return nil
}
func (this *Events) Validate() error {
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}
