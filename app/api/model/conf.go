package model

type Conf struct {
    Base
    Uid                         uint                        `json:"uid"`
    Pid                         uint                        `json:"pid"`
    Path                        string                      `json:"path" binding:"required"`
    Name                        string                      `gorm:"size:50" json:"name" binding:"required"`
}

func (*Conf) TableName() (string) {

    return "asset_conf"
}

func (*Conf) Create(mconf Conf) (uint) {
    Db.NewRecord(mconf)

    Db.Create(&mconf)

    Db.NewRecord(mconf)

    return mconf.Id
}

func (*Conf) FindAll() ([]Conf) {
    var Mconfs []Conf
    var Where = &Conf{}

    Db.Where(Where).Where("`status` = 0").Order("path asc").Find(&Mconfs)

    return Mconfs
}

func (*Conf) FindSetByUid(uid uint) ([]Conf) {
    var Mconfs []Conf
    var Where = &Conf{ Uid: uid }

    Db.Where(Where).Where("`status` = 0").Order("path asc").Find(&Mconfs)

    return Mconfs
}

func (*Conf) FindOneByIdUid(mconf Conf) (Conf) {
    var Mconf Conf
    var Mbase = Base{ Id: mconf.Id }
    var Where = &Conf{ Base: Mbase, Uid: mconf.Uid }

    Db.Where(Where).Where("`status` = 0").Order("path desc").First(&Mconf)

    return Mconf
}

func (*Conf) SaveOneByIdUid(mconf Conf) (int64) {
    var Mbase = Base{ Id: mconf.Id }
    var Where = &Conf{ Base: Mbase, Uid: mconf.Uid }

    return Db.Model(Conf{}).Where(Where).Where("`status` = 0").Updates(mconf).RowsAffected
}
