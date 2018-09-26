package resource

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

var stringTests = []struct {
	x ResId
	s string
}{
	{ResId{gvk: schema.GroupVersionKind{Group: "g", Version: "v", Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "g_v_k_ns_p_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{Version: "v", Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "_v_k_ns_p_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "__k_ns_p_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm", prefix: "p", namespace: "ns"}, "___ns_p_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm", prefix: "p"}, "____p_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm"}, "_____nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{}}, "_____.yaml"},
	{ResId{}, "_____.yaml"},
}

func TestString(t *testing.T) {
	for _, hey := range stringTests {
		if hey.x.String() != hey.s {
			t.Fatalf("Actual: %v,  Expected: '%s'", hey.x, hey.s)
		}
	}
}

var gvknStringTests = []struct {
	x ResId
	s string
}{
	{ResId{gvk: schema.GroupVersionKind{Group: "g", Version: "v", Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "g_v_k_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{Version: "v", Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "v_k_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{Kind: "k"},
		name: "nm", prefix: "p", namespace: "ns"}, "_k_nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm", prefix: "p", namespace: "ns"}, "__nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm", prefix: "p"}, "__nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{},
		name: "nm"}, "__nm.yaml"},
	{ResId{gvk: schema.GroupVersionKind{}}, "__.yaml"},
	{ResId{}, "__.yaml"},
}

func TestGvknString(t *testing.T) {
	for _, hey := range gvknStringTests {
		if hey.x.GvknString() != hey.s {
			t.Fatalf("Actual: %s,  Expected: '%s'", hey.x.GvknString(), hey.s)
		}
	}
}

var GvknEqualsTest = []struct {
	x1 ResId
	x2 ResId
}{
	{ResId{gvk: schema.GroupVersionKind{Group: "g", Version: "v", Kind: "k"},
		name: "nm", prefix: "AA", namespace: "X"},
		ResId{gvk: schema.GroupVersionKind{Group: "g", Version: "v", Kind: "k"},
			name: "nm", prefix: "BB", namespace: "Z"}},
	{ResId{gvk: schema.GroupVersionKind{Version: "v", Kind: "k"},
		name: "nm", prefix: "AA", namespace: "X"},
		ResId{gvk: schema.GroupVersionKind{Version: "v", Kind: "k"},
			name: "nm", prefix: "BB", namespace: "Z"}},
	{ResId{gvk: schema.GroupVersionKind{Kind: "k"},
		name: "nm", prefix: "AA", namespace: "X"},
		ResId{gvk: schema.GroupVersionKind{Kind: "k"},
			name: "nm", prefix: "BB", namespace: "Z"}},
	{ResId{name: "nm", prefix: "AA", namespace: "X"},
		ResId{name: "nm", prefix: "BB", namespace: "Z"}},
}

func TestEquals(t *testing.T) {
	for _, hey := range GvknEqualsTest {
		if !hey.x1.GvknEquals(hey.x2) {
			t.Fatalf("%v should equal %v", hey.x1, hey.x2)
		}
	}
}