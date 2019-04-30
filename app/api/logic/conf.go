package logic

import (
    ApiModel "life-b/app/api/model"
)

type Conf struct {}

type Tree struct {
    List                        map[uint]*Tree              `json:"list" binding:"required"`
    Conf                        ApiModel.Conf               `json:"conf" binding:"required"`
}

func (*Conf) GetList() ([]ApiModel.Conf) {
    var Mconf ApiModel.Conf

    return Mconf.FindAll()
}

func (*Conf) GetTree() (map[uint]*Tree) {
    var Mconf ApiModel.Conf
    var Trees = make(map[uint]*Tree)
    var TreesTemp = make(map[uint]*Tree)

    confs := Mconf.FindAll()

    for i := range confs {
        TreesTemp[confs[i].Id] = &Tree{ Conf: confs[i], List: make(map[uint]*Tree) }
    }

    for i := range confs {
        if _, ok := TreesTemp[confs[i].Pid]; ok {
            TreesTemp[confs[i].Pid].List[confs[i].Id] = TreesTemp[confs[i].Id]
        }
    }

    for i := range confs {
        if _, ok := TreesTemp[confs[i].Pid]; !ok {
            Trees[confs[i].Id] = &Tree{ Conf: confs[i], List: TreesTemp[confs[i].Id].List }
        }
    }

    return Trees
}

func (*Conf) GetListByUid(uid uint) ([]ApiModel.Conf) {
    var Mconf ApiModel.Conf

    return Mconf.FindSetByUid(uid)
}

func (*Conf) GetTreeByUid(uid uint) (map[uint]*Tree) {
    var Mconf ApiModel.Conf
    var Trees = make(map[uint]*Tree)
    var TreesTemp = make(map[uint]*Tree)

    confs := Mconf.FindSetByUid(uid)

    for i := range confs {
        TreesTemp[confs[i].Id] = &Tree{ Conf: confs[i], List: make(map[uint]*Tree) }
    }

    for i := range confs {
        if _, ok := TreesTemp[confs[i].Pid]; ok {
            TreesTemp[confs[i].Pid].List[confs[i].Id] = TreesTemp[confs[i].Id]
        }
    }

    for i := range confs {
        if _, ok := TreesTemp[confs[i].Pid]; !ok {
            Trees[confs[i].Id] = &Tree{ Conf: confs[i], List: TreesTemp[confs[i].Id].List }
        }
    }

    return Trees
}