# Go API client for goflagr

Flagr is a feature flagging, A/B testing and dynamic configuration microservice

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 0.1.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./goflagr"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConstraintApi* | [**CreateConstraint**](docs/ConstraintApi.md#createconstraint) | **Post** /flags/{flagID}/segments/{segmentID}/constraints | 
*ConstraintApi* | [**DeleteConstraint**](docs/ConstraintApi.md#deleteconstraint) | **Delete** /flags/{flagID}/segments/{segmentID}/constraints/{constraintID} | 
*ConstraintApi* | [**FindConstraints**](docs/ConstraintApi.md#findconstraints) | **Get** /flags/{flagID}/segments/{segmentID}/constraints | 
*ConstraintApi* | [**PutConstraint**](docs/ConstraintApi.md#putconstraint) | **Put** /flags/{flagID}/segments/{segmentID}/constraints/{constraintID} | 
*DistributionApi* | [**FindDistributions**](docs/DistributionApi.md#finddistributions) | **Get** /flags/{flagID}/segments/{segmentID}/distributions | 
*DistributionApi* | [**PutDistributions**](docs/DistributionApi.md#putdistributions) | **Put** /flags/{flagID}/segments/{segmentID}/distributions | 
*EvaluationApi* | [**PostEvaluation**](docs/EvaluationApi.md#postevaluation) | **Post** /evaluation | 
*EvaluationApi* | [**PostEvaluationBatch**](docs/EvaluationApi.md#postevaluationbatch) | **Post** /evaluation/batch | 
*FlagApi* | [**CreateFlag**](docs/FlagApi.md#createflag) | **Post** /flags | 
*FlagApi* | [**DeleteFlag**](docs/FlagApi.md#deleteflag) | **Delete** /flags/{flagID} | 
*FlagApi* | [**FindFlags**](docs/FlagApi.md#findflags) | **Get** /flags | 
*FlagApi* | [**GetFlag**](docs/FlagApi.md#getflag) | **Get** /flags/{flagID} | 
*FlagApi* | [**GetFlagSnapshots**](docs/FlagApi.md#getflagsnapshots) | **Get** /flags/{flagID}/snapshots | 
*FlagApi* | [**PutFlag**](docs/FlagApi.md#putflag) | **Put** /flags/{flagID} | 
*FlagApi* | [**SetFlagEnabled**](docs/FlagApi.md#setflagenabled) | **Put** /flags/{flagID}/enabled | 
*SegmentApi* | [**CreateSegment**](docs/SegmentApi.md#createsegment) | **Post** /flags/{flagID}/segments | 
*SegmentApi* | [**DeleteSegment**](docs/SegmentApi.md#deletesegment) | **Delete** /flags/{flagID}/segments/{segmentID} | 
*SegmentApi* | [**FindSegments**](docs/SegmentApi.md#findsegments) | **Get** /flags/{flagID}/segments | 
*SegmentApi* | [**PutSegment**](docs/SegmentApi.md#putsegment) | **Put** /flags/{flagID}/segments/{segmentID} | 
*SegmentApi* | [**PutSegmentsReorder**](docs/SegmentApi.md#putsegmentsreorder) | **Put** /flags/{flagID}/segments/reorder | 
*VariantApi* | [**CreateVariant**](docs/VariantApi.md#createvariant) | **Post** /flags/{flagID}/variants | 
*VariantApi* | [**DeleteVariant**](docs/VariantApi.md#deletevariant) | **Delete** /flags/{flagID}/variants/{variantID} | 
*VariantApi* | [**FindVariants**](docs/VariantApi.md#findvariants) | **Get** /flags/{flagID}/variants | 
*VariantApi* | [**PutVariant**](docs/VariantApi.md#putvariant) | **Put** /flags/{flagID}/variants/{variantID} | 


## Documentation For Models

 - [Constraint](docs/Constraint.md)
 - [CreateConstraintRequest](docs/CreateConstraintRequest.md)
 - [CreateFlagRequest](docs/CreateFlagRequest.md)
 - [CreateSegmentRequest](docs/CreateSegmentRequest.md)
 - [CreateVariantRequest](docs/CreateVariantRequest.md)
 - [Distribution](docs/Distribution.md)
 - [EvalContext](docs/EvalContext.md)
 - [EvalDebugLog](docs/EvalDebugLog.md)
 - [EvalResult](docs/EvalResult.md)
 - [EvaluationBatchRequest](docs/EvaluationBatchRequest.md)
 - [EvaluationBatchResponse](docs/EvaluationBatchResponse.md)
 - [EvaluationEntity](docs/EvaluationEntity.md)
 - [Flag](docs/Flag.md)
 - [FlagSnapshot](docs/FlagSnapshot.md)
 - [ModelError](docs/ModelError.md)
 - [PutDistributionsRequest](docs/PutDistributionsRequest.md)
 - [PutFlagRequest](docs/PutFlagRequest.md)
 - [PutSegmentReorderRequest](docs/PutSegmentReorderRequest.md)
 - [PutSegmentRequest](docs/PutSegmentRequest.md)
 - [PutVariantRequest](docs/PutVariantRequest.md)
 - [Segment](docs/Segment.md)
 - [SegmentDebugLog](docs/SegmentDebugLog.md)
 - [SetFlagEnabledRequest](docs/SetFlagEnabledRequest.md)
 - [Variant](docs/Variant.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



