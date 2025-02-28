// Code generated by ent, DO NOT EDIT.

package item

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldID, id))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldHash, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldTitle, v))
}

// Dt applies equality check predicate on the "dt" field. It's identical to DtEQ.
func Dt(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldDt, v))
}

// Cat applies equality check predicate on the "cat" field. It's identical to CatEQ.
func Cat(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldCat, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldSize, v))
}

// ExtID applies equality check predicate on the "ext_id" field. It's identical to ExtIDEQ.
func ExtID(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldExtID, v))
}

// Imdb applies equality check predicate on the "imdb" field. It's identical to ImdbEQ.
func Imdb(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldImdb, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldHash, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldTitle, v))
}

// DtEQ applies the EQ predicate on the "dt" field.
func DtEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldDt, v))
}

// DtNEQ applies the NEQ predicate on the "dt" field.
func DtNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldDt, v))
}

// DtIn applies the In predicate on the "dt" field.
func DtIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldDt, vs...))
}

// DtNotIn applies the NotIn predicate on the "dt" field.
func DtNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldDt, vs...))
}

// DtGT applies the GT predicate on the "dt" field.
func DtGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldDt, v))
}

// DtGTE applies the GTE predicate on the "dt" field.
func DtGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldDt, v))
}

// DtLT applies the LT predicate on the "dt" field.
func DtLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldDt, v))
}

// DtLTE applies the LTE predicate on the "dt" field.
func DtLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldDt, v))
}

// DtContains applies the Contains predicate on the "dt" field.
func DtContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldDt, v))
}

// DtHasPrefix applies the HasPrefix predicate on the "dt" field.
func DtHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldDt, v))
}

// DtHasSuffix applies the HasSuffix predicate on the "dt" field.
func DtHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldDt, v))
}

// DtEqualFold applies the EqualFold predicate on the "dt" field.
func DtEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldDt, v))
}

// DtContainsFold applies the ContainsFold predicate on the "dt" field.
func DtContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldDt, v))
}

// CatEQ applies the EQ predicate on the "cat" field.
func CatEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldCat, v))
}

// CatNEQ applies the NEQ predicate on the "cat" field.
func CatNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldCat, v))
}

// CatIn applies the In predicate on the "cat" field.
func CatIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldCat, vs...))
}

// CatNotIn applies the NotIn predicate on the "cat" field.
func CatNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldCat, vs...))
}

// CatGT applies the GT predicate on the "cat" field.
func CatGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldCat, v))
}

// CatGTE applies the GTE predicate on the "cat" field.
func CatGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldCat, v))
}

// CatLT applies the LT predicate on the "cat" field.
func CatLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldCat, v))
}

// CatLTE applies the LTE predicate on the "cat" field.
func CatLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldCat, v))
}

// CatContains applies the Contains predicate on the "cat" field.
func CatContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldCat, v))
}

// CatHasPrefix applies the HasPrefix predicate on the "cat" field.
func CatHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldCat, v))
}

// CatHasSuffix applies the HasSuffix predicate on the "cat" field.
func CatHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldCat, v))
}

// CatEqualFold applies the EqualFold predicate on the "cat" field.
func CatEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldCat, v))
}

// CatContainsFold applies the ContainsFold predicate on the "cat" field.
func CatContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldCat, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldSize, v))
}

// ExtIDEQ applies the EQ predicate on the "ext_id" field.
func ExtIDEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldExtID, v))
}

// ExtIDNEQ applies the NEQ predicate on the "ext_id" field.
func ExtIDNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldExtID, v))
}

// ExtIDIn applies the In predicate on the "ext_id" field.
func ExtIDIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldExtID, vs...))
}

// ExtIDNotIn applies the NotIn predicate on the "ext_id" field.
func ExtIDNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldExtID, vs...))
}

// ExtIDGT applies the GT predicate on the "ext_id" field.
func ExtIDGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldExtID, v))
}

// ExtIDGTE applies the GTE predicate on the "ext_id" field.
func ExtIDGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldExtID, v))
}

// ExtIDLT applies the LT predicate on the "ext_id" field.
func ExtIDLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldExtID, v))
}

// ExtIDLTE applies the LTE predicate on the "ext_id" field.
func ExtIDLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldExtID, v))
}

// ExtIDContains applies the Contains predicate on the "ext_id" field.
func ExtIDContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldExtID, v))
}

// ExtIDHasPrefix applies the HasPrefix predicate on the "ext_id" field.
func ExtIDHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldExtID, v))
}

// ExtIDHasSuffix applies the HasSuffix predicate on the "ext_id" field.
func ExtIDHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldExtID, v))
}

// ExtIDIsNil applies the IsNil predicate on the "ext_id" field.
func ExtIDIsNil() predicate.Item {
	return predicate.Item(sql.FieldIsNull(FieldExtID))
}

// ExtIDNotNil applies the NotNil predicate on the "ext_id" field.
func ExtIDNotNil() predicate.Item {
	return predicate.Item(sql.FieldNotNull(FieldExtID))
}

// ExtIDEqualFold applies the EqualFold predicate on the "ext_id" field.
func ExtIDEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldExtID, v))
}

// ExtIDContainsFold applies the ContainsFold predicate on the "ext_id" field.
func ExtIDContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldExtID, v))
}

// ImdbEQ applies the EQ predicate on the "imdb" field.
func ImdbEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldImdb, v))
}

// ImdbNEQ applies the NEQ predicate on the "imdb" field.
func ImdbNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldImdb, v))
}

// ImdbIn applies the In predicate on the "imdb" field.
func ImdbIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldImdb, vs...))
}

// ImdbNotIn applies the NotIn predicate on the "imdb" field.
func ImdbNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldImdb, vs...))
}

// ImdbGT applies the GT predicate on the "imdb" field.
func ImdbGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldImdb, v))
}

// ImdbGTE applies the GTE predicate on the "imdb" field.
func ImdbGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldImdb, v))
}

// ImdbLT applies the LT predicate on the "imdb" field.
func ImdbLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldImdb, v))
}

// ImdbLTE applies the LTE predicate on the "imdb" field.
func ImdbLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldImdb, v))
}

// ImdbContains applies the Contains predicate on the "imdb" field.
func ImdbContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldImdb, v))
}

// ImdbHasPrefix applies the HasPrefix predicate on the "imdb" field.
func ImdbHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldImdb, v))
}

// ImdbHasSuffix applies the HasSuffix predicate on the "imdb" field.
func ImdbHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldImdb, v))
}

// ImdbIsNil applies the IsNil predicate on the "imdb" field.
func ImdbIsNil() predicate.Item {
	return predicate.Item(sql.FieldIsNull(FieldImdb))
}

// ImdbNotNil applies the NotNil predicate on the "imdb" field.
func ImdbNotNil() predicate.Item {
	return predicate.Item(sql.FieldNotNull(FieldImdb))
}

// ImdbEqualFold applies the EqualFold predicate on the "imdb" field.
func ImdbEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldImdb, v))
}

// ImdbContainsFold applies the ContainsFold predicate on the "imdb" field.
func ImdbContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldImdb, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(sql.NotPredicates(p))
}
