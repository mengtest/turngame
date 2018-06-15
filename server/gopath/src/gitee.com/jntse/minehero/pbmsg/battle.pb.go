// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battle.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BattleUser struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Gold             *int32  `protobuf:"varint,2,opt,name=gold" json:"gold,omitempty"`
	Stepindex        *int32  `protobuf:"varint,3,opt,name=stepindex" json:"stepindex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BattleUser) Reset()                    { *m = BattleUser{} }
func (m *BattleUser) String() string            { return proto.CompactTextString(m) }
func (*BattleUser) ProtoMessage()               {}
func (*BattleUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BattleUser) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BattleUser) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *BattleUser) GetStepindex() int32 {
	if m != nil && m.Stepindex != nil {
		return *m.Stepindex
	}
	return 0
}

type GridItem struct {
	Index            *int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Id               *int32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Num              *int32 `protobuf:"varint,3,opt,name=num" json:"num,omitempty"`
	Gridtype         *int32 `protobuf:"varint,4,opt,name=gridtype" json:"gridtype,omitempty"`
	Control          *bool  `protobuf:"varint,5,opt,name=control" json:"control,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GridItem) Reset()                    { *m = GridItem{} }
func (m *GridItem) String() string            { return proto.CompactTextString(m) }
func (*GridItem) ProtoMessage()               {}
func (*GridItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GridItem) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *GridItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GridItem) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *GridItem) GetGridtype() int32 {
	if m != nil && m.Gridtype != nil {
		return *m.Gridtype
	}
	return 0
}

func (m *GridItem) GetControl() bool {
	if m != nil && m.Control != nil {
		return *m.Control
	}
	return false
}

// 发送个人信息到游戏房间
type BT_UploadGameUser struct {
	Roomid           *int64     `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Bin              *Serialize `protobuf:"bytes,2,opt,name=bin" json:"bin,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *BT_UploadGameUser) Reset()                    { *m = BT_UploadGameUser{} }
func (m *BT_UploadGameUser) String() string            { return proto.CompactTextString(m) }
func (*BT_UploadGameUser) ProtoMessage()               {}
func (*BT_UploadGameUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BT_UploadGameUser) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_UploadGameUser) GetBin() *Serialize {
	if m != nil {
		return m.Bin
	}
	return nil
}

// 房间存在，直接进入房间, 客户端参数留空
type BT_ReqEnterRoom struct {
	Roomid           *int64  `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Userid           *uint64 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
	Token            *string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_ReqEnterRoom) Reset()                    { *m = BT_ReqEnterRoom{} }
func (m *BT_ReqEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*BT_ReqEnterRoom) ProtoMessage()               {}
func (*BT_ReqEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BT_ReqEnterRoom) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_ReqEnterRoom) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_ReqEnterRoom) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 房间初始化
type BT_GameInit struct {
	Roomid           *int64      `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Ownerid          *uint64     `protobuf:"varint,2,opt,name=ownerid" json:"ownerid,omitempty"`
	Gamekind         *int32      `protobuf:"varint,3,opt,name=gamekind" json:"gamekind,omitempty"`
	Listitem         []*GridItem `protobuf:"bytes,4,rep,name=listitem" json:"listitem,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *BT_GameInit) Reset()                    { *m = BT_GameInit{} }
func (m *BT_GameInit) String() string            { return proto.CompactTextString(m) }
func (*BT_GameInit) ProtoMessage()               {}
func (*BT_GameInit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BT_GameInit) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_GameInit) GetOwnerid() uint64 {
	if m != nil && m.Ownerid != nil {
		return *m.Ownerid
	}
	return 0
}

func (m *BT_GameInit) GetGamekind() int32 {
	if m != nil && m.Gamekind != nil {
		return *m.Gamekind
	}
	return 0
}

func (m *BT_GameInit) GetListitem() []*GridItem {
	if m != nil {
		return m.Listitem
	}
	return nil
}

// 同步玩家数据
type BT_SendBattleUser struct {
	// optional int64  roomid = 1;
	Ownerid          *uint64 `protobuf:"varint,2,opt,name=ownerid" json:"ownerid,omitempty"`
	Gold             *uint32 `protobuf:"varint,3,opt,name=gold" json:"gold,omitempty"`
	Coupon           *uint32 `protobuf:"varint,4,opt,name=coupon" json:"coupon,omitempty"`
	Yuanbao          *uint32 `protobuf:"varint,5,opt,name=yuanbao" json:"yuanbao,omitempty"`
	Level            *uint32 `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
	Freestep         *int32  `protobuf:"varint,7,opt,name=freestep" json:"freestep,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_SendBattleUser) Reset()                    { *m = BT_SendBattleUser{} }
func (m *BT_SendBattleUser) String() string            { return proto.CompactTextString(m) }
func (*BT_SendBattleUser) ProtoMessage()               {}
func (*BT_SendBattleUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BT_SendBattleUser) GetOwnerid() uint64 {
	if m != nil && m.Ownerid != nil {
		return *m.Ownerid
	}
	return 0
}

func (m *BT_SendBattleUser) GetGold() uint32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *BT_SendBattleUser) GetCoupon() uint32 {
	if m != nil && m.Coupon != nil {
		return *m.Coupon
	}
	return 0
}

func (m *BT_SendBattleUser) GetYuanbao() uint32 {
	if m != nil && m.Yuanbao != nil {
		return *m.Yuanbao
	}
	return 0
}

func (m *BT_SendBattleUser) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *BT_SendBattleUser) GetFreestep() int32 {
	if m != nil && m.Freestep != nil {
		return *m.Freestep
	}
	return 0
}

// 游戏开始
type BT_GameStart struct {
	Roomid           *int64  `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Ownerid          *uint64 `protobuf:"varint,2,opt,name=ownerid" json:"ownerid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_GameStart) Reset()                    { *m = BT_GameStart{} }
func (m *BT_GameStart) String() string            { return proto.CompactTextString(m) }
func (*BT_GameStart) ProtoMessage()               {}
func (*BT_GameStart) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *BT_GameStart) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_GameStart) GetOwnerid() uint64 {
	if m != nil && m.Ownerid != nil {
		return *m.Ownerid
	}
	return 0
}

// 游戏结束
type BT_GameEnd struct {
	Roomid           *int64     `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Ownerid          *uint64    `protobuf:"varint,2,opt,name=ownerid" json:"ownerid,omitempty"`
	Reason           *string    `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	Bin              *Serialize `protobuf:"bytes,4,opt,name=bin" json:"bin,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *BT_GameEnd) Reset()                    { *m = BT_GameEnd{} }
func (m *BT_GameEnd) String() string            { return proto.CompactTextString(m) }
func (*BT_GameEnd) ProtoMessage()               {}
func (*BT_GameEnd) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BT_GameEnd) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_GameEnd) GetOwnerid() uint64 {
	if m != nil && m.Ownerid != nil {
		return *m.Ownerid
	}
	return 0
}

func (m *BT_GameEnd) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func (m *BT_GameEnd) GetBin() *Serialize {
	if m != nil {
		return m.Bin
	}
	return nil
}

// 通知客户端游戏结束
type BT_GameOver struct {
	Roomid           *int64 `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BT_GameOver) Reset()                    { *m = BT_GameOver{} }
func (m *BT_GameOver) String() string            { return proto.CompactTextString(m) }
func (*BT_GameOver) ProtoMessage()               {}
func (*BT_GameOver) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *BT_GameOver) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

// C2R 起跳检查
type BT_JumpPreCheck struct {
	Roomid           *int64  `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Userid           *uint64 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
	Token            *string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_JumpPreCheck) Reset()                    { *m = BT_JumpPreCheck{} }
func (m *BT_JumpPreCheck) String() string            { return proto.CompactTextString(m) }
func (*BT_JumpPreCheck) ProtoMessage()               {}
func (*BT_JumpPreCheck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *BT_JumpPreCheck) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_JumpPreCheck) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_JumpPreCheck) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// R2C 返回起跳检查
type BT_RetJumpPreCheck struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Errcode          *string `protobuf:"bytes,2,opt,name=errcode" json:"errcode,omitempty"`
	Dice             *int32  `protobuf:"varint,3,opt,name=dice" json:"dice,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_RetJumpPreCheck) Reset()                    { *m = BT_RetJumpPreCheck{} }
func (m *BT_RetJumpPreCheck) String() string            { return proto.CompactTextString(m) }
func (*BT_RetJumpPreCheck) ProtoMessage()               {}
func (*BT_RetJumpPreCheck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *BT_RetJumpPreCheck) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_RetJumpPreCheck) GetErrcode() string {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return ""
}

func (m *BT_RetJumpPreCheck) GetDice() int32 {
	if m != nil && m.Dice != nil {
		return *m.Dice
	}
	return 0
}

// C2R 请求行进, 参数客户端留空
type BT_ReqJumpStep struct {
	Roomid           *int64  `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Userid           *uint64 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
	Stepnum          *int32  `protobuf:"varint,3,opt,name=stepnum" json:"stepnum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_ReqJumpStep) Reset()                    { *m = BT_ReqJumpStep{} }
func (m *BT_ReqJumpStep) String() string            { return proto.CompactTextString(m) }
func (*BT_ReqJumpStep) ProtoMessage()               {}
func (*BT_ReqJumpStep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *BT_ReqJumpStep) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_ReqJumpStep) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_ReqJumpStep) GetStepnum() int32 {
	if m != nil && m.Stepnum != nil {
		return *m.Stepnum
	}
	return 0
}

// R2C 返回行进结果
type BT_RetJumpStep struct {
	// optional int64 roomid = 1;
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Stepindex        *int32  `protobuf:"varint,2,opt,name=stepindex" json:"stepindex,omitempty"`
	Fakelist         []int32 `protobuf:"varint,3,rep,name=fakelist" json:"fakelist,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_RetJumpStep) Reset()                    { *m = BT_RetJumpStep{} }
func (m *BT_RetJumpStep) String() string            { return proto.CompactTextString(m) }
func (*BT_RetJumpStep) ProtoMessage()               {}
func (*BT_RetJumpStep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *BT_RetJumpStep) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_RetJumpStep) GetStepindex() int32 {
	if m != nil && m.Stepindex != nil {
		return *m.Stepindex
	}
	return 0
}

func (m *BT_RetJumpStep) GetFakelist() []int32 {
	if m != nil {
		return m.Fakelist
	}
	return nil
}

// C2R 退出游戏房间，游戏结束,参数客户端留空
type BT_ReqQuitGameRoom struct {
	Roomid           *int64  `protobuf:"varint,1,opt,name=roomid" json:"roomid,omitempty"`
	Userid           *uint64 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BT_ReqQuitGameRoom) Reset()                    { *m = BT_ReqQuitGameRoom{} }
func (m *BT_ReqQuitGameRoom) String() string            { return proto.CompactTextString(m) }
func (*BT_ReqQuitGameRoom) ProtoMessage()               {}
func (*BT_ReqQuitGameRoom) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *BT_ReqQuitGameRoom) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *BT_ReqQuitGameRoom) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

// R2C 获得道具
type BT_PickItem struct {
	Userid           *uint64   `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Item             *GridItem `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BT_PickItem) Reset()                    { *m = BT_PickItem{} }
func (m *BT_PickItem) String() string            { return proto.CompactTextString(m) }
func (*BT_PickItem) ProtoMessage()               {}
func (*BT_PickItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *BT_PickItem) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *BT_PickItem) GetItem() *GridItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*BattleUser)(nil), "msg.BattleUser")
	proto.RegisterType((*GridItem)(nil), "msg.GridItem")
	proto.RegisterType((*BT_UploadGameUser)(nil), "msg.BT_UploadGameUser")
	proto.RegisterType((*BT_ReqEnterRoom)(nil), "msg.BT_ReqEnterRoom")
	proto.RegisterType((*BT_GameInit)(nil), "msg.BT_GameInit")
	proto.RegisterType((*BT_SendBattleUser)(nil), "msg.BT_SendBattleUser")
	proto.RegisterType((*BT_GameStart)(nil), "msg.BT_GameStart")
	proto.RegisterType((*BT_GameEnd)(nil), "msg.BT_GameEnd")
	proto.RegisterType((*BT_GameOver)(nil), "msg.BT_GameOver")
	proto.RegisterType((*BT_JumpPreCheck)(nil), "msg.BT_JumpPreCheck")
	proto.RegisterType((*BT_RetJumpPreCheck)(nil), "msg.BT_RetJumpPreCheck")
	proto.RegisterType((*BT_ReqJumpStep)(nil), "msg.BT_ReqJumpStep")
	proto.RegisterType((*BT_RetJumpStep)(nil), "msg.BT_RetJumpStep")
	proto.RegisterType((*BT_ReqQuitGameRoom)(nil), "msg.BT_ReqQuitGameRoom")
	proto.RegisterType((*BT_PickItem)(nil), "msg.BT_PickItem")
}

func init() { proto.RegisterFile("battle.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6f, 0xda, 0x3e,
	0x14, 0x17, 0x24, 0x14, 0x78, 0x10, 0xd2, 0xe6, 0x14, 0x7d, 0xab, 0xaf, 0x86, 0x7c, 0xe2, 0xc4,
	0xa4, 0x6a, 0xa7, 0x49, 0x93, 0x3a, 0x2a, 0x54, 0x75, 0x97, 0x75, 0x05, 0x2e, 0xbb, 0x20, 0x13,
	0xbf, 0x32, 0x8b, 0xd8, 0x4e, 0x1d, 0xa7, 0x5b, 0xf7, 0xd7, 0x4f, 0x36, 0x0e, 0x83, 0x32, 0x89,
	0xed, 0x18, 0xe7, 0xbd, 0x8f, 0x3f, 0xbf, 0x0c, 0xfd, 0x15, 0x35, 0x26, 0xc7, 0x71, 0xa1, 0x95,
	0x51, 0x49, 0x20, 0xca, 0xf5, 0x7f, 0x71, 0x89, 0x9a, 0xd3, 0x9c, 0xff, 0xf4, 0xa7, 0xe4, 0x03,
	0xc0, 0xc4, 0x4d, 0x2d, 0x4a, 0xd4, 0xc9, 0x00, 0xce, 0x2a, 0x3b, 0xc1, 0xd2, 0xc6, 0xb0, 0x31,
	0x0a, 0x93, 0x3e, 0x84, 0x6b, 0x95, 0xb3, 0xb4, 0x39, 0x6c, 0x8c, 0x5a, 0xc9, 0x05, 0x74, 0x4b,
	0x83, 0x05, 0x97, 0x0c, 0x7f, 0xa4, 0x81, 0x3d, 0x22, 0x0b, 0xe8, 0xdc, 0x6a, 0xce, 0xee, 0x0c,
	0x8a, 0x24, 0x82, 0xd6, 0xf6, 0x57, 0xc3, 0x4d, 0x03, 0x34, 0x79, 0xbd, 0xd9, 0x83, 0x40, 0x56,
	0x62, 0xbb, 0x93, 0x9c, 0x43, 0x67, 0xad, 0x39, 0x33, 0x2f, 0x05, 0xa6, 0xa1, 0x3b, 0x89, 0xa1,
	0x9d, 0x29, 0x69, 0xb4, 0xca, 0xd3, 0xd6, 0xb0, 0x31, 0xea, 0x90, 0x6b, 0xb8, 0x98, 0xcc, 0x97,
	0x8b, 0x22, 0x57, 0x94, 0xdd, 0x52, 0xb1, 0x23, 0xa7, 0x95, 0x12, 0x9e, 0x5c, 0x90, 0x5c, 0x42,
	0xb0, 0xe2, 0xd2, 0xdd, 0xd0, 0xbb, 0x1a, 0x8c, 0x45, 0xb9, 0x1e, 0xcf, 0x6a, 0x75, 0xe4, 0x1a,
	0xe2, 0xc9, 0x7c, 0xf9, 0x80, 0x4f, 0x53, 0x69, 0x50, 0x3f, 0x28, 0x25, 0x8e, 0xf6, 0x7f, 0x8b,
	0x6d, 0x3a, 0xb1, 0x11, 0xb4, 0x8c, 0xda, 0xa0, 0x74, 0x34, 0xbb, 0x84, 0x42, 0x6f, 0x32, 0x5f,
	0xda, 0xdb, 0xef, 0x24, 0x37, 0x47, 0xdb, 0x31, 0xb4, 0xd5, 0x77, 0xb9, 0xb7, 0x6e, 0x65, 0x51,
	0x81, 0x1b, 0x2e, 0x99, 0x17, 0xfa, 0x06, 0x3a, 0x39, 0x2f, 0x0d, 0x37, 0x28, 0xd2, 0x70, 0x18,
	0x8c, 0x7a, 0x57, 0x91, 0x63, 0x59, 0x3b, 0x46, 0x8c, 0x93, 0x39, 0x43, 0xc9, 0xf6, 0x32, 0x38,
	0x02, 0xae, 0x43, 0xb0, 0xa0, 0x91, 0xe5, 0x91, 0xa9, 0xaa, 0x50, 0xd2, 0x79, 0x17, 0xd9, 0xf1,
	0x97, 0x8a, 0xca, 0x15, 0x55, 0xce, 0xbb, 0xc8, 0xca, 0xc8, 0xf1, 0x19, 0xf3, 0xf4, 0xcc, 0x7d,
	0x9e, 0x43, 0xe7, 0x51, 0x23, 0xda, 0xe0, 0xd2, 0xb6, 0xcb, 0xec, 0x2d, 0xf4, 0xbd, 0xb0, 0x99,
	0xa1, 0xfa, 0xb4, 0x32, 0xf2, 0x15, 0xc0, 0x2f, 0x4c, 0x25, 0x3b, 0x6d, 0x84, 0x1d, 0x40, 0x5a,
	0x2a, 0x6f, 0x64, 0x9d, 0x53, 0xf8, 0xc7, 0x9c, 0xfe, 0xdf, 0xb9, 0xfc, 0xf9, 0xf9, 0x38, 0x63,
	0x1f, 0xe3, 0xa7, 0x4a, 0x14, 0xf7, 0x1a, 0x6f, 0xbe, 0x61, 0xb6, 0xf9, 0xd7, 0x18, 0x6f, 0x20,
	0x71, 0x45, 0x30, 0xaf, 0x41, 0x0e, 0x8a, 0x1e, 0x43, 0x1b, 0xb5, 0xce, 0x14, 0x43, 0x87, 0xd2,
	0xb5, 0xa6, 0x33, 0x9e, 0xa1, 0xaf, 0xf9, 0x47, 0x18, 0x6c, 0xdb, 0x64, 0x41, 0x66, 0x06, 0x8b,
	0x93, 0x2c, 0x62, 0x68, 0x5b, 0xcb, 0x77, 0xad, 0x27, 0x53, 0x0f, 0x61, 0xf6, 0x21, 0x0e, 0x38,
	0x1c, 0x3c, 0xaf, 0x66, 0xfd, 0x54, 0x1e, 0xe9, 0x06, 0x6d, 0x8b, 0xd2, 0x60, 0x18, 0x8c, 0x5a,
	0xe4, 0x9d, 0x97, 0xf3, 0xf4, 0xa5, 0xe2, 0xc6, 0xda, 0xf6, 0x37, 0xd5, 0x26, 0xef, 0x9d, 0xcb,
	0xf7, 0x3c, 0xdb, 0xb8, 0x97, 0xfa, 0xfa, 0xe6, 0x4b, 0x08, 0x5d, 0x49, 0xb7, 0x4f, 0xe9, 0xb0,
	0xa4, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x84, 0xc6, 0xcb, 0x46, 0x04, 0x00, 0x00,
}