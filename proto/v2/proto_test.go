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

	basePb := &models.Base{}

	extMap := proto.RegisteredExtensions((*models.Base)(basePb))

	var extSlice []*proto.ExtensionDesc
	for _, v := range extMap {
		st := reflect.TypeOf(v.ExtensionType).Elem()
		if strings.HasSuffix(st.Name(), "Rs") {
			continue
		}
		extSlice = append(extSlice, v)
	}

	sort.Slice(extSlice, func(i, j int) bool {
		return extSlice[i].Field < extSlice[j].Field
	})
	for _, v := range extSlice {
		fmt.Println(v.Field, v.Name)
		printFiled(reflect.TypeOf(v.ExtensionType).Elem())
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

func TestJsonToPb(t *testing.T) {

	oldStruct := &models.CrossMoveCityRq{
		MapId: proto.Int32(26),
		Type:  proto.Int32(1),
		Pos:   proto.Int32(11),
	}

	dataJ, err := json.Marshal(oldStruct)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dataJ))
	extMap := proto.RegisteredExtensions((*models.Base)(nil))
	extensionDescPrt := extMap[6021]

	//newStruct := newStruct(reflect.TypeOf(extensionDescPrt.ExtensionType).Elem())

	var fields []reflect.StructField
	etype := reflect.TypeOf(extensionDescPrt.ExtensionType).Elem()
	for i := 0; i < etype.NumField(); i++ {
		fields = append(fields, etype.Field(i))
	}

	newStruct := CreateStruct(fields)

	fmt.Printf("newStruct %T %v\n", newStruct, newStruct)

	json.Unmarshal(dataJ, &newStruct)

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

func newStruct(t reflect.Type) interface{} {
	// 只能创建一个map
	return reflect.New(t).Elem().Interface()
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
