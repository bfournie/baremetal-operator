package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCheckSettingIsValid(t *testing.T) {

	lower_bound := 1
	upper_bound := 20
	min_length := 1
	max_length := 16
	read_only := true

	fwSchema := &FirmwareSchema{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "myfwschema",
			Namespace: "myns",
		},
		Status: FirmwareSchemaStatus{
			Name: "VendorA-ModelT-Q456TY347111222333",

			Schema: make(map[string]SettingSchema),
		},
	}

	fwSchema.Status.Schema["AssetTag"] = SettingSchema{AttributeType: "String",
		MinLength: &min_length, MaxLength: &max_length}
	fwSchema.Status.Schema["ProcVirtualization"] = SettingSchema{AttributeType: "Enumeration",
		AllowableValues: []string{"Enabled", "Disabled"}}
	fwSchema.Status.Schema["NetworkBootRetryCount"] = SettingSchema{AttributeType: "Integer",
		LowerBound: &lower_bound, UpperBound: &upper_bound}
	fwSchema.Status.Schema["SerialNumber"] = SettingSchema{AttributeType: "String",
		MinLength: &min_length, MaxLength: &max_length, ReadOnly: &read_only}
	fwSchema.Status.Schema["QuietBoot"] = SettingSchema{AttributeType: "Boolean"}
	fwSchema.Status.Schema["SriovEnable"] = SettingSchema{} // fields intentionally excluded

	for _, tc := range []struct {
		Scenario string
		Name     string
		Value    string
		Expected bool
	}{
		{
			Scenario: "StringTypePass",
			Name:     "AssetTag",
			Value:    "NewServer",
			Expected: true,
		},
		{
			Scenario: "StringTypeFailUpper",
			Name:     "AssetTag",
			Value:    "NewServerPutInServiceIn2021",
			Expected: false,
		},
		{
			Scenario: "StringTypeFailLower",
			Name:     "AssetTag",
			Value:    "",
			Expected: false,
		},
		{
			Scenario: "EnumerationTypePass",
			Name:     "ProcVirtualization",
			Value:    "Disabled",
			Expected: true,
		},
		{
			Scenario: "EnumerationTypeFail",
			Name:     "ProcVirtualization",
			Value:    "Foo",
			Expected: false,
		},
		{
			Scenario: "IntegerTypePass",
			Name:     "NetworkBootRetryCount",
			Value:    "10",
			Expected: true,
		},
		{
			Scenario: "IntegerTypeFailUpper",
			Name:     "NetworkBootRetryCount",
			Value:    "42",
			Expected: false,
		},
		{
			Scenario: "IntegerTypeFailLower",
			Name:     "NetworkBootRetryCount",
			Value:    "0",
			Expected: false,
		},
		{
			Scenario: "BooleanTypePass",
			Name:     "QuietBoot",
			Value:    "true",
			Expected: true,
		},
		{
			Scenario: "BooleanTypeFail",
			Name:     "QuietBoot",
			Value:    "Enabled",
			Expected: false,
		},
		{
			Scenario: "ReadOnlyTypeFail",
			Name:     "SerialNumber",
			Value:    "42",
			Expected: false,
		},
		{
			Scenario: "MissingEnumerationField",
			Name:     "SriovEnable",
			Value:    "Disabled",
			Expected: true,
		},
		{
			Scenario: "UnknownSettingFail",
			Name:     "IceCream",
			Value:    "Vanilla",
			Expected: false,
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			actual := fwSchema.CheckSettingIsValid(tc.Name, tc.Value, fwSchema.Status.Schema)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
