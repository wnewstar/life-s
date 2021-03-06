package logic

import (
    "strings"
    "strconv"
    ModelFore "life/app/fore/model"
)

type Conf struct {}

type Tree struct {
    List                        map[uint]*Tree              `json:"list" binding:"required"`
    Conf                        ModelFore.Conf               `json:"conf" binding:"required"`
}

func (*Conf) GetTree() (map[uint]*Tree) {
    var Mconf ModelFore.Conf

    var Trees = make(map[uint]*Tree)
    var Temps = make(map[uint]*Tree)

    confs := Mconf.FindAll()

    for i := range confs {
        Temps[confs[i].Id] = &Tree{ Conf: confs[i], List: make(map[uint]*Tree) }
    }

    for i := range confs {
        if _, ok := Temps[confs[i].Pid]; ok {
            Temps[confs[i].Pid].List[confs[i].Id] = Temps[confs[i].Id]
        }
    }

    for i := range confs {
        if _, ok := Temps[confs[i].Pid]; !ok {
            Trees[confs[i].Id] = &Tree{ Conf: confs[i], List: Temps[confs[i].Id].List }
        }
    }

    return Trees
}

func (*Conf) GetList() (map[uint]*ModelFore.Conf) {
    var Mconf ModelFore.Conf

    var Confs = make(map[uint]*ModelFore.Conf)

    temps := Mconf.FindAll()

    for i := range temps {
        id := temps[i].Id
        Confs[id] = &temps[i]
        Confs[id].Path = strings.Join([]string{ Confs[id].Path, strconv.FormatUint(uint64(id), 10) }, "-")
    }

    return Confs
}

func (*Conf) GetTreeByUid(uid uint) (map[uint]*Tree) {
    var Mconf ModelFore.Conf

    var Trees = make(map[uint]*Tree)
    var Temps = make(map[uint]*Tree)

    confs := Mconf.FindSetByUid(uid)

    for i := range confs {
        Temps[confs[i].Id] = &Tree{ Conf: confs[i], List: make(map[uint]*Tree) }
    }

    for i := range confs {
        if _, ok := Temps[confs[i].Pid]; ok {
            Temps[confs[i].Pid].List[confs[i].Id] = Temps[confs[i].Id]
        }
    }

    for i := range confs {
        if _, ok := Temps[confs[i].Pid]; !ok {
            Trees[confs[i].Id] = &Tree{ Conf: confs[i], List: Temps[confs[i].Id].List }
        }
    }

    return Trees
}

func (*Conf) GetListByUid(uid uint) (map[uint]*ModelFore.Conf) {
    var Mconf ModelFore.Conf

    var Confs = make(map[uint]*ModelFore.Conf)

    temps := Mconf.FindSetByUid(uid)

    for i := range temps {
        id := temps[i].Id
        Confs[id] = &temps[i]
        Confs[id].Path = strings.Join([]string{ Confs[id].Path, strconv.FormatUint(uint64(id), 10) }, "-")
    }

    return Confs
}
