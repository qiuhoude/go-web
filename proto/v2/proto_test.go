package v2

import (
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"testing"
)

func TestBase(t *testing.T) {

	cmcPb := &models.CrossMoveCityRq{
		MapId: proto.Int32(26),
		Type:  proto.Int32(1),
		Pos:   proto.Int32(11),
	}

	basePb := models.Base{
		Cmd:  proto.Int32(models.E_CrossMoveCityRq_Ext.Field),
		Code: proto.Int32(200),
	}

	ext := &proto.ExtensionDesc{
		ExtendedType:  &basePb,
		ExtensionType: &cmcPb,
		Field:         6021,
		Name:          "CrossMoveCityRq.ext",
		Tag:           "bytes,6021,opt,name=ext",
		Filename:      "Game.proto",
	}

	proto.SetExtension(cmcPb, ext, nil)

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
