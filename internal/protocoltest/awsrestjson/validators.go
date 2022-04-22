// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpConstantQueryString struct {
}

func (*validateOpConstantQueryString) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConstantQueryString) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConstantQueryStringInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConstantQueryStringInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEndpointWithHostLabelOperation struct {
}

func (*validateOpEndpointWithHostLabelOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEndpointWithHostLabelOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEndpointWithHostLabelOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithFloatLabels struct {
}

func (*validateOpHttpRequestWithFloatLabels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithFloatLabels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithFloatLabelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithFloatLabelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithGreedyLabelInPath struct {
}

func (*validateOpHttpRequestWithGreedyLabelInPath) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithGreedyLabelInPath) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithGreedyLabelInPathInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithGreedyLabelInPathInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithLabelsAndTimestampFormat struct {
}

func (*validateOpHttpRequestWithLabelsAndTimestampFormat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithLabelsAndTimestampFormat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithLabelsAndTimestampFormatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithLabelsAndTimestampFormatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithLabels struct {
}

func (*validateOpHttpRequestWithLabels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithLabels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithLabelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithLabelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithRegexLiteral struct {
}

func (*validateOpHttpRequestWithRegexLiteral) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithRegexLiteral) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithRegexLiteralInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithRegexLiteralInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedBoolean struct {
}

func (*validateOpMalformedBoolean) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedBoolean) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedBooleanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedBooleanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedByte struct {
}

func (*validateOpMalformedByte) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedByte) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedByteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedByteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedDouble struct {
}

func (*validateOpMalformedDouble) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedDouble) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedDoubleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedDoubleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedFloat struct {
}

func (*validateOpMalformedFloat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedFloat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedFloatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedFloatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedInteger struct {
}

func (*validateOpMalformedInteger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedInteger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedIntegerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedIntegerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedLong struct {
}

func (*validateOpMalformedLong) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedLong) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedLongInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedLongInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedShort struct {
}

func (*validateOpMalformedShort) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedShort) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedShortInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedShortInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampBodyDateTime struct {
}

func (*validateOpMalformedTimestampBodyDateTime) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampBodyDateTime) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampBodyDateTimeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampBodyDateTimeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampBodyDefault struct {
}

func (*validateOpMalformedTimestampBodyDefault) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampBodyDefault) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampBodyDefaultInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampBodyDefaultInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampBodyHttpDate struct {
}

func (*validateOpMalformedTimestampBodyHttpDate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampBodyHttpDate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampBodyHttpDateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampBodyHttpDateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampHeaderDateTime struct {
}

func (*validateOpMalformedTimestampHeaderDateTime) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampHeaderDateTime) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampHeaderDateTimeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampHeaderDateTimeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampHeaderDefault struct {
}

func (*validateOpMalformedTimestampHeaderDefault) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampHeaderDefault) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampHeaderDefaultInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampHeaderDefaultInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampHeaderEpoch struct {
}

func (*validateOpMalformedTimestampHeaderEpoch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampHeaderEpoch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampHeaderEpochInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampHeaderEpochInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampPathDefault struct {
}

func (*validateOpMalformedTimestampPathDefault) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampPathDefault) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampPathDefaultInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampPathDefaultInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampPathEpoch struct {
}

func (*validateOpMalformedTimestampPathEpoch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampPathEpoch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampPathEpochInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampPathEpochInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampPathHttpDate struct {
}

func (*validateOpMalformedTimestampPathHttpDate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampPathHttpDate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampPathHttpDateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampPathHttpDateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampQueryDefault struct {
}

func (*validateOpMalformedTimestampQueryDefault) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampQueryDefault) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampQueryDefaultInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampQueryDefaultInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampQueryEpoch struct {
}

func (*validateOpMalformedTimestampQueryEpoch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampQueryEpoch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampQueryEpochInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampQueryEpochInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMalformedTimestampQueryHttpDate struct {
}

func (*validateOpMalformedTimestampQueryHttpDate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMalformedTimestampQueryHttpDate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MalformedTimestampQueryHttpDateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMalformedTimestampQueryHttpDateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpConstantQueryStringValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConstantQueryString{}, middleware.After)
}

func addOpEndpointWithHostLabelOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEndpointWithHostLabelOperation{}, middleware.After)
}

func addOpHttpRequestWithFloatLabelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithFloatLabels{}, middleware.After)
}

func addOpHttpRequestWithGreedyLabelInPathValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithGreedyLabelInPath{}, middleware.After)
}

func addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
}

func addOpHttpRequestWithLabelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithLabels{}, middleware.After)
}

func addOpHttpRequestWithRegexLiteralValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithRegexLiteral{}, middleware.After)
}

func addOpMalformedBooleanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedBoolean{}, middleware.After)
}

func addOpMalformedByteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedByte{}, middleware.After)
}

func addOpMalformedDoubleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedDouble{}, middleware.After)
}

func addOpMalformedFloatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedFloat{}, middleware.After)
}

func addOpMalformedIntegerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedInteger{}, middleware.After)
}

func addOpMalformedLongValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedLong{}, middleware.After)
}

func addOpMalformedShortValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedShort{}, middleware.After)
}

func addOpMalformedTimestampBodyDateTimeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampBodyDateTime{}, middleware.After)
}

func addOpMalformedTimestampBodyDefaultValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampBodyDefault{}, middleware.After)
}

func addOpMalformedTimestampBodyHttpDateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampBodyHttpDate{}, middleware.After)
}

func addOpMalformedTimestampHeaderDateTimeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampHeaderDateTime{}, middleware.After)
}

func addOpMalformedTimestampHeaderDefaultValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampHeaderDefault{}, middleware.After)
}

func addOpMalformedTimestampHeaderEpochValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampHeaderEpoch{}, middleware.After)
}

func addOpMalformedTimestampPathDefaultValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampPathDefault{}, middleware.After)
}

func addOpMalformedTimestampPathEpochValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampPathEpoch{}, middleware.After)
}

func addOpMalformedTimestampPathHttpDateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampPathHttpDate{}, middleware.After)
}

func addOpMalformedTimestampQueryDefaultValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampQueryDefault{}, middleware.After)
}

func addOpMalformedTimestampQueryEpochValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampQueryEpoch{}, middleware.After)
}

func addOpMalformedTimestampQueryHttpDateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMalformedTimestampQueryHttpDate{}, middleware.After)
}

func validateOpConstantQueryStringInput(v *ConstantQueryStringInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConstantQueryStringInput"}
	if v.Hello == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Hello"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointWithHostLabelOperationInput"}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithFloatLabelsInput(v *HttpRequestWithFloatLabelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithFloatLabelsInput"}
	if v.Float == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Float"))
	}
	if v.Double == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Double"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithGreedyLabelInPathInput(v *HttpRequestWithGreedyLabelInPathInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithGreedyLabelInPathInput"}
	if v.Foo == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Foo"))
	}
	if v.Baz == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Baz"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithLabelsAndTimestampFormatInput(v *HttpRequestWithLabelsAndTimestampFormatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithLabelsAndTimestampFormatInput"}
	if v.MemberEpochSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberEpochSeconds"))
	}
	if v.MemberHttpDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberHttpDate"))
	}
	if v.MemberDateTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberDateTime"))
	}
	if v.DefaultFormat == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DefaultFormat"))
	}
	if v.TargetEpochSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetEpochSeconds"))
	}
	if v.TargetHttpDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetHttpDate"))
	}
	if v.TargetDateTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetDateTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithLabelsInput(v *HttpRequestWithLabelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithLabelsInput"}
	if v.String_ == nil {
		invalidParams.Add(smithy.NewErrParamRequired("String_"))
	}
	if v.Short == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Short"))
	}
	if v.Integer == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Integer"))
	}
	if v.Long == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Long"))
	}
	if v.Float == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Float"))
	}
	if v.Double == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Double"))
	}
	if v.Boolean == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Boolean"))
	}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithRegexLiteralInput(v *HttpRequestWithRegexLiteralInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithRegexLiteralInput"}
	if v.Str == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Str"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedBooleanInput(v *MalformedBooleanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedBooleanInput"}
	if v.BooleanInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BooleanInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedByteInput(v *MalformedByteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedByteInput"}
	if v.ByteInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ByteInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedDoubleInput(v *MalformedDoubleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedDoubleInput"}
	if v.DoubleInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DoubleInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedFloatInput(v *MalformedFloatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedFloatInput"}
	if v.FloatInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FloatInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedIntegerInput(v *MalformedIntegerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedIntegerInput"}
	if v.IntegerInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IntegerInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedLongInput(v *MalformedLongInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedLongInput"}
	if v.LongInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LongInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedShortInput(v *MalformedShortInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedShortInput"}
	if v.ShortInPath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShortInPath"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampBodyDateTimeInput(v *MalformedTimestampBodyDateTimeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampBodyDateTimeInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampBodyDefaultInput(v *MalformedTimestampBodyDefaultInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampBodyDefaultInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampBodyHttpDateInput(v *MalformedTimestampBodyHttpDateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampBodyHttpDateInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampHeaderDateTimeInput(v *MalformedTimestampHeaderDateTimeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampHeaderDateTimeInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampHeaderDefaultInput(v *MalformedTimestampHeaderDefaultInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampHeaderDefaultInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampHeaderEpochInput(v *MalformedTimestampHeaderEpochInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampHeaderEpochInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampPathDefaultInput(v *MalformedTimestampPathDefaultInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampPathDefaultInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampPathEpochInput(v *MalformedTimestampPathEpochInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampPathEpochInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampPathHttpDateInput(v *MalformedTimestampPathHttpDateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampPathHttpDateInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampQueryDefaultInput(v *MalformedTimestampQueryDefaultInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampQueryDefaultInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampQueryEpochInput(v *MalformedTimestampQueryEpochInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampQueryEpochInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMalformedTimestampQueryHttpDateInput(v *MalformedTimestampQueryHttpDateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MalformedTimestampQueryHttpDateInput"}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
