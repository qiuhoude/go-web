package v2

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// 打印pb文件中所有协议
func TestBase(t *testing.T) {
	extMap := proto.RegisteredExtensions((*models.Base)(nil))

	var extSlice []*proto.ExtensionDesc
	for _, v := range extMap {
		//st := reflect.TypeOf(v.ExtensionType).Elem()
		//if strings.HasSuffix(st.Name(), "Rs") {
		//	continue
		//}
		extSlice = append(extSlice, v)
	}

	sort.Slice(extSlice, func(i, j int) bool {
		return extSlice[i].Field < extSlice[j].Field
	})
	for _, v := range extSlice {
		fmt.Println(v.Field, v.Name)
		printFiled(reflect.TypeOf(v.ExtensionType).Elem())
		//printFiledJson(v)
	}

}

func printFiled(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !strings.HasPrefix(field.Name, "XXX_") {
			jsonName := strings.Split(field.Tag.Get("json"), ",")[0]
			fmt.Println("\t name:", jsonName, " type:", field.Type)
		}
	}
}

func printFiledJson(pbDesc *proto.ExtensionDesc) {
	t := reflect.TypeOf(pbDesc.ExtensionType).Elem()
	var sb strings.Builder
	sb.WriteString("{")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !strings.HasPrefix(field.Name, "XXX_") {
			jsonName := strings.Split(field.Tag.Get("json"), ",")[0]

			_, _ = fmt.Fprintf(&sb, `"%s"`, jsonName)

		}
	}
	sb.WriteString("}")
}

func TestJsonToPb(t *testing.T) {
	twoIntSlice := make([]*models.TwoInt, 0)
	for i := 0; i < 2; i++ {
		twoIntSlice = append(twoIntSlice, &models.TwoInt{
			V1: proto.Int32(1),
			V2: proto.Int32(2),
		})
	}

	oldStruct := &models.GetTreasureRs{
		IdStatus: twoIntSlice,
		Status:   proto.Int32(1),
		Red:      proto.Bool(false),
	}

	dataJ, err := json.Marshal(oldStruct)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dataJ))
	extMap := proto.RegisteredExtensions((*models.Base)(nil))
	extensionDescPrt := extMap[1256]

	newStruct := NewStruct(reflect.TypeOf(extensionDescPrt.ExtensionType))

	//var fields []reflect.StructField
	//etype := reflect.TypeOf(extensionDescPrt.ExtensionType).Elem()
	//for i := 0; i < etype.NumField(); i++ {
	//	fields = append(fields, etype.Field(i))
	//}
	//

	fmt.Printf("newStruct %T %v\n", newStruct, newStruct)

	_ = json.Unmarshal(dataJ, &newStruct)

	fmt.Printf("oldStruct %T %v\n", oldStruct, oldStruct)
	fmt.Printf("newStruct %T %v\n", newStruct, newStruct)

}

// 创建结构体
func CreateStruct(fields []reflect.StructField) interface{} {
	var structType reflect.Type
	structType = reflect.StructOf(fields)
	so := reflect.New(structType)
	return so.Interface()
}

func NewStruct(t reflect.Type) proto.Message {
	// 只能创建一个map
	return reflect.New(t.Elem()).Interface().(proto.Message)
}

func TestProto(t *testing.T) {

	cmcPb := &models.CrossMoveCityRq{
		MapId: proto.Int32(26),
		Type:  proto.Int32(1),
		Pos:   proto.Int32(11),
	}
	_ = cmcPb

	data, err := proto.Marshal(cmcPb)
	if err != nil {
		t.Fatal("marshaling error: ", err)
	}
	newCmcPb := &models.CrossMoveCityRq{}
	err = proto.Unmarshal(data, newCmcPb)
	if err != nil {
		t.Fatal("unmarshaling error: ", err)
	}
	assert.Equal(t, cmcPb.MapId, newCmcPb.MapId)

	descs, err := proto.ExtensionDescs(cmcPb)
	fmt.Println(descs)

	descV1 := models.E_CrossMoveCityRq_Ext
	fmt.Println(descV1.Field)

}

func Test1(t *testing.T) {
	l1 := numToList(10000000000001)
	l2 := numToList(564)

	t.Log(listToNum(addTwoNumbers(l1, l2)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := listToNum(l1) + listToNum(l2)
	return numToList(s)
}

func numToList(num int) *ListNode {
	var rtl *ListNode
	var nl *ListNode
	h := &ListNode{Val: 0}
	s := num

	for ; s != 0; s /= 10 {
		c := s % 10
		if rtl == nil {
			h = &ListNode{Val: c}
			rtl = h
			nl = rtl
		} else {
			nl = &ListNode{Val: c}
			rtl.Next = nl
			rtl = nl
		}
	}
	return h
}

func listToNum(l *ListNode) int {
	var num int
	i := 1
	tp := l
	for ; tp != nil; tp = tp.Next {
		num += tp.Val * i
		i *= 10
	}
	return num
}
