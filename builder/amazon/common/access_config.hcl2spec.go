// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AccessConfig) HCL2Schema() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"AccessKey": &hcldec.AttrSpec{Name:"AccessKey", Type:cty.String, Required:false},
		"CustomEndpointEc2": &hcldec.AttrSpec{Name:"CustomEndpointEc2", Type:cty.String, Required:false},
		"DecodeAuthZMessages": &hcldec.AttrSpec{Name:"DecodeAuthZMessages", Type:cty.Bool, Required:false},
		"InsecureSkipTLSVerify": &hcldec.AttrSpec{Name:"InsecureSkipTLSVerify", Type:cty.Bool, Required:false},
		"MFACode": &hcldec.AttrSpec{Name:"MFACode", Type:cty.String, Required:false},
		"ProfileName": &hcldec.AttrSpec{Name:"ProfileName", Type:cty.String, Required:false},
		"RawRegion": &hcldec.AttrSpec{Name:"RawRegion", Type:cty.String, Required:false},
		"SecretKey": &hcldec.AttrSpec{Name:"SecretKey", Type:cty.String, Required:false},
		"SkipValidation": &hcldec.AttrSpec{Name:"SkipValidation", Type:cty.Bool, Required:false},
		"SkipMetadataApiCheck": &hcldec.AttrSpec{Name:"SkipMetadataApiCheck", Type:cty.Bool, Required:false},
		"Token": &hcldec.AttrSpec{Name:"Token", Type:cty.String, Required:false},
		"VaultAWSEngine": nil,
	}
	return hcldec.ObjectSpec(s)
}

func (*VaultAWSEngineOptions) HCL2Schema() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"Name": &hcldec.AttrSpec{Name:"Name", Type:cty.String, Required:false},
		"RoleARN": &hcldec.AttrSpec{Name:"RoleARN", Type:cty.String, Required:false},
		"TTL": &hcldec.AttrSpec{Name:"TTL", Type:cty.String, Required:false},
		"EngineName": &hcldec.AttrSpec{Name:"EngineName", Type:cty.String, Required:false},
	}
	return hcldec.ObjectSpec(s)
}

