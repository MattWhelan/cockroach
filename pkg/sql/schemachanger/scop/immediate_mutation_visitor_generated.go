// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by generate_visitor.go. DO NOT EDIT.

package scop

import "context"

// ImmediateMutationOp is an operation which can be visited by ImmediateMutationVisitor.
type ImmediateMutationOp interface {
	Op
	Visit(context.Context, ImmediateMutationVisitor) error
}

// ImmediateMutationVisitor is a visitor for ImmediateMutationOp operations.
type ImmediateMutationVisitor interface {
	NotImplemented(context.Context, NotImplemented) error
	NotImplementedForPublicObjects(context.Context, NotImplementedForPublicObjects) error
	UndoAllInTxnImmediateMutationOpSideEffects(context.Context, UndoAllInTxnImmediateMutationOpSideEffects) error
	MakeAbsentTempIndexDeleteOnly(context.Context, MakeAbsentTempIndexDeleteOnly) error
	MakeAbsentIndexBackfilling(context.Context, MakeAbsentIndexBackfilling) error
	SetAddedIndexPartialPredicate(context.Context, SetAddedIndexPartialPredicate) error
	MakeDeleteOnlyIndexWriteOnly(context.Context, MakeDeleteOnlyIndexWriteOnly) error
	MakeBackfilledIndexMerging(context.Context, MakeBackfilledIndexMerging) error
	MakeMergedIndexWriteOnly(context.Context, MakeMergedIndexWriteOnly) error
	MakeBackfillingIndexDeleteOnly(context.Context, MakeBackfillingIndexDeleteOnly) error
	MakeValidatedSecondaryIndexPublic(context.Context, MakeValidatedSecondaryIndexPublic) error
	MakeValidatedPrimaryIndexPublic(context.Context, MakeValidatedPrimaryIndexPublic) error
	MakePublicPrimaryIndexWriteOnly(context.Context, MakePublicPrimaryIndexWriteOnly) error
	MarkDescriptorAsPublic(context.Context, MarkDescriptorAsPublic) error
	InsertTemporarySchema(context.Context, InsertTemporarySchema) error
	InsertTemporarySchemaParent(context.Context, InsertTemporarySchemaParent) error
	MarkDescriptorAsDropped(context.Context, MarkDescriptorAsDropped) error
	DrainDescriptorName(context.Context, DrainDescriptorName) error
	AddDescriptorName(context.Context, AddDescriptorName) error
	SetNameInDescriptor(context.Context, SetNameInDescriptor) error
	MakeDeleteOnlyColumnWriteOnly(context.Context, MakeDeleteOnlyColumnWriteOnly) error
	MakePublicSecondaryIndexWriteOnly(context.Context, MakePublicSecondaryIndexWriteOnly) error
	MakeWriteOnlyIndexDeleteOnly(context.Context, MakeWriteOnlyIndexDeleteOnly) error
	RemoveDroppedIndexPartialPredicate(context.Context, RemoveDroppedIndexPartialPredicate) error
	MakeIndexAbsent(context.Context, MakeIndexAbsent) error
	MakeAbsentColumnDeleteOnly(context.Context, MakeAbsentColumnDeleteOnly) error
	UpsertColumnType(context.Context, UpsertColumnType) error
	AddColumnComputeExpression(context.Context, AddColumnComputeExpression) error
	RemoveColumnComputeExpression(context.Context, RemoveColumnComputeExpression) error
	MakeWriteOnlyColumnPublic(context.Context, MakeWriteOnlyColumnPublic) error
	MakePublicColumnWriteOnly(context.Context, MakePublicColumnWriteOnly) error
	MakeWriteOnlyColumnDeleteOnly(context.Context, MakeWriteOnlyColumnDeleteOnly) error
	RemoveDroppedColumnType(context.Context, RemoveDroppedColumnType) error
	MakeDeleteOnlyColumnAbsent(context.Context, MakeDeleteOnlyColumnAbsent) error
	AddOwnerBackReferenceInSequence(context.Context, AddOwnerBackReferenceInSequence) error
	AddSequenceOwner(context.Context, AddSequenceOwner) error
	RemoveOwnerBackReferenceInSequence(context.Context, RemoveOwnerBackReferenceInSequence) error
	RemoveSequenceOwner(context.Context, RemoveSequenceOwner) error
	RemoveCheckConstraint(context.Context, RemoveCheckConstraint) error
	RemoveColumnNotNull(context.Context, RemoveColumnNotNull) error
	AddCheckConstraint(context.Context, AddCheckConstraint) error
	MakeAbsentColumnNotNullWriteOnly(context.Context, MakeAbsentColumnNotNullWriteOnly) error
	MakePublicCheckConstraintValidated(context.Context, MakePublicCheckConstraintValidated) error
	MakePublicColumnNotNullValidated(context.Context, MakePublicColumnNotNullValidated) error
	MakeValidatedCheckConstraintPublic(context.Context, MakeValidatedCheckConstraintPublic) error
	MakeValidatedColumnNotNullPublic(context.Context, MakeValidatedColumnNotNullPublic) error
	AddForeignKeyConstraint(context.Context, AddForeignKeyConstraint) error
	MakeValidatedForeignKeyConstraintPublic(context.Context, MakeValidatedForeignKeyConstraintPublic) error
	MakePublicForeignKeyConstraintValidated(context.Context, MakePublicForeignKeyConstraintValidated) error
	RemoveForeignKeyConstraint(context.Context, RemoveForeignKeyConstraint) error
	RemoveForeignKeyBackReference(context.Context, RemoveForeignKeyBackReference) error
	AddUniqueWithoutIndexConstraint(context.Context, AddUniqueWithoutIndexConstraint) error
	MakeValidatedUniqueWithoutIndexConstraintPublic(context.Context, MakeValidatedUniqueWithoutIndexConstraintPublic) error
	MakePublicUniqueWithoutIndexConstraintValidated(context.Context, MakePublicUniqueWithoutIndexConstraintValidated) error
	RemoveUniqueWithoutIndexConstraint(context.Context, RemoveUniqueWithoutIndexConstraint) error
	RemoveSchemaParent(context.Context, RemoveSchemaParent) error
	AddSchemaParent(context.Context, AddSchemaParent) error
	AddIndexPartitionInfo(context.Context, AddIndexPartitionInfo) error
	AddColumnFamily(context.Context, AddColumnFamily) error
	AssertColumnFamilyIsRemoved(context.Context, AssertColumnFamilyIsRemoved) error
	AddColumnDefaultExpression(context.Context, AddColumnDefaultExpression) error
	RemoveColumnDefaultExpression(context.Context, RemoveColumnDefaultExpression) error
	AddColumnOnUpdateExpression(context.Context, AddColumnOnUpdateExpression) error
	RemoveColumnOnUpdateExpression(context.Context, RemoveColumnOnUpdateExpression) error
	AddTrigger(context.Context, AddTrigger) error
	SetTriggerName(context.Context, SetTriggerName) error
	SetTriggerEnabled(context.Context, SetTriggerEnabled) error
	SetTriggerTiming(context.Context, SetTriggerTiming) error
	SetTriggerEvents(context.Context, SetTriggerEvents) error
	SetTriggerTransition(context.Context, SetTriggerTransition) error
	SetTriggerWhen(context.Context, SetTriggerWhen) error
	SetTriggerFunctionCall(context.Context, SetTriggerFunctionCall) error
	SetTriggerForwardReferences(context.Context, SetTriggerForwardReferences) error
	RemoveTrigger(context.Context, RemoveTrigger) error
	UpdateTableBackReferencesInTypes(context.Context, UpdateTableBackReferencesInTypes) error
	UpdateTypeBackReferencesInTypes(context.Context, UpdateTypeBackReferencesInTypes) error
	RemoveBackReferenceInTypes(context.Context, RemoveBackReferenceInTypes) error
	RemoveBackReferenceInFunctions(context.Context, RemoveBackReferenceInFunctions) error
	UpdateTableBackReferencesInSequences(context.Context, UpdateTableBackReferencesInSequences) error
	RemoveBackReferencesInRelations(context.Context, RemoveBackReferencesInRelations) error
	AddTableConstraintBackReferencesInFunctions(context.Context, AddTableConstraintBackReferencesInFunctions) error
	RemoveTableConstraintBackReferencesFromFunctions(context.Context, RemoveTableConstraintBackReferencesFromFunctions) error
	AddTableColumnBackReferencesInFunctions(context.Context, AddTableColumnBackReferencesInFunctions) error
	RemoveTableColumnBackReferencesInFunctions(context.Context, RemoveTableColumnBackReferencesInFunctions) error
	AddTriggerBackReferencesInRoutines(context.Context, AddTriggerBackReferencesInRoutines) error
	RemoveTriggerBackReferencesInRoutines(context.Context, RemoveTriggerBackReferencesInRoutines) error
	SetColumnName(context.Context, SetColumnName) error
	SetIndexName(context.Context, SetIndexName) error
	SetConstraintName(context.Context, SetConstraintName) error
	DeleteDescriptor(context.Context, DeleteDescriptor) error
	RemoveUserPrivileges(context.Context, RemoveUserPrivileges) error
	RemoveJobStateFromDescriptor(context.Context, RemoveJobStateFromDescriptor) error
	SetJobStateOnDescriptor(context.Context, SetJobStateOnDescriptor) error
	UpsertTableComment(context.Context, UpsertTableComment) error
	RemoveTableComment(context.Context, RemoveTableComment) error
	UpsertTypeComment(context.Context, UpsertTypeComment) error
	RemoveTypeComment(context.Context, RemoveTypeComment) error
	UpsertDatabaseComment(context.Context, UpsertDatabaseComment) error
	RemoveDatabaseComment(context.Context, RemoveDatabaseComment) error
	UpsertSchemaComment(context.Context, UpsertSchemaComment) error
	RemoveSchemaComment(context.Context, RemoveSchemaComment) error
	UpsertIndexComment(context.Context, UpsertIndexComment) error
	RemoveIndexComment(context.Context, RemoveIndexComment) error
	UpsertColumnComment(context.Context, UpsertColumnComment) error
	RemoveColumnComment(context.Context, RemoveColumnComment) error
	UpsertConstraintComment(context.Context, UpsertConstraintComment) error
	RemoveConstraintComment(context.Context, RemoveConstraintComment) error
	AddColumnToIndex(context.Context, AddColumnToIndex) error
	RemoveColumnFromIndex(context.Context, RemoveColumnFromIndex) error
	RemoveObjectParent(context.Context, RemoveObjectParent) error
	CreateFunctionDescriptor(context.Context, CreateFunctionDescriptor) error
	SetFunctionName(context.Context, SetFunctionName) error
	SetFunctionVolatility(context.Context, SetFunctionVolatility) error
	SetFunctionLeakProof(context.Context, SetFunctionLeakProof) error
	SetFunctionNullInputBehavior(context.Context, SetFunctionNullInputBehavior) error
	SetFunctionBody(context.Context, SetFunctionBody) error
	SetFunctionSecurity(context.Context, SetFunctionSecurity) error
	UpdateFunctionTypeReferences(context.Context, UpdateFunctionTypeReferences) error
	UpdateFunctionRelationReferences(context.Context, UpdateFunctionRelationReferences) error
	UpdateTableBackReferencesInRelations(context.Context, UpdateTableBackReferencesInRelations) error
	SetObjectParentID(context.Context, SetObjectParentID) error
	UpdateUserPrivileges(context.Context, UpdateUserPrivileges) error
	UpdateOwner(context.Context, UpdateOwner) error
	CreateSchemaDescriptor(context.Context, CreateSchemaDescriptor) error
	CreateSequenceDescriptor(context.Context, CreateSequenceDescriptor) error
	SetSequenceOptions(context.Context, SetSequenceOptions) error
	InitSequence(context.Context, InitSequence) error
	CreateDatabaseDescriptor(context.Context, CreateDatabaseDescriptor) error
	AddNamedRangeZoneConfig(context.Context, AddNamedRangeZoneConfig) error
	AddDatabaseZoneConfig(context.Context, AddDatabaseZoneConfig) error
	DiscardZoneConfig(context.Context, DiscardZoneConfig) error
	DiscardTableZoneConfig(context.Context, DiscardTableZoneConfig) error
	AddTableZoneConfig(context.Context, AddTableZoneConfig) error
	AddIndexZoneConfig(context.Context, AddIndexZoneConfig) error
	AddPartitionZoneConfig(context.Context, AddPartitionZoneConfig) error
}

// Visit is part of the ImmediateMutationOp interface.
func (op NotImplemented) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.NotImplemented(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op NotImplementedForPublicObjects) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.NotImplementedForPublicObjects(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UndoAllInTxnImmediateMutationOpSideEffects) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UndoAllInTxnImmediateMutationOpSideEffects(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeAbsentTempIndexDeleteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeAbsentTempIndexDeleteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeAbsentIndexBackfilling) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeAbsentIndexBackfilling(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetAddedIndexPartialPredicate) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetAddedIndexPartialPredicate(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeDeleteOnlyIndexWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeDeleteOnlyIndexWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeBackfilledIndexMerging) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeBackfilledIndexMerging(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeMergedIndexWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeMergedIndexWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeBackfillingIndexDeleteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeBackfillingIndexDeleteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedSecondaryIndexPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedSecondaryIndexPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedPrimaryIndexPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedPrimaryIndexPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicPrimaryIndexWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicPrimaryIndexWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MarkDescriptorAsPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MarkDescriptorAsPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op InsertTemporarySchema) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.InsertTemporarySchema(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op InsertTemporarySchemaParent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.InsertTemporarySchemaParent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MarkDescriptorAsDropped) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MarkDescriptorAsDropped(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op DrainDescriptorName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.DrainDescriptorName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddDescriptorName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddDescriptorName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetNameInDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetNameInDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeDeleteOnlyColumnWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeDeleteOnlyColumnWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicSecondaryIndexWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicSecondaryIndexWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeWriteOnlyIndexDeleteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeWriteOnlyIndexDeleteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveDroppedIndexPartialPredicate) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveDroppedIndexPartialPredicate(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeIndexAbsent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeIndexAbsent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeAbsentColumnDeleteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeAbsentColumnDeleteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertColumnType) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertColumnType(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddColumnComputeExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddColumnComputeExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnComputeExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnComputeExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeWriteOnlyColumnPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeWriteOnlyColumnPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicColumnWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicColumnWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeWriteOnlyColumnDeleteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeWriteOnlyColumnDeleteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveDroppedColumnType) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveDroppedColumnType(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeDeleteOnlyColumnAbsent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeDeleteOnlyColumnAbsent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddOwnerBackReferenceInSequence) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddOwnerBackReferenceInSequence(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddSequenceOwner) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddSequenceOwner(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveOwnerBackReferenceInSequence) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveOwnerBackReferenceInSequence(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveSequenceOwner) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveSequenceOwner(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveCheckConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveCheckConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnNotNull) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnNotNull(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddCheckConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddCheckConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeAbsentColumnNotNullWriteOnly) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeAbsentColumnNotNullWriteOnly(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicCheckConstraintValidated) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicCheckConstraintValidated(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicColumnNotNullValidated) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicColumnNotNullValidated(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedCheckConstraintPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedCheckConstraintPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedColumnNotNullPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedColumnNotNullPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddForeignKeyConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddForeignKeyConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedForeignKeyConstraintPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedForeignKeyConstraintPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicForeignKeyConstraintValidated) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicForeignKeyConstraintValidated(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveForeignKeyConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveForeignKeyConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveForeignKeyBackReference) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveForeignKeyBackReference(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddUniqueWithoutIndexConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddUniqueWithoutIndexConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakeValidatedUniqueWithoutIndexConstraintPublic) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakeValidatedUniqueWithoutIndexConstraintPublic(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op MakePublicUniqueWithoutIndexConstraintValidated) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.MakePublicUniqueWithoutIndexConstraintValidated(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveUniqueWithoutIndexConstraint) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveUniqueWithoutIndexConstraint(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveSchemaParent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveSchemaParent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddSchemaParent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddSchemaParent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddIndexPartitionInfo) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddIndexPartitionInfo(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddColumnFamily) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddColumnFamily(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AssertColumnFamilyIsRemoved) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AssertColumnFamilyIsRemoved(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddColumnDefaultExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddColumnDefaultExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnDefaultExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnDefaultExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddColumnOnUpdateExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddColumnOnUpdateExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnOnUpdateExpression) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnOnUpdateExpression(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddTrigger) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddTrigger(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerEnabled) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerEnabled(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerTiming) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerTiming(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerEvents) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerEvents(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerTransition) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerTransition(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerWhen) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerWhen(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerFunctionCall) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerFunctionCall(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetTriggerForwardReferences) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetTriggerForwardReferences(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTrigger) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTrigger(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateTableBackReferencesInTypes) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateTableBackReferencesInTypes(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateTypeBackReferencesInTypes) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateTypeBackReferencesInTypes(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveBackReferenceInTypes) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveBackReferenceInTypes(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveBackReferenceInFunctions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveBackReferenceInFunctions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateTableBackReferencesInSequences) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateTableBackReferencesInSequences(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveBackReferencesInRelations) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveBackReferencesInRelations(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddTableConstraintBackReferencesInFunctions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddTableConstraintBackReferencesInFunctions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTableConstraintBackReferencesFromFunctions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTableConstraintBackReferencesFromFunctions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddTableColumnBackReferencesInFunctions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddTableColumnBackReferencesInFunctions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTableColumnBackReferencesInFunctions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTableColumnBackReferencesInFunctions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddTriggerBackReferencesInRoutines) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddTriggerBackReferencesInRoutines(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTriggerBackReferencesInRoutines) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTriggerBackReferencesInRoutines(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetColumnName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetColumnName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetIndexName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetIndexName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetConstraintName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetConstraintName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op DeleteDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.DeleteDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveUserPrivileges) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveUserPrivileges(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveJobStateFromDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveJobStateFromDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetJobStateOnDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetJobStateOnDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertTableComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertTableComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTableComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTableComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertTypeComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertTypeComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveTypeComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveTypeComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertDatabaseComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertDatabaseComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveDatabaseComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveDatabaseComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertSchemaComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertSchemaComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveSchemaComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveSchemaComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertIndexComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertIndexComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveIndexComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveIndexComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertColumnComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertColumnComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpsertConstraintComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpsertConstraintComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveConstraintComment) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveConstraintComment(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddColumnToIndex) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddColumnToIndex(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveColumnFromIndex) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveColumnFromIndex(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op RemoveObjectParent) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.RemoveObjectParent(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op CreateFunctionDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.CreateFunctionDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionName) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionName(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionVolatility) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionVolatility(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionLeakProof) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionLeakProof(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionNullInputBehavior) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionNullInputBehavior(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionBody) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionBody(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetFunctionSecurity) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetFunctionSecurity(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateFunctionTypeReferences) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateFunctionTypeReferences(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateFunctionRelationReferences) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateFunctionRelationReferences(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateTableBackReferencesInRelations) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateTableBackReferencesInRelations(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetObjectParentID) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetObjectParentID(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateUserPrivileges) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateUserPrivileges(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op UpdateOwner) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.UpdateOwner(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op CreateSchemaDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.CreateSchemaDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op CreateSequenceDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.CreateSequenceDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op SetSequenceOptions) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.SetSequenceOptions(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op InitSequence) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.InitSequence(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op CreateDatabaseDescriptor) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.CreateDatabaseDescriptor(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddNamedRangeZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddNamedRangeZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddDatabaseZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddDatabaseZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op DiscardZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.DiscardZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op DiscardTableZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.DiscardTableZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddTableZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddTableZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddIndexZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddIndexZoneConfig(ctx, op)
}

// Visit is part of the ImmediateMutationOp interface.
func (op AddPartitionZoneConfig) Visit(ctx context.Context, v ImmediateMutationVisitor) error {
	return v.AddPartitionZoneConfig(ctx, op)
}
