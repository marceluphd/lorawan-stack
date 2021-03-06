// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/gateway.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/struct"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/protobuf/field_mask"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *GatewayBrand) Validate() error {
	return nil
}
func (this *GatewayModel) Validate() error {
	return nil
}
func (this *GatewayVersionIdentifiers) Validate() error {
	return nil
}
func (this *GatewayRadio) Validate() error {
	if this.TxConfiguration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TxConfiguration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TxConfiguration", err)
		}
	}
	return nil
}
func (this *GatewayRadio_TxConfiguration) Validate() error {
	return nil
}
func (this *GatewayVersion) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayVersionIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayVersionIdentifiers", err)
	}
	for _, item := range this.Radios {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Radios", err)
			}
		}
	}
	return nil
}
func (this *Gateway) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayIdentifiers", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.CreatedAt)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.UpdatedAt)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.ContactInfo {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ContactInfo", err)
			}
		}
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayVersionIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayVersionIdentifiers", err)
	}
	for _, item := range this.Antennas {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(item)); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Antennas", err)
		}
	}
	return nil
}
func (this *Gateways) Validate() error {
	for _, item := range this.Gateways {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Gateways", err)
			}
		}
	}
	return nil
}
func (this *GetGatewayRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayIdentifiers", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.FieldMask)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("FieldMask", err)
	}
	return nil
}
func (this *GetGatewayIdentifiersForEUIRequest) Validate() error {
	return nil
}
func (this *ListGatewaysRequest) Validate() error {
	if this.Collaborator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Collaborator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Collaborator", err)
		}
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.FieldMask)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("FieldMask", err)
	}
	return nil
}
func (this *CreateGatewayRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Gateway)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Gateway", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Collaborator)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Collaborator", err)
	}
	return nil
}
func (this *UpdateGatewayRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Gateway)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Gateway", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.FieldMask)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("FieldMask", err)
	}
	return nil
}
func (this *CreateGatewayAPIKeyRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayIdentifiers", err)
	}
	return nil
}
func (this *UpdateGatewayAPIKeyRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayIdentifiers", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.APIKey)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("APIKey", err)
	}
	return nil
}
func (this *SetGatewayCollaboratorRequest) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.GatewayIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayIdentifiers", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Collaborator)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Collaborator", err)
	}
	return nil
}
func (this *GatewayAntenna) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Location)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Location", err)
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GatewayStatus) Validate() error {
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.Time)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("Time", err)
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.BootTime)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("BootTime", err)
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.AntennaLocations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AntennaLocations", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Advanced != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Advanced); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Advanced", err)
		}
	}
	return nil
}
func (this *GatewayConnectionStats) Validate() error {
	if this.ConnectedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectedAt", err)
		}
	}
	if this.LastStatusReceivedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastStatusReceivedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastStatusReceivedAt", err)
		}
	}
	if this.LastStatus != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastStatus); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastStatus", err)
		}
	}
	if this.LastUplinkReceivedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUplinkReceivedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUplinkReceivedAt", err)
		}
	}
	if this.LastDownlinkReceivedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastDownlinkReceivedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastDownlinkReceivedAt", err)
		}
	}
	return nil
}
